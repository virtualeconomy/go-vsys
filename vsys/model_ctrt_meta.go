package vsys

import (
	"fmt"
)

type CtrtMetaLangCode string

func (c CtrtMetaLangCode) String() string {
	return fmt.Sprintf("%T(%s)", c, string(c))
}

type CtrtMetaLangVer uint32

func (c CtrtMetaLangVer) String() string {
	return fmt.Sprintf("%T(%d)", c, c)
}

type CtrtMetaBytes []byte

func NewCtrtMetaBytesFromBytes(b []byte) (CtrtMetaBytes, error) {
	l, err := UnpackUInt16(b[:2])
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaBytesFromBytes: %w", err)
	}
	return CtrtMetaBytes(b[2 : 2+l]), nil
}

func (b CtrtMetaBytes) LenBytes() Bytes {
	return PackUInt16(uint16(len(b)))
}

func (b CtrtMetaBytes) Serialize() Bytes {
	return append(b.LenBytes(), []byte(b)...)
}

func (b CtrtMetaBytes) ByteSlice() []byte {
	return []byte(b)
}

func (b CtrtMetaBytes) Bytes() Bytes {
	return Bytes(b)
}

func (b CtrtMetaBytes) String() string {
	return b.Bytes().B58Str().Str()
}

func (b CtrtMetaBytes) Size() int { return len(b) }

type CtrtMetaBytesList []CtrtMetaBytes

func NewCtrtMetaBytesListFromBytes(b []byte, withBytesLen bool) (CtrtMetaBytesList, error) {
	if withBytesLen {
		l, err := UnpackUInt16(b[:2])
		if err != nil {
			return nil, fmt.Errorf("NewCtrtMetaBytesListFromBytes: %w", err)
		}
		b = b[2 : 2+l]
	}

	itemsCnt, err := UnpackUInt16(b[:2])
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaBytesListFromBytes: %w", err)
	}
	b = b[2:]

	items := make([]CtrtMetaBytes, 0, itemsCnt)

	for i := 0; i < int(itemsCnt); i++ {
		l, err := UnpackUInt16(b[:2])
		if err != nil {
			return nil, fmt.Errorf("NewCtrtMetaBytesListFromBytes: %w", err)
		}

		item, err := NewCtrtMetaBytesFromBytes(b)

		if err != nil {
			return nil, fmt.Errorf("NewCtrtMetaBytesListFromBytes: %w", err)
		}

		items = append(items, item)
		b = b[2+l:]
	}

	return CtrtMetaBytesList(items), nil
}

func (c CtrtMetaBytesList) CtrtMetaBytesSlice() []CtrtMetaBytes {
	return []CtrtMetaBytes(c)
}

func (c CtrtMetaBytesList) Serialize(withBytesLen bool) Bytes {
	b := PackUInt16(uint16(len(c.CtrtMetaBytesSlice())))

	for _, bu := range c {
		b = append(b, bu.Serialize()...)
	}

	if withBytesLen {
		b = append(PackUInt16(uint16(len(b))), b...)
	}

	return b
}

func (c CtrtMetaBytesList) Size() int {
	size := len(c)
	for _, bu := range c {
		size += bu.Size()
	}
	return size
}

func (c CtrtMetaBytesList) String() string {
	s := make([]string, 0, len(c))
	for _, b := range c {
		s = append(s, b.String())
	}
	return fmt.Sprint(s)
}

type CtrtMetaTriggers struct {
	CtrtMetaBytesList
}

func NewCtrtMetaTriggersFromBytes(b []byte) (*CtrtMetaTriggers, error) {
	cbl, err := NewCtrtMetaBytesListFromBytes(b, true)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaTriggersFromBytes: %w", err)
	}
	return &CtrtMetaTriggers{cbl}, nil
}

func (c *CtrtMetaTriggers) Serialize() Bytes {
	return c.CtrtMetaBytesList.Serialize(true)
}

func (c *CtrtMetaTriggers) String() string {
	return fmt.Sprintf("%T(%s)", c, c.CtrtMetaBytesList.String())
}

type CtrtMetaDescriptors struct {
	CtrtMetaBytesList
}

func NewCtrtMetaDescriptorsFromBytes(b []byte) (*CtrtMetaDescriptors, error) {
	cbl, err := NewCtrtMetaBytesListFromBytes(b, true)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaDescriptorsFromBytes: %w", err)
	}
	return &CtrtMetaDescriptors{cbl}, nil
}

func (c *CtrtMetaDescriptors) Serialize() Bytes {
	return c.CtrtMetaBytesList.Serialize(true)
}

func (c *CtrtMetaDescriptors) String() string {
	return fmt.Sprintf("%T(%s)", c, c.CtrtMetaBytesList.String())
}

type CtrtMetaStateVars struct {
	CtrtMetaBytesList
}

func NewCtrtMetaStateVarsFromBytes(b []byte) (*CtrtMetaStateVars, error) {
	cbl, err := NewCtrtMetaBytesListFromBytes(b, true)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaStateVarsFromBytes: %w", err)
	}
	return &CtrtMetaStateVars{cbl}, nil
}

func (c *CtrtMetaStateVars) Serialize() Bytes {
	return c.CtrtMetaBytesList.Serialize(true)
}

func (c *CtrtMetaStateVars) String() string {
	return fmt.Sprintf("%T(%s)", c, c.CtrtMetaBytesList.String())
}

type CtrtMetaStateMap struct {
	CtrtMetaBytesList
}

func NewEmptyCtrtMetaStateMap() *CtrtMetaStateMap {
	return &CtrtMetaStateMap{
		CtrtMetaBytesList(make([]CtrtMetaBytes, 0)),
	}
}

func NewCtrtMetaStateMapFromBytes(b []byte) (*CtrtMetaStateMap, error) {
	cbl, err := NewCtrtMetaBytesListFromBytes(b, true)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaStateMapFromBytes: %w", err)
	}
	return &CtrtMetaStateMap{cbl}, nil
}

func (c *CtrtMetaStateMap) Serialize() Bytes {
	return c.CtrtMetaBytesList.Serialize(true)
}

func (c *CtrtMetaStateMap) String() string {
	return fmt.Sprintf("%T(%s)", c, c.CtrtMetaBytesList.String())
}

type CtrtMetaTextual struct {
	CtrtMetaBytesList
}

func NewCtrtMetaTextualFromBytes(b []byte) (*CtrtMetaTextual, error) {
	cbl, err := NewCtrtMetaBytesListFromBytes(b, false)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaTextualFromBytes: %w", err)
	}
	return &CtrtMetaTextual{cbl}, nil
}

func (c *CtrtMetaTextual) Serialize() Bytes {
	return c.CtrtMetaBytesList.Serialize(false)
}

func (c *CtrtMetaTextual) String() string {
	return fmt.Sprintf("%T(%s)", c, c.CtrtMetaBytesList.String())
}

const (
	LANG_CODE_BYTES_LEN = 4
	LANG_VER_BYTES_LEN  = 4
)

type CtrtMeta struct {
	LangCode    CtrtMetaLangCode
	LangVer     CtrtMetaLangVer
	Triggers    *CtrtMetaTriggers
	Descriptors *CtrtMetaDescriptors
	StateVars   *CtrtMetaStateVars
	StateMap    *CtrtMetaStateMap
	Textual     *CtrtMetaTextual
}

func NewCtrtMeta(b []byte) (*CtrtMeta, error) {
	langCode := CtrtMetaLangCode(b[:LANG_CODE_BYTES_LEN])
	b = b[LANG_CODE_BYTES_LEN:]

	langVer, err := UnpackUInt32(b[:LANG_VER_BYTES_LEN])
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMeta: %w", err)
	}
	b = b[LANG_VER_BYTES_LEN:]

	triggersLen, err := UnpackUInt16(b[:2])
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMeta: %w", err)
	}
	triggers, err := NewCtrtMetaTriggersFromBytes(b)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMeta: %w", err)
	}
	b = b[2+triggersLen:]

	descriptorsLen, err := UnpackUInt16(b[:2])
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMeta: %w", err)
	}
	descriptors, err := NewCtrtMetaDescriptorsFromBytes(b)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMeta: %w", err)
	}
	b = b[2+descriptorsLen:]

	stateVarsLen, err := UnpackUInt16(b[:2])
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMeta: %w", err)
	}
	stateVars, err := NewCtrtMetaStateVarsFromBytes(b)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMeta: %w", err)
	}
	b = b[2+stateVarsLen:]

	var stateMap *CtrtMetaStateMap

	if langVer == 1 {
		stateMap = NewEmptyCtrtMetaStateMap()
	} else {
		stateMapLen, err := UnpackUInt16(b[:2])
		if err != nil {
			return nil, fmt.Errorf("NewCtrtMeta: %w", err)
		}
		stateMap, err = NewCtrtMetaStateMapFromBytes(b)
		if err != nil {
			return nil, fmt.Errorf("NewCtrtMeta: %w", err)
		}
		b = b[2+stateMapLen:]
	}

	textual, err := NewCtrtMetaTextualFromBytes(b)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMeta: %w", err)
	}

	return &CtrtMeta{
		LangCode:    langCode,
		LangVer:     CtrtMetaLangVer(langVer),
		Triggers:    triggers,
		Descriptors: descriptors,
		StateVars:   stateVars,
		StateMap:    stateMap,
		Textual:     textual,
	}, nil
}

func NewCtrtMetaFromB58Str(s string) (*CtrtMeta, error) {
	b, err := NewBytesFromB58Str(s)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaFromB58Str: %w", err)
	}
	ctrtMeta, err := NewCtrtMeta(b)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaFromB58Str: %w", err)
	}

	return ctrtMeta, nil
}

func NewCtrtMetaForNFTCtrt() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"VJodouhmnHVDwtkBZ2NdgahT7NAgNE9EpWoZApzobhpua2nDL9D3sbHSoRRk8bEFeme2BHrXPdcq5VNJcPdGMUD54Smwatyx74cPJyet6bCWmLciHE2jGw9u5TmatjdpFSjGKegh76GvJstK3VaLagvsJJMaaKM9MNXYtgJyDr1Zw7U9PXV7N9TQnSsqz6EHMgDvd8aTDqEG7bxxAotkAgeh4KHqnk6Ga117q5AJctJcbUtD99iUgPmJrC8vzX85TEXgHRY1psW7D6daeExfVVrEPHFHrU6XfhegKv9vRbJBGL861U4Qg6HWbWxbuitgtKoBazSp7VofDtrZebq2NSpZoXCAZC8DRiaysanAqyCJZf7jJ8NfXtWej8L9vg8PVs65MrEmK8toadcyCA2UGzg6pQKrMKQEUahruBiS7zuo62eWwJBxUD1fQ1RGPk9BbMDk9FQQxXu3thSJPnKktq3aJhD9GNFpvyEAaWigp5nfjgH5doVTQk1PgoxeXRAWQNPztjNvZWv6iD85CoZqfCWdJbAXPrWvYW5FsRLW1xJ4ELRUfReMAjCGYuFWdA3CZyefpiDEWqVTe5SA6J6XeUppRyXKpKQTc6upesoAGZZ2NtFDryq22izC6D5p1i98YpC6Dk1qcKevaANKHH8TfFoQT717nrQEY2aLoWrA1ip2t5etdZjNVFmghxXEeCAGy3NcLDFHmAfcBZhHKeJHp8H8HbiMRtWe3wmwKX6mPx16ahnd3dMGCsxAZfjQcy4J1HpuCm7rHMULkixUFYRYqx85c7UpLcijLRybE1MLRjEZ5SEYtazNuiZBwq1KUcNipzrxta9Rpvt2j4WyMadxPf5r9YeAaJJp42PiC6SGfyjHjRQN4K3pohdQRbbG4HQ95NaWCy7CAwbpXRCh9NDMMQ2cmTfB3KFW2M",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForNFTCtrt: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForAtomicSwapCtrt() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"4CrYZXauEHTHvUcNbU2qxvYSgdxPkSBum4PAUfytuZu7Nn56L59op72uKJUBMnF8dk8dLb5k63M9236s8S2yH4FTeWFP4zjfpkx9HGwjAuh6n6WJyxWE1S5HHH2cQLy4xk5B4iMpQKyHQwrQDn3zWwQQPsrfnwaHX1F3V2zKHKx15QYATS784BGfz9NeY72Ntdz2Cgsf6MLQE1YKdgdRfpryCwadqs5xchALCPYLNg6ECSxzPDa4XdS8ywTWzRpiajTGZA1z9YoQZiUMYBwM8S2G4ttZJkgrWTqpXuxberLv3CWZm7kp8bwvg577p8kJ7zAugTgaBU9vzSRFzi3fWtGEP1TPuMCjLSQfskepjoLXbPHyVMmvLZGbjx2AwCyGikdXBdLJWhheL6rnveiXJQfV6zfgF9zeMTpg9GE5SRstGHFetCZwfe3qCPV6vUWrobmWusQ9rDkj5uUXVpjwmBseynCnKNS1CZKDnBDy6mWBDPHNCtuDdYCamqaSEh1nx9ykk4vVJggzPJR8awFMHh5iKPRL9LGhuqbqs4rDPVsg7BCrdaszTGEBEHjfqF51K8PF9kUnPQJvGkf58MrLj2SAArizmZYcnpGMwdfYqGxrjz7xaJGZVAqvFbWFDk3x18ozp58PwFM1fdAn1dn15fKCsiQoqZBtVTxSd4GRJ2tFvBzgUJjig6hqhHqCqobCbpes8LoTdtDCHE5Co3YBnrYN19vbESu2W6LMpwrPPgd1YUeHx8AxR9evasFYrCjmnvBkEyefu5n66yTPYNXfjAk646dHmWYJiUPp1oWDXMjfDJ4xif4BXhRwBtfwgHoDhU2dMV6E7cPVppXxeVL2UsFCbqsafpNcDmhsrGEDAWmxJ3V8KymyuNugM1CHTEiwcTb7GXd4dD3UznDVoJEVrmBveETvCuGVNfGZ4zGZnURyoJHzMkDKPWFQhqgVYLoRuLg4MtquRAaSEKixtXiSJZFKZvQTzMbJC2ie3bnyQoX3x2C9pPpzp3uFKc1eGpgafgi8KoyiqiCMJvfzi8v8DiyTZ9QPENAtwToUpf6vsn1C4HhDzGb9otfigtVuh9JuzsZkJbd4r2rU8sUcKWZcaLF1uX4EdZiEfiW3aV5cm1L7oEJX2w4rQbNiFZWGUpS31WS6mYtWkSTnQupp7rggs8sQxcdWK8WamLgonF4mhXkY12Y2U9AXDJifMKr7mzxiFxZumPWxGn8A1PtTp34wcuhykNMesekwDgWGRCWca9w3YDkeinoD2QmV5ivF2GfHTKhCVH5pkGmBZczeVMA2ZTWb5DTM5qQA9vRy43aJipwmYH73ssbdF7N96678x4hsdcFXXJooRbDtuEY9UkhFPtFMjzD7D5uvXzN4qTPFSyoumwH3ag6cmZMxxQdHNJAm7vitgDpRy3HM174KpjE7uUQXtVvMKEYeAWus24vwW6M4i7APsVg6FeJTgGJJHAHFJFJ4YrZ1fmzgGFnugfp9g4hMuo9G76dzzkZetLhweJCggXBRVpNeRzQ9xmtuDN3wmiyQ1bLSx2ZtNcmWqzbSDsUnCezXtbF4CURyp2djUKo2DRza78CHpmUgHHVai8JrAxPwS6gB8mBg",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForAtomicSwapCtrt: %w", err)
	}
	return cm, nil
}

func (c *CtrtMeta) Serialize() Bytes {
	size := LANG_CODE_BYTES_LEN +
		LANG_VER_BYTES_LEN +
		c.Triggers.Size() +
		c.Descriptors.Size() +
		c.StateVars.Size() +
		c.StateMap.Size() +
		c.Textual.Size()

	b := make([]byte, 0, size)

	b = append(b, []byte(c.LangCode)...)
	b = append(b, PackUInt32(uint32(c.LangVer))...)
	b = append(b, c.Triggers.Serialize()...)
	b = append(b, c.Descriptors.Serialize()...)
	b = append(b, c.StateVars.Serialize()...)

	if c.LangVer != 1 {
		b = append(b, c.StateMap.Serialize()...)
	}

	b = append(b, c.Textual.Serialize()...)

	return b
}

func (c *CtrtMeta) String() string {
	return fmt.Sprintf("%T(%+v)", c, *c)
}
