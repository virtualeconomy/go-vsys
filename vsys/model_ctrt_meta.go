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
	CTRT_META_LANG_CODE_BYTE_LEN  = 4
	CTRT_META_LANG_VER_BYTES_LEN  = 4
	CTRT_META_CHECKSUM_LEN        = 4
	CTRT_META_TOKEN_ADDR_VER      = 132
	CTRT_META_TOKEN_IDX_BYTES_LEN = 4
	CTRT_META_CTRT_ADDR_VER       = 6
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
	langCode := CtrtMetaLangCode(b[:CTRT_META_LANG_CODE_BYTE_LEN])
	b = b[CTRT_META_LANG_CODE_BYTE_LEN:]

	langVer, err := UnpackUInt32(b[:CTRT_META_LANG_VER_BYTES_LEN])
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMeta: %w", err)
	}
	b = b[CTRT_META_LANG_VER_BYTES_LEN:]

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

func NewCtrtMetaForTokCtrtWithoutSplit() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"3GQnJtxDQc3zFuUwXKbrev1TL7VGxk5XNZ7kUveKK6BsneC1zTSTRjgBTdDrksHtVMv6nwy9Wy6MHRgydAJgEegDmL4yx7tdNjdnU38b8FrCzFhA1aRNxhEC3ez7JCi3a5dgVPr93hS96XmSDnHYvyiCuL6dggahs2hKXjdz4SGgyiUUP4246xnELkjhuCF4KqRncUDcZyWQA8UrfNCNSt9MRKTj89sKsV1hbcGaTcX2qqqSU841HyokLcoQSgmaP3uBBMdgSYVtovPLEFmpXFMoHWXAxQZDaEtZcHPkrhJyG6CdTgkNLUQKWtQdYzjxCc9AsUGMJvWrxWMi6RQpcqYk3aszbEyAh4r4fcszHHAJg64ovDgMNUDnWQWJerm5CjvN76J2MVN6FqQkS9YrM3FoHFTj1weiRbtuTc3mCR4iMcu2eoxcGYRmUHxKiRoZcWnWMX2mzDw31SbvHqqRbF3t44kouJznTyJM6z1ruiyQW6LfFZuV6VxsKLX3KQ46SxNsaJoUpvaXmVj2hULoGKHpwPrTVzVpzKvYQJmz19vXeZiqQ2J3tVcSFH17ahSzwRkXYJ5HP655FHqTr6Vvt8pBt8N5vixJdYtfx7igfKX4aViHgWkreAqBK3trH4VGJ36e28RJP8Xrt6NYG2icsHsoERqHik7GdjPAmXpnffDL6P7NBfyKWtp9g9C289TDGUykS8CNiW9L4sbUabdrqsdkdPRjJHzzrb2gKTf2vB56rZmreTUbJ53KsvpZht5bixZ59VbCNZaHfZyprvzzhyTAudAmhp8Nrks7SV1wTySZdmfLyw7vsNmTEi3hmuPmYqExp4PoLPUwT4TYt2doYUX1ds3CesnRSjFqMhXnLmTgYXsAXvvT2E6PWTY5nPCycQv5pozvQuw1onFtGwY9n5s2VFjxS9W6FkCiqyyZAhCXP5o44wkmD5SVqyqoL5HmgNc8SJL7uMMMDDwecy7Sh9vvt3RXirH7F7bpUv3VsaepVGCHLfDp9GMG59ZiWK9Rmzf66e8Tw4unphu7gFNZuqeBk2YjCBj3i4eXbJvBEgCRB51FATRQY9JUzdMv9Mbkaq4DW69AgdqbES8aHeoax1UDDBi3raM8WpP2cKVEqoeeCGYM2vfN6zBAh7Tu3M4NcNFJmkNtd8Mpc2Md1kxRsusVzHiYxnsZjo",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForTokCtrtWithoutSplit: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForTokCtrtWithSplit() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"3dPGAWbTw4srh5hmMiRUhHtcxmXXLUooKGAfnmz11j5NruyJpBzZpgvADMdZS7Mevy5MAHqFbfHYdfqaAe1JEpLWt1pJWLHZBV62zUhLGmVLXUP5UDvSs24jsBRHqZMC71ciE1uYtgydKxCoFJ3rAgsYqp7GDeTU2PXS5ygDmL6WXmbAYPS8jE4sfNUbJVwpvL1cTw4nnjnJvmLET8VmQybxFt415RemV3MFPeYZay5i5gMmyZa63bjzK1uMZAVWA9TpF5YQ1NTZjPaRPvQGYVY4kY9L4LFJvUG2bib1QaNh7wUAQnTzJfRYJoy1aegFGFZFnBGp9GugH4fHAY69vGmZQnhDw3jU45G9odFyXo3T5Ww4R5szegbjCUKdUGpXf9vY2cKEMJ7i8eCkFVG1dDFZeVov1KMjkVNV8rDBDYfcp3oSGNWQQvGSUT5iGUvDRN8phy1UpR3A9uMVebvjLnVzPx9RyqQ8HaXLM8vPhLuWLoh5hk1Zi1n9nwz55XvKDYjP6eeB55yK5vpg8xjaYDnw9bjYV7ZmS7LAsHvXfnwi8y2W6vk2hGvs4rtR1vNRZSQMPGRRSuwCRJL1yngH6uHWwm2ajWxc684jApuoLdyjZomfCtdpabSyU3kp9Lrn8zT8BVY332sJPQU6gTQi8ke9s9dBxCae4cfSQM6HhuBmFc5KKWHCVG4bm4KZRYbMtidw8ZZnjaAMtcGq7k3Se6GXaTxdS3GcuttB3VB7njypyzuqAcfCdYb9ht8Y1WuTCZ1aLsXsL6eydfk2WLJVrqYpbTk6AchV5gMAEopvc3qXvzrDCedjtNsDmA56Lh6PxrrKr8aV8Wzz8aMaQ88YsVBpE8J4cDkxzo31AojhzEGVBKLmpb3bjmsaw9VkpB6yL8ngYs8eJMSPdM289TSMaEmG4eHt1jezpHTKxkuB9cwqcvhGNLWuv8KXQkik5pRMXV67Qs2FvjpzeJ81z2hnVh1wCtsa6M6qAG1gsqLHa1AVMRzsowafC99uDexwWMBS2RqsZWZBXJcUiNVULjApSnoBREYfHYEpjJ152hnTYZCAwpZMWEkVdBQpZ3zk8gbfLxB4fWMfKgJJucbKPGp1K56u7P8MHQu9aNb9dEof2mwX8rTHjk8jSQ7kXVX4Mf1JqMRWWftkV3GmU1nqYhxRGu4FjDNAomwTr5epHpcMF6P5oiXcLWh5BFQVmGYKz129oizAyUJBsZdxr2WZEGDieLxUg8cve25g28oTuCVENST4z1ZsFAN9wTa1",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForTokCtrtWithSplit: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForTokCtrtWithoutSplitV2Whitelist() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"7BekqFZ2yZqjiQFFnsxL4CDRFWCjHdZvFXQd6sxAgEktxwn5kkR6vkV27SFC7VmhuMysVfunZWTtHAqPjg4gGz72pha6TMUerSUSXSn7BHaVexyQJoUfqDT5bdr3XVpok1mU2gT29mwtJ6BibizpAEgZZncZauDnvqrWWdkCmRP8VXpPBiPEaUZuq9eRusrUcc5YHshhN6BVkArN84tarVQH3pTRmiekdQveuxFw4r4weXUxwEGCkYX3Zqeqc4mmRsajVCQwV5DuGTEwaBVWiAAfHLGPFgJF6w6aP3d22tdBRLqZ2Y4G5WHdhMunNDEZ2E79w7gbwqDXtz3eVfGtyET5NZEJGmM2S8pZSn2MPjvfPAYZMa9Zd4WXnPLZng1pxjYvrpqPDy27VQu1rhvxXMNPVMdP9QyCQSoExZUot1FmskS1NcmzKfguwsSWR1Z1py58iVDKm8t7x7RnaP7avcjtvixJQkPGg7qaxBKfRQ26vFePWeNdkbJwQJvqComvjEg3hEYjQrysk3j3M9QWEgXQzRqTPTFEVCTJSbdpL2GyYXYC4cLcB81UzJuWf2zoERNPdfpHwumoaaaSutfg7dccbWRaqogrBf6u9PfANQm9TsFca37UHhxvsq8WZdu71NQCY1V7w9NKKLbHF7MjjyCs6w2TM4Ej9Tyj8hFR4qo3MosgSbmQt298aEB3qQHVF8FshVwGg2vqAK7PNBHE7KgBgXQJiVRc4X1XZvWQt4uASvMowRECURoMZ17z2s3LnDrQYVqYedfzjJXxwsWXQkoQp51WWkFfp7QStBtfEhdUx15wtD8sjDdNrda8n3P6sNrN8J7NXxH4JPE7DzLLCjPSbn5Yc2jzomULSRiQN2yzC5qE43XiHB89VFqTHTduCFbP3Pom3uc5iBgjW9ky8LyPBMcsqQZSv99adjgbKpeaGPtJN6iUQ9mae1ddw6SBKTxZVZvqK6k7dJBjJ5UsFDyXLWkm8jogkRCFBfXPxmxyB5ihqk2wnsWNEbKEz6sg6RJqy5SR9A8r3QEx8FZt5z4DJpHyUAoi6KKVHEJfRvdjtjSDrayG2WUrBCgTTHsyGZEnuXLRXpy7XmdzFSwKSr4p7NPbAqt44yHdgjycn2MY5X1P9rneBdh4LukH3syRAarjmTSZr67QexRE4cca5fnxUZJ2zYNWRynqWmZy6aCBLBQziP81bHHbN5WP9MMseovCvzTpMso9TB3QLSRkCphJpyvv9qLN4tpFB9r9g3UGhTqqJFvxJDcLwR485AqLymM91kMjTvodniJ4coymUeE3MjGf2P67z4UiBDBxnzWbkCzmaPpkWFY9125hg9SovQrJnn9zzpF5smp7oiHhjrkzyi2G4qWVidtaWi6TipZFXwb8z6TSSjZkaj4SWexgnE2bUKeJS9P1xYwVSX39At735bqhfKCNP29n7UzX7bMwQiTWWK8bCiCpYSXfcfUpxtbYXdHgGMEZzpzawS9H5UeFiw31rS5Caps7QQJmMeetAuDa8tsiMJ9QauABLfJ4G6Hjkn5GM9jH9yXJWj2boH1U4ErVQXbr9KvmSsSsLeLLc3XeKQaczMtLroQax4D5estuP3Cy1gfqhbTsEWL2HkF7dUKDnuLmzsjv3kZXF9PMhcVR1Qj9j8KaYWYqKYV5TxXkrPrzSVa1yYEjU71A6ZYW327vgFJYFUJmx9vqTGym3yRiSoJiaYVfgf8iLwqS1EKSTMiisxE8hCHfKiew4YmiCTxPkq7pc5tHrKkogoRX7GdDnX93BsxGACu9nEbXwDZERLFLexrnRKpWDjqR2Z6CLWhXNPDJYMcUQ5rfGAhgu4ZK16q1",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForTokCtrtWithSplit: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForTokCtrtWithoutSplitV2Blacklist() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"2wsw3fMnDpB5PpXoJxJeuE9RkRNzQqZrV35hBa366PhG9Sb3sPeBNeYQo8CuExtT8GpKuc84PLMsevNoodw7YGVf24PKstuzhM96H2gQoawx4BVNZwy3UFyWn156SyZakSvJPXz521p1nzactXZod1Qnn7BWYXFYCU3JFe1LGy35Sg6aXwKz6swFmBtPg1vBeQsUq1TJ5GXkDksaUYjB8ix9ScNNG8faB1mCCMWwfrcr6PyBA7YeHsTLD86zuviak6HQEQQi9kqVr4XhnDJnZyiTKGcNDo49KZyTyvkPmkFyDEhLf9DYrJM3niePqtDQ9unJj52Bku7f47hrxo83eSh3UPncyq8Hti2Ffhgb8ZFCFdnPyRDEZ1YbKFGAsJL3h3GdPFoVdnYySmnVJWrm6fVUdGgkA5ijMeqEUpXte1m7MFYCJ1wQchjebpLk3NnZzrT8FysUJVUgUzmkoSniF2UPEPXuF9cyWFWGGoZjfDWqarPMi7miqdCPQMMw4QRvSWkB3gVyeZykAvKYzXm8wYGV6HDbipZeVoyZ1UVeR6E5C4VZQmjs4GupAR9EuT5mt1ALFT4HyAMX6RCRxjeHoSgnnUJcEiRHapAYSene174RvVkRGLTtonWTYnsXUrtPD6xks4GdpQWQv89EdNWFEtmMfyVvUEFuTPGXUS5TuqYxCzg8Gor5WjPip2wDmoMYQ3wikJoRpYSfRVw88RHQPBmkHrpeHYWkAx6N7Yk4WwgBF9SVVtEWnWmPVVbuH2bQrvks4iGL8DnmEiLMs6JuFsg3a3cMHqbdvQgfu72XYKFqQzzDbDhaqFKpR3bxgMMiJvGbPuydPk9DCsG5KpqZepkkD6RGhWTQzga9G6y6ryctoGZPBHpFRwirALkksarQSEuGryhatvnjqG9U14zyW2KvJYrErMyUVy3wNK5wRqAKMjE6hFPdoH9Cn6TYQLebVTBoYTfimn5gBmgnKqBtXSfUxiwrjWujQPGxgtbNCL1RXRNRJ8nrtcpphQyRVZ8JVeubYq1zM7G1AUurEyAQi64rcbsimGptcXMAvt9TbwDjpUGRWvF6dyw1XijcukfZBQh1fG5C8peumkGnP8PemmYWKP7qsifNc44PqnNG5qYVivwtK4sz2h3B6pwneX8XNYtGSjVJCb6gJ7oDG45shocvALKNu7LwfJxXT7MPAdx7CjbHU5B3qs71wJphwkc4yWa6hHTamPTGRFGuhJa4kFfeGMctE1WZrFe47L32fKZkSxaX1sguoi5w9UPHw6udJiKPYENSSbASYpfS9q8suCs1bbq8jdMhCwoGMDZaA4MNAW1Q6sLSX6ezZ436AMbVnXZLQW8jdBaX8rvRSMJu8fdYU9PHq4MkoczxNz5jNvRiTX9jTpN1Z1P5rtgnf6XN9vzTLdqsvwZcXqvSdBwdTVgk7qn9uNjuFZEgSmA6rnPhSu6TMxJLmjKP93uqiNmXsj1NKtqBZiHjrRaUzA4pAFEyfZTdo8oaDH7umSBU2s9ff5Cruds7cYFopLm2KavHH33S7BczL7FMXAcqrESiPUzhUhHbkBKHGiCAUMVE8zxo6Eo85W2PGn6D39MaUfahEmzq8zxmrDQdmagx5EQZUev3fNCFzTzU4zpY1sra5ZPknXJkyKKfj4r9xy9Kfd8s5hsiKFyX6V1Kc2T1Ehpdkobwb7Wc8V1n1GaeL7jRgvhVg1inPaWZ3zyqNBjxnzqtLpZor3VdXLo6SikzWNahCMLNMXaoBvmJDEJUazC9qGxin7SC3YWCTAyoskJRhVMp592ehmpruu2azeCHBF2rzP6LabikVfkBSeAzGQKVeiEkU3devRNpjNM4YDXQDm9wbkPKWrqBK4SRdo44PRYG3XwNhu2gpNX8b9AuirrbRPiaJ1tJ7rzodHzLheMyUMXRB9nYx8JgrhkZzPZa4oUxo8JUNuKZnn7Ku7fEt5y",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForTokCtrtWithSplit: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForVEscrowCtrt() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"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",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForVEscrowCtrt: %w", err)
	}
	return cm, nil
}

func (c *CtrtMeta) Serialize() Bytes {
	size := CTRT_META_LANG_CODE_BYTE_LEN +
		CTRT_META_LANG_VER_BYTES_LEN +
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
