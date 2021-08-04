package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimSitesDeleteReader is a Reader for the DcimSitesDelete structure.
type DcimSitesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimSitesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimSitesDeleteNoContent creates a DcimSitesDeleteNoContent with default headers values
func NewDcimSitesDeleteNoContent() *DcimSitesDeleteNoContent {
	return &DcimSitesDeleteNoContent{}
}

/* DcimSitesDeleteNoContent describes a response with status code 204, with default header values.

DcimSitesDeleteNoContent dcim sites delete no content
*/
type DcimSitesDeleteNoContent struct {
}

func (o *DcimSitesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/sites/{id}/][%d] dcimSitesDeleteNoContent ", 204)
}

func (o *DcimSitesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
