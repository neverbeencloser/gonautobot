package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortsReadReader is a Reader for the DcimConsolePortsRead structure.
type DcimConsolePortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortsReadOK creates a DcimConsolePortsReadOK with default headers values
func NewDcimConsolePortsReadOK() *DcimConsolePortsReadOK {
	return &DcimConsolePortsReadOK{}
}

/* DcimConsolePortsReadOK describes a response with status code 200, with default header values.

DcimConsolePortsReadOK dcim console ports read o k
*/
type DcimConsolePortsReadOK struct {
	Payload *models.ConsolePort
}

func (o *DcimConsolePortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcimConsolePortsReadOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortsReadOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
