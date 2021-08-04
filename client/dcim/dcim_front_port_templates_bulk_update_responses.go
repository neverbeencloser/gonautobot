package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortTemplatesBulkUpdateReader is a Reader for the DcimFrontPortTemplatesBulkUpdate structure.
type DcimFrontPortTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortTemplatesBulkUpdateOK creates a DcimFrontPortTemplatesBulkUpdateOK with default headers values
func NewDcimFrontPortTemplatesBulkUpdateOK() *DcimFrontPortTemplatesBulkUpdateOK {
	return &DcimFrontPortTemplatesBulkUpdateOK{}
}

/* DcimFrontPortTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesBulkUpdateOK dcim front port templates bulk update o k
*/
type DcimFrontPortTemplatesBulkUpdateOK struct {
	Payload *models.FrontPortTemplate
}

func (o *DcimFrontPortTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/][%d] dcimFrontPortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortTemplatesBulkUpdateOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
