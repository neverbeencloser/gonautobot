package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimFrontPortsDeleteReader is a Reader for the DcimFrontPortsDelete structure.
type DcimFrontPortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimFrontPortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortsDeleteNoContent creates a DcimFrontPortsDeleteNoContent with default headers values
func NewDcimFrontPortsDeleteNoContent() *DcimFrontPortsDeleteNoContent {
	return &DcimFrontPortsDeleteNoContent{}
}

/* DcimFrontPortsDeleteNoContent describes a response with status code 204, with default header values.

DcimFrontPortsDeleteNoContent dcim front ports delete no content
*/
type DcimFrontPortsDeleteNoContent struct {
}

func (o *DcimFrontPortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/front-ports/{id}/][%d] dcimFrontPortsDeleteNoContent ", 204)
}

func (o *DcimFrontPortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
