package tenancy

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyTenantsDeleteReader is a Reader for the TenancyTenantsDelete structure.
type TenancyTenantsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyTenantsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantsDeleteNoContent creates a TenancyTenantsDeleteNoContent with default headers values
func NewTenancyTenantsDeleteNoContent() *TenancyTenantsDeleteNoContent {
	return &TenancyTenantsDeleteNoContent{}
}

/* TenancyTenantsDeleteNoContent describes a response with status code 204, with default header values.

TenancyTenantsDeleteNoContent tenancy tenants delete no content
*/
type TenancyTenantsDeleteNoContent struct {
}

func (o *TenancyTenantsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenants/{id}/][%d] tenancyTenantsDeleteNoContent ", 204)
}

func (o *TenancyTenantsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
