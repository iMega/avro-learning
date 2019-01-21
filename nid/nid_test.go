package nid

import (
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNID(t *testing.T) {
	r := NewNID()
	r1 := NewNID()
	assert.NotEqual(t, r, r1)
}

func TestUID_Scan(t *testing.T) {
	r := []byte{0x43, 0xa1, 0xf3, 0x83, 0x8c, 0x37, 0x71, 0x84, 0xa3, 0xf5, 0x56, 0x2a, 0x99, 0x9c, 0x99, 0x7b}
	exp, _ := DecodeString("43a1f3838c377184a3f5562a999c997b")

	var n NID
	err := n.Scan(r)

	assert.Nil(t, err)
	assert.Equal(t, n, exp)
}

func TestUID_Value(t *testing.T) {
	u, _ := DecodeString("43a1f3838c377184a3f5562a999c997b")
	exp := []byte{0x43, 0xa1, 0xf3, 0x83, 0x8c, 0x37, 0x71, 0x84, 0xa3, 0xf5, 0x56, 0x2a, 0x99, 0x9c, 0x99, 0x7b}
	dr, err := u.Value()

	assert.Nil(t, err)
	assert.Equal(t, dr, driver.Value(exp))
}

func TestUID_ScanError(t *testing.T) {
	r := 123
	var u NID
	err := u.Scan(r)

	assert.Error(t, err)
	assert.Equal(t, NID([16]byte{}), u)
}

func TestUID_ValueError(t *testing.T) {
	_, err := DecodeString("123")
	assert.Error(t, err)
}

func BenchmarkNewNID(b *testing.B) {
	b.SetBytes(2)
	for i := 0; i < b.N; i++ {
		NewNID()
	}
}
