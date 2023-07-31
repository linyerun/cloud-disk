package resp_code_msg

var msgMap = map[uint]string{
	Success:            "操作成功",
	ParamsError:        "请求参数有误",
	SendEmailError:     "邮件发送失败",
	EmailSendManyError: "该邮件地址不能短期内频繁发送",
	SaveDataError:      "数据保存失败",
	TokenGenerateError: "生成token失败",
	TokenParseError:    "解析token失败",
	LoginError:         "账号或密码错误",
	TokenError:         "token错误或者过期",
	GetDataError:       "您所提供的参数获取数据失败",
	DirError:           "您的文件夹存在文件关联",
	FileError:          "您的文件不存在文件关联",
	ParamLenError:      "参数存在长度错误",
	SystemError:        "系统错误",
}

func GetMsgByCode(code uint) string {
	val, ok := msgMap[code]
	if ok {
		return val
	}
	return ""
}
