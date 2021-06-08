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

// GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsReader is a Reader for the GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachments structure.
type GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK creates a GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK with default headers values
func NewGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK() *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK {
	return &GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK{}
}

/* GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK describes a response with status code 200, with default header values.

List of ticket attachments
*/
type GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK struct {
	Payload *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets/{ticketFid}/posts/{ticketPostTimestamp}/attachments][%d] getCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK  %+v", 200, o.Payload)
}
func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK) GetPayload() *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault creates a GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault with default headers values
func NewGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault(code int) *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault {
	return &GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault{
		_statusCode: code,
	}
}

/* GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault describes a response with status code -1, with default header values.

Error
*/
type GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid tickets ticket fid posts ticket post timestamp attachments default response
func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets/{ticketFid}/posts/{ticketPostTimestamp}/attachments][%d] GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachments default  %+v", o._statusCode, o.Payload)
}
func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody get customers customer fid tickets ticket fid posts ticket post timestamp attachments o k body
swagger:model GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody
*/
type GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody struct {
	models.Envelope

	// data
	Data *models.Attachments `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO0
	var getCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO0

	// GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1
	var dataGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1 struct {
		Data *models.Attachments `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO0)
	var dataGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1 struct {
		Data *models.Attachments `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1, errGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1)
	if errGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid tickets ticket fid posts ticket post timestamp attachments o k body
func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get customers customer fid tickets ticket fid posts ticket post timestamp attachments o k body based on the context it is used
func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidTicketsTicketFidPostsTicketPostTimestampAttachmentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
