package common

import (
	"ApiModule/business"
)

// Mengembalikan respons status dari permintaan
func NewBusinessErrorResponse(err error) (int, ControllerResponse) {
	return errorMapping(err)
}

func errorMapping(err error) (int, ControllerResponse) {
	switch err {
	default:
		return InternalServerErrorResponse()

	case business.ErrNotFound:
		return NotFoundResponse()

	case business.ErrInvalidSpec:
		return BadRequestResponse()

	case business.ErrConflig:
		return ConfligResponse()

	}
}
