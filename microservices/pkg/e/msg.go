package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ERROR_STATUS_UNAUTHORIZED: "Unauthorized",

	ERROR_INVALID_PARAMS: "Invalid Params",
	ERROR_NOT_FOUND_USER: "User could not be found",

	ERROR_GET_USERS:                  "Users could not retrieved",
	ERROR_CREATE_USER:                "User could not be created",
	ERROR_GET_USER:                   "User could not retrieved",
	ERROR_UPDATE_USER:                "User could not be updated",
	ERROR_DELETE_USER:                "User could not be deleted",
	ERROR_REGISTER:                   "Register error",
	ERROR_REGISTER_USER:              "User could not be registered",
	ERROR_LOGIN:                      "User could not be logged in",
	ERROR_INVALID_TYPE:               "Invalid user type",
	ERROR_INVALID_KEY:                "Invalid key",
	ERROR_GENERATE_HASH:              "Hash could not be generated",
	ERROR_WRONG_USERNAME_OR_PASSWORD: "Username or password is incorrect",
	ERROR_GET_TODOS:                  "Todos could not be retrieved",
	ERROR_CREATE_TODO:                "Todo could not be created",
	ERROR_GET_TODO:                   "Todo could not be retrieved",
	ERROR_UPDATE_TODO:                "Todo could not be updated",
	ERROR_DELETE_TODO:                "Todo could not be deleted",

	ERROR_TOKENHOURLIFESPAN_CONVERTION: "Token Hour Lifespan could not be converted",
	ERROR_UNEXPECTED_SIGNING_METHOD:    "Unxpected signing method",
	ERROR_TOKEN_GENERATION_FAIL:        "Token could not be generated",
	ERROR_NOT_FOUND_TOKEN:              "Token could not be found",
	ERROR_EXTRACT_TOKEN:                "Token could not be extracted",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
