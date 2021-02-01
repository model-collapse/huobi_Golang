package market

import (
	"github.com/model-collapse/huobi_golang/pkg/model/base"
)

type SubscribeDepthResponse struct {
	base.WebSocketResponseBase
	Data *Depth
	Tick *Depth
}
