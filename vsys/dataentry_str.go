package vsys

import "fmt"

type DeStr struct {
	Idx DeIdx

	Data Str
}

func NewDeStr(s Str) *DeStr {
	return &DeStr{
		Idx:  5,
		Data: s,
	}
}

func NewDeStrFromString(s string) *DeStr {
	return NewDeStr(Str(s))
}

func NewDeStrFromBytesGeneric(b []byte) (DataEntry, error) {
	l, err := UnpackUInt16(b[1 : 1+2])
	if err != nil {
		return nil, fmt.Errorf("NewDeStrFromBytesGeneric: %w", err)
	}
	s := b[3 : 3+l]
	if err != nil {
		return nil, fmt.Errorf("NewDeStrFromBytesGeneric: %w", err)
	}
	return NewDeStrFromString(string(s)), nil
}

func NewDeStrFromBytes(b []byte) *DeStr {
	s, _ := NewDeStrFromBytesGeneric(b)
	return s.(*DeStr)
}

func (s *DeStr) IdxBytes() Bytes {
	return s.Idx.Serialize()
}

func (s *DeStr) DataBytes() Bytes {
	return s.Data.Bytes()
}

func (s *DeStr) LenBytes() Bytes {
	return PackUInt16(uint16(len(s.DataBytes())))
}

func (s *DeStr) Serialize() Bytes {
	size := len(s.IdxBytes()) +
		len(s.LenBytes()) +
		len(s.DataBytes())

	b := make([]byte, 0, size)
	b = append(b, s.IdxBytes()...)
	b = append(b, s.LenBytes()...)
	b = append(b, s.DataBytes()...)

	return b
}

func (s *DeStr) Size() int {
	return 1 + len(s.DataBytes())
}

func (s *DeStr) String() string {
	return fmt.Sprintf("%T(%+v)", s, s.Data)
}
