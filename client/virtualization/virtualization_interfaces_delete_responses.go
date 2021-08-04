package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationInterfacesDeleteReader is a Reader for the VirtualizationInterfacesDelete structure.
type VirtualizationInterfacesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationInterfacesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationInterfacesDeleteNoContent creates a VirtualizationInterfacesDeleteNoContent with default headers values
func NewVirtualizationInterfacesDeleteNoContent() *VirtualizationInterfacesDeleteNoContent {
	return &VirtualizationInterfacesDeleteNoContent{}
}

/* VirtualizationInterfacesDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationInterfacesDeleteNoContent virtualization interfaces delete no content
*/
type VirtualizationInterfacesDeleteNoContent struct {
}

func (o *VirtualizationInterfacesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/interfaces/{id}/][%d] virtualizationInterfacesDeleteNoContent ", 204)
}

func (o *VirtualizationInterfacesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
