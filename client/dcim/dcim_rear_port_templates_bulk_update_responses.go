package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRearPortTemplatesBulkUpdateReader is a Reader for the DcimRearPortTemplatesBulkUpdate structure.
type DcimRearPortTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortTemplatesBulkUpdateOK creates a DcimRearPortTemplatesBulkUpdateOK with default headers values
func NewDcimRearPortTemplatesBulkUpdateOK() *DcimRearPortTemplatesBulkUpdateOK {
	return &DcimRearPortTemplatesBulkUpdateOK{}
}

/* DcimRearPortTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimRearPortTemplatesBulkUpdateOK dcim rear port templates bulk update o k
*/
type DcimRearPortTemplatesBulkUpdateOK struct {
	Payload *models.RearPortTemplate
}

func (o *DcimRearPortTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-port-templates/][%d] dcimRearPortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRearPortTemplatesBulkUpdateOK) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
