package ctype

import "encoding/json"

type SignStatus int

const (
	SignedByQq     SignStatus = 1
	SignedByEmail  SignStatus = 2
	SignedByGithub SignStatus = 3
)

func (s SignStatus) MarshalJson() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignedByQq:
		str = "QQ"
	case SignedByEmail:
		str = "邮箱"
	case SignedByGithub:
		str = "GitHub"
	default:
		str = "其他"
	}
	return str
}
