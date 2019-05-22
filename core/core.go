package core

import (
	"errors"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/evenfound/even-go/ipfs"
	"github.com/evenfound/even-go/net"
	rep "github.com/evenfound/even-go/net/repointer"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-ipfs/namesys"
	"github.com/ipfs/go-ipfs/repo"
	libp2p "github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-peer"

	"path"
	"sync"
	"time"

	//"github.com/OpenBazaar/multiwallet"
	//"github.com/btcsuite/btcutil/hdkeychain"
	//"github.com/evenfound/even-go/ipfs"
	//"github.com/evenfound/even-go/namesys"
	//ret "github.com/evenfound/even-go/net/retriever"
	//"github.com/evenfound/even-go/repo"
	//sto "github.com/evenfound/even-go/storage"
	"github.com/ipfs/go-ipfs/core"
	"github.com/op/go-logging"
	"golang.org/x/net/proxy"
)

const (
	// VERSION - current version
	VERSION = "0.13.0"
	// USERAGENT - user-agent header string
	USERAGENT = "/openbazaar-go:" + VERSION + "/"
)

var log = logging.MustGetLogger("core")

// Node - ob node
var Node *EvenNode

var inflightPublishRequests int

// EvenNode - represent ob node which encapsulates ipfsnode, wallet etc
type EvenNode struct {
	// IPFS node object
	IpfsNode *core.IpfsNode

	/* The roothash of the node directory inside the openbazaar repo.
	   This directory hash is published on IPNS at our peer ID making
	   the directory publicly viewable on the network. */
	RootHash string

	// The path to the openbazaar repo in the file system
	RepoPath string

	// The EvenNetwork network service for direct communication between peers
	Service net.NetworkService

	// Database for storing node specific data
	Datastore repo.Datastore

	// Websocket channel used for pushing data to the UI

	// TODO: commented mode refactoring
	//Broadcast chan repo.Notifier

	// A map of cryptocurrency wallets
	// TODO: commented mode refactoring
	//Multiwallet multiwallet.MultiWallet

	// Storage for our outgoing messages
	// TODO: commented mode refactoring
	//MessageStorage sto.OfflineMessagingStorage

	// A service that periodically checks the dht for outstanding messages
	// TODO: commented mode refactoring
	//MessageRetriever *ret.MessageRetriever

	// OfflineMessageFailoverTimeout is the duration until the protocol
	// will stop looking for the peer to send a direct message and failover to
	// sending an offline message
	OfflineMessageFailoverTimeout time.Duration

	// A service that periodically republishes active pointers
	PointerRepublisher *rep.PointerRepublisher

	// Used to resolve domains to EvenNetwork IDs
	NameSystem *namesys.NameSystem

	// A service that periodically fetches and caches the bitcoin exchange rates
	//ExchangeRates wallet.ExchangeRates

	// Optional nodes to push user data to
	PushNodes []peer.ID

	// The user-agent for this node
	UserAgent string

	// A dialer for Tor if available
	TorDialer proxy.Dialer

	// Manage blocked peers
	BanManager *net.BanManager

	// Allow other nodes to push data to this node for storage
	AcceptStoreRequests bool

	// Last ditch API to find records that dropped out of the DHT
	IPNSBackupAPI string

	// RecordAgingNotifier is a worker that walks the cases datastore to
	// notify the user as disputes age past certain thresholds
	// TODO: commented mode refactoring
	//RecordAgingNotifier *recordAgingNotifier

	// Generic pubsub interface
	Pubsub ipfs.Pubsub

	// The master private key derived from the mnemonic
	MasterPrivateKey *hdkeychain.ExtendedKey

	TestnetEnable        bool
	RegressionTestEnable bool
}

// PublishLock seedLock - Unpin the current node repo, re-add it, then publish to IPNS
var PublishLock sync.Mutex
var seedLock sync.Mutex

// InitalPublishComplete - indicate publish completion
var InitalPublishComplete bool // = false

// TestNetworkEnabled indicates whether the node is operating with test parameters
func (n *EvenNode) TestNetworkEnabled() bool { return n.TestnetEnable }

// RegressionNetworkEnabled indicates whether the node is operating with regression parameters
func (n *EvenNode) RegressionNetworkEnabled() bool { return n.RegressionTestEnable }

// SeedNode - publish to IPNS
func (n *EvenNode) SeedNode() error {
	seedLock.Lock()
	ipfs.UnPinDir(n.IpfsNode, n.RootHash)
	var aerr error
	var rootHash string
	// There's an IPFS bug on Windows that might be related to the Windows indexer that could cause this to fail
	// If we fail the first time, let's retry a couple times before giving up.
	for i := 0; i < 3; i++ {
		rootHash, aerr = ipfs.AddDirectory(n.IpfsNode, path.Join(n.RepoPath, "root"))
		if aerr == nil {
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
	if aerr != nil {
		seedLock.Unlock()
		return aerr
	}
	n.RootHash = rootHash
	seedLock.Unlock()
	InitalPublishComplete = true
	go n.publish(rootHash)
	return nil
}

func (n *EvenNode) publish(hash string) {
	// TODO: commented mode refactoring
	// Multiple publishes may have been queued
	// We only need to publish the most recent
	//PublishLock.Lock()
	//defer PublishLock.Unlock()
	//if hash != n.RootHash {
	//	return
	//}
	//
	//if inflightPublishRequests == 0 {
	//	n.Broadcast <- repo.StatusNotification{Status: "publishing"}
	//}
	//
	//err := n.sendToPushNodes(hash)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//
	//inflightPublishRequests++
	//err = ipfs.Publish(n.IpfsNode, hash)
	//
	//inflightPublishRequests--
	//if inflightPublishRequests == 0 {
	//	if err != nil {
	//		log.Error(err)
	//		n.Broadcast <- repo.StatusNotification{Status: "error publishing"}
	//	} else {
	//		n.Broadcast <- repo.StatusNotification{Status: "publish complete"}
	//	}
	//}
	return
}

func (n *EvenNode) sendToPushNodes(hash string) error {
	// TODO: commented mode refactoring
	//id, err := cid.Decode(hash)
	//if err != nil {
	//	return err
	//}
	//
	//var graph []cid.Cid
	//if len(n.PushNodes) > 0 {
	//	graph, err = ipfs.FetchGraph(n.IpfsNode, id)
	//	if err != nil {
	//		return err
	//	}
	//	pointers, err := n.Datastore.Pointers().GetByPurpose(ipfs.MESSAGE)
	//	if err != nil {
	//		return err
	//	}
	//	// Check if we're seeding any outgoing messages and add their CIDs to the graph
	//	for _, p := range pointers {
	//		if len(p.Value.Addrs) > 0 {
	//			s, err := p.Value.Addrs[0].ValueForProtocol(ma.P_IPFS)
	//			if err != nil {
	//				continue
	//			}
	//			c, err := cid.Decode(s)
	//			if err != nil {
	//				continue
	//			}
	//			graph = append(graph, *c)
	//		}
	//	}
	//}
	//for _, p := range n.PushNodes {
	//	go n.retryableSeedStoreToPeer(p, hash, graph)
	//}
	//
	//return nil
	return errors.New("sendToPushNodes not implemented")
}

func (n *EvenNode) retryableSeedStoreToPeer(pid peer.ID, graphHash string, graph []cid.Cid) {
	// TODO: commented mode refactoring
	//var retryTimeout = 2 * time.Second
	//for {
	//	if graphHash != n.RootHash {
	//		log.Errorf("root hash has changed, aborting push to %s", pid.Pretty())
	//		return
	//	}
	//	err := n.SendStore(pid.Pretty(), graph)
	//	if err != nil {
	//		if retryTimeout > 60*time.Second {
	//			log.Errorf("error pushing to peer %s: %s", pid.Pretty(), err.Error())
	//			return
	//		}
	//		log.Errorf("error pushing to peer %s...backing off: %s", pid.Pretty(), err.Error())
	//		time.Sleep(retryTimeout)
	//		retryTimeout *= 2
	//		continue
	//	}
	//	return
	//}
	return
}

// SetUpRepublisher - periodic publishing to IPNS
func (n *EvenNode) SetUpRepublisher(interval time.Duration) {
	// TODO: commented mode refactoring
	//if interval == 0 {
	//	return
	//}
	//ticker := time.NewTicker(interval)
	//go func() {
	//	for range ticker.C {
	//		n.UpdateFollow()
	//		n.SeedNode()
	//	}
	//}()
	return
}

/*EncryptMessage This is a placeholder until the libsignal is operational.
  For now we will just encrypt outgoing offline messages with the long lived identity key.
  Optionally you may provide a public key, to avoid doing an IPFS lookup */
func (n *EvenNode) EncryptMessage(peerID peer.ID, peerKey *libp2p.PubKey, message []byte) (ct []byte, rerr error) {
	// TODO: commented mode refactoring
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//if peerKey == nil {
	//	var pubKey libp2p.PubKey
	//	keyval, err := n.IpfsNode.Repo.Datastore().Get(dshelp.NewKeyFromBinary([]byte(KeyCachePrefix + peerID.Pretty())))
	//	if err != nil {
	//		pubKey, err = routing.GetPublicKey(n.IpfsNode.Routing, ctx, []byte(peerID))
	//		if err != nil {
	//			log.Errorf("Failed to find public key for %s", peerID.Pretty())
	//			return nil, err
	//		}
	//	} else {
	//		pubKey, err = libp2p.UnmarshalPublicKey(keyval.([]byte))
	//		if err != nil {
	//			log.Errorf("Failed to find public key for %s", peerID.Pretty())
	//			return nil, err
	//		}
	//	}
	//	peerKey = &pubKey
	//}
	//if peerID.MatchesPublicKey(*peerKey) {
	//	ciphertext, err := net.Encrypt(*peerKey, message)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return ciphertext, nil
	//}
	//log.Errorf("peer public key and id do not match for peer: %s", peerID.Pretty())
	//return nil, errors.New("peer public key and id do not match")
	return nil, errors.New("EncryptMessage not implemented")
}

// IPFSIdentityString - IPFS identifier
func (n *EvenNode) IPFSIdentityString() string {
	return n.IpfsNode.Identity.Pretty()
}

// createSlugFor Create a slug from a string
func createSlugFor(slugName string) string {
	// TODO: commented mode refactoring
	//l := SentenceMaxCharacters - SlugBuffer
	//if len(slugName) < SentenceMaxCharacters-SlugBuffer {
	//	l = len(slugName)
	//}
	//return url.QueryEscape(sanitize.Path(strings.ToLower(slugName[:l])))
	return ""
}
