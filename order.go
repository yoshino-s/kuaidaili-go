package kuaidailigo

import (
	"context"
	"strings"

	"github.com/go-viper/mapstructure/v2"
)

const (
	EndpointDevGetIPWhitelist     Endpoint = "dev.kdlapi.com/api/getipwhitelist"
	EndpointDevCloseOrder         Endpoint = "dev.kdlapi.com/api/closeorder"
	EndpointDevSetIPWhitelist     Endpoint = "dev.kdlapi.com/api/setipwhitelist"
	EndpointDevAddWhiteIP         Endpoint = "dev.kdlapi.com/api/addwhiteip"
	EndpointDevDelWhiteIP         Endpoint = "dev.kdlapi.com/api/delwhiteip"
	EndpointGetProxyAuthorization Endpoint = "dev.kdlapi.com/api/getproxyauthorization"
	EndpointGetOrderInfo          Endpoint = "dev.kdlapi.com/api/getorderinfo"
	EndpointGetSecretToken        Endpoint = "dev.kdlapi.com/api/getsecrettoken"
)

type OrderClient struct {
	*BaseClient
}

func NewOrderClient(secretID, secretKey string) *OrderClient {
	return &OrderClient{
		BaseClient: newClient(secretID, secretKey),
	}
}

func (c *OrderClient) GetProxyAuthorization(ctx context.Context) (string, string, error) {
	var res struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.getBaseRes(ctx, "POST", EndpointGetProxyAuthorization, map[string]string{
		"plaintext": "1",
	}, &res)
	if err != nil {
		return "", "", err
	}
	return res.Username, res.Password, nil
}

type OrderInfo struct {
	OrderID           string         `json:"order_id"`
	PayType           PayType        `json:"pay_type"`
	Product           ProductType    `json:"product"`
	Status            OrderStatus    `json:"status"`
	ExpireTime        Time           `json:"expire_time"`
	IsAutoRenew       bool           `json:"is_auto_renew"`
	AutoRenewTimeType TimeType       `json:"auto_renew_time_type"`
	FreezeAmount      float64        `json:"freeze_amount"`
	UnitPrice         float64        `json:"unit_price"`
	Extra             map[string]any `json:"extra"`
}

func (c *OrderClient) Info(ctx context.Context) (*OrderInfo, error) {
	orderInfo := OrderInfo{}
	r := map[string]any{}
	err := c.getBaseRes(ctx, "GET", EndpointGetOrderInfo, nil, &r)
	mapstructure.Decode(r, &orderInfo)
	orderInfo.Extra = r

	if err != nil {
		return nil, err
	}
	return &orderInfo, nil
}

func (c *OrderClient) Close(ctx context.Context) error {
	err := c.getBaseRes(ctx, "GET", EndpointDevCloseOrder, nil, nil)
	return err
}

func (c *OrderClient) GetIPWhitelist(ctx context.Context) ([]string, error) {
	var res struct {
		IPWhitelist []string `json:"ipwhitelist"`
	}
	err := c.getBaseRes(ctx, "GET", EndpointDevGetIPWhitelist, nil, &res)
	if err != nil {
		return nil, err
	}
	return res.IPWhitelist, nil
}

func (c *OrderClient) SetIPWhitelist(ctx context.Context, ipWhitelist []string) error {
	err := c.getBaseRes(ctx, "POST", EndpointDevSetIPWhitelist, map[string]string{
		"iplist": strings.Join(ipWhitelist, ","),
	}, nil)
	return err
}

func (c *OrderClient) AddIPWhitelist(ctx context.Context, ipWhitelist []string) error {
	err := c.getBaseRes(ctx, "POST", EndpointDevAddWhiteIP, map[string]string{
		"iplist": strings.Join(ipWhitelist, ","),
	}, nil)
	return err
}

func (c *OrderClient) DelIPWhitelist(ctx context.Context, ipWhitelist []string) error {
	err := c.getBaseRes(ctx, "POST", EndpointDevDelWhiteIP, map[string]string{
		"iplist": strings.Join(ipWhitelist, ","),
	}, nil)
	return err
}
