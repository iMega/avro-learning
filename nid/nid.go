package nid

import (
	"bytes"
	"database/sql/driver"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
)

type NID [16]byte

func NewNID() NID {
	var n NID
	uuidv4 := uuid.New()
	j := bytes.Join([][]byte{uuidv4[6:8], uuidv4[4:6], uuidv4[0:4], uuidv4[8:10], uuidv4[10:]}, []byte{})
	copy(n[:], j)

	return NID(n)
}

// String encode NID to string
func (n NID) String() string {
	return hex.EncodeToString(n[:])
}

// DecodeString decode string to NID
func DecodeString(str string) (NID, error) {
	var n NID

	b, err := hex.DecodeString(str)
	if err != nil {
		return NID([16]byte{}), err
	}
	copy(n[:], b)

	return NID(n), nil
}

// Value encode unique identification for store to db
func (n NID) Value() (driver.Value, error) {
	return n[:], nil
}

// Scan decode unique identification to get from db
func (n *NID) Scan(src interface{}) error {
	var nid NID
	s, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("error scan uid %v", src)
	}
	copy(nid[:], s)
	*n = NID(nid)
	return nil
}
