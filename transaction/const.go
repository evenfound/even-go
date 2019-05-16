package transaction

// FileFormat represents transaction file format.
type FileFormat int

// SaveDirection represents transaction save direction: IPFS or local.
type SaveDirection int

const (
	version byte = 0x01
)

const (
	unspecifiedTransaction byte = iota
	metaTx
	outputTx
	inputSpendTx
	inputChangeTx
	deployContractTx
	invokeContractTx
)

const (
	unspecifiedFileFormat FileFormat = iota
	jsonFile
	zjsonFile
	ubjsonFile
	gobFile
)

const (
	unspecifiedSaveDirection SaveDirection = iota
	localSave
	ipfsSave
)

const (
	// TagNewReg is tag for initial transaction.
	TagNewReg         = "NEW_REG"
	TagContractDeploy = "CONTRACT_DEPLOY"
	TagContractInvoke = "CONTRACT_INVOKE"
)

var (
	metaID           = id{metaTx, version}
	outputID         = id{outputTx, version}
	inputSpendID     = id{inputSpendTx, version}
	inputChangeID    = id{inputChangeTx, version}
	contractDeployID = id{deployContractTx, version}
	contractInvokeID = id{invokeContractTx, version}

	fileFormatStrings = [...]string{
		"UNKNOWN",
		"JSON",
		"ZLIBJSON",
		"UBJSON",
		"GOB",
	}
)

// String satisfies interface Stringer.
func (f FileFormat) String() string {
	return fileFormatStrings[f]
}
