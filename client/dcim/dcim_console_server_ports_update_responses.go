package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsoleServerPortsUpdateReader is a Reader for the DcimConsoleServerPortsUpdate structure.
type DcimConsoleServerPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortsUpdateOK creates a DcimConsoleServerPortsUpdateOK with default headers values
func NewDcimConsoleServerPortsUpdateOK() *DcimConsoleServerPortsUpdateOK {
	return &DcimConsoleServerPortsUpdateOK{}
}

/* DcimConsoleServerPortsUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsUpdateOK dcim console server ports update o k
*/
type DcimConsoleServerPortsUpdateOK struct {
	Payload *models.ConsoleServerPort
}

func (o *DcimConsoleServerPortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortsUpdateOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
