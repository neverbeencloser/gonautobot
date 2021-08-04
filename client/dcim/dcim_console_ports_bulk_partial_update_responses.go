package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortsBulkPartialUpdateReader is a Reader for the DcimConsolePortsBulkPartialUpdate structure.
type DcimConsolePortsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortsBulkPartialUpdateOK creates a DcimConsolePortsBulkPartialUpdateOK with default headers values
func NewDcimConsolePortsBulkPartialUpdateOK() *DcimConsolePortsBulkPartialUpdateOK {
	return &DcimConsolePortsBulkPartialUpdateOK{}
}

/* DcimConsolePortsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortsBulkPartialUpdateOK dcim console ports bulk partial update o k
*/
type DcimConsolePortsBulkPartialUpdateOK struct {
	Payload *models.ConsolePort
}

func (o *DcimConsolePortsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-ports/][%d] dcimConsolePortsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortsBulkPartialUpdateOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
