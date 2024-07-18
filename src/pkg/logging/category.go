package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Internal        Category = "Internal"
	Database        Category = "Database"
	CacheService    Category = "CacheService"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
	Io              Category = "Io"
)

const (
	// General
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"

	// Postgres
	Migration SubCategory = "Migration"
	Select    SubCategory = "Select"
	Rollback  SubCategory = "Rollback"
	Update    SubCategory = "Update"
	Delete    SubCategory = "Delete"
	Insert    SubCategory = "Insert"

	// Internal
	Api                 SubCategory = "Api"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"
	FailedToCreateUser  SubCategory = "FailedToCreateUser"

	// Validation
	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"

	// IO
	RemoveFile SubCategory = "RemoveFile"
	UploadFile SubCategory = "UploadFile"

	RecoverError SubCategory = "Recover Error"
)

const (
	AppName       ExtraKey = "AppName"
	LoggerName    ExtraKey = "Logger"
	ClientIp      ExtraKey = "ClientIp"
	HostIp        ExtraKey = "HostIp"
	Method        ExtraKey = "Method"
	StatusCode    ExtraKey = "StatusCode"
	BodySize      ExtraKey = "BodySize"
	Path          ExtraKey = "Path"
	Latency       ExtraKey = "Latency"
	RequestBody   ExtraKey = "RequestBody"
	RequestHeader ExtraKey = "RequestHeader"
	ResponseBody  ExtraKey = "ResponseBody"
	ErrorMessage  ExtraKey = "ErrorMessage"
)
