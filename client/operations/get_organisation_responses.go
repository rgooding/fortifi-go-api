// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// GetOrganisationReader is a Reader for the GetOrganisation structure.
type GetOrganisationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganisationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganisationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganisationOK creates a GetOrganisationOK with default headers values
func NewGetOrganisationOK() *GetOrganisationOK {
	return &GetOrganisationOK{}
}

/*GetOrganisationOK handles this case with default header values.

Organisation Information
*/
type GetOrganisationOK struct {
	Payload *models.GetOrganisationOKBody
}

func (o *GetOrganisationOK) Error() string {
	return fmt.Sprintf("[GET /organisation][%d] getOrganisationOK  %+v", 200, o.Payload)
}

func (o *GetOrganisationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetOrganisationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
