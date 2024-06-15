package constant

import "time"

const UserTableName = "user"
const InterfaceInfoTableName = "interfaceInfo"
const UserInterfaceInfoTableName = "user_interface_info"

const PatternStr = "/[`~!@#$%^&*()_\\-+=<>?:\"{}|,.\\/;'\\\\[\\]·~！@#￥%……&*（）——\\-+={}|《》？：“”【】、；‘'，。、]/"

const Salt = "gyu"

const (
	MemberRole     = 0
	AdminRole      = 1
	UsernameMinLen = 6
	PasswordMinLen = 8
)

// InterfaceInfoStatus
const (
	Online  = 1
	Offline = 0
)

const (
	DefaultTopNLimit = 3
	BlankString      = ""
	BlankInt         = 0
)

// Jwt

const KeyJwtUserId = "jwtUserId"
const TokenPrefixStr = "login:token:"
const TokenExpireTime = time.Hour * 24 * 7

// Redis Key

const KeyUserId = "user_id"
const KeyUserRole = "user_role"
const KeyUsername = "username"
const KeyAvatarUrl = "avatar_url"

// CORS
const (
	AllowOrigin        = "Access-Control-Allow-Origin"
	AllOrigins         = "*"
	AllowMethods       = "Access-Control-Allow-Methods"
	AllowHeaders       = "Access-Control-Allow-Headers"
	AllowCredentials   = "Access-Control-Allow-Credentials"
	AllowExposeHeaders = "Access-Control-Expose-Headers"
	Headers            = "Content-Type, Content-Length, Origin, Access-Control-Allow-Origin, Access-Control-Allow-Headers"
	Methods            = "GET, OPTIONS, POST, PATCH, PUT, DELETE"
	PostMethod         = "POST"
	True               = "true"
)

// 测试环境的网关 host

const GatewayHost = "http://localhost:8123"
const GatewayUrl = "/api/invoke"
