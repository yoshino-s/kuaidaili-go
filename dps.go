package kuaidailigo

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	EndpointDPSGetIPBalance    Endpoint = "dps.kdlapi.com/api/getipbalance"
	EndpointDPSGetDPSValidTime Endpoint = "dps.kdlapi.com/api/getdpsvalidtime"
	EndpointDPSCheckDPSValid   Endpoint = "dps.kdlapi.com/api/checkdpsvalid"
	EndpointDPSGetDPS          Endpoint = "dps.kdlapi.com/api/getdps"
)

type DPSOrderClient struct {
	*OrderClient
}

func (c *DPSOrderClient) IPBalance(ctx context.Context) (int, error) {
	var res struct {
		Balance int `json:"balance"`
	}
	err := c.getBaseRes(ctx, "GET", EndpointDPSGetIPBalance, nil, &res)
	if err != nil {
		return 0, err
	}
	return res.Balance, nil
}

func (c *DPSOrderClient) ValidTime(ctx context.Context, ipToValid []string) (map[string]time.Duration, error) {
	res := map[string]time.Duration{}
	err := c.getBaseRes(ctx, "GET", EndpointDPSGetDPSValidTime, map[string]string{
		"proxy": strings.Join(ipToValid, ","),
	}, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *DPSOrderClient) CheckValid(ctx context.Context, ipToCheck []string) (map[string]bool, error) {
	res := map[string]bool{}
	err := c.getBaseRes(ctx, "GET", EndpointDPSCheckDPSValid, map[string]string{
		"proxy": strings.Join(ipToCheck, ","),
	}, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *DPSOrderClient) GetDPS(ctx context.Context, num int, protocol ProxyProtocol, extra map[string]string) ([]*url.URL, error) {
	var res struct {
		List []string `json:"proxy_list"`
	}
	params := map[string]string{
		"num": fmt.Sprint(num),
		"pt":  fmt.Sprint(protocol),
	}
	for k, v := range extra {
		params[k] = v
	}
	err := c.getBaseRes(ctx, "POST", EndpointDPSGetDPS, params, &res)
	if err != nil {
		return nil, err
	}

	protocolStr := "http"
	if protocol == ProxyProtocolSocks {
		protocolStr = "socks5"
	}

	result := make([]*url.URL, len(res.List))
	for i, proxy := range res.List {
		u, err := url.Parse(fmt.Sprintf("%s://%s", protocolStr, proxy))
		if err != nil {
			return nil, err
		}
		result[i] = u
	}

	return result, nil
}
