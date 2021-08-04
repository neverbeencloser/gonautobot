package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRackRolesDeleteReader is a Reader for the DcimRackRolesDelete structure.
type DcimRackRolesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRackRolesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackRolesDeleteNoContent creates a DcimRackRolesDeleteNoContent with default headers values
func NewDcimRackRolesDeleteNoContent() *DcimRackRolesDeleteNoContent {
	return &DcimRackRolesDeleteNoContent{}
}

/* DcimRackRolesDeleteNoContent describes a response with status code 204, with default header values.

DcimRackRolesDeleteNoContent dcim rack roles delete no content
*/
type DcimRackRolesDeleteNoContent struct {
}

func (o *DcimRackRolesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-roles/{id}/][%d] dcimRackRolesDeleteNoContent ", 204)
}

func (o *DcimRackRolesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
