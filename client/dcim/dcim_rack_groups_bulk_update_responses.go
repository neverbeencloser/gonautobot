package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackGroupsBulkUpdateReader is a Reader for the DcimRackGroupsBulkUpdate structure.
type DcimRackGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackGroupsBulkUpdateOK creates a DcimRackGroupsBulkUpdateOK with default headers values
func NewDcimRackGroupsBulkUpdateOK() *DcimRackGroupsBulkUpdateOK {
	return &DcimRackGroupsBulkUpdateOK{}
}

/* DcimRackGroupsBulkUpdateOK describes a response with status code 200, with default header values.

DcimRackGroupsBulkUpdateOK dcim rack groups bulk update o k
*/
type DcimRackGroupsBulkUpdateOK struct {
	Payload *models.RackGroup
}

func (o *DcimRackGroupsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-groups/][%d] dcimRackGroupsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRackGroupsBulkUpdateOK) GetPayload() *models.RackGroup {
	return o.Payload
}

func (o *DcimRackGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
