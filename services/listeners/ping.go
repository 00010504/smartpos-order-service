package listeners

import (
	"context"
	"genproto/common"
)

func (c *orderService) Ping(ctx context.Context, message *common.PingPong) (*common.PingPong, error) {
	return &common.PingPong{}, nil
}
