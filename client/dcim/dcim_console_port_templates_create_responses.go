package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortTemplatesCreateReader is a Reader for the DcimConsolePortTemplatesCreate structure.
type DcimConsolePortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimConsolePortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesCreateCreated creates a DcimConsolePortTemplatesCreateCreated with default headers values
func NewDcimConsolePortTemplatesCreateCreated() *DcimConsolePortTemplatesCreateCreated {
	return &DcimConsolePortTemplatesCreateCreated{}
}

/* DcimConsolePortTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimConsolePortTemplatesCreateCreated dcim console port templates create created
*/
type DcimConsolePortTemplatesCreateCreated struct {
	Payload *models.ConsolePortTemplate
}

func (o *DcimConsolePortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/console-port-templates/][%d] dcimConsolePortTemplatesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimConsolePortTemplatesCreateCreated) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
