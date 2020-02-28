// Code generated by go-swagger; DO NOT EDIT.

package orders

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

// PostOrdersOrderFidProductsReader is a Reader for the PostOrdersOrderFidProducts structure.
type PostOrdersOrderFidProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrdersOrderFidProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOrdersOrderFidProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostOrdersOrderFidProductsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostOrdersOrderFidProductsOK creates a PostOrdersOrderFidProductsOK with default headers values
func NewPostOrdersOrderFidProductsOK() *PostOrdersOrderFidProductsOK {
	return &PostOrdersOrderFidProductsOK{}
}

/*PostOrdersOrderFidProductsOK handles this case with default header values.

Product added to the order successfully
*/
type PostOrdersOrderFidProductsOK struct {
	Payload *PostOrdersOrderFidProductsOKBody
}

func (o *PostOrdersOrderFidProductsOK) Error() string {
	return fmt.Sprintf("[POST /orders/{orderFid}/products][%d] postOrdersOrderFidProductsOK  %+v", 200, o.Payload)
}

func (o *PostOrdersOrderFidProductsOK) GetPayload() *PostOrdersOrderFidProductsOKBody {
	return o.Payload
}

func (o *PostOrdersOrderFidProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostOrdersOrderFidProductsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrdersOrderFidProductsDefault creates a PostOrdersOrderFidProductsDefault with default headers values
func NewPostOrdersOrderFidProductsDefault(code int) *PostOrdersOrderFidProductsDefault {
	return &PostOrdersOrderFidProductsDefault{
		_statusCode: code,
	}
}

/*PostOrdersOrderFidProductsDefault handles this case with default header values.

Error
*/
type PostOrdersOrderFidProductsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post orders order fid products default response
func (o *PostOrdersOrderFidProductsDefault) Code() int {
	return o._statusCode
}

func (o *PostOrdersOrderFidProductsDefault) Error() string {
	return fmt.Sprintf("[POST /orders/{orderFid}/products][%d] PostOrdersOrderFidProducts default  %+v", o._statusCode, o.Payload)
}

func (o *PostOrdersOrderFidProductsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostOrdersOrderFidProductsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostOrdersOrderFidProductsOKBody post orders order fid products o k body
swagger:model PostOrdersOrderFidProductsOKBody
*/
type PostOrdersOrderFidProductsOKBody struct {
	models.Envelope

	// data
	Data *models.OrderAddProducts `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostOrdersOrderFidProductsOKBody) UnmarshalJSON(raw []byte) error {
	// PostOrdersOrderFidProductsOKBodyAO0
	var postOrdersOrderFidProductsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postOrdersOrderFidProductsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postOrdersOrderFidProductsOKBodyAO0

	// PostOrdersOrderFidProductsOKBodyAO1
	var dataPostOrdersOrderFidProductsOKBodyAO1 struct {
		Data *models.OrderAddProducts `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostOrdersOrderFidProductsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostOrdersOrderFidProductsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostOrdersOrderFidProductsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postOrdersOrderFidProductsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postOrdersOrderFidProductsOKBodyAO0)

	var dataPostOrdersOrderFidProductsOKBodyAO1 struct {
		Data *models.OrderAddProducts `json:"data,omitempty"`
	}

	dataPostOrdersOrderFidProductsOKBodyAO1.Data = o.Data

	jsonDataPostOrdersOrderFidProductsOKBodyAO1, errPostOrdersOrderFidProductsOKBodyAO1 := swag.WriteJSON(dataPostOrdersOrderFidProductsOKBodyAO1)
	if errPostOrdersOrderFidProductsOKBodyAO1 != nil {
		return nil, errPostOrdersOrderFidProductsOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostOrdersOrderFidProductsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post orders order fid products o k body
func (o *PostOrdersOrderFidProductsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostOrdersOrderFidProductsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOrdersOrderFidProductsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOrdersOrderFidProductsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOrdersOrderFidProductsOKBody) UnmarshalBinary(b []byte) error {
	var res PostOrdersOrderFidProductsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
