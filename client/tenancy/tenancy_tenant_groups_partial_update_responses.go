package tenancy

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// TenancyTenantGroupsPartialUpdateReader is a Reader for the TenancyTenantGroupsPartialUpdate structure.
type TenancyTenantGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantGroupsPartialUpdateOK creates a TenancyTenantGroupsPartialUpdateOK with default headers values
func NewTenancyTenantGroupsPartialUpdateOK() *TenancyTenantGroupsPartialUpdateOK {
	return &TenancyTenantGroupsPartialUpdateOK{}
}

/* TenancyTenantGroupsPartialUpdateOK describes a response with status code 200, with default header values.

TenancyTenantGroupsPartialUpdateOK tenancy tenant groups partial update o k
*/
type TenancyTenantGroupsPartialUpdateOK struct {
	Payload *models.TenantGroup
}

func (o *TenancyTenantGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyTenantGroupsPartialUpdateOK) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
