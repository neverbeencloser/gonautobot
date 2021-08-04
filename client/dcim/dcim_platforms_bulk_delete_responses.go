package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPlatformsBulkDeleteReader is a Reader for the DcimPlatformsBulkDelete structure.
type DcimPlatformsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPlatformsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPlatformsBulkDeleteNoContent creates a DcimPlatformsBulkDeleteNoContent with default headers values
func NewDcimPlatformsBulkDeleteNoContent() *DcimPlatformsBulkDeleteNoContent {
	return &DcimPlatformsBulkDeleteNoContent{}
}

/* DcimPlatformsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPlatformsBulkDeleteNoContent dcim platforms bulk delete no content
*/
type DcimPlatformsBulkDeleteNoContent struct {
}

func (o *DcimPlatformsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/][%d] dcimPlatformsBulkDeleteNoContent ", 204)
}

func (o *DcimPlatformsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
