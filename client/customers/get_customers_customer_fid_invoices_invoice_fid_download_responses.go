// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// GetCustomersCustomerFidInvoicesInvoiceFidDownloadReader is a Reader for the GetCustomersCustomerFidInvoicesInvoiceFidDownload structure.
type GetCustomersCustomerFidInvoicesInvoiceFidDownloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidInvoicesInvoiceFidDownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidInvoicesInvoiceFidDownloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidInvoicesInvoiceFidDownloadOK creates a GetCustomersCustomerFidInvoicesInvoiceFidDownloadOK with default headers values
func NewGetCustomersCustomerFidInvoicesInvoiceFidDownloadOK() *GetCustomersCustomerFidInvoicesInvoiceFidDownloadOK {
	return &GetCustomersCustomerFidInvoicesInvoiceFidDownloadOK{}
}

/*GetCustomersCustomerFidInvoicesInvoiceFidDownloadOK handles this case with default header values.

Download URL
*/
type GetCustomersCustomerFidInvoicesInvoiceFidDownloadOK struct {
	Payload *models.InvoiceDownloadResponse
}

func (o *GetCustomersCustomerFidInvoicesInvoiceFidDownloadOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/invoices/{invoiceFid}/download][%d] getCustomersCustomerFidInvoicesInvoiceFidDownloadOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidInvoicesInvoiceFidDownloadOK) GetPayload() *models.InvoiceDownloadResponse {
	return o.Payload
}

func (o *GetCustomersCustomerFidInvoicesInvoiceFidDownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoiceDownloadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault creates a GetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault with default headers values
func NewGetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault(code int) *GetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault {
	return &GetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault{
		_statusCode: code,
	}
}

/*GetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault handles this case with default header values.

Error
*/
type GetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid invoices invoice fid download default response
func (o *GetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/invoices/{invoiceFid}/download][%d] GetCustomersCustomerFidInvoicesInvoiceFidDownload default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidInvoicesInvoiceFidDownloadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}