package kuaidailigo

type ProxyProtocol int

const (
	ProxyProtocolHTTP  ProxyProtocol = 1
	ProxyProtocolSocks ProxyProtocol = 2
)

type OrderStatus string

const (
	OrderStatusAll       OrderStatus = "ALL"       // 全部
	OrderStatusValid     OrderStatus = "VALID"     // 有效
	OrderStatusWaitPay   OrderStatus = "WAIT_PAY"  // 待支付
	OrderStatusOpening   OrderStatus = "OPENING"   // 开通中
	OrderStatusExpired   OrderStatus = "EXPIRED"   // 已过期
	OrderStatusOwing     OrderStatus = "OWING"     // 欠费暂停
	OrderStatusClosed    OrderStatus = "CLOSED"    // 已关闭
	OrderStatusForbidden OrderStatus = "FORBIDDEN" // 被封禁
)

type PayType string

const (
	PayTypePrePay   PayType = "PRE_PAY"    // 包年包月(预付费)
	PayTypePrePayIP PayType = "PRE_PAY_IP" // 按IP付费(预付费)
	PayTypePostPay  PayType = "POST_PAY"   // 按量付费(后付费)
)

type ProductType string

const (
	ProductTypeTPSProductType ProductType = "TPS_PRO" // 普通代理
	ProductTypeTPS            ProductType = "TPS"     // 隧道代理
	ProductTypeDPS            ProductType = "DPS"     // 私密代理
	ProductTypeKPS            ProductType = "KPS"     // 独享代理
)

type TimeType string

const (
	TimeTypeDay   TimeType = "DAY"   // 天
	TimeTypeWeek  TimeType = "WEEK"  // 周
	TimeTypeMonth TimeType = "MONTH" // 月
	TimeTypeYear  TimeType = "YEAR"  // 年
)
