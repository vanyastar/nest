package catsDtos

import "example/app/common/validators"

type SetCatNameNoValidation struct {
	Name string `json:"name" validate:"required"`
}

// SetCatNameNoValidation doesn't have a validation method.

type SetCatName struct {
	Name string `json:"name" validate:"required"`
}

// Validate checks the fields of the setCatNameDto instance against predefined validation rules.
// This method is automatically invoked after processing the request containing this DTO.
// It returns an error if any validation rules are violated, or nil if the instance is valid.
func (dto *SetCatName) Validate() error {
	return validators.V10_validator.Struct(dto)
}
