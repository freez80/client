package v3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/screwdriver-cd/client/models"
)

// GetV3JobsIDReader is a Reader for the GetV3JobsID structure.
type GetV3JobsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetV3JobsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV3JobsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV3JobsIDOK creates a GetV3JobsIDOK with default headers values
func NewGetV3JobsIDOK() *GetV3JobsIDOK {
	return &GetV3JobsIDOK{}
}

/*GetV3JobsIDOK handles this case with default header values.

Successful
*/
type GetV3JobsIDOK struct {
	Payload *models.Model2
}

func (o *GetV3JobsIDOK) Error() string {
	return fmt.Sprintf("[GET /v3/jobs/{id}][%d] getV3JobsIdOK  %+v", 200, o.Payload)
}

func (o *GetV3JobsIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}