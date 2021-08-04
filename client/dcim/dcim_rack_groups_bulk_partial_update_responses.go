package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackGroupsBulkPartialUpdateReader is a Reader for the DcimRackGroupsBulkPartialUpdate structure.
type DcimRackGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackGroupsBulkPartialUpdateOK creates a DcimRackGroupsBulkPartialUpdateOK with default headers values
func NewDcimRackGroupsBulkPartialUpdateOK() *DcimRackGroupsBulkPartialUpdateOK {
	return &DcimRackGroupsBulkPartialUpdateOK{}
}

/* DcimRackGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRackGroupsBulkPartialUpdateOK dcim rack groups bulk partial update o k
*/
type DcimRackGroupsBulkPartialUpdateOK struct {
	Payload *models.RackGroup
}

func (o *DcimRackGroupsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-groups/][%d] dcimRackGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRackGroupsBulkPartialUpdateOK) GetPayload() *models.RackGroup {
	return o.Payload
}

func (o *DcimRackGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
