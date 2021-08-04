package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationClusterGroupsBulkDeleteReader is a Reader for the VirtualizationClusterGroupsBulkDelete structure.
type VirtualizationClusterGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClusterGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsBulkDeleteNoContent creates a VirtualizationClusterGroupsBulkDeleteNoContent with default headers values
func NewVirtualizationClusterGroupsBulkDeleteNoContent() *VirtualizationClusterGroupsBulkDeleteNoContent {
	return &VirtualizationClusterGroupsBulkDeleteNoContent{}
}

/* VirtualizationClusterGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClusterGroupsBulkDeleteNoContent virtualization cluster groups bulk delete no content
*/
type VirtualizationClusterGroupsBulkDeleteNoContent struct {
}

func (o *VirtualizationClusterGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-groups/][%d] virtualizationClusterGroupsBulkDeleteNoContent ", 204)
}

func (o *VirtualizationClusterGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
