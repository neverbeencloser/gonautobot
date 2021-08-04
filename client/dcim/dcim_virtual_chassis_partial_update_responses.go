package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimVirtualChassisPartialUpdateReader is a Reader for the DcimVirtualChassisPartialUpdate structure.
type DcimVirtualChassisPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimVirtualChassisPartialUpdateOK creates a DcimVirtualChassisPartialUpdateOK with default headers values
func NewDcimVirtualChassisPartialUpdateOK() *DcimVirtualChassisPartialUpdateOK {
	return &DcimVirtualChassisPartialUpdateOK{}
}

/* DcimVirtualChassisPartialUpdateOK describes a response with status code 200, with default header values.

DcimVirtualChassisPartialUpdateOK dcim virtual chassis partial update o k
*/
type DcimVirtualChassisPartialUpdateOK struct {
	Payload *models.VirtualChassis
}

func (o *DcimVirtualChassisPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimVirtualChassisPartialUpdateOK) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
