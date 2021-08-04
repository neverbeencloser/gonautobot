package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimManufacturersBulkDeleteReader is a Reader for the DcimManufacturersBulkDelete structure.
type DcimManufacturersBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimManufacturersBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimManufacturersBulkDeleteNoContent creates a DcimManufacturersBulkDeleteNoContent with default headers values
func NewDcimManufacturersBulkDeleteNoContent() *DcimManufacturersBulkDeleteNoContent {
	return &DcimManufacturersBulkDeleteNoContent{}
}

/* DcimManufacturersBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimManufacturersBulkDeleteNoContent dcim manufacturers bulk delete no content
*/
type DcimManufacturersBulkDeleteNoContent struct {
}

func (o *DcimManufacturersBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/manufacturers/][%d] dcimManufacturersBulkDeleteNoContent ", 204)
}

func (o *DcimManufacturersBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
