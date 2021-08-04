package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerOutletsBulkDeleteReader is a Reader for the DcimPowerOutletsBulkDelete structure.
type DcimPowerOutletsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerOutletsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletsBulkDeleteNoContent creates a DcimPowerOutletsBulkDeleteNoContent with default headers values
func NewDcimPowerOutletsBulkDeleteNoContent() *DcimPowerOutletsBulkDeleteNoContent {
	return &DcimPowerOutletsBulkDeleteNoContent{}
}

/* DcimPowerOutletsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerOutletsBulkDeleteNoContent dcim power outlets bulk delete no content
*/
type DcimPowerOutletsBulkDeleteNoContent struct {
}

func (o *DcimPowerOutletsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlets/][%d] dcimPowerOutletsBulkDeleteNoContent ", 204)
}

func (o *DcimPowerOutletsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
