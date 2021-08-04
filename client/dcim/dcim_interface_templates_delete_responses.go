package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimInterfaceTemplatesDeleteReader is a Reader for the DcimInterfaceTemplatesDelete structure.
type DcimInterfaceTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInterfaceTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesDeleteNoContent creates a DcimInterfaceTemplatesDeleteNoContent with default headers values
func NewDcimInterfaceTemplatesDeleteNoContent() *DcimInterfaceTemplatesDeleteNoContent {
	return &DcimInterfaceTemplatesDeleteNoContent{}
}

/* DcimInterfaceTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimInterfaceTemplatesDeleteNoContent dcim interface templates delete no content
*/
type DcimInterfaceTemplatesDeleteNoContent struct {
}

func (o *DcimInterfaceTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesDeleteNoContent ", 204)
}

func (o *DcimInterfaceTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
