package util

const ERROR_CODE_RUNTIME int = 1000              //运行时异常
const ERROR_CODE_PARAM_IS_EMPTY int = 1001       //参数为空
const ERROR_CODE_OBJECT_NOT_EXIST int = 1002     //对象不存在
const ERROR_CODE_OBJECT_ALREADY_EXIST int = 1003 //对象已经存在
const ERROR_CODE_PARAM_ILLEGAL int = 1004        //非法参数

const ERROR_CODE_DATABASE_CONNECT_FAILED int = 1100 //连接数据库失败
const ERROR_CODE_DATABASE_QUERY_FAILED int = 1101   //数据查询失败

var ERROR_INFO_MAP map[int]string

func init() {
	ERROR_INFO_MAP[ERROR_CODE_RUNTIME] = "运行时异常"
	ERROR_INFO_MAP[ERROR_CODE_PARAM_IS_EMPTY] = "参数为空"
	ERROR_INFO_MAP[ERROR_CODE_OBJECT_NOT_EXIST] = "对象不存在"
	ERROR_INFO_MAP[ERROR_CODE_OBJECT_ALREADY_EXIST] = "对象已经存在"
	ERROR_INFO_MAP[ERROR_CODE_PARAM_ILLEGAL] = "非法参数"
	ERROR_INFO_MAP[ERROR_CODE_DATABASE_CONNECT_FAILED] = "连接数据库失败"
	ERROR_INFO_MAP[ERROR_CODE_DATABASE_QUERY_FAILED] = "数据查询失败"

}
