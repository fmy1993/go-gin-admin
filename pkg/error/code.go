package error

//枚举HTTP状态码
const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	//根据业务自定义的http错误码
	ERROR_EXIST_TAG          = 10001
	ERROR_NOT_EXIST_TAG      = 10002
	ERROR_NOT_EXIIST_ARTICLE = 10003
	//使用jwt的相关错误
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)
