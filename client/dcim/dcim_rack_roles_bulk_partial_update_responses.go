package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackRolesBulkPartialUpdateReader is a Reader for the DcimRackRolesBulkPartialUpdate structure.
type DcimRackRolesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackRolesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackRolesBulkPartialUpdateOK creates a DcimRackRolesBulkPartialUpdateOK with default headers values
func NewDcimRackRolesBulkPartialUpdateOK() *DcimRackRolesBulkPartialUpdateOK {
	return &DcimRackRolesBulkPartialUpdateOK{}
}

/* DcimRackRolesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRackRolesBulkPartialUpdateOK dcim rack roles bulk partial update o k
*/
type DcimRackRolesBulkPartialUpdateOK struct {
	Payload *models.RackRole
}

func (o *DcimRackRolesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-roles/][%d] dcimRackRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRackRolesBulkPartialUpdateOK) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
