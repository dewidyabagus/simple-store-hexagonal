package business

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrConflig = errors.New("Conflig")

	ErrInternalServer = errors.New("Internal_Server_Error")

	ErrHasBeenModified = gorm.ErrInvalidData

	ErrNotFound = gorm.ErrRecordNotFound

	ErrInvalidSpec = errors.New("Given_Spec_Is_Not_Valid")

	ErrPasswordMisMatch = errors.New("Wrong_Password")

	ErrDeleted = errors.New("Object_Deleted")

	ErrLoginFailed = errors.New("Login_Failed")

	ErrNotHavePermission = errors.New("Not_Have_Permission")
)
