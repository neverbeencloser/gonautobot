package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamIPAddressesReadReader is a Reader for the IpamIPAddressesRead structure.
type IpamIPAddressesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPAddressesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamIPAddressesReadOK creates a IpamIPAddressesReadOK with default headers values
func NewIpamIPAddressesReadOK() *IpamIPAddressesReadOK {
	return &IpamIPAddressesReadOK{}
}

/* IpamIPAddressesReadOK describes a response with status code 200, with default header values.

IpamIPAddressesReadOK ipam Ip addresses read o k
*/
type IpamIPAddressesReadOK struct {
	Payload *models.IPAddress
}

func (o *IpamIPAddressesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/ip-addresses/{id}/][%d] ipamIpAddressesReadOK  %+v", 200, o.Payload)
}
func (o *IpamIPAddressesReadOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
