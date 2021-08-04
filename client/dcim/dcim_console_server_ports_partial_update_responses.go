package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsoleServerPortsPartialUpdateReader is a Reader for the DcimConsoleServerPortsPartialUpdate structure.
type DcimConsoleServerPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortsPartialUpdateOK creates a DcimConsoleServerPortsPartialUpdateOK with default headers values
func NewDcimConsoleServerPortsPartialUpdateOK() *DcimConsoleServerPortsPartialUpdateOK {
	return &DcimConsoleServerPortsPartialUpdateOK{}
}

/* DcimConsoleServerPortsPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsPartialUpdateOK dcim console server ports partial update o k
*/
type DcimConsoleServerPortsPartialUpdateOK struct {
	Payload *models.ConsoleServerPort
}

func (o *DcimConsoleServerPortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortsPartialUpdateOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
