package vsys

import "fmt"

// FastHashResp is the response of POST /utils/hash/fast
type FastHashResp struct {
	Msg  Str `json:"message"`
	Hash Str `json:"hash"`
}

// FastHash gets the FastCryptographicHash of the given message.
func (na *NodeAPI) FastHash(msg string) (*FastHashResp, error) {
	res := &FastHashResp{}
	resp, err := na.R().
		SetBody(msg).
		SetResult(res).
		Post("/utils/hash/fast")

	if err != nil {
		return nil, fmt.Errorf("FastHash: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("FastHash: %s", resp.String())
	}
	return res, nil
}
