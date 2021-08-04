package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortTemplatesBulkPartialUpdateReader is a Reader for the DcimConsolePortTemplatesBulkPartialUpdate structure.
type DcimConsolePortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesBulkPartialUpdateOK creates a DcimConsolePortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimConsolePortTemplatesBulkPartialUpdateOK() *DcimConsolePortTemplatesBulkPartialUpdateOK {
	return &DcimConsolePortTemplatesBulkPartialUpdateOK{}
}

/* DcimConsolePortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesBulkPartialUpdateOK dcim console port templates bulk partial update o k
*/
type DcimConsolePortTemplatesBulkPartialUpdateOK struct {
	Payload *models.ConsolePortTemplate
}

func (o *DcimConsolePortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-port-templates/][%d] dcimConsolePortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortTemplatesBulkPartialUpdateOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
