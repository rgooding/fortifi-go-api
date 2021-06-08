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

// PostEntitiesEntityFidConfigSectionNameReader is a Reader for the PostEntitiesEntityFidConfigSectionName structure.
type PostEntitiesEntityFidConfigSectionNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEntitiesEntityFidConfigSectionNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostEntitiesEntityFidConfigSectionNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostEntitiesEntityFidConfigSectionNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostEntitiesEntityFidConfigSectionNameOK creates a PostEntitiesEntityFidConfigSectionNameOK with default headers values
func NewPostEntitiesEntityFidConfigSectionNameOK() *PostEntitiesEntityFidConfigSectionNameOK {
	return &PostEntitiesEntityFidConfigSectionNameOK{}
}

/* PostEntitiesEntityFidConfigSectionNameOK describes a response with status code 200, with default header values.

Config Item Saved
*/
type PostEntitiesEntityFidConfigSectionNameOK struct {
}

func (o *PostEntitiesEntityFidConfigSectionNameOK) Error() string {
	return fmt.Sprintf("[POST /entities/{entityFid}/config/{sectionName}][%d] postEntitiesEntityFidConfigSectionNameOK ", 200)
}

func (o *PostEntitiesEntityFidConfigSectionNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostEntitiesEntityFidConfigSectionNameDefault creates a PostEntitiesEntityFidConfigSectionNameDefault with default headers values
func NewPostEntitiesEntityFidConfigSectionNameDefault(code int) *PostEntitiesEntityFidConfigSectionNameDefault {
	return &PostEntitiesEntityFidConfigSectionNameDefault{
		_statusCode: code,
	}
}

/* PostEntitiesEntityFidConfigSectionNameDefault describes a response with status code -1, with default header values.

Error
*/
type PostEntitiesEntityFidConfigSectionNameDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post entities entity fid config section name default response
func (o *PostEntitiesEntityFidConfigSectionNameDefault) Code() int {
	return o._statusCode
}

func (o *PostEntitiesEntityFidConfigSectionNameDefault) Error() string {
	return fmt.Sprintf("[POST /entities/{entityFid}/config/{sectionName}][%d] PostEntitiesEntityFidConfigSectionName default  %+v", o._statusCode, o.Payload)
}
func (o *PostEntitiesEntityFidConfigSectionNameDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostEntitiesEntityFidConfigSectionNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
