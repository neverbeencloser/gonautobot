package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationClustersBulkDeleteReader is a Reader for the VirtualizationClustersBulkDelete structure.
type VirtualizationClustersBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClustersBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClustersBulkDeleteNoContent creates a VirtualizationClustersBulkDeleteNoContent with default headers values
func NewVirtualizationClustersBulkDeleteNoContent() *VirtualizationClustersBulkDeleteNoContent {
	return &VirtualizationClustersBulkDeleteNoContent{}
}

/* VirtualizationClustersBulkDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClustersBulkDeleteNoContent virtualization clusters bulk delete no content
*/
type VirtualizationClustersBulkDeleteNoContent struct {
}

func (o *VirtualizationClustersBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/clusters/][%d] virtualizationClustersBulkDeleteNoContent ", 204)
}

func (o *VirtualizationClustersBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
