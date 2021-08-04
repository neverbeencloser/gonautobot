package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimVirtualChassisUpdateReader is a Reader for the DcimVirtualChassisUpdate structure.
type DcimVirtualChassisUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimVirtualChassisUpdateOK creates a DcimVirtualChassisUpdateOK with default headers values
func NewDcimVirtualChassisUpdateOK() *DcimVirtualChassisUpdateOK {
	return &DcimVirtualChassisUpdateOK{}
}

/* DcimVirtualChassisUpdateOK describes a response with status code 200, with default header values.

DcimVirtualChassisUpdateOK dcim virtual chassis update o k
*/
type DcimVirtualChassisUpdateOK struct {
	Payload *models.VirtualChassis
}

func (o *DcimVirtualChassisUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimVirtualChassisUpdateOK) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
