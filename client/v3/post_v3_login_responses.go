package v3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// PostV3LoginReader is a Reader for the PostV3Login structure.
type PostV3LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostV3LoginReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewPostV3LoginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostV3LoginDefault creates a PostV3LoginDefault with default headers values
func NewPostV3LoginDefault(code int) *PostV3LoginDefault {
	return &PostV3LoginDefault{
		_statusCode: code,
	}
}

/*PostV3LoginDefault handles this case with default header values.

Successful
*/
type PostV3LoginDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the post v3 login default response
func (o *PostV3LoginDefault) Code() int {
	return o._statusCode
}

func (o *PostV3LoginDefault) Error() string {
	return fmt.Sprintf("[POST /v3/login][%d] postV3Login default  %+v", o._statusCode, o.Payload)
}

func (o *PostV3LoginDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
