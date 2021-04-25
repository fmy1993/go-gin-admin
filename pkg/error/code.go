/*
 * @Descripttion:
 * @version:
 * @Author: fmy1993
 * @Date: 2021-04-07 11:48:21
 * @LastEditors: fmy1993
 * @LastEditTime: 2021-04-25 11:45:19
 */
package error

//枚举HTTP状态码
const (
	SUCCESS = 200
	//服务器错误
	ERROR = 500
	//参数错误
	INVALID_PARAMS = 400
	//根据业务自定义的http错误码
	//标签已经存在
	ERROR_EXIST_TAG = 10001
	//标签不存在
	ERROR_NOT_EXIST_TAG = 10002
	//文章不存在
	ERROR_NOT_EXIIST_ARTICLE = 10003
	//使用jwt的相关错误
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)
