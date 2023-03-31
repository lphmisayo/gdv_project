package res

type ErrCode int

// err_code 错误处理包
const (
	SettingError ErrCode = 3001 //系统配置错误
	ParamError   ErrCode = 3002 //参数错误
)

var (
	ErrMap = map[ErrCode]string{
		SettingError: "系统配置错误",
		ParamError:   "参数错误",
	}
)
