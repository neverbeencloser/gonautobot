package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortTemplatesBulkUpdateReader is a Reader for the DcimConsolePortTemplatesBulkUpdate structure.
type DcimConsolePortTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesBulkUpdateOK creates a DcimConsolePortTemplatesBulkUpdateOK with default headers values
func NewDcimConsolePortTemplatesBulkUpdateOK() *DcimConsolePortTemplatesBulkUpdateOK {
	return &DcimConsolePortTemplatesBulkUpdateOK{}
}

/* DcimConsolePortTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesBulkUpdateOK dcim console port templates bulk update o k
*/
type DcimConsolePortTemplatesBulkUpdateOK struct {
	Payload *models.ConsolePortTemplate
}

func (o *DcimConsolePortTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/][%d] dcimConsolePortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortTemplatesBulkUpdateOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
