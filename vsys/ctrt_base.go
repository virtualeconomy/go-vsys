package vsys

import "fmt"

type Ctrt struct {
	CtrtId *CtrtId
	Chain  *Chain
}

func (c *Ctrt) String() string {
	return fmt.Sprintf("%T(%+v)", c, *c)
}

func (c *Ctrt) QueryDBKey(dbKey Bytes) (*CtrtDataResp, error) {
	resp, err := c.Chain.NodeAPI.GetCtrtData(
		c.CtrtId.B58Str().Str(),
		dbKey.B58Str().Str(),
	)
	if err != nil {
		return nil, fmt.Errorf("QueryDBKey: %w", err)
	}
	return resp, nil
}
