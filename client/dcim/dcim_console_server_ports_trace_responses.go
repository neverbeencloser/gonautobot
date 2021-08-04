package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsoleServerPortsTraceReader is a Reader for the DcimConsoleServerPortsTrace structure.
type DcimConsoleServerPortsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortsTraceOK creates a DcimConsoleServerPortsTraceOK with default headers values
func NewDcimConsoleServerPortsTraceOK() *DcimConsoleServerPortsTraceOK {
	return &DcimConsoleServerPortsTraceOK{}
}

/* DcimConsoleServerPortsTraceOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsTraceOK dcim console server ports trace o k
*/
type DcimConsoleServerPortsTraceOK struct {
	Payload *models.ConsoleServerPort
}

func (o *DcimConsoleServerPortsTraceOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-ports/{id}/trace/][%d] dcimConsoleServerPortsTraceOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortsTraceOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
