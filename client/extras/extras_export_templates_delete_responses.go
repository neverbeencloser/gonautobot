package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasExportTemplatesDeleteReader is a Reader for the ExtrasExportTemplatesDelete structure.
type ExtrasExportTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasExportTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasExportTemplatesDeleteNoContent creates a ExtrasExportTemplatesDeleteNoContent with default headers values
func NewExtrasExportTemplatesDeleteNoContent() *ExtrasExportTemplatesDeleteNoContent {
	return &ExtrasExportTemplatesDeleteNoContent{}
}

/* ExtrasExportTemplatesDeleteNoContent describes a response with status code 204, with default header values.

ExtrasExportTemplatesDeleteNoContent extras export templates delete no content
*/
type ExtrasExportTemplatesDeleteNoContent struct {
}

func (o *ExtrasExportTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/{id}/][%d] extrasExportTemplatesDeleteNoContent ", 204)
}

func (o *ExtrasExportTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
