package error

//直接实例化，同一个包里不同文件的变量可以直接使用
var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	//实际业务定义的错误
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIIST_ARTICLE:       "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

/*
@code 根据错误码获得定义好的消息
利用go map自带的key检查机制，将所有没定义的错误码归类为500 error
*/
func GetMsg(code int) string {
	//从map中取值会自动得到一个bool类型的值，定义为ok
	//对其进行检测是为了防止传入空key 会返回 ""
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
