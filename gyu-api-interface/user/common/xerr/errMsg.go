package xerr

var mapCodMsg map[uint32]string

func init() {
	mapCodMsg = make(map[uint32]string)
	mapCodMsg[SUCCESS] = "success"
	mapCodMsg[ERROR] = "error"
	mapCodMsg[UnknownError] = "未知错误"
	mapCodMsg[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	mapCodMsg[ParamFormatError] = "参数格式错误"
	mapCodMsg[RequestParamError] = "参数缺失或不规范"
}

func GetMsgByCode(errCode uint32) string {
	if msg, ok := mapCodMsg[errCode]; ok {
		return msg
	}
	return "服务器开小差啦,稍后再来试一试"
}
