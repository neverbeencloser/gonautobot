package tenancy

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// TenancyTenantGroupsUpdateReader is a Reader for the TenancyTenantGroupsUpdate structure.
type TenancyTenantGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantGroupsUpdateOK creates a TenancyTenantGroupsUpdateOK with default headers values
func NewTenancyTenantGroupsUpdateOK() *TenancyTenantGroupsUpdateOK {
	return &TenancyTenantGroupsUpdateOK{}
}

/* TenancyTenantGroupsUpdateOK describes a response with status code 200, with default header values.

TenancyTenantGroupsUpdateOK tenancy tenant groups update o k
*/
type TenancyTenantGroupsUpdateOK struct {
	Payload *models.TenantGroup
}

func (o *TenancyTenantGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyTenantGroupsUpdateOK) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
