package business

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrConflig = errors.New("Conflig")

	ErrInternalServer = errors.New("Internal_Server_Error")

	ErrNotFound = gorm.ErrRecordNotFound

	ErrInvalidSpec = errors.New("Given_Spec_Is_Not_Valid")
)
