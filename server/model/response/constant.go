package response

type State struct {
	Code  int         `json:"code,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

var (
	Success = State{Code: 200}

	CreateDataError = R{State: State{1001, "创建失败"}}
	DeleteDataError = R{State: State{1002, "删除失败"}}
	UpdateDataError = R{State: State{1003, "更新失败"}}
	SelectDataError = R{State: State{1004, "查询失败"}}

	TokenCreateError                  = R{State: State{3001, "token创建失败"}}
	UserDataDoesNotExist              = R{State: State{3002, "用户数据不存在"}}
	TokenAuthorizationError           = R{State: State{3003, "token验证失败"}}
	CaptchaVerificationIsRequired     = R{State: State{3004, "需要验证码验证"}}
	VerificationCodeCreateError       = R{State: State{3005, "验证码生成失败"}}
	VerificationCodeVerificationError = R{State: State{3006, "验证码验证失败"}}

	ParameterBindingError    = R{State: State{Code: 5000, Error: "参数绑定错误"}}
	ParameterValidationError = R{State: State{Code: 5001, Error: "参数校验错误"}}
)
