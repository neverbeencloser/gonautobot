package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationClustersDeleteReader is a Reader for the VirtualizationClustersDelete structure.
type VirtualizationClustersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClustersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClustersDeleteNoContent creates a VirtualizationClustersDeleteNoContent with default headers values
func NewVirtualizationClustersDeleteNoContent() *VirtualizationClustersDeleteNoContent {
	return &VirtualizationClustersDeleteNoContent{}
}

/* VirtualizationClustersDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClustersDeleteNoContent virtualization clusters delete no content
*/
type VirtualizationClustersDeleteNoContent struct {
}

func (o *VirtualizationClustersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/clusters/{id}/][%d] virtualizationClustersDeleteNoContent ", 204)
}

func (o *VirtualizationClustersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
