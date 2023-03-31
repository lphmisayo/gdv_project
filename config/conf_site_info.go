package config

type SiteInfo struct {
	CreateAt  string `yaml:"create_at" json:"create_at"`
	Title     string `yaml:"title" json:"title"`
	Email     string `yaml:"email" json:"email"`
	Slogan    string `yaml:"slogan" json:"slogan"`
	GithubUrl string `yaml:"github_url" json:"github_url"`
}
