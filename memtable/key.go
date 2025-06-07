package memtable

import "bytes"

type ByteKey []byte

type Key interface {
	Less(other Key) bool
	Equal(other Key) bool
}

// ByteKey implements the Keyer interface.
var _ Key = ByteKey{}

func (k ByteKey) Less(other Key) bool {
	otherBytes, ok := other.(ByteKey)
	if !ok {
		return false // or handle the error as needed
	}
	return bytes.Compare(k, otherBytes) < 0

}

func (k ByteKey) Equal(other Key) bool {
	otherBytes, ok := other.(ByteKey)
	if !ok {
		return false // or handle the error as needed
	}
	return bytes.Equal(k, otherBytes)
}
