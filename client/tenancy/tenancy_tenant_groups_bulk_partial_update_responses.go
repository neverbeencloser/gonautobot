package tenancy

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// TenancyTenantGroupsBulkPartialUpdateReader is a Reader for the TenancyTenantGroupsBulkPartialUpdate structure.
type TenancyTenantGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantGroupsBulkPartialUpdateOK creates a TenancyTenantGroupsBulkPartialUpdateOK with default headers values
func NewTenancyTenantGroupsBulkPartialUpdateOK() *TenancyTenantGroupsBulkPartialUpdateOK {
	return &TenancyTenantGroupsBulkPartialUpdateOK{}
}

/* TenancyTenantGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

TenancyTenantGroupsBulkPartialUpdateOK tenancy tenant groups bulk partial update o k
*/
type TenancyTenantGroupsBulkPartialUpdateOK struct {
	Payload *models.TenantGroup
}

func (o *TenancyTenantGroupsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyTenantGroupsBulkPartialUpdateOK) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
