package kuaidailigo

import (
	"context"
	"strconv"
)

type AccountClient struct {
	*BaseClient
}

func NewAccountClient(secretID, secretKey string) *AccountClient {
	return &AccountClient{
		BaseClient: newClient(secretID, secretKey),
	}
}

func (client *AccountClient) GetAccountBalance(ctx context.Context) (float64, error) {
	var res struct {
		Balance string `json:"balance"`
	}
	err := client.getBaseRes(ctx, "GET", GetAccountBalance, nil, &res)
	if err != nil {
		return -1, err
	}
	balance, err := strconv.ParseFloat(res.Balance, 64)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (client *AccountClient) GetOrderClient(ctx context.Context, orderId string) (*OrderClient, error) {
	var res struct {
		SecretID  string `json:"secret_id"`
		SecretKey string `json:"secret_key"`
	}
	err := client.getBaseRes(ctx, "GET", GetOrderSecret, map[string]string{
		"orderid": orderId,
	}, &res)
	if err != nil {
		return nil, err
	}

	return NewOrderClient(res.SecretID, res.SecretKey), nil
}

type GetAccountOrdersRequest struct {
	PayType PayType     `json:"pay_type,omitempty"`
	Product ProductType `json:"product,omitempty"`
	Status  OrderStatus `json:"status,omitempty"`
}

type GetAccountOrdersResponse []struct {
	OrderID    string      `json:"orderid"`
	PayType    PayType     `json:"pay_type"`
	Product    ProductType `json:"product"`
	Status     OrderStatus `json:"status"`
	ExpireTime Time        `json:"expire_time"`
}

func (client *AccountClient) GetAccountOrders(ctx context.Context, req GetAccountOrdersRequest) (GetAccountOrdersResponse, error) {
	var res GetAccountOrdersResponse
	err := client.getBaseRes(ctx, "GET", GetAccountOrders, toParams(req), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type isCreateOrderParams interface {
	Type() ProductType
	isCreateOrderParams()
}

type CreateOrderRequest struct {
	IsNotify          bool                `json:"is_notify,omitempty"`
	IsCoupon          bool                `json:"is_coupon,omitempty"`
	CreateOrderParams isCreateOrderParams `json:"-"`
	PayType           PayType             `json:"pay_type,omitempty"`
	Prepay            struct {
		TimeType      TimeType `json:"time_type,omitempty"`
		TimeNumber    int      `json:"time_number,omitempty"`
		IsAutoRenew   bool     `json:"is_auto_renew,omitempty"`
		AutoRenewUnit TimeType `json:"auto_renew_unit,omitempty"`
	}
}

func (client *AccountClient) CreateOrder(ctx context.Context, req CreateOrderRequest) (*OrderClient, error) {
	var res struct {
		OrderID string `json:"orderid"`
	}
	params := toParams(req)
	params["product"] = string(req.CreateOrderParams.Type())
	err := client.getBaseRes(ctx, "GET", CreateOrder, params, &res)
	if err != nil {
		return nil, err
	}
	return client.GetOrderClient(ctx, res.OrderID)
}
