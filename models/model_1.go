package models

import (
	"github.com/go-openapi/errors"
	// "github.com/go-openapi/strfmt"
	strfmt "github.com/go-swagger/go-swagger/strfmt"

)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Model1 model 1

swagger:model Model 1
*/
type Model1 struct {
	GetBuild
}

// Validate validates this model 1
func (m *Model1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.GetBuild.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
