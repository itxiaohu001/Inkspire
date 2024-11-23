package dto

type Basic struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var SuccessBasic = Basic{
	Code:    200,
	Message: "success",
}
