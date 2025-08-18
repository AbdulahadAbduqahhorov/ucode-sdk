package ucodesdk

import (
	"google.golang.org/api/option"
)

type Config struct {
	AppId          string
	BaseURL        string
	BotToken       string
	AccountIds     []string
	FunctionName   string
	FirebaseConfig option.ClientOption
}

func (cfg *Config) SetAppId(appId string) {
	cfg.AppId = appId
}

func (cfg *Config) SetBaseUrl(url string) {
	cfg.BaseURL = url
}

func (cfg *Config) SetBotToken(token string) {
	cfg.BotToken = token
}
