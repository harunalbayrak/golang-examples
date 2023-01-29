package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	ERROR_INVALID_PARAMS:           "invalid params",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "auth token failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token timeout",
	ERROR_AUTH_TOKEN:               "auth token",
	ERROR_STATUS_UNAUTHORIZED:      "Unauthorized",
	ERROR_NOT_FOUND:                "Not found",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
