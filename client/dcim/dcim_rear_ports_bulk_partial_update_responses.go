package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRearPortsBulkPartialUpdateReader is a Reader for the DcimRearPortsBulkPartialUpdate structure.
type DcimRearPortsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortsBulkPartialUpdateOK creates a DcimRearPortsBulkPartialUpdateOK with default headers values
func NewDcimRearPortsBulkPartialUpdateOK() *DcimRearPortsBulkPartialUpdateOK {
	return &DcimRearPortsBulkPartialUpdateOK{}
}

/* DcimRearPortsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRearPortsBulkPartialUpdateOK dcim rear ports bulk partial update o k
*/
type DcimRearPortsBulkPartialUpdateOK struct {
	Payload *models.RearPort
}

func (o *DcimRearPortsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rear-ports/][%d] dcimRearPortsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRearPortsBulkPartialUpdateOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
