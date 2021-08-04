package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRearPortTemplatesBulkPartialUpdateReader is a Reader for the DcimRearPortTemplatesBulkPartialUpdate structure.
type DcimRearPortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortTemplatesBulkPartialUpdateOK creates a DcimRearPortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimRearPortTemplatesBulkPartialUpdateOK() *DcimRearPortTemplatesBulkPartialUpdateOK {
	return &DcimRearPortTemplatesBulkPartialUpdateOK{}
}

/* DcimRearPortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRearPortTemplatesBulkPartialUpdateOK dcim rear port templates bulk partial update o k
*/
type DcimRearPortTemplatesBulkPartialUpdateOK struct {
	Payload *models.RearPortTemplate
}

func (o *DcimRearPortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rear-port-templates/][%d] dcimRearPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRearPortTemplatesBulkPartialUpdateOK) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
