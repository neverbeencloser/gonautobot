package tenancy

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyTenantGroupsDeleteReader is a Reader for the TenancyTenantGroupsDelete structure.
type TenancyTenantGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyTenantGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantGroupsDeleteNoContent creates a TenancyTenantGroupsDeleteNoContent with default headers values
func NewTenancyTenantGroupsDeleteNoContent() *TenancyTenantGroupsDeleteNoContent {
	return &TenancyTenantGroupsDeleteNoContent{}
}

/* TenancyTenantGroupsDeleteNoContent describes a response with status code 204, with default header values.

TenancyTenantGroupsDeleteNoContent tenancy tenant groups delete no content
*/
type TenancyTenantGroupsDeleteNoContent struct {
}

func (o *TenancyTenantGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsDeleteNoContent ", 204)
}

func (o *TenancyTenantGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
