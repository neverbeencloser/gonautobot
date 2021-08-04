package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsoleServerPortsReadReader is a Reader for the DcimConsoleServerPortsRead structure.
type DcimConsoleServerPortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortsReadOK creates a DcimConsoleServerPortsReadOK with default headers values
func NewDcimConsoleServerPortsReadOK() *DcimConsoleServerPortsReadOK {
	return &DcimConsoleServerPortsReadOK{}
}

/* DcimConsoleServerPortsReadOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsReadOK dcim console server ports read o k
*/
type DcimConsoleServerPortsReadOK struct {
	Payload *models.ConsoleServerPort
}

func (o *DcimConsoleServerPortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsReadOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortsReadOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
