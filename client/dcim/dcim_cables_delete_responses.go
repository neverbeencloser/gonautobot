package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimCablesDeleteReader is a Reader for the DcimCablesDelete structure.
type DcimCablesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimCablesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimCablesDeleteNoContent creates a DcimCablesDeleteNoContent with default headers values
func NewDcimCablesDeleteNoContent() *DcimCablesDeleteNoContent {
	return &DcimCablesDeleteNoContent{}
}

/* DcimCablesDeleteNoContent describes a response with status code 204, with default header values.

DcimCablesDeleteNoContent dcim cables delete no content
*/
type DcimCablesDeleteNoContent struct {
}

func (o *DcimCablesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cables/{id}/][%d] dcimCablesDeleteNoContent ", 204)
}

func (o *DcimCablesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
