package vsys

import (
	"fmt"
	"sync"

	"github.com/imroc/req/v3"
)

// NodeAPI is the go client for VSYS node API.
type NodeAPI struct {
	*req.Client
	once sync.Once
}

// NewNodeAPI creates a NodeAPI object.
func NewNodeAPI(host string) *NodeAPI {
	return &NodeAPI{
		Client: req.C().
			SetBaseURL(host).
			SetCommonHeader("Content-Type", "application/json"),
	}
}

// SetAPIKey sets the API key used for requests.
func (na *NodeAPI) SetAPIKey(key string) {
	na.once.Do(func() {
		na.SetCommonHeader("api_key", key)
	})
}

// Get sends a simple GET HTTP request to the endpoint.
func (na *NodeAPI) Get(edpt string) (*req.Response, error) {
	resp, err := na.R().Get(edpt)
	if err != nil {
		return nil, fmt.Errorf("Get: %w", err)
	}

	return resp, nil
}

// Post sends a simple POST HTTP request to the endpoint.
func (na *NodeAPI) Post(edpt, body string) (*req.Response, error) {
	resp, err := na.R().SetBody(body).Post(edpt)
	if err != nil {
		return nil, fmt.Errorf("Post: %w", err)
	}

	return resp, nil
}

func (na *NodeAPI) String() string {
	return fmt.Sprintf("%T(%s)", na, na.BaseURL)
}

type TxProof struct {
	ProofType Str `json:"proofType"`
	PubKey    Str `json:"publicKey"`
	Addr      Str `json:"address"`
	Signature Str `json:"signature"`
}

type TxBasic struct {
	Type      TxType        `json:"type"`
	Id        Str           `json:"id"`
	Fee       VSYS          `json:"fee"`
	FeeScale  VSYS          `json:"feeScale"`
	Timestamp VSYSTimestamp `json:"timestamp"`
	Proofs    []TxProof     `json:"proofs"`
}

type CtrtMetaResp struct {
	LangCode       CtrtMetaLangCode `json:"languageCode"`
	LangVer        CtrtMetaLangVer  `json:"languageVersion"`
	Triggers       []Str            `json:"triggers"`
	Descriptors    []Str            `json:"descriptors"`
	StateVariables []Str            `json:"stateVariables"`
	Textual        struct {
		Triggers       Str `json:"triggers"`
		Descriptors    Str `json:"descriptors"`
		StateVariables Str `json:"stateVariables"`
	} `json:"textual"`
}
