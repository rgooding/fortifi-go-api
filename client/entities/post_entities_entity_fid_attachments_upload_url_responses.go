// Code generated by go-swagger; DO NOT EDIT.

package entities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fortifi/go-api/models"
)

// PostEntitiesEntityFidAttachmentsUploadURLReader is a Reader for the PostEntitiesEntityFidAttachmentsUploadURL structure.
type PostEntitiesEntityFidAttachmentsUploadURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEntitiesEntityFidAttachmentsUploadURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostEntitiesEntityFidAttachmentsUploadURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostEntitiesEntityFidAttachmentsUploadURLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostEntitiesEntityFidAttachmentsUploadURLOK creates a PostEntitiesEntityFidAttachmentsUploadURLOK with default headers values
func NewPostEntitiesEntityFidAttachmentsUploadURLOK() *PostEntitiesEntityFidAttachmentsUploadURLOK {
	return &PostEntitiesEntityFidAttachmentsUploadURLOK{}
}

/*
PostEntitiesEntityFidAttachmentsUploadURLOK describes a response with status code 200, with default header values.

Attachment upload URL
*/
type PostEntitiesEntityFidAttachmentsUploadURLOK struct {
	Payload *PostEntitiesEntityFidAttachmentsUploadURLOKBody
}

// IsSuccess returns true when this post entities entity fid attachments upload Url o k response has a 2xx status code
func (o *PostEntitiesEntityFidAttachmentsUploadURLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post entities entity fid attachments upload Url o k response has a 3xx status code
func (o *PostEntitiesEntityFidAttachmentsUploadURLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post entities entity fid attachments upload Url o k response has a 4xx status code
func (o *PostEntitiesEntityFidAttachmentsUploadURLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post entities entity fid attachments upload Url o k response has a 5xx status code
func (o *PostEntitiesEntityFidAttachmentsUploadURLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post entities entity fid attachments upload Url o k response a status code equal to that given
func (o *PostEntitiesEntityFidAttachmentsUploadURLOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post entities entity fid attachments upload Url o k response
func (o *PostEntitiesEntityFidAttachmentsUploadURLOK) Code() int {
	return 200
}

func (o *PostEntitiesEntityFidAttachmentsUploadURLOK) Error() string {
	return fmt.Sprintf("[POST /entities/{entityFid}/attachments/uploadUrl][%d] postEntitiesEntityFidAttachmentsUploadUrlOK  %+v", 200, o.Payload)
}

func (o *PostEntitiesEntityFidAttachmentsUploadURLOK) String() string {
	return fmt.Sprintf("[POST /entities/{entityFid}/attachments/uploadUrl][%d] postEntitiesEntityFidAttachmentsUploadUrlOK  %+v", 200, o.Payload)
}

func (o *PostEntitiesEntityFidAttachmentsUploadURLOK) GetPayload() *PostEntitiesEntityFidAttachmentsUploadURLOKBody {
	return o.Payload
}

func (o *PostEntitiesEntityFidAttachmentsUploadURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostEntitiesEntityFidAttachmentsUploadURLOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEntitiesEntityFidAttachmentsUploadURLDefault creates a PostEntitiesEntityFidAttachmentsUploadURLDefault with default headers values
func NewPostEntitiesEntityFidAttachmentsUploadURLDefault(code int) *PostEntitiesEntityFidAttachmentsUploadURLDefault {
	return &PostEntitiesEntityFidAttachmentsUploadURLDefault{
		_statusCode: code,
	}
}

/*
PostEntitiesEntityFidAttachmentsUploadURLDefault describes a response with status code -1, with default header values.

Error
*/
type PostEntitiesEntityFidAttachmentsUploadURLDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post entities entity fid attachments upload URL default response has a 2xx status code
func (o *PostEntitiesEntityFidAttachmentsUploadURLDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post entities entity fid attachments upload URL default response has a 3xx status code
func (o *PostEntitiesEntityFidAttachmentsUploadURLDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post entities entity fid attachments upload URL default response has a 4xx status code
func (o *PostEntitiesEntityFidAttachmentsUploadURLDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post entities entity fid attachments upload URL default response has a 5xx status code
func (o *PostEntitiesEntityFidAttachmentsUploadURLDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post entities entity fid attachments upload URL default response a status code equal to that given
func (o *PostEntitiesEntityFidAttachmentsUploadURLDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post entities entity fid attachments upload URL default response
func (o *PostEntitiesEntityFidAttachmentsUploadURLDefault) Code() int {
	return o._statusCode
}

func (o *PostEntitiesEntityFidAttachmentsUploadURLDefault) Error() string {
	return fmt.Sprintf("[POST /entities/{entityFid}/attachments/uploadUrl][%d] PostEntitiesEntityFidAttachmentsUploadURL default  %+v", o._statusCode, o.Payload)
}

func (o *PostEntitiesEntityFidAttachmentsUploadURLDefault) String() string {
	return fmt.Sprintf("[POST /entities/{entityFid}/attachments/uploadUrl][%d] PostEntitiesEntityFidAttachmentsUploadURL default  %+v", o._statusCode, o.Payload)
}

func (o *PostEntitiesEntityFidAttachmentsUploadURLDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostEntitiesEntityFidAttachmentsUploadURLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostEntitiesEntityFidAttachmentsUploadURLOKBody post entities entity fid attachments upload URL o k body
swagger:model PostEntitiesEntityFidAttachmentsUploadURLOKBody
*/
type PostEntitiesEntityFidAttachmentsUploadURLOKBody struct {
	models.Envelope

	// data
	Data *models.AttachmentURL `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostEntitiesEntityFidAttachmentsUploadURLOKBody) UnmarshalJSON(raw []byte) error {
	// PostEntitiesEntityFidAttachmentsUploadURLOKBodyAO0
	var postEntitiesEntityFidAttachmentsUploadURLOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postEntitiesEntityFidAttachmentsUploadURLOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postEntitiesEntityFidAttachmentsUploadURLOKBodyAO0

	// PostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1
	var dataPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1 struct {
		Data *models.AttachmentURL `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostEntitiesEntityFidAttachmentsUploadURLOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postEntitiesEntityFidAttachmentsUploadURLOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postEntitiesEntityFidAttachmentsUploadURLOKBodyAO0)
	var dataPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1 struct {
		Data *models.AttachmentURL `json:"data,omitempty"`
	}

	dataPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1.Data = o.Data

	jsonDataPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1, errPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1 := swag.WriteJSON(dataPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1)
	if errPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1 != nil {
		return nil, errPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostEntitiesEntityFidAttachmentsUploadURLOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post entities entity fid attachments upload URL o k body
func (o *PostEntitiesEntityFidAttachmentsUploadURLOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostEntitiesEntityFidAttachmentsUploadURLOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postEntitiesEntityFidAttachmentsUploadUrlOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postEntitiesEntityFidAttachmentsUploadUrlOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post entities entity fid attachments upload URL o k body based on the context it is used
func (o *PostEntitiesEntityFidAttachmentsUploadURLOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostEntitiesEntityFidAttachmentsUploadURLOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postEntitiesEntityFidAttachmentsUploadUrlOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postEntitiesEntityFidAttachmentsUploadUrlOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostEntitiesEntityFidAttachmentsUploadURLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostEntitiesEntityFidAttachmentsUploadURLOKBody) UnmarshalBinary(b []byte) error {
	var res PostEntitiesEntityFidAttachmentsUploadURLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
