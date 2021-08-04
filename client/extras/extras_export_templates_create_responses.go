package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasExportTemplatesCreateReader is a Reader for the ExtrasExportTemplatesCreate structure.
type ExtrasExportTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasExportTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasExportTemplatesCreateCreated creates a ExtrasExportTemplatesCreateCreated with default headers values
func NewExtrasExportTemplatesCreateCreated() *ExtrasExportTemplatesCreateCreated {
	return &ExtrasExportTemplatesCreateCreated{}
}

/* ExtrasExportTemplatesCreateCreated describes a response with status code 201, with default header values.

ExtrasExportTemplatesCreateCreated extras export templates create created
*/
type ExtrasExportTemplatesCreateCreated struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/export-templates/][%d] extrasExportTemplatesCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasExportTemplatesCreateCreated) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
