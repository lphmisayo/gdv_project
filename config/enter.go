package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"QQ"`
	Email    Email    `yaml:"email"`
	Qiniu    Qiniu    `yaml:"qiniu"`
	Jwt      Jwt      `yaml:"jwt"`
}
