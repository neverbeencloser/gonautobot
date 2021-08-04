package tenancy

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// TenancyTenantGroupsBulkUpdateReader is a Reader for the TenancyTenantGroupsBulkUpdate structure.
type TenancyTenantGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantGroupsBulkUpdateOK creates a TenancyTenantGroupsBulkUpdateOK with default headers values
func NewTenancyTenantGroupsBulkUpdateOK() *TenancyTenantGroupsBulkUpdateOK {
	return &TenancyTenantGroupsBulkUpdateOK{}
}

/* TenancyTenantGroupsBulkUpdateOK describes a response with status code 200, with default header values.

TenancyTenantGroupsBulkUpdateOK tenancy tenant groups bulk update o k
*/
type TenancyTenantGroupsBulkUpdateOK struct {
	Payload *models.TenantGroup
}

func (o *TenancyTenantGroupsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyTenantGroupsBulkUpdateOK) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
