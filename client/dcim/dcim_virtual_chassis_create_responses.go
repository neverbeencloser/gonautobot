package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimVirtualChassisCreateReader is a Reader for the DcimVirtualChassisCreate structure.
type DcimVirtualChassisCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimVirtualChassisCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimVirtualChassisCreateCreated creates a DcimVirtualChassisCreateCreated with default headers values
func NewDcimVirtualChassisCreateCreated() *DcimVirtualChassisCreateCreated {
	return &DcimVirtualChassisCreateCreated{}
}

/* DcimVirtualChassisCreateCreated describes a response with status code 201, with default header values.

DcimVirtualChassisCreateCreated dcim virtual chassis create created
*/
type DcimVirtualChassisCreateCreated struct {
	Payload *models.VirtualChassis
}

func (o *DcimVirtualChassisCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/virtual-chassis/][%d] dcimVirtualChassisCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimVirtualChassisCreateCreated) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
