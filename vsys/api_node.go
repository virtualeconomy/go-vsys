package vsys

import (
	"fmt"
	"time"
)

// NodeStatusResp represents response from GET /node/status
type NodeStatusResp struct {
	BlockchainHeight Height        `json:"blockchainHeight"`
	StateHeight      Height        `json:"stateHeight"`
	UpdatedTimestamp VSYSTimestamp `json:"updatedTimestamp"`
	UpdatedDate      time.Time     `json:"updatedDate"`
}

// GetNodeStatus gets the status of the node.
func (na *NodeAPI) GetNodeStatus() (*NodeStatusResp, error) {
	res := &NodeStatusResp{}
	resp, err := na.R().SetResult(res).Get("/node/status")
	if err != nil {
		return nil, fmt.Errorf("GetNodeStatus: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetNodeStatus: %s", resp.String())
	}
	return res, nil
}

func (n *NodeStatusResp) String() string {
	return fmt.Sprintf("%T(%+v)", n, *n)
}

// NodeVersionResp represents response from GET /node/version
type NodeVersionResp struct {
	Version Str `json:"version"`
}

// GetNodeVersion gets the version of the node.
func (na *NodeAPI) GetNodeVersion() (*NodeVersionResp, error) {
	res := &NodeVersionResp{}
	resp, err := na.R().SetResult(res).Get("/node/version")
	if err != nil {
		return nil, fmt.Errorf("GetNodeVersion: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetNodeVersion: %s", resp.String())
	}
	return res, nil
}
