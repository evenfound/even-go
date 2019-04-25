// Copyright (c) 2018-2019 The Even Foundation developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package ltcnet

import (
	"bytes"
	"encoding/hex"
	"fmt"

	v "github.com/evenfound/even-go/mbnd/common"
	"github.com/go-errors/errors"
	"github.com/ltcsuite/ltcd/blockchain"
	"github.com/ltcsuite/ltcd/btcjson"
	"github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcd/chaincfg/chainhash"
	"github.com/ltcsuite/ltcd/database"
	"github.com/ltcsuite/ltcd/txscript"
	"github.com/ltcsuite/ltcd/wire"
	"github.com/ltcsuite/ltcutil"
)

const (
	numRequestedLimit  = 1000
	maxProtocolVersion = 70002
)

type (
	// retrievedTx defines the retrieved transactions struct.
	retrievedTx struct {
		txBytes []byte
		blkHash *chainhash.Hash // Only set when transaction is in a block.
		tx      *ltcutil.Tx
	}
)

// Fetch balances implementation
func fetchBalance(addr string) (*v.Balance, error) {

	res := &v.Balance{
		Count: int(0),
		Value: float64(0),
	}

	err = db.View(func(dbTx database.Tx) error {

		addrList := make(map[string]ltcutil.Address)
		if _, ok := addrList[addr]; ok != true {
			addrList[addr], err = ltcutil.DecodeAddress(addr, chainConfig.ChainParams)
			if err != nil {
				return err
			}
		}

		for addrHash, addr := range addrList {

			listTx, err := fetchTransactions(addr, &[]string{addrHash}, true, true)
			if err != nil {
				return err
			}

			var incoming []btcjson.Vout
			var outgoing []btcjson.VinPrevOut

			for _, txs := range listTx {
				incoming = append(incoming, txs.Vout...)
				outgoing = append(outgoing, txs.Vin...)
			}

			res.Count = len(incoming)

			for _, in := range incoming {
				res.Value += in.Value
			}

			for _, out := range outgoing {
				res.Value -= out.PrevOut.Value
			}
		}

		return nil
	})

	return res, err
}

// Search raw transactions implementation
func fetchTransactions(
	addr ltcutil.Address,
	filterAddrs *[]string,
	vinExtra bool,
	reverse bool) ([]btcjson.SearchRawTransactionsResult, error) {

	// Respond with an error if the address index is not enabled.
	if addrIndex == nil {
		return nil, errors.New("Address index must be enabled (--addrindex)")
	}

	// Including the extra previous output information requires the transaction index.
	// Currently the address index relies on the transaction index, so this check is redundant,
	// but it's better to be safe in case the address index is ever changed to not rely on it.
	if vinExtra && txIndex == nil {
		return nil, errors.New("Transaction index must be enabled (--txindex)")
	}

	numToSkip := 0
	numRequested := numRequestedLimit

	// Add transactions from mempool first if client asked for reverse order.
	// Otherwise, they will be added last (as needed depending on the requested counts).
	//
	// NOTE: This code doesn't sort by dependency.
	// This might be something to do in the future for the client's convenience, or leave it to the client.
	numSkipped := uint32(0)
	addrTxns := make([]retrievedTx, 0, numRequested)
	if reverse {
		// Transactions in the mempool are not in a block header yet, so the block header field in the retieved
		// transaction struct is left nil.
		mpTxns, mpSkipped := fetchMempoolTxnsForAddr(addr, uint32(numToSkip), uint32(numRequested))
		numSkipped += mpSkipped
		for _, tx := range mpTxns {
			addrTxns = append(addrTxns, retrievedTx{tx: tx})
		}
	}

	var err error
	// Fetch transactions from the database in the desired order if more are needed.
	if len(addrTxns) < numRequested {

		err = db.View(func(dbTx database.Tx) error {
			regions, dbSkipped, err := addrIndex.TxRegionsForAddress(dbTx,
				addr,
				uint32(numToSkip)-numSkipped,
				uint32(numRequested-len(addrTxns)),
				reverse)
			if err != nil {
				return err
			}

			// Load the raw transaction bytes from the database.
			serializedTxns, err := dbTx.FetchBlockRegions(regions)
			if err != nil {
				return err
			}

			// Add the transaction and the hash of the block it is contained in to the list.
			// Note that the transaction is left serialized here since the caller might have requested non-verbose
			// output and hence there would be no point in deserializing it just to reserialize it later.
			for i, serializedTx := range serializedTxns {
				addrTxns = append(addrTxns, retrievedTx{
					txBytes: serializedTx,
					blkHash: regions[i].Hash,
				})
			}
			numSkipped += dbSkipped

			return nil
		})
		if err != nil {
			return nil, errors.New("Failed to load address index entries")
		}

	}

	// Add transactions from mempool last if client did not Request reverse order and the number of results
	// is still under the number requested.
	if !reverse && len(addrTxns) < numRequested {

		// Transactions in the mempool are not in a block header yet, so the block header field in
		// the retieved transaction struct is left nil.
		mpTxns, mpSkipped := fetchMempoolTxnsForAddr(addr,
			uint32(numToSkip)-numSkipped,
			uint32(numRequested-len(addrTxns)))

		numSkipped += mpSkipped
		for _, tx := range mpTxns {
			addrTxns = append(addrTxns, retrievedTx{tx: tx})
		}
	}

	// Address has never been used if neither source yielded any results.
	if len(addrTxns) == 0 {
		return nil, errors.New("No information available about address")
	}

	// Serialize all of the transactions to hex.
	hexTxns := make([]string, len(addrTxns))
	for i := range addrTxns {
		// Simply encode the raw bytes to hex when the retrieved
		// transaction is already in serialized form.
		rtx := &addrTxns[i]
		if rtx.txBytes != nil {
			hexTxns[i] = hex.EncodeToString(rtx.txBytes)
			continue
		}

		// Serialize the transaction first and convert to hex when the
		// retrieved transaction is the deserialized structure.
		hexTxns[i], err = messageToHex(rtx.tx.MsgTx())
		if err != nil {
			return nil, err
		}
	}

	// Normalize the provided filter addresses (if any) to ensure there are no duplicates.
	filterAddrMap := make(map[string]struct{})
	if filterAddrs != nil && len(*filterAddrs) > 0 {
		for _, addr := range *filterAddrs {
			filterAddrMap[addr] = struct{}{}
		}
	}

	// The verbose flag is set, so generate the JSON object and return it.
	best := net.BestSnapshot()
	txList := make([]btcjson.SearchRawTransactionsResult, len(addrTxns))
	for i := range addrTxns {

		var mtx *wire.MsgTx

		// The deserialized transaction is needed, so deserialize the retrieved transaction
		// if it's in serialized form (which will be the case when it was lookup up from the database).
		// Otherwise, use the existing deserialized transaction.
		rtx := &addrTxns[i]
		if rtx.tx == nil {
			// Deserialize the transaction.
			mtx = new(wire.MsgTx)
			err := mtx.Deserialize(bytes.NewReader(rtx.txBytes))
			if err != nil {
				return nil, errors.New("Failed to deserialize transaction")
			}
		} else {
			mtx = rtx.tx.MsgTx()
		}

		result := &txList[i]
		result.Hex = hexTxns[i]
		result.Txid = mtx.TxHash().String()
		result.Vin, err = createVinListPrevOut(mtx, chainConfig.ChainParams, vinExtra, filterAddrMap)
		if err != nil {
			return nil, err
		}

		result.Vout = createVoutList(mtx, chainConfig.ChainParams, filterAddrMap)
		result.Version = mtx.Version
		result.LockTime = mtx.LockTime

		// Transactions grabbed from the mempool aren't yet in a block, so conditionally fetch block details here.
		// This will be reflected in the final JSON output (mempool won't have confirmations or block information).
		var (
			blkHeader  *wire.BlockHeader
			blkHashStr string
			blkHeight  int32
		)

		if blkHash := rtx.blkHash; blkHash != nil {
			// Fetch the header from net.
			header, err := net.HeaderByHash(blkHash)
			if err != nil {
				return nil, errors.New("Block not found")
			}

			// Get the block height from net.
			height, err := net.BlockHeightByHash(blkHash)
			if err != nil {
				return nil, errors.New("Failed to obtain block height")
			}

			blkHeader = &header
			blkHashStr = blkHash.String()
			blkHeight = height
		}

		// Add the block information to the result if there is any.
		if blkHeader != nil {
			// This is not a typo, they are identical in Bitcoin Core as well.
			result.Time = blkHeader.Timestamp.Unix()
			result.Blocktime = blkHeader.Timestamp.Unix()
			result.BlockHash = blkHashStr
			result.Confirmations = uint64(1 + best.Height - blkHeight)
		}
	}

	return txList, nil
}

// fetchMempoolTxnsForAddr queries the address index for all unconfirmed transactions that involve the provided address.
// The results will be limited by the number to skip and the number requested.
func fetchMempoolTxnsForAddr(addr ltcutil.Address,
	numToSkip,
	numRequested uint32) ([]*ltcutil.Tx, uint32) {
	// There are no entries to return when there are less available than the number being skipped.
	mpTxns := addrIndex.UnconfirmedTxnsForAddress(addr)
	numAvailable := uint32(len(mpTxns))
	if numToSkip > numAvailable {
		return nil, numAvailable
	}

	// Filter the available entries based on the number to skip and number
	// requested.
	rangeEnd := numToSkip + numRequested
	if rangeEnd > numAvailable {
		rangeEnd = numAvailable
	}

	return mpTxns[numToSkip:rangeEnd], numToSkip
}

// fetchInputTxos fetches the outpoints from all transactions referenced by the inputs to the passed transaction
// by checking the transaction mempool first then the transaction index for those already mined into blocks.
func fetchInputTxos(tx *wire.MsgTx) (map[wire.OutPoint]wire.TxOut, error) {

	mp := txMemPool
	originOutputs := make(map[wire.OutPoint]wire.TxOut)

	for txInIndex, txIn := range tx.TxIn {
		// Attempt to fetch and use the referenced transaction from the memory pool.
		origin := &txIn.PreviousOutPoint
		originTx, err := mp.FetchTransaction(&origin.Hash)
		if err == nil {
			txOuts := originTx.MsgTx().TxOut
			if origin.Index >= uint32(len(txOuts)) {
				msg := "Unable to find output %v referenced from transaction %s:%d"
				return nil, errors.New(fmt.Sprintf(msg, origin, tx.TxHash(), txInIndex))
			}

			originOutputs[*origin] = *txOuts[origin.Index]
			continue
		}

		// Look up the location of the transaction.
		blockRegion, err := txIndex.TxBlockRegion(&origin.Hash)
		if err != nil {
			return nil, errors.New("Failed to retrieve transaction location")
		}
		if blockRegion == nil {
			return nil, errors.New("Failed to retrieve transaction info")
		}

		// Load the raw transaction bytes from the database.
		var txBytes []byte
		err = db.View(func(dbTx database.Tx) error {
			var err error
			txBytes, err = dbTx.FetchBlockRegion(blockRegion)
			return err
		})
		if err != nil {
			return nil, errors.New("Failed to load transaction info")
		}

		// Deserialize the transaction
		var msgTx wire.MsgTx
		err = msgTx.Deserialize(bytes.NewReader(txBytes))
		if err != nil {
			return nil, errors.New("Failed to deserialize transaction")
		}

		// Add the referenced output to the map.
		if origin.Index >= uint32(len(msgTx.TxOut)) {
			msg := "Unable to find output %v referenced from transaction %s:%d"
			return nil, errors.New(fmt.Sprintf(msg, origin, tx.TxHash(), txInIndex))
		}

		originOutputs[*origin] = *msgTx.TxOut[origin.Index]
	}

	return originOutputs, nil
}

// messageToHex serializes a message to the wire protocol encoding using the latest protocol version
// returns a hex-encoded string of the result.
func messageToHex(msg wire.Message) (string, error) {

	var buf bytes.Buffer

	if err := msg.BtcEncode(&buf, maxProtocolVersion, wire.WitnessEncoding); err != nil {
		return "", errors.New(fmt.Sprintf("Failed to encode msg of type %T", msg))
	}

	return hex.EncodeToString(buf.Bytes()), nil
}

// createVinListPrevOut returns a slice of JSON objects for the inputs of the passed transaction.
func createVinListPrevOut(mtx *wire.MsgTx,
	chainParams *chaincfg.Params,
	vinExtra bool,
	filterAddrMap map[string]struct{}) ([]btcjson.VinPrevOut, error) {

	// Coinbase transactions only have a single txin by definition.
	if blockchain.IsCoinBaseTx(mtx) {
		// Only include the transaction if the filter map is empty because a coinbase input has no addresses
		// and so would never match a non-empty filter.
		if len(filterAddrMap) != 0 {
			return nil, nil
		}

		txIn := mtx.TxIn[0]
		vinList := make([]btcjson.VinPrevOut, 1)
		vinList[0].Coinbase = hex.EncodeToString(txIn.SignatureScript)
		vinList[0].Sequence = txIn.Sequence
		return vinList, nil
	}

	// Use a dynamically sized list to accommodate the address filter.
	vinList := make([]btcjson.VinPrevOut, 0, len(mtx.TxIn))

	// Lookup all of the referenced transaction outputs needed to populate
	// the previous output information if requested.
	var originOutputs map[wire.OutPoint]wire.TxOut
	if vinExtra || len(filterAddrMap) > 0 {
		var err error
		originOutputs, err = fetchInputTxos(mtx)
		if err != nil {
			return nil, err
		}
	}

	for _, txIn := range mtx.TxIn {
		// The disassembled string will contain [error] inline
		// if the script doesn't fully parse, so ignore the
		// error here.
		disbuf, _ := txscript.DisasmString(txIn.SignatureScript)

		// Create the basic input entry without the additional optional
		// previous output details which will be added later if
		// requested and available.
		prevOut := &txIn.PreviousOutPoint
		vinEntry := btcjson.VinPrevOut{
			Txid:     prevOut.Hash.String(),
			Vout:     prevOut.Index,
			Sequence: txIn.Sequence,
			ScriptSig: &btcjson.ScriptSig{
				Asm: disbuf,
				Hex: hex.EncodeToString(txIn.SignatureScript),
			},
		}

		if len(txIn.Witness) != 0 {
			vinEntry.Witness = witnessToHex(txIn.Witness)
		}

		// Add the entry to the list now if it already passed the filter
		// since the previous output might not be available.
		passesFilter := len(filterAddrMap) == 0
		if passesFilter {
			vinList = append(vinList, vinEntry)
		}

		// Only populate previous output information if requested and
		// available.
		if len(originOutputs) == 0 {
			continue
		}
		originTxOut, ok := originOutputs[*prevOut]
		if !ok {
			continue
		}

		// Ignore the error here since an error means the script
		// couldn't parse and there is no additional information about
		// it anyways.
		_, addrs, _, _ := txscript.ExtractPkScriptAddrs(
			originTxOut.PkScript, chainParams)

		// Encode the addresses while checking if the address passes the
		// filter when needed.
		encodedAddrs := make([]string, len(addrs))
		for j, addr := range addrs {
			encodedAddr := addr.EncodeAddress()
			encodedAddrs[j] = encodedAddr

			// No need to check the map again if the filter already
			// passes.
			if passesFilter {
				continue
			}
			if _, exists := filterAddrMap[encodedAddr]; exists {
				passesFilter = true
			}
		}

		// Ignore the entry if it doesn't pass the filter.
		if !passesFilter {
			continue
		}

		// Add entry to the list if it wasn't already done above.
		if len(filterAddrMap) != 0 {
			vinList = append(vinList, vinEntry)
		}

		// Update the entry with previous output information if
		// requested.
		if vinExtra {
			vinListEntry := &vinList[len(vinList)-1]
			vinListEntry.PrevOut = &btcjson.PrevOut{
				Addresses: encodedAddrs,
				Value:     ltcutil.Amount(originTxOut.Value).ToBTC(),
			}
		}
	}

	return vinList, nil
}

// createVoutList returns a slice of JSON objects for the outputs of the passed transaction.
func createVoutList(mtx *wire.MsgTx,
	chainParams *chaincfg.Params,
	filterAddrMap map[string]struct{}) []btcjson.Vout {

	voutList := make([]btcjson.Vout, 0, len(mtx.TxOut))

	for i, v := range mtx.TxOut {
		// The disassembled string will contain [error] inline if the script doesn't fully parse,
		// so ignore the error here.
		disbuf, _ := txscript.DisasmString(v.PkScript)

		// Ignore the error here since an error means the script couldn't parse and there is no additional information
		// about it anyways.
		scriptClass, addrs, reqSigs, _ := txscript.ExtractPkScriptAddrs(v.PkScript, chainParams)

		// Encode the addresses while checking if the address passes the filter when needed.
		passesFilter := len(filterAddrMap) == 0
		encodedAddrs := make([]string, len(addrs))

		for j, addr := range addrs {

			encodedAddr := addr.EncodeAddress()
			encodedAddrs[j] = encodedAddr

			// No need to check the map again if the filter already passes.
			if passesFilter {
				continue
			}

			if _, exists := filterAddrMap[encodedAddr]; exists {
				passesFilter = true
			}
		}

		if !passesFilter {
			continue
		}

		var vout btcjson.Vout

		vout.N = uint32(i)
		vout.Value = ltcutil.Amount(v.Value).ToBTC()
		vout.ScriptPubKey.Addresses = encodedAddrs
		vout.ScriptPubKey.Asm = disbuf
		vout.ScriptPubKey.Hex = hex.EncodeToString(v.PkScript)
		vout.ScriptPubKey.Type = scriptClass.String()
		vout.ScriptPubKey.ReqSigs = int32(reqSigs)

		voutList = append(voutList, vout)
	}

	return voutList
}

// witnessToHex formats the passed witness stack as a slice of hex-encoded strings to be used in a JSON Response.
func witnessToHex(witness wire.TxWitness) []string {

	// Ensure nil is returned when there are no entries versus an empty slice
	// so it can properly be omitted as necessary.
	if len(witness) == 0 {
		return nil
	}

	result := make([]string, 0, len(witness))
	for _, wit := range witness {
		result = append(result, hex.EncodeToString(wit))
	}

	return result
}
