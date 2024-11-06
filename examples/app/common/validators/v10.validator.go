package validators

import "github.com/go-playground/validator/v10"

// Using it for DTO validation
var V10_validator *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
