package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasExportTemplatesBulkUpdateReader is a Reader for the ExtrasExportTemplatesBulkUpdate structure.
type ExtrasExportTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasExportTemplatesBulkUpdateOK creates a ExtrasExportTemplatesBulkUpdateOK with default headers values
func NewExtrasExportTemplatesBulkUpdateOK() *ExtrasExportTemplatesBulkUpdateOK {
	return &ExtrasExportTemplatesBulkUpdateOK{}
}

/* ExtrasExportTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasExportTemplatesBulkUpdateOK extras export templates bulk update o k
*/
type ExtrasExportTemplatesBulkUpdateOK struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/export-templates/][%d] extrasExportTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasExportTemplatesBulkUpdateOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
