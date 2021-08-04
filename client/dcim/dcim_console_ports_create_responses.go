package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortsCreateReader is a Reader for the DcimConsolePortsCreate structure.
type DcimConsolePortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimConsolePortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortsCreateCreated creates a DcimConsolePortsCreateCreated with default headers values
func NewDcimConsolePortsCreateCreated() *DcimConsolePortsCreateCreated {
	return &DcimConsolePortsCreateCreated{}
}

/* DcimConsolePortsCreateCreated describes a response with status code 201, with default header values.

DcimConsolePortsCreateCreated dcim console ports create created
*/
type DcimConsolePortsCreateCreated struct {
	Payload *models.ConsolePort
}

func (o *DcimConsolePortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/console-ports/][%d] dcimConsolePortsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimConsolePortsCreateCreated) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
