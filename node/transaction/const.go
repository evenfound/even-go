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
)

const (
	unspecifiedFileFormat FileFormat = iota
	jsonFile
	zlibFile
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
	TagNewReg = "NEW_REG"
)

var (
	metaID        = id{metaTx, version}
	outputID      = id{outputTx, version}
	inputSpendID  = id{inputSpendTx, version}
	inputChangeID = id{inputChangeTx, version}

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
