package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackRolesBulkUpdateReader is a Reader for the DcimRackRolesBulkUpdate structure.
type DcimRackRolesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackRolesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackRolesBulkUpdateOK creates a DcimRackRolesBulkUpdateOK with default headers values
func NewDcimRackRolesBulkUpdateOK() *DcimRackRolesBulkUpdateOK {
	return &DcimRackRolesBulkUpdateOK{}
}

/* DcimRackRolesBulkUpdateOK describes a response with status code 200, with default header values.

DcimRackRolesBulkUpdateOK dcim rack roles bulk update o k
*/
type DcimRackRolesBulkUpdateOK struct {
	Payload *models.RackRole
}

func (o *DcimRackRolesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/][%d] dcimRackRolesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRackRolesBulkUpdateOK) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
