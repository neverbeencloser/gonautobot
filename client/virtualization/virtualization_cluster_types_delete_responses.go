package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationClusterTypesDeleteReader is a Reader for the VirtualizationClusterTypesDelete structure.
type VirtualizationClusterTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClusterTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterTypesDeleteNoContent creates a VirtualizationClusterTypesDeleteNoContent with default headers values
func NewVirtualizationClusterTypesDeleteNoContent() *VirtualizationClusterTypesDeleteNoContent {
	return &VirtualizationClusterTypesDeleteNoContent{}
}

/* VirtualizationClusterTypesDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClusterTypesDeleteNoContent virtualization cluster types delete no content
*/
type VirtualizationClusterTypesDeleteNoContent struct {
}

func (o *VirtualizationClusterTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesDeleteNoContent ", 204)
}

func (o *VirtualizationClusterTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
