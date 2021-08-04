package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortTemplatesReadReader is a Reader for the DcimConsolePortTemplatesRead structure.
type DcimConsolePortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesReadOK creates a DcimConsolePortTemplatesReadOK with default headers values
func NewDcimConsolePortTemplatesReadOK() *DcimConsolePortTemplatesReadOK {
	return &DcimConsolePortTemplatesReadOK{}
}

/* DcimConsolePortTemplatesReadOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesReadOK dcim console port templates read o k
*/
type DcimConsolePortTemplatesReadOK struct {
	Payload *models.ConsolePortTemplate
}

func (o *DcimConsolePortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesReadOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortTemplatesReadOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
