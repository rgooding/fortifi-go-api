// Code generated by go-swagger; DO NOT EDIT.

package marketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// DeleteAdvertisersAdvertiserFidApprovedReader is a Reader for the DeleteAdvertisersAdvertiserFidApproved structure.
type DeleteAdvertisersAdvertiserFidApprovedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAdvertisersAdvertiserFidApprovedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAdvertisersAdvertiserFidApprovedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteAdvertisersAdvertiserFidApprovedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAdvertisersAdvertiserFidApprovedOK creates a DeleteAdvertisersAdvertiserFidApprovedOK with default headers values
func NewDeleteAdvertisersAdvertiserFidApprovedOK() *DeleteAdvertisersAdvertiserFidApprovedOK {
	return &DeleteAdvertisersAdvertiserFidApprovedOK{}
}

/*DeleteAdvertisersAdvertiserFidApprovedOK handles this case with default header values.

Advertiser no longer approved
*/
type DeleteAdvertisersAdvertiserFidApprovedOK struct {
}

func (o *DeleteAdvertisersAdvertiserFidApprovedOK) Error() string {
	return fmt.Sprintf("[DELETE /advertisers/{advertiserFid}/approved][%d] deleteAdvertisersAdvertiserFidApprovedOK ", 200)
}

func (o *DeleteAdvertisersAdvertiserFidApprovedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAdvertisersAdvertiserFidApprovedDefault creates a DeleteAdvertisersAdvertiserFidApprovedDefault with default headers values
func NewDeleteAdvertisersAdvertiserFidApprovedDefault(code int) *DeleteAdvertisersAdvertiserFidApprovedDefault {
	return &DeleteAdvertisersAdvertiserFidApprovedDefault{
		_statusCode: code,
	}
}

/*DeleteAdvertisersAdvertiserFidApprovedDefault handles this case with default header values.

Error
*/
type DeleteAdvertisersAdvertiserFidApprovedDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the delete advertisers advertiser fid approved default response
func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) Error() string {
	return fmt.Sprintf("[DELETE /advertisers/{advertiserFid}/approved][%d] DeleteAdvertisersAdvertiserFidApproved default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
