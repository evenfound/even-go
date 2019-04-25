package transaction

import (
	"bytes"
	"errors"

	"github.com/evenfound/even-go/core"
)

// Builder builds new transaction.
type Builder struct {
	t *transaction
}

// NewBuilderMeta constructs builder for meta transaction.
func NewBuilderMeta(tag string) *Builder {
	return &Builder{
		t: newTransaction(metaID, 0, tag),
	}
}

// NewBuilderOutput constructs builder for output transaction.
func NewBuilderOutput(tag string, value value) *Builder {
	return &Builder{
		t: newTransaction(outputID, value, tag),
	}
}

// NewBuilderInputSpend constructs builder for input spend transaction.
func NewBuilderInputSpend(tag string, value value) *Builder {
	return &Builder{
		t: newTransaction(inputSpendID, value, tag),
	}
}

// NewBuilderInputChange constructs builder for input change transaction.
func NewBuilderInputChange(tag string, value value) *Builder {
	return &Builder{
		t: newTransaction(inputChangeID, value, tag),
	}
}

// SetAddress sets address of an account.
func (b *Builder) SetAddress(a string) *Builder {
	b.t.Address = address(a)
	return b
}

// SetMessage sets message — signature of trunk.
func (b *Builder) SetMessage(m message) *Builder {
	b.t.Message = m
	return b
}

// SetSource sets source — address of some resource.
func (b *Builder) SetSource(s Hash) *Builder {
	b.t.Source = s
	return b
}

// SetTrunk sets trunk — address of the last verified transaction.
func (b *Builder) SetTrunk(trunk Hash) *Builder {
	b.t.Trunk = trunk
	return b
}

// AddTwig adds another candidate transaction.
func (b *Builder) AddTwig(h Hash) *Builder {
	b.t.Branch = append(b.t.Branch, newTwig(h))
	return b
}

// SaveLocal writes completed transaction into a file and returns it's filename.
func (b Builder) SaveLocal(format FileFormat) (string, error) {
	if format == unspecifiedFileFormat {
		format = jsonFile
	}

	ser, err := b.t.serialize(format)
	if err != nil {
		return "", err
	}

	// Prepend format header
	ser = append([]byte(format.String()), ser...)

	filename := format.String() + ".tr"
	return saveLocal(filename, ser)
}

// Save writes completed transaction into IPFS and returns it's hash.
func (b Builder) Save() (string, error) {
	format := zjsonFile
	stream, err := b.t.serialize(format)
	if err != nil {
		return "", err
	}

	// Prepend format header
	stream = append([]byte(format.String()), stream...)

	path := b.t.generateFileName()
	if err := core.CreateNewFile(path, stream); err != nil {
		return "", err
	}

	stat, err := core.FileStat(path)
	if err != nil {
		return "", err
	}

	return stat.Cid, nil
}

// Load reads transaction from a local file.
func Load(filename string) (Interface, error) {
	stream, err := load(filename)
	if err != nil {
		return nil, err
	}

	header := []byte(jsonFile.String())
	if bytes.HasPrefix(stream, header) {
		stream = stream[len(header):]
		return fromJSON(stream)
	}

	header = []byte(ubjsonFile.String())
	if bytes.HasPrefix(stream, header) {
		stream = stream[len(header):]
		return fromUBJSON(stream)
	}

	header = []byte(gobFile.String())
	if bytes.HasPrefix(stream, header) {
		stream = stream[len(header):]
		return fromGOB(stream)
	}

	return nil, errors.New("unknown file format")
}
