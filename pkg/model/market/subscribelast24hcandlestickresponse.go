package market

import (
	"github.com/model-collapse/huobi_golang/pkg/model/base"
)

type SubscribeLast24hCandlestickResponse struct {
	base.WebSocketResponseBase
	Data *Candlestick
	Tick *Candlestick
}
