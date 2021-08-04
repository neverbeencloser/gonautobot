package tenancy

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// TenancyTenantsBulkUpdateReader is a Reader for the TenancyTenantsBulkUpdate structure.
type TenancyTenantsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantsBulkUpdateOK creates a TenancyTenantsBulkUpdateOK with default headers values
func NewTenancyTenantsBulkUpdateOK() *TenancyTenantsBulkUpdateOK {
	return &TenancyTenantsBulkUpdateOK{}
}

/* TenancyTenantsBulkUpdateOK describes a response with status code 200, with default header values.

TenancyTenantsBulkUpdateOK tenancy tenants bulk update o k
*/
type TenancyTenantsBulkUpdateOK struct {
	Payload *models.Tenant
}

func (o *TenancyTenantsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenants/][%d] tenancyTenantsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyTenantsBulkUpdateOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *TenancyTenantsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
