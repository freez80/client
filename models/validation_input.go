package models

import "github.com/go-openapi/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*ValidationInput validation input

swagger:model Validation Input
*/
type ValidationInput struct {

	/* yaml
	 */
	Yaml *string `json:"yaml,omitempty"`
}

// Validate validates this validation input
func (m *ValidationInput) Validate(formats strfmt.Registry) error {
	return nil
}
