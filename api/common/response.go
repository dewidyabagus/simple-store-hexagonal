package common

type responseCode string

type ControllerResponse struct {
	Code    responseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}
}
