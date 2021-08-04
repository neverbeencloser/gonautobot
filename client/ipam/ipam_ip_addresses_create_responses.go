package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamIPAddressesCreateReader is a Reader for the IpamIPAddressesCreate structure.
type IpamIPAddressesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamIPAddressesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamIPAddressesCreateCreated creates a IpamIPAddressesCreateCreated with default headers values
func NewIpamIPAddressesCreateCreated() *IpamIPAddressesCreateCreated {
	return &IpamIPAddressesCreateCreated{}
}

/* IpamIPAddressesCreateCreated describes a response with status code 201, with default header values.

IpamIPAddressesCreateCreated ipam Ip addresses create created
*/
type IpamIPAddressesCreateCreated struct {
	Payload *models.IPAddress
}

func (o *IpamIPAddressesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/ip-addresses/][%d] ipamIpAddressesCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamIPAddressesCreateCreated) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
