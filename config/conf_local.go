package config

type Local struct {
	Path      string `yaml:"path" json:"path"`
	StorePath string `yaml:"store_path" json:"store_path"`
	Size      int64  `yaml:"size" json:"size"`
}
