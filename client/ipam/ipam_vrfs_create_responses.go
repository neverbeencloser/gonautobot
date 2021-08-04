package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVrfsCreateReader is a Reader for the IpamVrfsCreate structure.
type IpamVrfsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamVrfsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVrfsCreateCreated creates a IpamVrfsCreateCreated with default headers values
func NewIpamVrfsCreateCreated() *IpamVrfsCreateCreated {
	return &IpamVrfsCreateCreated{}
}

/* IpamVrfsCreateCreated describes a response with status code 201, with default header values.

IpamVrfsCreateCreated ipam vrfs create created
*/
type IpamVrfsCreateCreated struct {
	Payload *models.VRF
}

func (o *IpamVrfsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vrfs/][%d] ipamVrfsCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamVrfsCreateCreated) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
