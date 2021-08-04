package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVlansReadReader is a Reader for the IpamVlansRead structure.
type IpamVlansReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlansReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlansReadOK creates a IpamVlansReadOK with default headers values
func NewIpamVlansReadOK() *IpamVlansReadOK {
	return &IpamVlansReadOK{}
}

/* IpamVlansReadOK describes a response with status code 200, with default header values.

IpamVlansReadOK ipam vlans read o k
*/
type IpamVlansReadOK struct {
	Payload *models.VLAN
}

func (o *IpamVlansReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/vlans/{id}/][%d] ipamVlansReadOK  %+v", 200, o.Payload)
}
func (o *IpamVlansReadOK) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
