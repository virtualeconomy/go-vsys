package vsys

import (
	"encoding/json"
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

func (c *CtrtMetaTriggers) UnmarshalJSON(b []byte) error {
	var arr []string
	if err := json.Unmarshal(b, &arr); err != nil {
		return fmt.Errorf("UnmarshalJSON: CtrtMetaTriggers cannot unmarshal: %w", err)
	}
	arrb := make([]CtrtMetaBytes, len(arr))
	for i := range arr {
		d, err := NewBytesFromB58Str(arr[i])
		if err != nil {
			return fmt.Errorf("UnmarshalJSON: CtrtMetaTriggers cannot unmarshal: %w", err)
		}
		arrb[i] = CtrtMetaBytes(d)
	}
	c.CtrtMetaBytesList = arrb
	return nil
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

func (c *CtrtMetaDescriptors) UnmarshalJSON(b []byte) error {
	var arr []string
	if err := json.Unmarshal(b, &arr); err != nil {
		return fmt.Errorf("UnmarshalJSON: CtrtMetaDescriptors cannot unmarshal: %w", err)
	}
	arrb := make([]CtrtMetaBytes, len(arr))
	for i := range arr {
		d, err := NewBytesFromB58Str(arr[i])
		if err != nil {
			return fmt.Errorf("UnmarshalJSON: CtrtMetaDescriptors cannot unmarshal: %w", err)
		}
		arrb[i] = CtrtMetaBytes(d)
	}
	c.CtrtMetaBytesList = arrb
	return nil
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

func (c *CtrtMetaStateVars) UnmarshalJSON(b []byte) error {
	var arr []string
	if err := json.Unmarshal(b, &arr); err != nil {
		return fmt.Errorf("UnmarshalJSON: CtrtMetaStateVars cannot unmarshal: %w", err)
	}
	arrb := make([]CtrtMetaBytes, len(arr))
	for i := range arr {
		d, err := NewBytesFromB58Str(arr[i])
		if err != nil {
			return fmt.Errorf("UnmarshalJSON: CtrtMetaStateVars cannot unmarshal: %w", err)
		}
		arrb[i] = CtrtMetaBytes(d)
	}
	c.CtrtMetaBytesList = arrb
	return nil
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

func (c *CtrtMetaStateMap) UnmarshalJSON(b []byte) error {
	var arr []string
	if err := json.Unmarshal(b, &arr); err != nil {
		return fmt.Errorf("UnmarshalJSON: CtrtMetaStateMap cannot unmarshal: %w", err)
	}
	arrb := make([]CtrtMetaBytes, len(arr))
	for i := range arr {
		d, err := NewBytesFromB58Str(arr[i])
		if err != nil {
			return fmt.Errorf("UnmarshalJSON: CtrtMetaStateMap cannot unmarshal: %w", err)
		}
		arrb[i] = CtrtMetaBytes(d)
	}
	c.CtrtMetaBytesList = arrb
	return nil
}

func (c *CtrtMetaStateMap) String() string {
	return fmt.Sprintf("%T(%s)", c, c.CtrtMetaBytesList.String())
}

type CtrtMetaTextual struct {
	CtrtMetaBytesList
}

type CtrtMetaTextualJSON struct {
	Triggers    string `json:"triggers"`
	Descriptors string `json:"descriptors"`
	StateVars   string `json:"stateVariables"`
	StateMap    string `json:"stateMaps"`
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

func (c *CtrtMetaTextualJSON) String() string {
	return fmt.Sprintf("%T(%+v)", c, *c)
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
	LangCode    CtrtMetaLangCode     `json:"languageCode"`
	LangVer     CtrtMetaLangVer      `json:"languageVersion"`
	Triggers    *CtrtMetaTriggers    `json:"triggers"`
	Descriptors *CtrtMetaDescriptors `json:"descriptors"`
	StateVars   *CtrtMetaStateVars   `json:"stateVariables"`
	StateMap    *CtrtMetaStateMap    `json:"stateMaps"`
	Textual     *CtrtMetaTextual
}

type CtrtMetaJSON struct {
	CtrtMeta
	Textual *CtrtMetaTextualJSON `json:"textual"`
}

func (c *CtrtMetaJSON) GetCtrtMeta() (*CtrtMeta, error) {
	cm := &c.CtrtMeta
	l := 3
	if c.LangVer == 2 {
		l++
	} else {
		c.StateMap = NewEmptyCtrtMetaStateMap()
	}
	b := PackUInt16(uint16(l))

	var err error
	bytes, err := B58Decode(c.Textual.Triggers)
	if err != nil {
		return nil, err
	}
	litem := PackUInt16(uint16(len(bytes)))
	b = append(b, litem...)
	b = append(b, bytes...)

	bytes, err = B58Decode(c.Textual.Descriptors)
	if err != nil {
		return nil, err
	}
	litem = PackUInt16(uint16(len(bytes)))
	b = append(b, litem...)
	b = append(b, bytes...)

	bytes, err = B58Decode(c.Textual.StateVars)
	if err != nil {
		return nil, err
	}
	litem = PackUInt16(uint16(len(bytes)))
	b = append(b, litem...)
	b = append(b, bytes...)
	
	if c.LangVer == 2 {
		bytes, err = B58Decode(c.Textual.StateMap)
		if err != nil {
			return nil, err
		}
		litem = PackUInt16(uint16(len(bytes)))
		b = append(b, litem...)
		b = append(b, bytes...)
	}

	cm.Textual, err = NewCtrtMetaTextualFromBytes(b)
	if err != nil {
		return nil, err
	}
	return cm, nil
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

func NewCtrtMetaForNFTCtrtV2Whitelist() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"3g9JzsVg6kPLJKHuWAbMKgiH2aeZt5VTTdrVNeVBQuviDGJnyLrPB4FHtt6Np2rKXy2ZCZftZ1SkNRifVAasWGF5dYt1zagnDrgE52Forq9QyXq2vmyq8NUMVuLfHFDgUC7d7tJPSVZdmhDNzc3cR9WcobXqcR3x923wmTZp63ztxgzdk4cV39TJLoTBLFguFKjqetkU7WUmP6ivMfcvDzMBzgq48fjJ1AYn5fxt31ZV6tAorCQ4w2zfekL8aUEhePgR66RXSBggiqQhTcw7dGg8xkGtRh3wkAVEbFuZa78R1C9cUUytbYM5fi17AE5q9UEgegxMMpZgsk9YNHs4mx4NPLj6Rz5DK3QwbeUbaVWceSqssYS6GodJ41bEm84x3aQrqQK33tHSPRy9uAr9ku773fZuHWPEeNoEDdsnUVsxCKQ7AyM5K1JVFRFwMABGGAnkYsFV23pfLFktBSvAJkzo8Hi6Wss7ZEBgSDeCJJohqoxmsR7L8kcfjRwy3Rb7VU76LMuqGrBfb39uUy5qdxRqAMFtwE4imkxxX6akuR7RMd3RmKQ2W7TXMuWZNyJHd4c17ZJrSCQNAXQ2iKXxSbUoDUmetuCud81SQonTjomq9RsGqRvaV2iGjHUb4wvUuKhodE4dF8xrNWXQxfPpwed1mUEuUPmhppY7Lg7p5EJyXVYDr4ybdsmYohDFgTDbGs3mZBmgUpEVAUC4vJrXqWWv8gjw8j5xabF6QfbtcWrbrVu4sTtMGzybVAoeB4b1x3Rkd67ABWnmzHfDxMopfb21TSDGpWLnSQeRn2gA2jnLUokb8FXUHG5qttmLNzG7RY1XRmC7TKRQ3X5JqGbHbN4rhUxU8iQUKpACWsyGuEP8VrUNvx41sMEbfReZ8ay7v2cQEtmw5uFfXMmAcsQBrRdxsHTaN5Cpu7Ak1pRvZzQKKesWuHLuUgNStdqVpHih4cTk1YzoJJ34spDa7FYhzTWTSVJBwHvYy5WQxrXnXAXBmMeNVroX8x9gT38LeqJ2z4KoAWnj2o1waKB8TC1JXet7sXHttGWDs7YHJHNEy5CcWkVCPnt5xVTq9ZwPkc4EhLQDWortL35e75vyQR3F3tW2Pr89UiPSNWEXxC5L8apavKVyv9zUcWUwShd5bdcfKa1CnLSMhW9DE6CT4APWKuPdxW9hLgkYZziJtN4WebcbA5PbG8hrkhU2E7easz3pRJQ49vhMtSf7tKTf9NDwZuuZ9ix9q5TZMzYvNbg5rk9P6uoPLRZk61J2LpQv8K7YLBrcWSduPsxWWjiCvxL7bW8vA8gWQocxfuXiM5i3wdA1zLx8As3Ydufi2S3nk23BwRjZhjhh7BEq7p1nwpqP97PqqW2CpMJspEHdHCzRR3fBJw6mLdSGAYeia22r2uJm1o73WrPFTt9vQwCLXMKS3WMd3GpRmR36n3C9Ed7xdnFcRDYZBgLis63UEvczGvH9HS8MMHkoAXE3wuahEzYZEd1NxJXSXFhe2h6DJbABXQKMMkZdPQmGJkDhBPTh9nZ9DgGHhnnitxQ5ESfxqvqxwuVubAXTt3psg8LS2B16mjDGh9",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForNFTCtrtV2Whitelist: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForNFTCtrtV2Blacklist() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"3g9JzsVg6kPLJKHuWAbMKgiH2aeZt5VTTdrVNeVBQuviDGJnyLrPB4FHtt6Np2rKXy2ZCZftZ1SkNRifVAasWGF5dYt1zagnDrgE52Forq9QyXq2vmyq8NUMVuLfHFDgUC7d7tJPSVZdmhDNzc3cR9WcobXqcR3x923wmTZp63ztxgzdk4cV39TJLoTBLFguFKjqetkU7WUmP6ivMfcvDzMBzgq48fjJ1AYn5fxt31ZV6tAorCQ4w2zfekL8aUEhePgR66RXSBggiqQhTcw7dGg8xkGtRh3wkAVEbFuZa78R1Bw8Fc7fND3crHRj8pY66QYiaksdHixYVm4R68ez9K1ndEZq1ShQBs5DbvyoFGc4Dr1Yosv5VKJbqaB5fu7ZZ8SvB5RVYqSsN9tTTmUinNmJ4v63DWvH2N7WnFq8JYPL4RpEpnvBYnSUdAxN44skS45uVi5F4bkueAXbgUeoir82hTgLvgnf573Ziw9Mon4STtfhP8Y5DKTqA2gM44MmVkNWW7WwNDXerdYwD65QMG7BSSU9UhH6eNvay2LYXNph9heAWYwKcQPJnA7niSZto23XaFoU8kGRUoDNvofQw1XJkdTgVgLt5yz8HbGxnXT5AdKa3YNyAnq4KgXjU4W3Xj8xWqpYHX54C8GQF7poCM4E5XNDXbgExoK3bS4WHkbmwJJJzJ6MtsiyZnmSYGs7HhfcueFH4SjjNKevcntrC4Kenc6tygSWrSzefdSC78XrQ5bgSp24wKoX4WxUUUky8KB9NvWGHYF3x8Bg59HwH67haNB9wejM8Jj5a88XoVTYAqMh6z8zuZUqANshYRaxjxYLaV2VATrTKM13zMARaBVoDRFKtYiE8CmFKeequ9HdWix6CmCEtKQdCC4UmtYJ1Ch4qpfjKyMP4Bd7YbKLg928ZHFiLN2Uq1KLfbn1V83Xe1xPGwkX1TCsJpBXyqmsByaYUckFgkCNNvkpuAs1dA8HLLrnd1Tx6zT99vDaPUr2k9nLQ6o1hjPyK1EPBVg5zxrnaSP446m54CemwNPa1UECFx6sEhrL1EbL1yQR7cfMnrr82z9iSiSMZMubfEhPyuD58TYjSRGd1XRSnhjo1tBwN2k27RsNtdhAmH2u57eCfDQpnMUnBkSZj71o2Kk5cMfMxNWLBYr1w7Ma8yJriQYNedNo5fG5XVubmmd5H7YpVAjPKWVVru3SQXR7AHcLv834pCQD7EjYEbNdFeheaDiA1yp7amZrig3cd6jabMPoDSvP1GxX8HrUnv4hCvSmDivGpFvcGJnGbNuSHTP8qHTAf8jVFeMpeMiLH9rP9qcpMAhh9mAzmj5pVhZZBuiWFor8empJoKGv2RcUFRALEFDXoYaPrri7oCypNeWS4eiVum8fm5hx3CMY9N2HMqMrokCCTHceiHYKfgjYRnXaJsJUs28rPyqqxAaxUj3qNpaB2D6x6nc4fKLSZyuUCgZSmRPPBWWugRNRDxppG6ecA1hkNZDX2NQY9erhuMYX9jhVCLb6NLVe5euWFkvBjF4Y7qfpKM1uLSZvxd4gmA5VGA99vKFkYUwvBB5TNPnupdECD9",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForNFTCtrtV2Blacklist: %w", err)
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
		"neYvWKcRQc7czFuzGcHiQrZDaFXXjyX3TeD43tojTqu8vRgdDaF7B5wJupyvKn7RMFQrb5dRMzf87VPa6kSk5v4zWYQAvDqvf34uByuekBA3CHwyBUvFmN2LGUx3ktTGcf5k1zH79jGnY1waSXqsB82348aSpKyzUiKvFko1DFM87FS6SxntjFYVyaZtCqvyd3NMRPZXaZUqLuHHJUNd63zhxMoYA6QokeoDnCM4HWXx3tvz9KYP1L8MpkusEac6yv5FFqhKkzBSwBPkSH73VtGYdFNpeBuTCeracN4WbAWnDrt8jD4cnUYNDQxuPTuczuZ8UApc3wYcM6Vp7LNgtZr5X9WxrarU4N8AsDXMKuwrRDQ3nZprW3BZFARjRhZs9TBqUkazXAbm5k3jfYqEncMPGBbmbr3HdeohsCv8t9uWPT7YBVr27ykaVDHc7NSxxFCHVefqGYQV25AwwGE7ax6MiZCwAibbuZS2hwXXKnTHY89K8jp7hqva9WvMtXtHaDyoXiJdrAaUto63F9bkWrJzVMkdsdqdm4BMF6Kgg7q4H7fwyKNeDxjYeYVT3SYhhntCaKqnNmCpENScHCwCEiJAM9S9ZTHqE3so8kt31rmx92xbD2pQgNHRzSVenDng7DxGJr1sHnzciX6cQRbgaVWycDJiqax79KRxWhPAnYyJQgh1RHHKz8utpocsFgYm8rkiwnzY3biA48EA2FaTqno26N4W58nU14g8xAG4wjGZof9NAMNqc5zBpKTDYov53xLEEJArhbrntyyAYUiWdpqZznzJoaXf6EitKZihRXBuGCgCZ8dpZCwwsfnpEmhNBZEyxGZ3h1P4aG9UVTSu35UNSK6sqvstZtz1bGiYycY4dxdUqQVVzgLAsVMUkaWu7ETKPbw4CJv72oDN48LgBFmLKdtCrkLipVf79CFS6xVUWJ7usK4XxtmnWDjGWvYQNZ62QSWCwTy9SXZDZQMk1qRXYBsfbfpusXGPM4ofT9F7D1GrmEevNZ9pqmkLdgchhm5iwR4hnbsZ7hzJoprLMUG8wtbDKpZeuNYTio4KfRRhAYJmYqNEL15hBfw2yWYDYttUQPe3VYcVE13tWFxuLpjwNgdycHVZxfoLMYoRUMKyCdz2sXuTPQ9EbCF6BEM29ncp2JEiZxJ3unPnwTu3vVUb4ad236qkQ3CfubCAkNLw7huZquMAktPbWEVCPAAp4USXeH84Yt91z3LqtBCx6f6B9UFwrCtNQ1gX8NpmA4sBwJE5iECsc81JKVFVMNhdDZB7wb2HaRyuuiRWhAbQujJGs7dQ8rnaXff9TR4cNK21L36uALBP8iKicb1JRzT8t6idopvSJLphAK3qBQa4Tc3UJpLNsJgMPuGfcvy4pjYH5tL2t59JGibinsmL6fJcYvhXWxSsrZviF5sJxstUvzGmZjxZ3gxcQpGT9CumdE6UJqkrNrUoqt7ZDW4RpPH3fYknrbNsM3gra2R9v2Vc53SQsu1w5SWxTHBqxCAxndBddzM7jhZBvgheJaN3eNh3NdM9WZDZHzWheqYhSDNXQJfMqvNzNq3GBar2Gt8aY1fqZovsFtt16bhfvPXsTixStEnDoiQSy4QgorEryppCckbySpf3pstFtm9i3w5NHCZ4K9eybaWCdN2mKZK2Npv2e3Rj2uchPFWRMWfMyzEcLAPyWjXdF7tPUbrPfp6xK9i3FfpJbKtzaA4VpYx68hWExRe4NiKHteHENTWth8dEqz4GkbJyDaXnJgmRzppw5csdVY8at9rSHqPjq9jXvY7WV1Cfva1rhtrDkFGZ2peoBUGi1U418EVsh5vX8XVHmdf359BU8W3Uk1ChXa8hc67dbz4aMkR6scehz3FxYE3DCUwJ8k9wPxGkrQri4hKUzzoKpCmboeSPYjyiJYrcmSACRifUUEnqVavA38Xe4NSaPxCZeFzwbtEKhLLjdScNosBRZt3kVPPoUWmDVatdzeTtpvTd8KAysju8ruCqD51nU2sUd9yiZbBV3TNRSDsz6BW87nZRewfvPdyx7WQniyyE3Kfww7Q8enAk57KRiSizVaKB3waK68rE76fXzHjCGfkU3UXp9pUsFx41u7BQtpw8VJDWnqzTGyzppntLG4PVh9cQWsGh1dQeapQ5Kx4jFSdGGaePUuXcdfDZ9eXS6SrQEgd9ZKVFdTEAVTeVG2abwcLKoSdF9H8sBtaresTokAUJZfynY5tvnVmCKLaPHT1rBAoAZWennU2XEF6HS2AoHHdCd8JsAfypfpUxTdNGVdQ4JNLNbtPVj6yJw46dYbXjb4HbuKi1bsJjERL3f3HESF5xogqFADA1ApJRsDisSHtqCZhUZXCX8nX7wU9T3hSVp75bnthWZ86TmXPfPkEUnsQryMfGo5sbveJYMP4XUT9TuXphdx34oahDpensuwXvft6BbnfwwSdYjsFdDuRtieUaLad359shRy4KkEyRwPDvVhEc1itqcyWTGzXZs4f7xvJxU3AjzEEjQoBWDYELGhafy2wEoRCSEmMxFPyaumuqyPXiD1usjTXyMsYPRy9pc62c2G5BWB6JEG3z1N82Ps8VH9EcioDrh14EHAuYQ4f82tCqmrx5QWQL8XBiQLofEy8LKDEgKrkYZeFi7nkvnxfezMfVpq7CdGta65opj5C8q43YrN3Gqvu4Bfr97pehzrNxbijLqH31rx1n1aejg5QEiSTT4ajhkbPzZQN6PEbVtHeoaZFw5ZrUdkhV5uage2z3wYKRPTuMmre6dFBevgaH5m9abtTKzM1ZkuTx4nHipV7TnCsRU7ivGgvrfbUcypa8M6FzdjTwvmyjXXnpNivT9waXyxuNMQPgwDt9jFcdP1DkV3utSiE5EGkgUTYgmhDrNwkUpVzFBV5epmee5vqNmbrSfUXvtpv7VWwx9EZq1mK4hxZKTXoMtaAJ7ia87KDKwTcy89gW1iRh1XfA6h9uKdAUz2vhc2xPSxbLEasdWnrZ66GfrQQFfsqzgGb7T7VCzNCMuAFTn9Ziq3qJ9BuBuT8tEnmoFkhitEexeFjaUS9bh53kbnudFK9HzC4KZ8DsLwBUxygnvS7RQjWfSFcv4DJBKVmjN7iBFyCnk6AuY5oXqZSn9JW9yhKyNpBqyxNfDRujNc4jfQku6R8dCZMFcz2EimxQAWV76cFK1HZtRAZcZxoKrLHk9QmgETkXkdcScbQVBkUGa92s5cjUoD5JzEovb612neaZPRK7Z2nCMAeLjUutVqrqrUpY1RprM6DNTvK91hCgGJEiEfeoAnJDrAt474NY6wLp5th4L2J2YA5hBDabjFeWBy8u9ZwxxPyG5vyHKmgLqwkyXeKaCwoEzQjWPFnXmY4eW66bSqXq2Uzgt3v4a1vqmaMNCeUsMsNtG3GhL3tLgtA669E3VcGKk51HLrdE7yu5mPX5NEng9JkydtRBseP3wJyfSFgW9LU5eNo6Dv8W6xt4ZMV8piGPmDvCm9Ue6gQyTTfUwXHjaC3fXPGz7mL1DoxreMqRf8ajqz66iwHibujaW25kR5ENoNvH7tASBQFesXny1oBkwdQkyYFBDE5qJZqt9qd1YFC7g738C7E2HBFfmFvTG8cXUCaeDVdcvzm3eQCVv3b8drauKQeQR3prJDtdt1Diingsg9MhL4TPuEg4T6eu9UeqrVpg7CURNFPhBMEEtdLT4CTjTVzv6oRHw7TqguMKGaSUWyDBrPrbExPq28zCCSdcFwoSm91Az9KDYnYuXdS52ZBMyASifUVoFMqWeEELR2vc4hG1wpvBKT3qHv8gTiCGTtxP6cjkoAJGszM5xbLo62HywyVQu8AKer36QbC1SQkJwGAioHuTjoKKJDMyGrEBtTsTkbH1Btk4pjPBXPMQvjcAQzVPHcRjMWeNVnmmdrx1PP5fU5PKeB8Ww5c75e8dDQrnK6m7Fa4wjaPQMetTgP4ESfGxXgioEbm3mn7e3nwma2rMxW1RqrzyyE8V3ZmHf6qmRQFdpJAvdfWHDwWn5e1t9sn3j292vwmPD27Z2JLQZXUYK7t6LPrjdeqgqf2GRhkYbv8PSM6pKCmGXsXgnabvjhfEH2ep5bD7N92oBWTVxPfBCY983RgcdbFeD3eVaUMXm4xe6jm7jbprEi1ZdjwGJvdLNvrDavHGRnM9ujtmcbiCH3vrkCd348WnGaL6CWjAfwPeEK6PwL3XR3rc1hJ2EPydekHxPtXAUn52WtTf24SqyAuTqBr8AdWdcXDUixd2rnBNDA8DmmDgRCHdqsL5cQdYDiv7RWEtHP6RkXh4A8StsXU6gwJjpK7ESYe1WLHaHiutAwtBEknKecSxywB3ShbQHa3kmY5LXuHwCam8M2P3s3MMeGcKUsadxLqKwRt7JG3Fy9dUuwzaD2gtLkde8VBcqakzcCAgtrsiC5z2Nohtrb1yBNH581TwzTwK3YyyN7Fn1EpHLzzZTWiziAJwwDommXn3VQbW2LgMn2jcuhNtQbnp4mFupHyvMfkfSTUAWLxvWYseacMYPTDK4jfpghukDnGkF589Mfz7sLFcEAsVYLas6kAo3P9DSi7kgthoaKXqtwuiva6YB4CtZYtpcBfvaSYzgvq48nvzMEWKyZCTQEFEe4TRZFyTrEPGygfJVTPCigeQDTbjCXc2DscpDLpfChk9wS5CgYxhyweUJi8T8uqBz5AZkzTj5wPm2Rx1kunfnCJdjXoRYeSpSRKeqh5RQbTHcBgZLKW722pvxEgCyrNKmMLdBjv3d3nmJ3B4Wfjs6Pei8hM2ouMosNnT9Czy9WX2zHpNzYso4JPwhFWDaxMnU1ToWY37dXviptwsLKmmsLujjpwjCp1npRowUJsmuQiVpdqPbPn5ACdBiQEnt4SbeY5933DVP2JpeL2NorUByaMJM4QR6QxSzoKRo1HKHy39wcJdcYFQ3XphebR3c2tHyvjPuzMw9FKkW6jABmBWL9PmRjde7rgFFnThEVKt94n1pKoFjRb1BKcoDqrc4jvKVevu48WVK85AiqBnuhD26zybQtsMFgSTf364B95eoVBk1fSsDkXHkvvquBVZ4yC4tiFd2rXsnBr4R2syTD99wmoh61PpXwN2BAifqMVbhD99WxJtCt2qdthKWhCprqKzJcLPj8KN35MqgboYNPrFCihoS6jyUQRFPzaNBcqkaKrurtMaWTe1LAG1DMvAUiBGjPuHb7rPvuC4jjSNsBJL8TMeC149ni1jn1UriEnZqPrB9tLuHHcP42D7WtztqbyRcwvA3EQRJT9UhbY1zfkg7Wdq9ZwKkb3Wzo4MwFxGu5VUzzDPCSUMAdRny5c5dejFeJrK687kDT6HwidwzYRLgY1CVmSK1VrcUwPxNxQQ58etAQFuw8PiigBTnwQZaiu2z81uyqpUJ9KYhnzHjLC5YwYg14XEmVpQKCs6rW6SxVDD4JqU8GvuAx1Tig73FCwjvR8Miz7K77pUsyVtJ9s9c36qGm8aC2wRTvHP68H1HfQj9z2NVcswfyFd1LoL7wqn16FLqEY1hvaK4kBpWZDV72rmrgZqGDb6ufFQ6rvhk6LfM7W9GVtDciwCWdxTuFHVJQUUHsDWbRq9kxrny42ogTC5R6CXPUo6xLbSEevN6k7N9Zwmc5QY4ZevHcmJYS5ztQ8CDbA3F6b3jZiF1nFPFCCZeAUjhH6ACV9bnvVFX6NYPhEHpw9sznzeTQSiHSUWwqo1VTGsGVuoB42mSXiVhjZ9D4LKMc58AHsxq5EKzwm2hC7zHtPwCcgzYcSBS6mdLXYvPSUYx6jCE6GdRaR989p4own8XRC33YU1kG7m2FVq8gMikVUKH53Xk4u4G5PcZ44rrcRv7qJGmvq3a2e8EhKETdE4stoUs3H8StG834q2R6uLGqHsXMJ4LbB477EKwj62dm5BZsMgLnWv8txz2VUZpSwRosncB5Jp7obwrJ7ihSRWFQjFJirH9LcwwmPwEipSGNAAE18F78pN6kxbUkLpjEqTKh8eu1rvWgqozV35JajWWgodpdFN8nGEFBTx4SJW5R9RZfoo8ScVNAafCG9xDXKxgUGMj82WjsfJYvFyqDTUsYRy49jZomALXyeN4SjrU5yehhqXMvrHEEKdFcmAsYme336yFRdQ2hBQvPSE2b2tnWe3pr9zxHTYQFXjKo7QEF6N62cmMPve9rhvJEWMdcXBDrwDEFySKsJSeMPWzuc9v3rL5qScNpMbp3KGCW6nBBW27E83TSiAUtkY9FKCo3gdbgpTqrS5QSZ531Eqp1KFnaB46C6idScLortbyFquQ6si9FUVJK6GqQWZYFnzh8v8DeYdPE6z9C4Fb2Svuf6Gvh7Lwd853eDChAWUZsQwYjmZka2esqjv5cfprNxm7G7AAVg8DEhiChExkY5eTCm5NVQDiq23jiYcqjMmsFZ3eWEA6tGPi3KVMTkB2ttMVARk72AyRw14Gfb56bDXTbwEnUN2f3zHSNaARfz8mS6SbkRZ7nKtSZsqL5GmjqYL71yrhutwrpgv1rqT4XgqgPJSu6hXpnDo7VXvCjmQkLeMvdjSjBsEgn2BLFKJ3DTTssGbuTWyeS2pDVpv9TCxbeFjYmqndJtVWhKbGoeMCQ1FvijSwjL5kobeoVCBqDvEjEVkHsmTdXiRTysuEvipVQfSzGPXjSx2pKh6M4ejGNjnev18hgvaNYaLMoU84CMpYQ7gzuZXPkhFReNvwMycoMCRoMyyracAzSsp9apni7AVTbs4hBT8L7jBq4Ttce8ewqMtPdzRhrHpip6d1RP5pCQ7DSgYCtAi9kbsiXMCuafHjHmJbSbfdkcgs61svjNTGH9xLjBMxpEpRCPTg28dgTqNMh2UY6vknGNhFzw8hdryGVmkrWtFHhaVMEx25M3egmbLEmm6or6haM4EJvDtUDus5Hgxda5toxz2Mzgi2or7HJAU4Mef3pWWixkpSQcBBDDKwJas6xQkny6Dw52mmyJkiyqVhCWtRwHXw1JSKkdgfEdY68nmTuTYCxMkNcDCXQRyw2SSivfwW3G53dcm4si8rquYAk4Y4Ekq6MaHN8aqv3a6BJ7tNEFVQSyDvJYtnA2Fn9eXtXm1eV97dL7BYgwMyPurvay2YiaTMcUXPHh3xHUePq38M1A4fQXSiBxhi1nb4VDDWbr3FhaTsk2aPJL7ALLrAFcvZWJr1WeCDyH2WNWD3mFcqiykQwauNcUCqrmrsyLVUpFXHicqLh6SMdxLneXcNfAPhi8dKvxrm5UkToSamHbbxZyDQFqm59rzX95VABSurbFe38YfEWgPQAhMFuuCy9yNsAAdp4n9mVjPsxZTfUk6QcAL6qa5SFwj7Xb8frUfiYzLYjWBm6CqUyrbocDFWryiieALKLKuJ4nnHF5Tcd2rWBydd4sRLb6WvNoy36BRdjRkohb5MXLRkccgjVVHFhjqkLiKkF6bNyCRmzbChesKUPPWhD3j2wcbDFfq8UmqxL1dndy5sV1GXN1EPs8QyckywYVKr7u3aBrw8qokLevTGoos1WcvLFiZQEkqrfsjKVzJdq52Tu5x7SdTMHFUUw16TKagZNtLYFNP2ZbqqLbDuBJjM5A8qkaYRQ93iqGJ8T4MkCJPRBqCxbzEG9NjQsfKdwgsVLryXA1MV4PeMANjk94fBKyJuCm9CMUtSoaDGNDs2XcUhQRdeAqhjrpc5FN15AHHGz7t2vySQXu2aYfZ4TwL5X9ZFQfrZgQjGwwqKJC3BTiSD3RdzEbTXYVTQhtKUAaZdbzzXbpipP7qpAetZhuRZbyLchdcvqGPXyHVAhn5YTbVmYqChzsUaK6jhrcnCHV37HyBR2HAQG8BMkwJffcm8uD259JSYMmrKbgvQggXcXdCfh2bu3qHgZvbwsgF9vkjAwWhsJz2BGdRDSRGhtqDc8hjcYRSBMizzFEpQytET4KRUJqHPNhVgfeuDiPPRivH1s1D",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForVEscrowCtrt: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForVSwapCtrt() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"2EREWfvy1baLsqLRsHBYn6c9iwJYeL4DecvXDL41sRFSD73orM4cD1U7ejH9NMKr8Np3FYsB1EUvxSs6o8q11GwbPY5XHC9W2kf1stZu4Bhj3qm67Ma4v9Arcb5gL17J1m6JfYALn8dB1sG1svR52vLh6x3BLcSkfYG6HNth8yLnXbEciJHUxmfkFT1g32Q2RwXAzXMnbKtGKxuhmhYKL3zVcZ4GwsU4XpSLpMhAQS6Z211gzRDwxVKWcKa1zaSaDurEWztnWGmVMainGN8UNJ5zQvxv6Y6pH1cUtKCsnDKE6KX4Py7tpMpm8ckHGfQ6hB3RUWe3YozyFGRaC37FhzawBbMS685zbE7k8hFF8jokPZEGQDtNonfwrRymQbppEFJ9MhsYxn7xXATsF3RTEB99iqUybWcfKfqtbmhpLv7KeMQQctA6NPeMmF2Xv1sZpNup2SEPnVC5jJpWutuvFRp1gBm4QBVZDDjyrvH7UnDkbniE8ohh85fbFjr7sRY8xjdtbTrVdYc8xJ4krdxZkwREuof2Nz3wpKtji56LsjVcQeLwbZh95aFQbqXp3oqQCy351Ejy3cLGh1ftANmbH1Sy7gR3fdi3wWFnAxSHR1GUED432fQTxXWxSezjRJZdCiMT5dHcKFVZcjU8W7skpuq5r7jLSon8naqG4x4ZLLXsFj1E1SZdbyDCBMGq9JaBRWQ3vPxWPgXn3PKXNdaqWyfd4n4Y2h7neXoX9pfUEEF34tJjc45eJiZMcwVSj6D6LxfSQRXa4hj9KevZAdL7WTTgkgjmno57pxuBtB5USGv6xi9UAYvLKTHokhY21fwQGJcYBB3uNsLXrSqFiKZjnFS4ygKeCeCK8e6EP43KXoRTGaJbcjZ3gfoSKFkm8MLy1xxk9agJiBRBraSn5jsPRYU6EwkbsgPvu153TbAzNSRYUUpWVi5xJsNiXcfTAZjCwfP89uLjDxCSXVxqfiYS4W7X2ac6yKn3dxMppxpeNGV3tYRwwTRL8LTmxKm7RLyxVFtmrfoo5bHREEYB9NLovafB8NpMMuxe7trwyDSR6eByYGewyNTQ1d7c2XxWEN1wHM7K6o1dq2pxfL7yEK7xhg59NMohnRRYvb8zQX694DBTEWLRxcyGreB4ZVBF93HqnCnTCKBSU5FcfzPTyVxFM4EXNna8wPvBGxZ8ZVUQWyHUpch528pRXJNSw2sLxwhHD3NyytxBsG3pTXa2Lwzi73Hi5s1ztcbjHYrBFcKzFPF8AMmq58Hz1Hp2VvjAX8uQ31puSiu8ReM6xDo82iiYrrvYJXSbUkY1HADZXYVawAUmKhCqBoZaxqFAUFPoYD5jfUZWkkK4E2kUbdkCFtpvaxNMNsx5KeNvtXT5XwfDTCEsPiTBD2xfUUT6oybzowqvxpLwZ77GBrmrYbQfb7BiGN9eSZ4jp52j5EgmRZeJ7CX3QXXXXE4dfbuqAtek6izRHsPfEjmUhzPZy88qbPR6ptawucGpEHDmvLnzMosfaJcEFCZVfWVSNJi7ZKod8DQM7oZUCWHETEJ8NY2akjAhAgzk6ufg8MR3KukH8Gx14soHJzVyBKn2hgCPvanMBe3V8eTZDA66SZtWxvSmEHQqaKsKgQHSbEiQHtJMcHzun7F8h59AMz3tw6trgWtLcjeKMSNpCvXBiCETiNj8r3jw8xwxomGK5JeZchj9DPKVWJYhhyaqnuNJMGdT57Lrqes6gK7Z24Z1beHn62JUhfKQZZddx9TJd7mf7rWwGPaTzFnZQdqtjHiFpm51K6KMrP8hrBW6bYD7LSc1EKipos3Vyg5Zr73TgvAFGPrc7uvFuMLpBCBSSWLdajMXGDkPibxMcXcHvkAKee695C2AJvBrWTL6b5MEAbnjwuWcpv8K4z6tgFADjouTPn6in6eGQnXounCBnukHQFdT6JwPJV2KGEPbdfskfJkm75QDtBzFTCgAreoXmuXcdrAeFYAn7smrfZNU6qWaAY4pujPUizyBDk7YiLXWfgbtEv6F8MjMG1RMaCP7KyGbXsSXCLvPgLJ48rkc5MpQKqxBZW3EuvgS4F13BNG1FYfrxwAosmYb9ZEEN2mSESnC8LxDJcMi8k8GapkHqnsmDUVYv1kzLBJPHSc7pEnNfw17pvtbbaorMgWUQULeKnoCmz8EtB2fSsgosrqE1FnaByTmPtSFXqTfZLVWXPvMMsGuL9fMGAeoEye3HghyUqLybUSXrUoQxzTZp3CmkJvoAJ2vEqWzRddhqJN2nZBEhMcJ4u5Q7PkTUtxKa12VXnXv4MzMF4yYLfyCwwpcgFjM9fXbd1ELjLJzqoaHsz88xLwn1Ng3jPsGqJaQuftUgmmVAX8niViTdoZcSeQ3mSzpgN7w6k1eh6bofGjvNRXE7xP4yE3umUz3uLin1SZeWqbHkYxSkEBfkzxF5WFML8h9SUqn61rWVFPTm6o7rwtQuEwiWphxJPWbLVFw4C62QNkqJcXKHPimTVnAEMjQCCs8ZvcLzaqEzEXxUK3C5fQqoyRMScXi9tDEzQpTTE23J5NEwhixzu9J2MJXBAZGUkpaEyYLyr4t5kjVx9bvgzQ9SD4ZT58NsTnXTmZFoWJGjgmxmqMPXJ8iEugj425shZWmqhMhc4kUrGtzJgYbrYhfyJwZrrfZDuHJMpGxXDeCYzzkQ2pKxz91xCyfTU1s3Qir9cK9UrRnNm3yEZz2cgh1Q2w68k95oDdQRdVdNYCntbHkBqhPue8KL4QzngueZCYqqc9bE31BhB4Xygmab5MJJcpazKAgCFiAUKdfS5yjGk12Yc2n16SaPUoaWSmRHexKc3Te5g7a3dmkaLApHDuJoGq4o5qMH7uKZsBPWDiRHAGpZ7c6yorya7MFxiJmH4XgSu3EujbHwqpwxvH3sTDw3FMRAMJADh9MB25AvX415J9ufwZutfMoe67mNUd5cWZoLB7BLuVYT1Ya4JbXEE92A7Xk3dF6cVQYfRSQf36NYzE2Z5BufZcNETDNAwApd3zqnshEsHx8PEnebD1Yrtc8juRHinUrDfqw44CjYDYkWiBgpEV34UeBMrdzeQoob14biBzDK79PDn7hKshxhRL72bTSJB3qpHNLRbRtqkg3srE5YcCJJe8RSUHxKz71GkuHURa4746bzuDHBLQvTBPaMugsWqQdGWtox2ANGvCXdAtWkvMXcVfztkGi1g857SZP3xK4U6BNzTvxJVJviHqF2vyL9bzji3pWhWbQw5mXyCPjDBUES8Tj9CTxdbN85MKYyQbEfmn1e2wUgAzRf6TtMcrQfU9S27c8pF1zz23eRxWB59y6vm7pkUSzgCX7TKSkT3PmEmuXVgtGHU3cTaRdgoFNttQfuseYSKmuU3znRjTEuPcku8htqnGePZuYm43hNLCmEFRw5PbgvLbuimCKp8jvrZtpJnQvRNWhqh8gm8ob92qdzyPKaMbgknRYvReTsFg66azobyjngk3ZsuV368fMbrykEn9GDot7CBtbGFcfP4nQM8JHm8RMnVVd63j4gFinXUbho4R1bKQ5s4TfECvzvRcWfcLCuCqQaW5BLMoXc9oxW9WSS1ApWFUfyJwmMX4X9KiVu3WQhUFzAtxqc3r4gFg9rS42PfkiuTBiLBGtAkYDdjBvtErjhN7AXodhmNMoBsJSqphY9caqP2D4ceje8ygLFyhz7SNGcZoKF9apMsMq9nYmyxj3btfDGophMo1k6J9cKHM7HqAnM2Lx9sZa1M9b3LhtBW7vmu1vYLZUebevQSQdXJXgcCXotf2PXLBUkwXvSYfmkV7w3G7NA5aggWhDXu7t9tR5MwzSFCjDPdtmHCN9VeYnEFjofFYxAed2MyP5sJAiTsLoSQ7EsnNcZa6b1ZbcTxZADJoWg2kKkMnaxRJexJ9M8K5DLQB5Kk1sa7Hn1YRNSgoqWdKqLQaqmM5D3SvT7bXBCxmLNpShYr7GpPhyBEHk8p7QkRpBstE7D3k1D8eWj7nEAPF5Vy9fnkftgjYQQTgpZbnDs1n7JjbGAXBpwBxB1P9pwv2z2duyEWqbbwgjcd5ujMCuNRQbx7WeQvFVrqw7Zp6xgTgBE4hYWgZVhsYgoYyLnBL4HZEoccbYsNHHYRs1J5eVJwLQj8748FFQdcMk9YUHv6FQUj4Wy2baEivHWi3WGpEdi63oSNEeWkbEi7FBqRiATiik3AsRfhjE5oanx4eTrv4Y1nvGrEwQEvbADAoKuRP74KH1qtt4c35MtiDP9RWWpoPaZg58PqfghLXb5XQfvemxh9SWkXh2AnjGu1XJz2Qvp36rbu6hBNkMH5v6fppd7QwpTVcBA4VbKYNq8U9dNPKiKKdtXsLAheTvf77D1K54FAEZmNas9Y85jEC7jWt8vCgBAeGWggbsgzCu1cDmW6YiEqbyctirrLLLr4FvB1Rp8tudJ9pjsFV6e94n1woBGkR8EypWuEeDdhw4wFYHiNbjX7N3sCLfPAVMvqP6sRr9MttJzacxdB8WKntpjzerCqTquoPR3aRg6iNLABNYqHxiPQEuekt2yThbpk9BxYn1M3vtyfB1TCedT8vHBu75kTMaiPk3udg99YxfAE3DACdqwuxj8zkF8Y54kfenem3s979ZDmTrtjmnicCib3SknWNC9v44oSiDQRFc2ofaRDdF6FLacWFxcrhnAWuqqiuwXyTwCvuvEEM7SnbYBcEsvs83BGug5Cw9ceifTDf4FcfHrVoaNVo9KEHdLHZeJGAQALZBHoaT4mPzjeUuzanzrED68yDQtTsRXZCnGzbcaRMVc1GeXLmXLpnMgiLEp9szDyNqT7jMS7nod1v58C8dp4C7uXsUnB3z2iVAqggL8R6oUUfxrzunhN5DcmhCtUik3hwWjJYN64RAMsvki5gVNdz7YDkFyD1L1fXhdTqrbndnT428Ya69o5DA4Rakdg92xhz3QiM3SMh8wNwEBdZfSEtxZ3ESLUhZwMQiArUtR7D6t8Yp5ShXFyXiHSRWWbWqEQHh9KR5RpXJf9P8NqckUdX1rXX1ZvukYBd8M65RooMfZAjckseZTWJgZrU78iNeawUgfjmFsXhXUQF2yVMG64bnb9EkCemScA18LuvvmiAZbtW36dbvVMy91JieCtXpfpRiVwh2VaGv89ZzSd9To3YwkNGjNsMqhNBgovNyTkW9FDPHKmT5vfNmB7GXXhkX7n4m8yUcX6scoVUx3wgMcrx2Jm49woFBvfswTah4oaUwfcw3Xgta3Uyr4NGaoPXmPkPUAEFSqgoWFWSQaPFdb5uRgVKAkSFmPM4LABmZuXhztLMdmzBVVk7rLFLQgp2cFgkcqMrhBSbyi7EFiMEEGBT7b2vtRdR555o3AwTeJrnBxK6X269xyLW85Qb3aq2anV1REicPcgirshmtvtpAEfU6ixC8Q4k2NsP6LkiuszwJeeU8LmJRsZeni12pMmSR2ndDFrBFBfnhErTrJvHriETatn5Hq8hyRkKwVguVNYeG3DDtNNrDChDH5KoLx7XEgiSe6rifVrCDfJreszfpAdTd79rx9o26J1fAPRxfdAE2zUnbJKzHgioLKNWWQDDBo9UngA6mcxcYi7A7BSDyChwei1uFVvGCEW6raZGjyHzsEF5RwBaKeZMbJqE93KcxekhPnS1LREyDw15XpCQ6gvevowyUZ64iRHEa1is5HPtxY38qcXXp5ifFgtFmQx1LHFkMo4jQmQADNt8caPyg2zyqj6Sy96w73oMKig84QhE26qZKUgEEeAiuVDVQUemBFz5cgbNHsHGPcDGbX6pDbL9b6XASWLZyNrgLPHviW85Eso4tDEuBab9uxSuvEHW1hYXMNYwPH47Ma6qYgewSaPT5S91NXHdqW7PYKtr3TNnDjHhBitKJGaRHs3u7JPk81DfYyszeqc2R1znHpcJs7v4vF4PVU9FmHtrSNr42Tbx7L1kTy8NHWzGJasMw2pVuZyLWj2BdzyfTaJUcrWQSsRpcv2UoJRMzfxpLbFSie99MyzsLsCGYbtiaB95sFALkRfS4gpitHwHLceiD6FfCfFYdAiM8iGh5TNin7RSkMZxynWfqnAWHenibowQy9E1TdnVYsXjK4hZ49L3xfZEd7TadQQib9FvyM1APBrKJ8dRe4JizRWQtGvxtxEpxAAuqeDCgv11qUgtCQ6csdLVjzGaLKG34KYZS8P1iJ6vnmC9y3r36e8mbnJrxXevEbkbViPAjDMd5RjHR3f54YcoiwiQTTEs5py5rf217v4j9vXSeiaQb8gdj855CHp6AgwKLmLgbV7jP1gTb6N8PZdvnfWwnVbWrwLTKjeAcJ2WzMfPwXXAA24hkmvX2RQSAZYXMgmA5WtdNcZLFyRLvj3eCZwzgTahBeyjHcpGwLWHtMY3m9TSGg1N4CQsQJEfgKdXKYSZPiD2kQ12NQp2z2Qqg4Bz8aCGN8Qy7Sva218MxYFGTHUEFS9obsLo3ijF2hMH56wuHUqNhQrhPFefmaJgEF4KGJf8dWnGJTxSefEAbvS3CaVFRxHYdJvmZYfBgiXRPT3VbxZbWaWy118fhaFc97fHF1ZJ1WxUZUgUDzHLY8H8BTxLsnjB65cvZsUNk5yHP6rK8aZCQT7dAHvsbyGkZMmaDJscMCK8uRsKMtMr3q3boaSfNreNxunmAxAhJe3KW5x22AoG86QDBQvj2bbKTnXGN7Wd7ZmRncKnUsaknwiks6vypEJHVBRSHAppujCVJ7DWY5FnhBV3V421wKCWaHRsDLrVc3DCDhttcApQ89hjnEYhKYuSuBZQxzTAPzgpCPKAxTBfhK5bK2XjbjvrtkxzbkkKbe4rw253PFpmb4LPTPK5VTZEJyWyjX1XfZPxReuDqNst9oBtRmkcaAkhEuwUh9y6Vp8dYgdUwcwBtNchkYVNHBWTXSqqKu2tECfciZoDQGSXGj43wjTRpgQq1JWgrFUZmGpb5Y1q54JgP73CdZQ7hdy7RqipsExUTvxWyrdozaPWeD2xX6WA8Bb1KgUQgLmHTU8izRxpMQdymAumWfFU5ToW98uBQSPJ7o3cAXnMmGe6sWjEceYmkemZCAEvp7JxqVKuPYZrw1NcC855wckC8BLEQzjPsP1Yjcdf9VTtNrto6dFt3yX3xRs2jXKLxQqVWwhMqVAyo2RRBmEMD4ub3VJcj5rsZDzQnHTmvYjfRLTLV8ZvmuQFA9ckadQ8m5FkWR34HKPN7GyH4QAG81ttp5urkeH3EmZ2jU2aXFCPkNSGEvQtH97sjz6hLGjZF72mMLQb6q3ERv1HY37Xe1nNcuimYyWG6Vaj9PKQ6CjuynMdvQ1yaZtremYjWeRy8aHZ4ywcAb8Dt6HzA5VZBKaiki3yqPWJFKUhSoQfd1Roqb3TZdQj1esXLh8hnuFjjeWVp5sriu85tfP59SywWzLo5ytHcGmaEPcsY333uQJqZsTwM2bJP71VWmJ2vTUoQdHq98J6qqSH9rPcw1WCGhaKVdtnBKpH6qYjf8YtWArdQ7q9QEBKLQbKTdoQzyp59vZTUEBu3fqJwQuoBZk62rFfG9KcXX1J6vvKMgsm3w8Y6sD1P46fDKFBeimV524q1vfMwFbBZQg9RTECVHoGjeQj3SrRVch5MdAdZ7qjD4mQC5YBnUWWy5GDHJjzj1VJVmxojrFc4xbVRwUDeLk1Z6KRBYER3dto9nc5q15xkXniDy6M2m1jqBtqc5JuEhNjWKb4GA5fyQwG7SPCmh8bXikjN4EYXsizrSQV9ZLJx7914zwvXefxq3aLY2bySypsvKSVdEcG9n3aP1bGGaCo3MnCnKFs3jmjLodHCoAv2EG3Rnkhddb6mkvEGR7kpMreGjRUjAdPHit3Pa74i7q1WUNSCP2th7F3UjHQZhYTsHUxhSAgTXTieJid3N3ssCdfPs2g4S5s4iPcfqosBUiBiPtbQ6vX52kYEhEehJte4F9rjm3zu2MvXJwehueqrctjXHh5RiqAfP8zt7pG2yGDAp17xBC1BqMCk2giq3L3dNWjj6E682MEJMHnUv1buRsdptLqJE2CL6YaiwcQ41WSiipkzQEEYjb1UZhpbpeCnAxVsFVEMTQKEi8CtVwan7nhuZzaSbrpFZZZ6yZFNbSpD6PcqUu5vsybvrmZUQsAKeQdo1wgPCKG2EYSk6LnLaapkBAm8sWbMWdv7iEngRqere6W3nLYwKxuuYTaF23Uc41szDsH3BdMsbDmc3Yxjf4knsV6zmNegvjK7anrXCr9Qsr8fvPXn3SMiNiQgQqTPLmYgfViUEChMy3xSqZXvLHRaCEJzg3xvstvjRtiZ66P4ZhvBQ6LrSAcmAkm7iY1jjSwL5tJygGgrpqj2yKcAXcYyy7hSVqrwoViKtccKWT98Jbffw4w3WpUiBjBnEQporLriYvhKN1eDaQpz8CjKVxXRhPvaLCq48EZyAcmyrD39cCKhzUVgSD5wnwKS3NnY1ZtoRp5qR1tAFpFtMTpgKMSKX8NLLgPo4c2U5aMAvcQ6JP9BCYEHuxqBsF7aeQB49hQ23aJmj5covNHA8peRpmWZ3TBWE46vKL2XNg6c94XNZv3PZDyeo46BqnkFY8PHGgTrxdp4JZ8pj5DX6TSMdxMVbG6LNWravLUGqBdveHyHVxb7Uhba4MbHY6tx3Kiuf3AQbXUjRcoChCzWVkxATRRcnzQSw9DLR94BBmZH4eCWcuhSnMJoyALc1FQDGgY8S9PKcK2j5L6JqeJKsv2cPEicJieVYvJbhgQvcB2WzwN1ZPQGtMYZJFVV8SpjHEYxQEtBYyuBaHb7aAgUqbwDqjshvWRaZimSJwzXBwasma2HMc3iaB6RqZVF62fZF8Zz7t1N9CZtWqizTSbY5eKq7UJRsP1aQ8Yhj5WZcj7SubS",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForVSwapCtrt: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForVOptionCtrt() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"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",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForVOptionCtrt: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForPayChanCtrt() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"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",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForPayChan: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForVStableSwapCtrt() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"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",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForVStableSwapCtrt: %w", err)
	}
	return cm, nil
}

func NewCtrtMetaForLockCtrt() (*CtrtMeta, error) {
	cm, err := NewCtrtMetaFromB58Str(
		"4Qgfi31k6qfLxTguJg8AeYzmmgaCTJCEPQyAdoRUUSrFDc91PhkdU6C8QQSsNCFc2xEud2XnuQ4YNJ51HgdNtBdnxZcU5Rnqdzyop41Ck81v4nRKkHpTdTrfD8vTur2w4mTFeTFKVzGvGjpHXUVvT47vZiKLBHSB7FHHpGf69bu8DQGXWu6xnZZkn9v2Rfc9mByhwVLSNghNdRhrQwRWPFJ9Qt7Yb8N8WdmcUCAC6PrC3Ha3Z9w7dyf6CsKcCMS6JmB2gvNQitm9jqAfjRxDdqPBUR6TtyjSdmHP9BZRGgiVCaQH7X8fbJZVWSib4RXvFoSrqY4SfVftDY3PU4hXASaRWbaheB8m4VgM4mA8nKDbZvRWZtZ4cHdWeNFyVPs6HxHQZHrQ3GZGNPjmBSyAkGRFS7i5dK8aYWQDEYu1Xijk63UFAWuf6tRdR44ZgRjWGUZJtdQBDFB38XaU8LSFEj2eaC1yNqZ6nnGeRXDzS1q3YKsGyJTqaDDMHvPHiHonGn76JQHAZN7eGU7biaSLxoikW4MaTPSfmcTmDyPGJyJNHjc8MrpV8aQSaGGyDkf1a9MpoJcyEjsPFQbxYzSJVqFEFg2oUL7Z8VUtJK2kYcWDz7w8UiiQqe3uuQnKDGb1nJ5Ad3W8ZPfVP6YHbJrnBKZXMMypNoveokVvxZMCkSNYDsoBxJzrwFvm5DcDJbePQU6VbeZ5SzQw9XTAw4DZpxkQm9RwRE9PXPqogpp9P6LhaiUa6ZD1cWUAHypjWLJ2Rds96oap3biBp5aESunuh99HByoXg5Aa7EQ3FrEvmeq9TLVFYpJraZyW",
	)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtMetaForLockCtrt: %w", err)
	}
	return cm, nil
}

func (c *CtrtMeta) Serialize() Bytes {
	size := CTRT_META_LANG_CODE_BYTE_LEN +
		CTRT_META_LANG_VER_BYTES_LEN +
		c.Triggers.Size() +
		c.Descriptors.Size() +
		c.StateVars.Size() +
		c.Textual.Size()

	if c.LangVer != 1 {
		size += c.StateMap.Size()
	}

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
