package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimFrontPortTemplatesDeleteReader is a Reader for the DcimFrontPortTemplatesDelete structure.
type DcimFrontPortTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimFrontPortTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortTemplatesDeleteNoContent creates a DcimFrontPortTemplatesDeleteNoContent with default headers values
func NewDcimFrontPortTemplatesDeleteNoContent() *DcimFrontPortTemplatesDeleteNoContent {
	return &DcimFrontPortTemplatesDeleteNoContent{}
}

/* DcimFrontPortTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimFrontPortTemplatesDeleteNoContent dcim front port templates delete no content
*/
type DcimFrontPortTemplatesDeleteNoContent struct {
}

func (o *DcimFrontPortTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimFrontPortTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
