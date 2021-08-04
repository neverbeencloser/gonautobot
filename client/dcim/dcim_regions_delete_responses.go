package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRegionsDeleteReader is a Reader for the DcimRegionsDelete structure.
type DcimRegionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRegionsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRegionsDeleteNoContent creates a DcimRegionsDeleteNoContent with default headers values
func NewDcimRegionsDeleteNoContent() *DcimRegionsDeleteNoContent {
	return &DcimRegionsDeleteNoContent{}
}

/* DcimRegionsDeleteNoContent describes a response with status code 204, with default header values.

DcimRegionsDeleteNoContent dcim regions delete no content
*/
type DcimRegionsDeleteNoContent struct {
}

func (o *DcimRegionsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/regions/{id}/][%d] dcimRegionsDeleteNoContent ", 204)
}

func (o *DcimRegionsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
