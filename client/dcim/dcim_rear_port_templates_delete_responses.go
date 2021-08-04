package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRearPortTemplatesDeleteReader is a Reader for the DcimRearPortTemplatesDelete structure.
type DcimRearPortTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRearPortTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortTemplatesDeleteNoContent creates a DcimRearPortTemplatesDeleteNoContent with default headers values
func NewDcimRearPortTemplatesDeleteNoContent() *DcimRearPortTemplatesDeleteNoContent {
	return &DcimRearPortTemplatesDeleteNoContent{}
}

/* DcimRearPortTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimRearPortTemplatesDeleteNoContent dcim rear port templates delete no content
*/
type DcimRearPortTemplatesDeleteNoContent struct {
}

func (o *DcimRearPortTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimRearPortTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
