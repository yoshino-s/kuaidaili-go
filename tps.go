package kuaidailigo

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type IPPoolType int

const (
	IPPoolDefault IPPoolType = 1
	IPPoolPro     IPPoolType = 2
)

type CreateTPSParams struct {
	// 换IP周期(秒)
	// 有效取值：
	// • 0: 每次请求换IP
	// • 15: 15秒
	// • 30: 30秒
	// • 60: 1分钟
	// • 300：5分钟
	// • 600：10分钟
	// • 900：15分钟
	// • 1800：30分钟
	// • 3600：1小时
	// • 21600：6小时
	// • 43200：12小时
	// • 86400：24小时
	Period int `json:"period,omitempty"`
	// 并发请求数
	// 取值范围：5~1000次/s (默认值: 5)
	// 注：tps.period >= 60 默认无限并发
	MaxRps int `json:"max_rps,omitempty"`
	// 带宽规格
	// 0 <=tps.period <= 1800，取值范围：3~300Mb/s (默认值: 3)
	// tps.period = 3600，取值范围：3~15 (默认值: 3)
	// tps.period >= 21600，取值固定为3
	// 注：购买时长小于1个月仅支持200M以下带宽
	MaxBandwidth  int `json:"max_bandwidth,omitempty"`
	BandwidthType int `json:"bandwidth_type,omitempty"`
	// IP资源池
	// 1: 企业IP池（默认）
	// 2: 星辰IP池
	// 注：tps.period>=300时 tps.ip_pool仅可为1
	IPPool IPPoolType `json:"ip_pool,omitempty"`
}

func (CreateTPSParams) Type() ProductType {
	return ProductTypeTPS
}

func (CreateTPSParams) isCreateOrderParams() {}

type TPSOrderClient struct {
	*OrderClient
}

func (c *TPSOrderClient) CurrentIP(ctx context.Context) (string, []string, error) {
	var res struct {
		CurrentIP     string   `json:"current_ip"`
		CurrentIPList []string `json:"current_ip_list"`
	}
	err := c.getBaseRes(ctx, "GET", TpsCurrentIP, nil, &res)
	if err != nil {
		return "", nil, err
	}
	return res.CurrentIP, res.CurrentIPList, nil
}

func (c *TPSOrderClient) ChangeIP(ctx context.Context, changeIP []string) (string, []string, error) {
	var res struct {
		NewIP     string   `json:"new_ip"`
		NewIPList []string `json:"new_ip_list"`
	}
	err := c.getBaseRes(ctx, "GET", ChangeTpsIP, map[string]string{
		"changeip": strings.Join(changeIP, ","),
	}, &res)
	if err != nil {
		return "", nil, err
	}
	return res.NewIP, res.NewIPList, nil
}

func (c *TPSOrderClient) GetProxy(ctx context.Context, protocol ProxyProtocol) (*url.URL, error) {
	var res struct {
		ProxyList []string `json:"proxy_list"`
	}
	err := c.getBaseRes(ctx, "GET", GetTpsIp, map[string]string{
		"num":    "1",
		"format": "json",
		"pt":     fmt.Sprintf("%d", protocol),
		"sep":    "1",
	}, &res)
	if err != nil {
		return nil, err
	}
	protocolStr := "http"
	if protocol == ProxyProtocolSocks {
		protocolStr = "socks5"
	}

	u, err := url.Parse(fmt.Sprintf("%s://%s", protocolStr, res.ProxyList[0]))
	if err != nil {
		return nil, err
	}
	return u, nil
}
