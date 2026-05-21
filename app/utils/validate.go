package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors struct {
	errs []FieldError
}

func (ve *ValidationErrors) Add(field, message string) {
	ve.errs = append(ve.errs, FieldError{Field: field, Message: message})
}

func (ve *ValidationErrors) HasErrors() bool {
	return len(ve.errs) > 0
}

func (ve *ValidationErrors) Error() string {
	b, _ := json.Marshal(ve.errs)
	return string(b)
}

func (ve *ValidationErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(ve.errs)
}

type Validatable interface {
	Validate() error
}

// ParseAndValidate decodes the request body into T and runs its Validate().
// Returns a *ValidationErrors on validation failure so callers can pass it
// directly to utils.ValidationError(w, err).
func ParseAndValidate[T Validatable](r *http.Request) (*T, error) {
	body, err := ParseBody[T](r)
	if err != nil {
		return nil, err
	}
	if err := (*body).Validate(); err != nil {
		return nil, err
	}
	return body, nil
}

// ValidateSlice runs Validate() on every item and collects all field errors
// with their index prefix, e.g. "[0].product_id".
func ValidateSlice[T Validatable](items []T) error {
	ve := &ValidationErrors{}
	for i, item := range items {
		if err := item.Validate(); err != nil {
			if verrs, ok := err.(*ValidationErrors); ok {
				for _, fe := range verrs.errs {
					ve.Add(fmt.Sprintf("[%d].%s", i, fe.Field), fe.Message)
				}
			} else {
				ve.Add(fmt.Sprintf("[%d]", i), err.Error())
			}
		}
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}
