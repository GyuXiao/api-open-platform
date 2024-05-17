package constant

//全局错误码

const (
	SUCCESS           uint32 = 0
	ERROR             uint32 = 1
	UnknownError      uint32 = 100000
	ServerCommonError uint32 = 100001
	ParamFormatError  uint32 = 100002
	RequestParamError uint32 = 100003
)

const JSON = "json"

const RequestTimeout = 5
