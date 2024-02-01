package dto

type (
	// Define common response struct
	Common struct {
		Status  string      `json:"status"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
