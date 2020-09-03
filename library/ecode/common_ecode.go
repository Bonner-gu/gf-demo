package ecode

// All common ecode
var (
	OK = add(0) // 正确

	NotModified         = add(-304) // 木有改动
	TemporaryRedirect   = add(-307) // 撞车跳转
	RequestErr          = add(-400) // 请求错误
	Unauthorized        = add(-401) // 未认证 = NotLogin
	AccessDenied        = add(-403) // 访问权限不足
	NothingFound        = add(-404) // 啥都木有
	MethodNotAllowed    = add(-405) // 不支持该方法
	Conflict            = add(-409) // 冲突
	Canceled            = add(-498) // 客户端取消请求
	ServerErr           = add(-500) // 服务器错误
	ServiceUnavailable  = add(-503) // 过载保护,服务暂不可用
	Deadline            = add(-504) // 服务调用超时
	LimitExceed         = add(-509) // 超出限制
	ParamErr            = add(-600) // 参数错误
	CaptchaCheckErr     = add(-601) // 验证码校验失败
	CaptchaCreateErr    = add(-602) // 验证码创建失败
	MysqlManageErr      = add(-603) // db操作失败
	RedisManageErr      = add(-604) // redis操作失败
	TokenMakeErr        = add(-605) // token生成失败
	TokenInvalid        = add(-606) // token无效
	UploadFileErr       = add(-607) // 文件上传失败
	PublishingFailedErr = add(-608) // 版本发布失败

	PasswordInvalid = add(-4000) //密码不正确

	AdminNoExist = add(2000) //管理员不存在
)
