package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationClusterTypesBulkDeleteReader is a Reader for the VirtualizationClusterTypesBulkDelete structure.
type VirtualizationClusterTypesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClusterTypesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterTypesBulkDeleteNoContent creates a VirtualizationClusterTypesBulkDeleteNoContent with default headers values
func NewVirtualizationClusterTypesBulkDeleteNoContent() *VirtualizationClusterTypesBulkDeleteNoContent {
	return &VirtualizationClusterTypesBulkDeleteNoContent{}
}

/* VirtualizationClusterTypesBulkDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClusterTypesBulkDeleteNoContent virtualization cluster types bulk delete no content
*/
type VirtualizationClusterTypesBulkDeleteNoContent struct {
}

func (o *VirtualizationClusterTypesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-types/][%d] virtualizationClusterTypesBulkDeleteNoContent ", 204)
}

func (o *VirtualizationClusterTypesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
