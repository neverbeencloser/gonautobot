package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVlansPartialUpdateReader is a Reader for the IpamVlansPartialUpdate structure.
type IpamVlansPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlansPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlansPartialUpdateOK creates a IpamVlansPartialUpdateOK with default headers values
func NewIpamVlansPartialUpdateOK() *IpamVlansPartialUpdateOK {
	return &IpamVlansPartialUpdateOK{}
}

/* IpamVlansPartialUpdateOK describes a response with status code 200, with default header values.

IpamVlansPartialUpdateOK ipam vlans partial update o k
*/
type IpamVlansPartialUpdateOK struct {
	Payload *models.VLAN
}

func (o *IpamVlansPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlans/{id}/][%d] ipamVlansPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlansPartialUpdateOK) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
