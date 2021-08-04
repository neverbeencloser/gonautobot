package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRegionsBulkDeleteReader is a Reader for the DcimRegionsBulkDelete structure.
type DcimRegionsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRegionsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRegionsBulkDeleteNoContent creates a DcimRegionsBulkDeleteNoContent with default headers values
func NewDcimRegionsBulkDeleteNoContent() *DcimRegionsBulkDeleteNoContent {
	return &DcimRegionsBulkDeleteNoContent{}
}

/* DcimRegionsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRegionsBulkDeleteNoContent dcim regions bulk delete no content
*/
type DcimRegionsBulkDeleteNoContent struct {
}

func (o *DcimRegionsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/regions/][%d] dcimRegionsBulkDeleteNoContent ", 204)
}

func (o *DcimRegionsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
