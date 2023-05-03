package handler

import "WebServer/utils"

type ApiMessage struct {
	Err     int32
	Message string
	Data    string
}

var (
	UNAUTHORIZE_MSG = ApiMessage{Err: -100, Message: "Unauthorized request", Data: ""}
	SUCCESS_MSG     = ApiMessage{Err: 0, Message: "Success", Data: ""}
)

func GetApiMessage(api ApiMessage) string {
	res, err := utils.Marshal(api)
	if err != nil {
		return "{\"Err\": -100, \"Message\": \"Cannot parse ApiMessage\", \"Data\": \"\"}"
	}
	return res
}
