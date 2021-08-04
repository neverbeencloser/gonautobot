package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInventoryItemsReadReader is a Reader for the DcimInventoryItemsRead structure.
type DcimInventoryItemsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInventoryItemsReadOK creates a DcimInventoryItemsReadOK with default headers values
func NewDcimInventoryItemsReadOK() *DcimInventoryItemsReadOK {
	return &DcimInventoryItemsReadOK{}
}

/* DcimInventoryItemsReadOK describes a response with status code 200, with default header values.

DcimInventoryItemsReadOK dcim inventory items read o k
*/
type DcimInventoryItemsReadOK struct {
	Payload *models.InventoryItem
}

func (o *DcimInventoryItemsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/inventory-items/{id}/][%d] dcimInventoryItemsReadOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemsReadOK) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
