package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimSitesBulkDeleteReader is a Reader for the DcimSitesBulkDelete structure.
type DcimSitesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimSitesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimSitesBulkDeleteNoContent creates a DcimSitesBulkDeleteNoContent with default headers values
func NewDcimSitesBulkDeleteNoContent() *DcimSitesBulkDeleteNoContent {
	return &DcimSitesBulkDeleteNoContent{}
}

/* DcimSitesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimSitesBulkDeleteNoContent dcim sites bulk delete no content
*/
type DcimSitesBulkDeleteNoContent struct {
}

func (o *DcimSitesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/sites/][%d] dcimSitesBulkDeleteNoContent ", 204)
}

func (o *DcimSitesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
