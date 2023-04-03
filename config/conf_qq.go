package config

import "fmt"

type QQ struct {
	AppId    string `yaml:"app_id" json:"app_id"`
	Key      string `yaml:"key" json:"key"`
	Redirect string `yaml:"redirect" json:"redirect"` //登录之后的回调地址
}

func (q QQ) GetPath() string {
	if q.Key == "" || q.AppId == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_url=%s", q.AppId, q.Redirect)
}
