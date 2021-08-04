package tenancy

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyTenantsBulkDeleteReader is a Reader for the TenancyTenantsBulkDelete structure.
type TenancyTenantsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyTenantsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantsBulkDeleteNoContent creates a TenancyTenantsBulkDeleteNoContent with default headers values
func NewTenancyTenantsBulkDeleteNoContent() *TenancyTenantsBulkDeleteNoContent {
	return &TenancyTenantsBulkDeleteNoContent{}
}

/* TenancyTenantsBulkDeleteNoContent describes a response with status code 204, with default header values.

TenancyTenantsBulkDeleteNoContent tenancy tenants bulk delete no content
*/
type TenancyTenantsBulkDeleteNoContent struct {
}

func (o *TenancyTenantsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenants/][%d] tenancyTenantsBulkDeleteNoContent ", 204)
}

func (o *TenancyTenantsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
