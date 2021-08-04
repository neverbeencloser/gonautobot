package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationClusterGroupsDeleteReader is a Reader for the VirtualizationClusterGroupsDelete structure.
type VirtualizationClusterGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClusterGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsDeleteNoContent creates a VirtualizationClusterGroupsDeleteNoContent with default headers values
func NewVirtualizationClusterGroupsDeleteNoContent() *VirtualizationClusterGroupsDeleteNoContent {
	return &VirtualizationClusterGroupsDeleteNoContent{}
}

/* VirtualizationClusterGroupsDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClusterGroupsDeleteNoContent virtualization cluster groups delete no content
*/
type VirtualizationClusterGroupsDeleteNoContent struct {
}

func (o *VirtualizationClusterGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsDeleteNoContent ", 204)
}

func (o *VirtualizationClusterGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
