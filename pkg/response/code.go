package response

type Code int

const (
	Success Code = 0
	Failed  Code = 10000 + iota
	UnAuthed
	AuthExpired
	InternalErr
	RequestMethodErr
	RequestParamErr
	RequestFormErr
	AccountInfoErr
)

var CodeMap = map[Code]string{
	Success:          "请求成功",
	Failed:           "请求失败",
	UnAuthed:         "用户未认证",
	AuthExpired:      "会话过期，请重新登录",
	InternalErr:      "服务器内部错误",
	RequestMethodErr: "请求方式错误",
	RequestParamErr:  "请求参数错误",
	RequestFormErr:   "请求表单错误",
	AccountInfoErr:   "账号信息有误",
}

// Msg 返回错误码对应的说明
func (c Code) Msg() string {
	if v, ok := CodeMap[c]; ok {
		return v
	}
	return ``
}
