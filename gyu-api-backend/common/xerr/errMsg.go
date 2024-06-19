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

	// JWT
	mapCodMsg[NotLogin] = "未登录"
	mapCodMsg[LoginExpired] = "登录过期"
	mapCodMsg[TokenExpire] = "token 已过期，请重新登陆"
	mapCodMsg[TokenNotValidYet] = "token 无效，请重新登陆"
	mapCodMsg[TokenMalformed] = "token 不正确，请重新登陆"
	mapCodMsg[TokenInvalid] = "这不是一个 token，请重新登陆"
	mapCodMsg[TokenCreateFail] = "token 创建失败"
	mapCodMsg[PermissionDenied] = "权限不足"
	mapCodMsg[TokenParseError] = "token 解析错误"
	mapCodMsg[TokenInsertError] = "向缓存中插入 token 错误"
	mapCodMsg[TokenGetFromCacheError] = "从缓存中获取 token 错误"

	// Encryption
	mapCodMsg[EncryptionError] = "encrypt 加密错误"
	mapCodMsg[DecodeMd5Error] = "md5 解码错误"

	// DB
	mapCodMsg[RecordDuplicateError] = "数据库记录重复"
	mapCodMsg[RecordNotFoundError] = "数据库未找到记录"
	mapCodMsg[RecordUpdateError] = "数据库更新记录错误"
	mapCodMsg[RecordDeleteError] = "数据库删除记录错误"
	mapCodMsg[RecordCreateError] = "数据库创建记录错误"
	mapCodMsg[RecordCountError] = "数据库统计记录错误"
	mapCodMsg[SearchUserError] = "数据库检索用户错误"
	mapCodMsg[CreateUserError] = "数据库创建用户错误"

	// Redis
	mapCodMsg[KeyExpireError] = "设置 key 过期时间错误"
	mapCodMsg[KeyDelError] = "删除 key 错误"
	mapCodMsg[KeyInsertError] = "插入 key 错误"

	// User
	mapCodMsg[UserNotExistError] = "用户不存在"
	mapCodMsg[UserExistError] = "用户已经存在"
	mapCodMsg[UserLoginError] = "用户登陆错误"
	mapCodMsg[UserRegisterError] = "用户注册错误"
	mapCodMsg[UserPasswordError] = "用户密码错误"
	mapCodMsg[UserIdNotExistError] = "UserId 不存在"
	mapCodMsg[UserNotLoginError] = "用户当前状态并没有登陆"
	mapCodMsg[SearchUserByAccessKeyError] = "通过 accessKey 检索用户信息错误"
	mapCodMsg[AccessKeyNotExistError] = "accessKey 不存在"

	// InterfaceInfo
	mapCodMsg[AddInterfaceInfoError] = "添加接口信息错误"
	mapCodMsg[SearchInterfaceInfoPageListError] = "查询接口列表错误"
	mapCodMsg[SearchInterfaceInfoError] = "查询接口错误"
	mapCodMsg[UpdateInterfaceInfoError] = "更新接口信息错误"
	mapCodMsg[DeleteInterfaceInfoError] = "删除接口信息错误"
	mapCodMsg[UpdateInterfaceInfoStatusError] = "更新接口状态错误"
	mapCodMsg[InterfaceInfoOfflineError] = "接口未上线错误"
	mapCodMsg[SearchTopNInvokeInterfaceInfoError] = "搜索调用次数 TopN 的接口信息错误"

	// UserInterfaceInfo
	mapCodMsg[CreateUserInterfaceInfoError] = "添加用户接口关系信息错误"
	mapCodMsg[SearchUserInterfaceInfoError] = "查询用户接口关系信息错误"
	mapCodMsg[InvokeSuccessUpdateError] = "用户接口调用成功更新次数统计错误"
	mapCodMsg[InvokeInterfaceLeftNumNonPositiveError] = "用户的剩余接口调用次数不大于0"

	// SDK
	mapCodMsg[SDKNewClientError] = "SDK 创建客户端错误"
	mapCodMsg[SDKSendRequestError] = "SDK 客户端发起请求错误"

	// JSON
	mapCodMsg[JSONMarshalError] = "JSON 序列化错误"
	mapCodMsg[JSONUnmarshalError] = "JSON 反序列化错误"

}

func GetMsgByCode(errCode uint32) string {
	if msg, ok := mapCodMsg[errCode]; ok {
		return msg
	}
	return "服务器开小差啦,稍后再来试一试"
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := mapCodMsg[errCode]; ok {
		return true
	}
	return false
}
