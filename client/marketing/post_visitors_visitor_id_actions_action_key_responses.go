// Code generated by go-swagger; DO NOT EDIT.

package marketing

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

// PostVisitorsVisitorIDActionsActionKeyReader is a Reader for the PostVisitorsVisitorIDActionsActionKey structure.
type PostVisitorsVisitorIDActionsActionKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVisitorsVisitorIDActionsActionKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostVisitorsVisitorIDActionsActionKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostVisitorsVisitorIDActionsActionKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVisitorsVisitorIDActionsActionKeyOK creates a PostVisitorsVisitorIDActionsActionKeyOK with default headers values
func NewPostVisitorsVisitorIDActionsActionKeyOK() *PostVisitorsVisitorIDActionsActionKeyOK {
	return &PostVisitorsVisitorIDActionsActionKeyOK{}
}

/* PostVisitorsVisitorIDActionsActionKeyOK describes a response with status code 200, with default header values.

Action Tracked
*/
type PostVisitorsVisitorIDActionsActionKeyOK struct {
	Payload *PostVisitorsVisitorIDActionsActionKeyOKBody
}

func (o *PostVisitorsVisitorIDActionsActionKeyOK) Error() string {
	return fmt.Sprintf("[POST /visitors/{visitorId}/actions/{actionKey}][%d] postVisitorsVisitorIdActionsActionKeyOK  %+v", 200, o.Payload)
}
func (o *PostVisitorsVisitorIDActionsActionKeyOK) GetPayload() *PostVisitorsVisitorIDActionsActionKeyOKBody {
	return o.Payload
}

func (o *PostVisitorsVisitorIDActionsActionKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostVisitorsVisitorIDActionsActionKeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVisitorsVisitorIDActionsActionKeyDefault creates a PostVisitorsVisitorIDActionsActionKeyDefault with default headers values
func NewPostVisitorsVisitorIDActionsActionKeyDefault(code int) *PostVisitorsVisitorIDActionsActionKeyDefault {
	return &PostVisitorsVisitorIDActionsActionKeyDefault{
		_statusCode: code,
	}
}

/* PostVisitorsVisitorIDActionsActionKeyDefault describes a response with status code -1, with default header values.

Error
*/
type PostVisitorsVisitorIDActionsActionKeyDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post visitors visitor ID actions action key default response
func (o *PostVisitorsVisitorIDActionsActionKeyDefault) Code() int {
	return o._statusCode
}

func (o *PostVisitorsVisitorIDActionsActionKeyDefault) Error() string {
	return fmt.Sprintf("[POST /visitors/{visitorId}/actions/{actionKey}][%d] PostVisitorsVisitorIDActionsActionKey default  %+v", o._statusCode, o.Payload)
}
func (o *PostVisitorsVisitorIDActionsActionKeyDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostVisitorsVisitorIDActionsActionKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostVisitorsVisitorIDActionsActionKeyOKBody post visitors visitor ID actions action key o k body
swagger:model PostVisitorsVisitorIDActionsActionKeyOKBody
*/
type PostVisitorsVisitorIDActionsActionKeyOKBody struct {
	models.Envelope

	// data
	Data *models.VisitorPostAction `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostVisitorsVisitorIDActionsActionKeyOKBody) UnmarshalJSON(raw []byte) error {
	// PostVisitorsVisitorIDActionsActionKeyOKBodyAO0
	var postVisitorsVisitorIDActionsActionKeyOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postVisitorsVisitorIDActionsActionKeyOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postVisitorsVisitorIDActionsActionKeyOKBodyAO0

	// PostVisitorsVisitorIDActionsActionKeyOKBodyAO1
	var dataPostVisitorsVisitorIDActionsActionKeyOKBodyAO1 struct {
		Data *models.VisitorPostAction `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostVisitorsVisitorIDActionsActionKeyOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostVisitorsVisitorIDActionsActionKeyOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostVisitorsVisitorIDActionsActionKeyOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postVisitorsVisitorIDActionsActionKeyOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postVisitorsVisitorIDActionsActionKeyOKBodyAO0)
	var dataPostVisitorsVisitorIDActionsActionKeyOKBodyAO1 struct {
		Data *models.VisitorPostAction `json:"data,omitempty"`
	}

	dataPostVisitorsVisitorIDActionsActionKeyOKBodyAO1.Data = o.Data

	jsonDataPostVisitorsVisitorIDActionsActionKeyOKBodyAO1, errPostVisitorsVisitorIDActionsActionKeyOKBodyAO1 := swag.WriteJSON(dataPostVisitorsVisitorIDActionsActionKeyOKBodyAO1)
	if errPostVisitorsVisitorIDActionsActionKeyOKBodyAO1 != nil {
		return nil, errPostVisitorsVisitorIDActionsActionKeyOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostVisitorsVisitorIDActionsActionKeyOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post visitors visitor ID actions action key o k body
func (o *PostVisitorsVisitorIDActionsActionKeyOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostVisitorsVisitorIDActionsActionKeyOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postVisitorsVisitorIdActionsActionKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post visitors visitor ID actions action key o k body based on the context it is used
func (o *PostVisitorsVisitorIDActionsActionKeyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostVisitorsVisitorIDActionsActionKeyOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postVisitorsVisitorIdActionsActionKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostVisitorsVisitorIDActionsActionKeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostVisitorsVisitorIDActionsActionKeyOKBody) UnmarshalBinary(b []byte) error {
	var res PostVisitorsVisitorIDActionsActionKeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
