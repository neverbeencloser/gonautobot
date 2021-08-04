package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimVirtualChassisBulkUpdateReader is a Reader for the DcimVirtualChassisBulkUpdate structure.
type DcimVirtualChassisBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimVirtualChassisBulkUpdateOK creates a DcimVirtualChassisBulkUpdateOK with default headers values
func NewDcimVirtualChassisBulkUpdateOK() *DcimVirtualChassisBulkUpdateOK {
	return &DcimVirtualChassisBulkUpdateOK{}
}

/* DcimVirtualChassisBulkUpdateOK describes a response with status code 200, with default header values.

DcimVirtualChassisBulkUpdateOK dcim virtual chassis bulk update o k
*/
type DcimVirtualChassisBulkUpdateOK struct {
	Payload *models.VirtualChassis
}

func (o *DcimVirtualChassisBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/][%d] dcimVirtualChassisBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimVirtualChassisBulkUpdateOK) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
