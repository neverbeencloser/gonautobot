package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRackGroupsDeleteReader is a Reader for the DcimRackGroupsDelete structure.
type DcimRackGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRackGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackGroupsDeleteNoContent creates a DcimRackGroupsDeleteNoContent with default headers values
func NewDcimRackGroupsDeleteNoContent() *DcimRackGroupsDeleteNoContent {
	return &DcimRackGroupsDeleteNoContent{}
}

/* DcimRackGroupsDeleteNoContent describes a response with status code 204, with default header values.

DcimRackGroupsDeleteNoContent dcim rack groups delete no content
*/
type DcimRackGroupsDeleteNoContent struct {
}

func (o *DcimRackGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-groups/{id}/][%d] dcimRackGroupsDeleteNoContent ", 204)
}

func (o *DcimRackGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
