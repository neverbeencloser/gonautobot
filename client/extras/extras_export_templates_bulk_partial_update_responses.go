package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasExportTemplatesBulkPartialUpdateReader is a Reader for the ExtrasExportTemplatesBulkPartialUpdate structure.
type ExtrasExportTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasExportTemplatesBulkPartialUpdateOK creates a ExtrasExportTemplatesBulkPartialUpdateOK with default headers values
func NewExtrasExportTemplatesBulkPartialUpdateOK() *ExtrasExportTemplatesBulkPartialUpdateOK {
	return &ExtrasExportTemplatesBulkPartialUpdateOK{}
}

/* ExtrasExportTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasExportTemplatesBulkPartialUpdateOK extras export templates bulk partial update o k
*/
type ExtrasExportTemplatesBulkPartialUpdateOK struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/][%d] extrasExportTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasExportTemplatesBulkPartialUpdateOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
