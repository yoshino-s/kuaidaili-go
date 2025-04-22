package kuaidailigo

// SignType 签名计算方式
type SignType = string

const (
	SignTypeSimple   SignType = "simple"
	SignTypeHmacSha1 SignType = "hmacsha1"
	SignTypeToken    SignType = "token"
)
