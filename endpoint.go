package kuaidailigo

import "net/url"

// Endpoint 接口地址
type Endpoint string

const (
	// GetOrderExpireTime 获取订单过期时间
	GetOrderExpireTime Endpoint = "dev.kdlapi.com/api/getorderexpiretime"

	// GetKpsProxy 获取独享代理
	GetKpsProxy Endpoint = "kps.kdlapi.com/api/getkps"
	// GetDpsProxy 获取私密代理
	GetDpsProxy Endpoint = "dps.kdlapi.com/api/getdps"
	// GetOpsProxyNormalOrVip 获取开放代理普通版和vip版代理
	GetOpsProxyNormalOrVip Endpoint = "dev.kdlapi.com/api/getproxy"
	// GetOpsProxySvip 获取开放代理Svip版
	GetOpsProxySvip Endpoint = "svip.kdlapi.com/api/getproxy"
	// GetOpsProxyEnt 获取开放代理企业版
	GetOpsProxyEnt Endpoint = "ent.kdlapi.com/api/getproxy"
	// CheckDpsValid 验证私密代理是否有效
	CheckDpsValid Endpoint = "dps.kdlapi.com/api/checkdpsvalid"
	// CheckOpsValid 验证开放代理是否有效
	CheckOpsValid Endpoint = "dev.kdlapi.com/api/checkopsvalid"
	// GetIPBalance 获取IP可用余额
	GetIPBalance Endpoint = "dps.kdlapi.com/api/getipbalance"
	// GetDpsValidTime 获取私密代理可用时间
	GetDpsValidTime Endpoint = "dps.kdlapi.com/api/getdpsvalidtime"
	// TpsCurrentIP 获取当前隧道代理IP
	TpsCurrentIP Endpoint = "tps.kdlapi.com/api/tpscurrentip"
	// ChangeTpsIP 更改当前隧道代理IP
	ChangeTpsIP Endpoint = "tps.kdlapi.com/api/changetpsip"
	//GetProxyAuthorization 获取代理鉴权信息
	GetProxyAuthorization Endpoint = "dev.kdlapi.com/api/getproxyauthorization"
	//GetTpsIP 获取隧道代理IP
	GetTpsIp Endpoint = "tps.kdlapi.com/api/gettps"

	//工具接口
	GetUA             Endpoint = "dev.kdlapi.com/api/getua"             //获取User Agent
	GetAreaCode       Endpoint = "dev.kdlapi.com/api/getareacode"       //获取指定地区编码
	GetAccountBalance Endpoint = "dev.kdlapi.com/api/getaccountbalance" //获取账户余额
	GetAccountOrders  Endpoint = "dev.kdlapi.com/api/getaccountorders"  //获取账户订单列表

	// 订单相关
	CreateOrder  Endpoint = "dev.kdlapi.com/api/createorder"  // 创建订单
	GetOrderInfo Endpoint = "dev.kdlapi.com/api/getorderinfo" // 获取订单信息
	SetAutoRenew Endpoint = "dev.kdlapi.com/api/setautorenew" // 开启/关闭自动续费
	QueryKpsCity Endpoint = "dev.kdlapi.com/api/querykpscity" // 查询独享代理城市信息

	GetSecretToken Endpoint = "auth.kdlapi.com/api/get_secret_token" // 获取密钥令牌

	GetOrderSecret Endpoint = "dev.kdlapi.com/api/getordersecret" // 获取订单密钥
)

func (e Endpoint) URL() *url.URL {
	u, _ := url.Parse("https://" + string(e))
	return u
}
