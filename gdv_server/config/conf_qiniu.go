package config

type Qiniu struct {
	Enable    bool   `yaml:"enable" json:"enable"`
	AccessKey string `yaml:"access_key" json:"access_key"`
	SecretKey string `yaml:"secret_key" json:"secret_key"`
	Bucket    string `yaml:"bucket" json:"bucket"` // 存储桶名字
	Cdn       string `yaml:"cdn" json:"cdn"`       //访问图片的地址的前缀
	Zone      string `yaml:"zone" json:"zone"`     //存储的地区
	Size      int    `yaml:"size" json:"size"`     //图片存储大小上限 单位是Mb
}
