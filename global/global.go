package global

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate
	Logger   *log.Logger
)
