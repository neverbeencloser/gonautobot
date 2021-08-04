package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsoleServerPortTemplatesBulkUpdateReader is a Reader for the DcimConsoleServerPortTemplatesBulkUpdate structure.
type DcimConsoleServerPortTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesBulkUpdateOK creates a DcimConsoleServerPortTemplatesBulkUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesBulkUpdateOK() *DcimConsoleServerPortTemplatesBulkUpdateOK {
	return &DcimConsoleServerPortTemplatesBulkUpdateOK{}
}

/* DcimConsoleServerPortTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesBulkUpdateOK dcim console server port templates bulk update o k
*/
type DcimConsoleServerPortTemplatesBulkUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/][%d] dcimConsoleServerPortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortTemplatesBulkUpdateOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
