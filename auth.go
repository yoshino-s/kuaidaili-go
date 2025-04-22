package kuaidailigo

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
)

func (c *BaseClient) signStr(method string, url *url.URL, query url.Values) string {
	key := []byte(c.secretKey)
	hash := hmac.New(sha1.New, key)
	hash.Write([]byte(method + url.Path + "?" + query.Encode()))
	sig := base64.StdEncoding.EncodeToString([]byte(string(hash.Sum(nil))))
	return sig
}
