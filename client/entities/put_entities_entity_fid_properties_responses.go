// Code generated by go-swagger; DO NOT EDIT.

package entities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutEntitiesEntityFidPropertiesReader is a Reader for the PutEntitiesEntityFidProperties structure.
type PutEntitiesEntityFidPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEntitiesEntityFidPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutEntitiesEntityFidPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutEntitiesEntityFidPropertiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutEntitiesEntityFidPropertiesOK creates a PutEntitiesEntityFidPropertiesOK with default headers values
func NewPutEntitiesEntityFidPropertiesOK() *PutEntitiesEntityFidPropertiesOK {
	return &PutEntitiesEntityFidPropertiesOK{}
}

/* PutEntitiesEntityFidPropertiesOK describes a response with status code 200, with default header values.

Properties Saved
*/
type PutEntitiesEntityFidPropertiesOK struct {
}

func (o *PutEntitiesEntityFidPropertiesOK) Error() string {
	return fmt.Sprintf("[PUT /entities/{entityFid}/properties][%d] putEntitiesEntityFidPropertiesOK ", 200)
}

func (o *PutEntitiesEntityFidPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEntitiesEntityFidPropertiesDefault creates a PutEntitiesEntityFidPropertiesDefault with default headers values
func NewPutEntitiesEntityFidPropertiesDefault(code int) *PutEntitiesEntityFidPropertiesDefault {
	return &PutEntitiesEntityFidPropertiesDefault{
		_statusCode: code,
	}
}

/* PutEntitiesEntityFidPropertiesDefault describes a response with status code -1, with default header values.

Error
*/
type PutEntitiesEntityFidPropertiesDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put entities entity fid properties default response
func (o *PutEntitiesEntityFidPropertiesDefault) Code() int {
	return o._statusCode
}

func (o *PutEntitiesEntityFidPropertiesDefault) Error() string {
	return fmt.Sprintf("[PUT /entities/{entityFid}/properties][%d] PutEntitiesEntityFidProperties default  %+v", o._statusCode, o.Payload)
}
func (o *PutEntitiesEntityFidPropertiesDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutEntitiesEntityFidPropertiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
