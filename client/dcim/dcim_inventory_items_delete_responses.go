package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimInventoryItemsDeleteReader is a Reader for the DcimInventoryItemsDelete structure.
type DcimInventoryItemsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInventoryItemsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInventoryItemsDeleteNoContent creates a DcimInventoryItemsDeleteNoContent with default headers values
func NewDcimInventoryItemsDeleteNoContent() *DcimInventoryItemsDeleteNoContent {
	return &DcimInventoryItemsDeleteNoContent{}
}

/* DcimInventoryItemsDeleteNoContent describes a response with status code 204, with default header values.

DcimInventoryItemsDeleteNoContent dcim inventory items delete no content
*/
type DcimInventoryItemsDeleteNoContent struct {
}

func (o *DcimInventoryItemsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/{id}/][%d] dcimInventoryItemsDeleteNoContent ", 204)
}

func (o *DcimInventoryItemsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
