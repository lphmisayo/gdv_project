package ctype

import "encoding/json"

type Role int

const (
	PermissionUser   Role = 1
	PermissionAdmin  Role = 2
	PermissionVisit  Role = 3
	PermissionBanned Role = 4
)

func (s Role) MarshalJson() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionUser:
		str = "用户"
	case PermissionAdmin:
		str = "管理员"
	case PermissionVisit:
		str = "游客"
	case PermissionBanned:
		str = "封禁人员"
	default:
		str = "其他"
	}
	return str
}
