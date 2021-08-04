package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimInventoryItemsBulkDeleteReader is a Reader for the DcimInventoryItemsBulkDelete structure.
type DcimInventoryItemsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInventoryItemsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInventoryItemsBulkDeleteNoContent creates a DcimInventoryItemsBulkDeleteNoContent with default headers values
func NewDcimInventoryItemsBulkDeleteNoContent() *DcimInventoryItemsBulkDeleteNoContent {
	return &DcimInventoryItemsBulkDeleteNoContent{}
}

/* DcimInventoryItemsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimInventoryItemsBulkDeleteNoContent dcim inventory items bulk delete no content
*/
type DcimInventoryItemsBulkDeleteNoContent struct {
}

func (o *DcimInventoryItemsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/][%d] dcimInventoryItemsBulkDeleteNoContent ", 204)
}

func (o *DcimInventoryItemsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
