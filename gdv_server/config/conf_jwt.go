package config

type Jwt struct {
	Secret string `yaml:"secret" json:"secret"` //密钥
	Expire int    `yaml:"expire" json:"expire"` //过期时间 单位是小时
	Issuer string `yaml:"issuer" json:"issuer"` //颁发人
}
