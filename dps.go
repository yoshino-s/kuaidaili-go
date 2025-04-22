package kuaidailigo

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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

type DPSProxy struct {
	Location string
	Carrier  string
	Expire   time.Time
	Url      *url.URL
}

func (c *DPSOrderClient) GetDPS(ctx context.Context, num int, protocol ProxyProtocol, extra map[string]string) ([]DPSProxy, error) {
	var res struct {
		List []string `json:"proxy_list"`
	}
	params := map[string]string{
		"num":       fmt.Sprint(num),
		"pt":        fmt.Sprint(protocol),
		"f_loc":     "1",
		"f_et":      "1",
		"f_carrier": "1",
		"f_auth":    "1",
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

	result := make([]DPSProxy, len(res.List))
	for i, item := range res.List {
		parts := strings.Split(item, ",")

		u, err := url.Parse(fmt.Sprintf("%s://%s@%s", protocolStr, strings.TrimSpace(parts[1]), strings.TrimSpace(parts[0])))
		if err != nil {
			return nil, err
		}
		t, _ := strconv.Atoi(strings.TrimSpace(parts[3]))
		result[i] = DPSProxy{
			Location: strings.TrimSpace(parts[2]),
			Carrier:  strings.TrimSpace(parts[4]),
			Expire:   time.Now().Add(time.Duration(t) * time.Second),
			Url:      u,
		}
	}

	return result, nil
}
