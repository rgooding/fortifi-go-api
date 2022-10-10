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

// GetCustomersCustomerFidTicketsTicketFidPostsReader is a Reader for the GetCustomersCustomerFidTicketsTicketFidPosts structure.
type GetCustomersCustomerFidTicketsTicketFidPostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidTicketsTicketFidPostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidTicketsTicketFidPostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidTicketsTicketFidPostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidTicketsTicketFidPostsOK creates a GetCustomersCustomerFidTicketsTicketFidPostsOK with default headers values
func NewGetCustomersCustomerFidTicketsTicketFidPostsOK() *GetCustomersCustomerFidTicketsTicketFidPostsOK {
	return &GetCustomersCustomerFidTicketsTicketFidPostsOK{}
}

/*
GetCustomersCustomerFidTicketsTicketFidPostsOK describes a response with status code 200, with default header values.

Ticket Data
*/
type GetCustomersCustomerFidTicketsTicketFidPostsOK struct {
	Payload *GetCustomersCustomerFidTicketsTicketFidPostsOKBody
}

// IsSuccess returns true when this get customers customer fid tickets ticket fid posts o k response has a 2xx status code
func (o *GetCustomersCustomerFidTicketsTicketFidPostsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get customers customer fid tickets ticket fid posts o k response has a 3xx status code
func (o *GetCustomersCustomerFidTicketsTicketFidPostsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customers customer fid tickets ticket fid posts o k response has a 4xx status code
func (o *GetCustomersCustomerFidTicketsTicketFidPostsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customers customer fid tickets ticket fid posts o k response has a 5xx status code
func (o *GetCustomersCustomerFidTicketsTicketFidPostsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get customers customer fid tickets ticket fid posts o k response a status code equal to that given
func (o *GetCustomersCustomerFidTicketsTicketFidPostsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets/{ticketFid}/posts][%d] getCustomersCustomerFidTicketsTicketFidPostsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsOK) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets/{ticketFid}/posts][%d] getCustomersCustomerFidTicketsTicketFidPostsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsOK) GetPayload() *GetCustomersCustomerFidTicketsTicketFidPostsOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidTicketsTicketFidPostsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidTicketsTicketFidPostsDefault creates a GetCustomersCustomerFidTicketsTicketFidPostsDefault with default headers values
func NewGetCustomersCustomerFidTicketsTicketFidPostsDefault(code int) *GetCustomersCustomerFidTicketsTicketFidPostsDefault {
	return &GetCustomersCustomerFidTicketsTicketFidPostsDefault{
		_statusCode: code,
	}
}

/*
GetCustomersCustomerFidTicketsTicketFidPostsDefault describes a response with status code -1, with default header values.

Error
*/
type GetCustomersCustomerFidTicketsTicketFidPostsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid tickets ticket fid posts default response
func (o *GetCustomersCustomerFidTicketsTicketFidPostsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get customers customer fid tickets ticket fid posts default response has a 2xx status code
func (o *GetCustomersCustomerFidTicketsTicketFidPostsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get customers customer fid tickets ticket fid posts default response has a 3xx status code
func (o *GetCustomersCustomerFidTicketsTicketFidPostsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get customers customer fid tickets ticket fid posts default response has a 4xx status code
func (o *GetCustomersCustomerFidTicketsTicketFidPostsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get customers customer fid tickets ticket fid posts default response has a 5xx status code
func (o *GetCustomersCustomerFidTicketsTicketFidPostsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get customers customer fid tickets ticket fid posts default response a status code equal to that given
func (o *GetCustomersCustomerFidTicketsTicketFidPostsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets/{ticketFid}/posts][%d] GetCustomersCustomerFidTicketsTicketFidPosts default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsDefault) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets/{ticketFid}/posts][%d] GetCustomersCustomerFidTicketsTicketFidPosts default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCustomersCustomerFidTicketsTicketFidPostsOKBody get customers customer fid tickets ticket fid posts o k body
swagger:model GetCustomersCustomerFidTicketsTicketFidPostsOKBody
*/
type GetCustomersCustomerFidTicketsTicketFidPostsOKBody struct {
	models.Envelope

	// data
	Data *models.TicketPosts `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidTicketsTicketFidPostsOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0
	var getCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0

	// GetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1
	var dataGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1 struct {
		Data *models.TicketPosts `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidTicketsTicketFidPostsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0)
	var dataGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1 struct {
		Data *models.TicketPosts `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1, errGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1)
	if errGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid tickets ticket fid posts o k body
func (o *GetCustomersCustomerFidTicketsTicketFidPostsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidTicketsTicketFidPostsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidTicketsTicketFidPostsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getCustomersCustomerFidTicketsTicketFidPostsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get customers customer fid tickets ticket fid posts o k body based on the context it is used
func (o *GetCustomersCustomerFidTicketsTicketFidPostsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidTicketsTicketFidPostsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidTicketsTicketFidPostsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getCustomersCustomerFidTicketsTicketFidPostsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidTicketsTicketFidPostsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidTicketsTicketFidPostsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidTicketsTicketFidPostsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
