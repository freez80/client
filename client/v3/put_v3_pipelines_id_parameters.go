package v3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/screwdriver-cd/client/models"
)

// NewPutV3PipelinesIDParams creates a new PutV3PipelinesIDParams object
// with the default values initialized.
func NewPutV3PipelinesIDParams() *PutV3PipelinesIDParams {
	var ()
	return &PutV3PipelinesIDParams{}
}

/*PutV3PipelinesIDParams contains all the parameters to send to the API endpoint
for the put v3 pipelines id operation typically these are written to a http.Request
*/
type PutV3PipelinesIDParams struct {

	/*Body*/
	Body *models.CreatePipeline
	/*ID
	  Identifier of this Pipeline

	*/
	ID string
}

// WithBody adds the body to the put v3 pipelines id params
func (o *PutV3PipelinesIDParams) WithBody(body *models.CreatePipeline) *PutV3PipelinesIDParams {
	o.Body = body
	return o
}

// WithID adds the id to the put v3 pipelines id params
func (o *PutV3PipelinesIDParams) WithID(id string) *PutV3PipelinesIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutV3PipelinesIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.CreatePipeline)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
