// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// PostCustomersCustomerFidPaymentMethodsPaypalInitialiseReader is a Reader for the PostCustomersCustomerFidPaymentMethodsPaypalInitialise structure.
type PostCustomersCustomerFidPaymentMethodsPaypalInitialiseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK creates a PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK with default headers values
func NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK() *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK {
	return &PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK{}
}

/* PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK describes a response with status code 200, with default header values.

Redirect instructions
*/
type PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK struct {
	Payload *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody
}

func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/paymentMethods/paypal/initialise][%d] postCustomersCustomerFidPaymentMethodsPaypalInitialiseOK  %+v", 200, o.Payload)
}
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK) GetPayload() *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody {
	return o.Payload
}

func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault creates a PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault with default headers values
func NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault(code int) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault {
	return &PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault{
		_statusCode: code,
	}
}

/* PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault describes a response with status code -1, with default header values.

Error
*/
type PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid payment methods paypal initialise default response
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/paymentMethods/paypal/initialise][%d] PostCustomersCustomerFidPaymentMethodsPaypalInitialise default  %+v", o._statusCode, o.Payload)
}
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody post customers customer fid payment methods paypal initialise o k body
swagger:model PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody
*/
type PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody struct {
	models.Envelope

	// data
	Data *models.PaypalInit `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody) UnmarshalJSON(raw []byte) error {
	// PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO0
	var postCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO0

	// PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1
	var dataPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1 struct {
		Data *models.PaypalInit `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO0)
	var dataPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1 struct {
		Data *models.PaypalInit `json:"data,omitempty"`
	}

	dataPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1.Data = o.Data

	jsonDataPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1, errPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1 := swag.WriteJSON(dataPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1)
	if errPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1 != nil {
		return nil, errPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post customers customer fid payment methods paypal initialise o k body
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCustomersCustomerFidPaymentMethodsPaypalInitialiseOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post customers customer fid payment methods paypal initialise o k body based on the context it is used
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCustomersCustomerFidPaymentMethodsPaypalInitialiseOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
