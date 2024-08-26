package zerror

type ZErrorCode string

func getErrMsg(errCode ZErrorCode) string {
	msg, ok := errorMsgs[errCode]
	if ok {
		return msg
	}
	return ""
}

// 错误码与之对应的错误消息
var errorMsgs = map[ZErrorCode]string{}
