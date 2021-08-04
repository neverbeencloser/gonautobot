package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfaceTemplatesBulkPartialUpdateReader is a Reader for the DcimInterfaceTemplatesBulkPartialUpdate structure.
type DcimInterfaceTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesBulkPartialUpdateOK creates a DcimInterfaceTemplatesBulkPartialUpdateOK with default headers values
func NewDcimInterfaceTemplatesBulkPartialUpdateOK() *DcimInterfaceTemplatesBulkPartialUpdateOK {
	return &DcimInterfaceTemplatesBulkPartialUpdateOK{}
}

/* DcimInterfaceTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesBulkPartialUpdateOK dcim interface templates bulk partial update o k
*/
type DcimInterfaceTemplatesBulkPartialUpdateOK struct {
	Payload *models.InterfaceTemplate
}

func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interface-templates/][%d] dcimInterfaceTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
