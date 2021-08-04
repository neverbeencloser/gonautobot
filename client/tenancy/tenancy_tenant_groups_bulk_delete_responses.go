package tenancy

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyTenantGroupsBulkDeleteReader is a Reader for the TenancyTenantGroupsBulkDelete structure.
type TenancyTenantGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyTenantGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantGroupsBulkDeleteNoContent creates a TenancyTenantGroupsBulkDeleteNoContent with default headers values
func NewTenancyTenantGroupsBulkDeleteNoContent() *TenancyTenantGroupsBulkDeleteNoContent {
	return &TenancyTenantGroupsBulkDeleteNoContent{}
}

/* TenancyTenantGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

TenancyTenantGroupsBulkDeleteNoContent tenancy tenant groups bulk delete no content
*/
type TenancyTenantGroupsBulkDeleteNoContent struct {
}

func (o *TenancyTenantGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkDeleteNoContent ", 204)
}

func (o *TenancyTenantGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
