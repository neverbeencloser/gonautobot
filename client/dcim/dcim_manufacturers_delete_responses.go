package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimManufacturersDeleteReader is a Reader for the DcimManufacturersDelete structure.
type DcimManufacturersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimManufacturersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimManufacturersDeleteNoContent creates a DcimManufacturersDeleteNoContent with default headers values
func NewDcimManufacturersDeleteNoContent() *DcimManufacturersDeleteNoContent {
	return &DcimManufacturersDeleteNoContent{}
}

/* DcimManufacturersDeleteNoContent describes a response with status code 204, with default header values.

DcimManufacturersDeleteNoContent dcim manufacturers delete no content
*/
type DcimManufacturersDeleteNoContent struct {
}

func (o *DcimManufacturersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/manufacturers/{id}/][%d] dcimManufacturersDeleteNoContent ", 204)
}

func (o *DcimManufacturersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
