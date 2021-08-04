package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasExportTemplatesPartialUpdateReader is a Reader for the ExtrasExportTemplatesPartialUpdate structure.
type ExtrasExportTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasExportTemplatesPartialUpdateOK creates a ExtrasExportTemplatesPartialUpdateOK with default headers values
func NewExtrasExportTemplatesPartialUpdateOK() *ExtrasExportTemplatesPartialUpdateOK {
	return &ExtrasExportTemplatesPartialUpdateOK{}
}

/* ExtrasExportTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasExportTemplatesPartialUpdateOK extras export templates partial update o k
*/
type ExtrasExportTemplatesPartialUpdateOK struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/{id}/][%d] extrasExportTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasExportTemplatesPartialUpdateOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
