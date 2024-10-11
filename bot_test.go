package feishu_test

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/feishu"
)

func TestNewWebHookBot(t *testing.T) {
	tt := zlsgo.NewTest(t)
	token := "4ebfb3bb-6057-43a1-a8c8-c4aa61135754"
	secret := "mXC0cCOqET3OYhpOVizRye"

	{
		bot := feishu.NewWebHookBot(token)
		err := bot.SendText("test")
		tt.Log(err)
	}

	{
		bot := feishu.NewWebHookBot(token, secret)
		err := bot.SendText("test")
		tt.Log(err)
	}
}
