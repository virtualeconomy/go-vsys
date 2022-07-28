package vsys

import (
	"fmt"
)

var idxMap = map[int]func([]byte) (DataEntry, error){
	1:  NewDePubKeyFromBytesGeneric,
	2:  NewDeAddrFromBytesGeneric,
	5:  NewDeStrFromBytesGeneric,
	11: NewDeBytesFromBytesGeneric,
}

type DataStack []DataEntry

func NewDataStackFromBytes(b []byte) (DataStack, error) {
	entriesCnt, err := UnpackUInt16(b[:2])
	if err != nil {
		return nil, fmt.Errorf("NewDataStackFromBytes: %w", err)
	}

	b = b[2:]

	entries := make([]DataEntry, 0, entriesCnt)
	for i := 0; i < int(entriesCnt); i++ {
		idx, err := UnpackUint8(b[:1])
		if err != nil {
			return nil, fmt.Errorf("NewDataStackFromBytes: %w", err)
		}

		deserializer, ok := idxMap[int(idx)]
		if !ok {
			return nil, fmt.Errorf("NewDataStackFromBytes: DataEntry type not identified. Byte %v at position %d", idx, i)
		}
		de, err := deserializer(b)
		if err != nil {
			return nil, fmt.Errorf("NewDataStackFromBytes: %w", err)
		}

		entries = append(entries, de)
		b = b[de.Size():]
	}
	return DataStack(entries), nil
}

func (d DataStack) Serialize() Bytes {
	lenBytes := PackUInt16(uint16(len(d)))
	size := len(lenBytes)
	for _, de := range d {
		size += de.Size()
	}

	b := make([]byte, 0, size)
	b = append(b, lenBytes...)

	for _, de := range d {
		b = append(b, de.Serialize()...)
	}
	return b
}

func (d DataStack) String() string {
	return fmt.Sprintf("%T(%s)", d, []DataEntry(d))
}
