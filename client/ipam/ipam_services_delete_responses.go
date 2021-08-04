package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamServicesDeleteReader is a Reader for the IpamServicesDelete structure.
type IpamServicesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamServicesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamServicesDeleteNoContent creates a IpamServicesDeleteNoContent with default headers values
func NewIpamServicesDeleteNoContent() *IpamServicesDeleteNoContent {
	return &IpamServicesDeleteNoContent{}
}

/* IpamServicesDeleteNoContent describes a response with status code 204, with default header values.

IpamServicesDeleteNoContent ipam services delete no content
*/
type IpamServicesDeleteNoContent struct {
}

func (o *IpamServicesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/services/{id}/][%d] ipamServicesDeleteNoContent ", 204)
}

func (o *IpamServicesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
