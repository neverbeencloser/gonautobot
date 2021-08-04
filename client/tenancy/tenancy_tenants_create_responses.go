package tenancy

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// TenancyTenantsCreateReader is a Reader for the TenancyTenantsCreate structure.
type TenancyTenantsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyTenantsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantsCreateCreated creates a TenancyTenantsCreateCreated with default headers values
func NewTenancyTenantsCreateCreated() *TenancyTenantsCreateCreated {
	return &TenancyTenantsCreateCreated{}
}

/* TenancyTenantsCreateCreated describes a response with status code 201, with default header values.

TenancyTenantsCreateCreated tenancy tenants create created
*/
type TenancyTenantsCreateCreated struct {
	Payload *models.Tenant
}

func (o *TenancyTenantsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /tenancy/tenants/][%d] tenancyTenantsCreateCreated  %+v", 201, o.Payload)
}
func (o *TenancyTenantsCreateCreated) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *TenancyTenantsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
