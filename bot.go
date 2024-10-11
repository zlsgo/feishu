package feishu

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"

	"github.com/sohaha/zlsgo/zhttp"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
)

type WebHookBot struct {
	token  string
	secret string
	url    string
}

func NewWebHookBot(token string, secret ...string) *WebHookBot {
	bot := &WebHookBot{token: token, url: "https://open.feishu.cn/open-apis/bot/v2/hook/" + token}
	if len(secret) > 0 {
		bot.secret = secret[0]
	}
	return bot
}

func (b *WebHookBot) genSecret(timestamp int64) (string, error) {
	stringToSign := ztype.ToString(timestamp) + "\n" + b.secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}

	return zstring.Bytes2String(zstring.Base64Encode(h.Sum(nil))), nil
}

func (b *WebHookBot) SendText(text string) (err error) {
	timestamp := ztime.UnixMicro(ztime.Clock()).Unix()
	body := ztype.Map{
		"timestamp": timestamp,
		"msg_type":  "text",
		"content": ztype.Map{
			"text": text,
		},
	}

	if b.secret != "" {
		body["sign"], err = b.genSecret(timestamp)
		if err != nil {
			return err
		}
	}

	resp, err := zhttp.Post(b.url, zhttp.BodyJSON(body))
	if err != nil {
		return err
	}

	json := resp.JSONs()
	code := json.Get("code").String()
	if code != "0" {
		return errors.New(code + ": " + json.Get("msg").String())
	}

	return nil
}
