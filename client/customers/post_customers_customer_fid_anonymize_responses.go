// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostCustomersCustomerFidAnonymizeReader is a Reader for the PostCustomersCustomerFidAnonymize structure.
type PostCustomersCustomerFidAnonymizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidAnonymizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerFidAnonymizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCustomersCustomerFidAnonymizeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidAnonymizeOK creates a PostCustomersCustomerFidAnonymizeOK with default headers values
func NewPostCustomersCustomerFidAnonymizeOK() *PostCustomersCustomerFidAnonymizeOK {
	return &PostCustomersCustomerFidAnonymizeOK{}
}

/*PostCustomersCustomerFidAnonymizeOK handles this case with default header values.

Anonymize Request
*/
type PostCustomersCustomerFidAnonymizeOK struct {
	Payload *PostCustomersCustomerFidAnonymizeOKBody
}

func (o *PostCustomersCustomerFidAnonymizeOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/anonymize][%d] postCustomersCustomerFidAnonymizeOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerFidAnonymizeOK) GetPayload() *PostCustomersCustomerFidAnonymizeOKBody {
	return o.Payload
}

func (o *PostCustomersCustomerFidAnonymizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCustomersCustomerFidAnonymizeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersCustomerFidAnonymizeDefault creates a PostCustomersCustomerFidAnonymizeDefault with default headers values
func NewPostCustomersCustomerFidAnonymizeDefault(code int) *PostCustomersCustomerFidAnonymizeDefault {
	return &PostCustomersCustomerFidAnonymizeDefault{
		_statusCode: code,
	}
}

/*PostCustomersCustomerFidAnonymizeDefault handles this case with default header values.

Error
*/
type PostCustomersCustomerFidAnonymizeDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid anonymize default response
func (o *PostCustomersCustomerFidAnonymizeDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidAnonymizeDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/anonymize][%d] PostCustomersCustomerFidAnonymize default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidAnonymizeDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostCustomersCustomerFidAnonymizeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCustomersCustomerFidAnonymizeOKBody post customers customer fid anonymize o k body
swagger:model PostCustomersCustomerFidAnonymizeOKBody
*/
type PostCustomersCustomerFidAnonymizeOKBody struct {
	models.Envelope

	// data
	Data *models.BoolMessage `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostCustomersCustomerFidAnonymizeOKBody) UnmarshalJSON(raw []byte) error {
	// PostCustomersCustomerFidAnonymizeOKBodyAO0
	var postCustomersCustomerFidAnonymizeOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postCustomersCustomerFidAnonymizeOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postCustomersCustomerFidAnonymizeOKBodyAO0

	// PostCustomersCustomerFidAnonymizeOKBodyAO1
	var dataPostCustomersCustomerFidAnonymizeOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostCustomersCustomerFidAnonymizeOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostCustomersCustomerFidAnonymizeOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostCustomersCustomerFidAnonymizeOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postCustomersCustomerFidAnonymizeOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postCustomersCustomerFidAnonymizeOKBodyAO0)

	var dataPostCustomersCustomerFidAnonymizeOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}

	dataPostCustomersCustomerFidAnonymizeOKBodyAO1.Data = o.Data

	jsonDataPostCustomersCustomerFidAnonymizeOKBodyAO1, errPostCustomersCustomerFidAnonymizeOKBodyAO1 := swag.WriteJSON(dataPostCustomersCustomerFidAnonymizeOKBodyAO1)
	if errPostCustomersCustomerFidAnonymizeOKBodyAO1 != nil {
		return nil, errPostCustomersCustomerFidAnonymizeOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostCustomersCustomerFidAnonymizeOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post customers customer fid anonymize o k body
func (o *PostCustomersCustomerFidAnonymizeOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostCustomersCustomerFidAnonymizeOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCustomersCustomerFidAnonymizeOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomersCustomerFidAnonymizeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomersCustomerFidAnonymizeOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomersCustomerFidAnonymizeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
