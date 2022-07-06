package vsys

import (
	"encoding/binary"
	"fmt"
)

func PackBool(val bool) []byte {
	if val {
		return []byte{1}
	}
	return []byte{0}
}

func PackUInt8(val uint8) []byte {
	return []byte{val}
}

func PackUInt16(val uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, val)
	return b
}

func PackUInt32(val uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, val)
	return b
}

func PackUInt64(val uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, val)
	return b
}

func UnpackBool(b []byte) (bool, error) {
	if len(b) != 1 {
		return false, fmt.Errorf("UnpackBool: the byte slice must be 1 byte long")
	}

	if b[0] > 1 {
		return false, fmt.Errorf("UnpackBool: the byte slice must contain either 0 or 1")
	}

	if b[0] == 1 {
		return true, nil
	}
	return false, nil
}

func UnpackUint8(b []byte) (uint8, error) {
	if len(b) != 1 {
		return 0, fmt.Errorf("UnpackUint8 the byte slice must be 1 byte long")
	}
	return uint8(b[0]), nil
}

func UnpackUInt16(b []byte) (uint16, error) {
	if len(b) != 2 {
		return 0, fmt.Errorf("UnpackUInt16: the byte slice must be 2 bytes long")
	}
	return binary.BigEndian.Uint16(b), nil
}

func UnpackUInt32(b []byte) (uint32, error) {
	if len(b) != 4 {
		return 0, fmt.Errorf("UnpackUInt32: the byte slice must be 4 bytes long")
	}
	return binary.BigEndian.Uint32(b), nil
}

func UnpackUInt64(b []byte) (uint64, error) {
	if len(b) != 8 {
		return 0, fmt.Errorf("UnpackUInt64: the byte slice must be 4 bytes long")
	}
	return binary.BigEndian.Uint64(b), nil
}
