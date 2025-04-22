package kuaidailigo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

type BaseClient struct {
	secretID  string
	secretKey string

	http *http.Client
	sign SignType

	secretToken string
}

func newClient(secretID, secretKey string, options ...WithOption) *BaseClient {
	r := &BaseClient{
		secretID:  secretID,
		secretKey: secretKey,

		http: &http.Client{},
		sign: SignTypeHmacSha1,
	}
	for _, option := range options {
		option(r)
	}

	return r
}

func (client *BaseClient) getParams(method string, ep Endpoint, kwargs map[string]string) url.Values {
	signType := client.sign
	params := url.Values{}
	params.Set("secret_id", client.secretID)
	params.Set("sign_type", string(signType))
	for k, v := range kwargs {
		params.Set(k, v)
	}

	if client.secretToken == "" && signType == SignTypeToken {
		signType = SignTypeHmacSha1
	}

	switch signType {
	case SignTypeSimple:
		params.Set("signature", client.secretKey)
	case SignTypeToken:
		secretToken, _ := client.getSecretToken(context.Background())
		client.secretToken = secretToken
		params.Set("signature", secretToken)
	case SignTypeHmacSha1:
		params.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
		var sign string
		sign = client.signStr(method, ep.URL(), params)
		params.Set("signature", sign)
	}
	return params
}

func (client *BaseClient) getSecretToken(ctx context.Context) (string, int64) {
	var resp struct {
		SecretToken string `json:"secret_token"`
		Expire      int64  `json:"expire"`
	}

	err := client.getBaseRes(ctx, "POST", GetSecretToken, map[string]string{
		"secret_id":  client.secretID,
		"secret_key": client.secretKey,
	}, &resp)

	if err != nil {
		return "", 0
	}

	return resp.SecretToken, resp.Expire
}

func (client *BaseClient) getBaseRes(ctx context.Context, method string, _endpoint Endpoint, _params map[string]string, response any) error {
	endpoint, err := url.Parse("https://" + string(_endpoint))
	if err != nil {
		return err
	}
	params := client.getParams(method, _endpoint, _params)

	var body io.Reader
	if method == http.MethodPost {
		body = strings.NewReader(params.Encode())
	} else {
		body = nil
		endpoint.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		endpoint.String(),
		body,
	)
	if err != nil {
		return err
	}
	r, err := client.http.Do(req)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		panic("fail to read body bytes")
	}
	bodyString := string(bodyBytes)

	if r.StatusCode != http.StatusOK {
		return errors.New("request status code error: " + strconv.Itoa(r.StatusCode) + ":" + bodyString)
	}

	code := gjson.Get(bodyString, "code")
	msg := gjson.Get(bodyString, "msg")
	if code.Int() != 0 {
		return fmt.Errorf("response status code error: %d, err msg: %s", code.Int(), msg.String())
	}

	data := gjson.Get(bodyString, "data").String()

	if response != nil {
		if err := json.Unmarshal([]byte(data), response); err != nil {
			fmt.Printf("data: %+v\n", data)
			fmt.Printf("err: %+v\n", err)
			return errors.New("fail to parse result content: " + bodyString)
		}
	}
	return nil
}
