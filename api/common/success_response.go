package common

import (
	"net/http"
)

const (
	Success responseCode = "200"
)

func SuccessResponseWithData(data interface{}) (int, ControllerResponse) {
	return http.StatusOK, ControllerResponse{Success, "success", data}
}

func SuccessResponseWithoutData() (int, ControllerResponse) {
	return http.StatusOK, ControllerResponse{Success, "success", map[string]interface{}{}}
}
