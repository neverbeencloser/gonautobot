package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVrfsDeleteReader is a Reader for the IpamVrfsDelete structure.
type IpamVrfsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVrfsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVrfsDeleteNoContent creates a IpamVrfsDeleteNoContent with default headers values
func NewIpamVrfsDeleteNoContent() *IpamVrfsDeleteNoContent {
	return &IpamVrfsDeleteNoContent{}
}

/* IpamVrfsDeleteNoContent describes a response with status code 204, with default header values.

IpamVrfsDeleteNoContent ipam vrfs delete no content
*/
type IpamVrfsDeleteNoContent struct {
}

func (o *IpamVrfsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/{id}/][%d] ipamVrfsDeleteNoContent ", 204)
}

func (o *IpamVrfsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
