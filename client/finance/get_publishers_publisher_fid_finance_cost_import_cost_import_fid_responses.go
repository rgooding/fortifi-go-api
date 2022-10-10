// Code generated by go-swagger; DO NOT EDIT.

package finance

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

// GetPublishersPublisherFidFinanceCostImportCostImportFidReader is a Reader for the GetPublishersPublisherFidFinanceCostImportCostImportFid structure.
type GetPublishersPublisherFidFinanceCostImportCostImportFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublishersPublisherFidFinanceCostImportCostImportFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPublishersPublisherFidFinanceCostImportCostImportFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPublishersPublisherFidFinanceCostImportCostImportFidOK creates a GetPublishersPublisherFidFinanceCostImportCostImportFidOK with default headers values
func NewGetPublishersPublisherFidFinanceCostImportCostImportFidOK() *GetPublishersPublisherFidFinanceCostImportCostImportFidOK {
	return &GetPublishersPublisherFidFinanceCostImportCostImportFidOK{}
}

/*
GetPublishersPublisherFidFinanceCostImportCostImportFidOK describes a response with status code 200, with default header values.

Cost Import Status
*/
type GetPublishersPublisherFidFinanceCostImportCostImportFidOK struct {
	Payload *GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody
}

// IsSuccess returns true when this get publishers publisher fid finance cost import cost import fid o k response has a 2xx status code
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get publishers publisher fid finance cost import cost import fid o k response has a 3xx status code
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get publishers publisher fid finance cost import cost import fid o k response has a 4xx status code
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get publishers publisher fid finance cost import cost import fid o k response has a 5xx status code
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get publishers publisher fid finance cost import cost import fid o k response a status code equal to that given
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOK) Error() string {
	return fmt.Sprintf("[GET /publishers/{publisherFid}/finance/costImport/{costImportFid}][%d] getPublishersPublisherFidFinanceCostImportCostImportFidOK  %+v", 200, o.Payload)
}

func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOK) String() string {
	return fmt.Sprintf("[GET /publishers/{publisherFid}/finance/costImport/{costImportFid}][%d] getPublishersPublisherFidFinanceCostImportCostImportFidOK  %+v", 200, o.Payload)
}

func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOK) GetPayload() *GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody {
	return o.Payload
}

func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublishersPublisherFidFinanceCostImportCostImportFidDefault creates a GetPublishersPublisherFidFinanceCostImportCostImportFidDefault with default headers values
func NewGetPublishersPublisherFidFinanceCostImportCostImportFidDefault(code int) *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault {
	return &GetPublishersPublisherFidFinanceCostImportCostImportFidDefault{
		_statusCode: code,
	}
}

/*
GetPublishersPublisherFidFinanceCostImportCostImportFidDefault describes a response with status code -1, with default header values.

Error
*/
type GetPublishersPublisherFidFinanceCostImportCostImportFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get publishers publisher fid finance cost import cost import fid default response
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get publishers publisher fid finance cost import cost import fid default response has a 2xx status code
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get publishers publisher fid finance cost import cost import fid default response has a 3xx status code
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get publishers publisher fid finance cost import cost import fid default response has a 4xx status code
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get publishers publisher fid finance cost import cost import fid default response has a 5xx status code
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get publishers publisher fid finance cost import cost import fid default response a status code equal to that given
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault) Error() string {
	return fmt.Sprintf("[GET /publishers/{publisherFid}/finance/costImport/{costImportFid}][%d] GetPublishersPublisherFidFinanceCostImportCostImportFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault) String() string {
	return fmt.Sprintf("[GET /publishers/{publisherFid}/finance/costImport/{costImportFid}][%d] GetPublishersPublisherFidFinanceCostImportCostImportFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody get publishers publisher fid finance cost import cost import fid o k body
swagger:model GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody
*/
type GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody struct {
	models.Envelope

	// data
	Data *models.BoolMessage `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody) UnmarshalJSON(raw []byte) error {
	// GetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO0
	var getPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO0

	// GetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1
	var dataGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO0)
	var dataGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}

	dataGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1.Data = o.Data

	jsonDataGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1, errGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1 := swag.WriteJSON(dataGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1)
	if errGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1 != nil {
		return nil, errGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetPublishersPublisherFidFinanceCostImportCostImportFidOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get publishers publisher fid finance cost import cost import fid o k body
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPublishersPublisherFidFinanceCostImportCostImportFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPublishersPublisherFidFinanceCostImportCostImportFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get publishers publisher fid finance cost import cost import fid o k body based on the context it is used
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPublishersPublisherFidFinanceCostImportCostImportFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPublishersPublisherFidFinanceCostImportCostImportFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody) UnmarshalBinary(b []byte) error {
	var res GetPublishersPublisherFidFinanceCostImportCostImportFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
