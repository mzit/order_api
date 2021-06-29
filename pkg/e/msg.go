package e

var MsgFlags = map[int]string{
	SUCCESS:        "OK",
	ERROR:          "Fail",
	INVALID_PARAMS: "请求参数错误",
}

func GetMsg(errorCode int) string {
	msg, ok := MsgFlags[errorCode]
	if ok {
		return msg
	}
	//找不到errorCode时统一返回Fail
	return MsgFlags[ERROR]
}
