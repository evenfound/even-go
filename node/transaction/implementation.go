package transaction

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/cesanta/ubjson"
)

type transaction struct {
	ID        id        `json:"id"`
	Address   address   `json:"address"`
	Message   message   `json:"message"`
	Source    Hash      `json:"source"`
	Value     value     `json:"value"`
	Timestamp timestamp `json:"timestamp"`
	Trunk     Hash      `json:"trunk"`
	Branch    twigs     `json:"branch"`
	Tag       string    `json:"tag"`
}

func newTransaction(id id, value value, tag string) *transaction {
	return &transaction{
		ID:        id,
		Value:     value,
		Timestamp: timestamp(time.Now()),
		Branch:    make([]twig, 0),
		Tag:       tag,
	}
}

// String satisfies interface Stringer.
func (t *transaction) String() string {
	return t.Tag
}

func (t *transaction) serialize(format FileFormat) ([]byte, error) {
	switch format {
	case jsonFile:
		return toJSON(t)
	case zlibFile:
		return toZlibJSON(t)
	case ubjsonFile:
		return toUBJSON(t)
	case gobFile:
		return toGOB(t)
	}
	msg := fmt.Sprintf("'%d' unknown file format (expected %d | %d | %d | %d)",
		format, jsonFile, zlibFile, ubjsonFile, gobFile)
	return nil, errors.New(msg)
}

func toJSON(t *transaction) ([]byte, error) {
	return json.Marshal(t)
}

func toZlibJSON(t *transaction) ([]byte, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return compress(data)
}

func toUBJSON(t *transaction) ([]byte, error) {
	return ubjson.Marshal(t)
}

func toGOB(t *transaction) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(t)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func fromJSON(stream []byte) (*transaction, error) {
	tran := new(transaction)
	if err := json.Unmarshal(stream, tran); err != nil {
		return nil, err
	}
	return tran, nil
}

func fromUBJSON(stream []byte) (*transaction, error) {
	tran := new(transaction)
	if err := ubjson.Unmarshal(stream, tran); err != nil {
		return nil, err
	}
	return tran, nil
}

func fromGOB(stream []byte) (*transaction, error) {
	tran := new(transaction)
	dec := gob.NewDecoder(bytes.NewBuffer(stream))
	if err := dec.Decode(tran); err != nil {
		return nil, err
	}
	return tran, nil
}
