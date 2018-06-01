package service

import (
	lb "git.shopex.cn/luban/go-sdk"
	"golang.org/x/net/context"
)

type Dealer struct {
}

func (c *Dealer) Deal(ctx context.Context, in *lb.DealRequest) (t *lb.Trade, err error) {

	return
}

func (c *Dealer) TradeInfo(ctx context.Context, in *lb.TradeInfoRequest) (lt *lb.Trade, err error) {

	return
}

func (c *Dealer) TradeList(ctx context.Context, in *lb.SearchRequest) (out *lb.TradeListResponse, err error) {

	return
}
func (c *Dealer) TradePay(ctx context.Context, in *lb.TradePayRequest) (out *lb.TradePayResponse, err error) {

	return
}

func (c *Dealer) TradeCancel(ctx context.Context, in *lb.TradeCancelRequest) (out *lb.TradeCancelResponse, err error) {

	return
}

func (c *Dealer) TradeConfirm(ctx context.Context, in *lb.TradeConfirmRequest) (out *lb.TradeCancelResponse, err error) {

	return
}
func (c *Dealer) TradeStatusConfirm(ctx context.Context, in *lb.TradeUpdateRequest) (out *lb.TradeUpdateResponse, err error) {

	return
}

func (c *Dealer) TradeShip(ctx context.Context, in *lb.TradeShipRequest) (out *lb.TradeCancelResponse, err error) {

	return
}

func (c *Dealer) TradeAfterSale(ctx context.Context, in *lb.TradeAfterSaleRequest) (out *lb.TradeCancelResponse, err error) {

	return
}
