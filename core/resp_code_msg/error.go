package resp_code_msg

const (
	ParamsError        = 401
	EmailSendManyError = 402
	LoginError         = 403
	TokenError         = 405
	GetDataError       = 406
	DirError           = 407
	FileError          = 408
	ParamLenError      = 409

	SystemError        = 500
	SendEmailError     = 501
	SaveDataError      = 502
	TokenGenerateError = 503
	TokenParseError    = 504
)
