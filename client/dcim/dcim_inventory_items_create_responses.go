package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInventoryItemsCreateReader is a Reader for the DcimInventoryItemsCreate structure.
type DcimInventoryItemsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimInventoryItemsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInventoryItemsCreateCreated creates a DcimInventoryItemsCreateCreated with default headers values
func NewDcimInventoryItemsCreateCreated() *DcimInventoryItemsCreateCreated {
	return &DcimInventoryItemsCreateCreated{}
}

/* DcimInventoryItemsCreateCreated describes a response with status code 201, with default header values.

DcimInventoryItemsCreateCreated dcim inventory items create created
*/
type DcimInventoryItemsCreateCreated struct {
	Payload *models.InventoryItem
}

func (o *DcimInventoryItemsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/inventory-items/][%d] dcimInventoryItemsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimInventoryItemsCreateCreated) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
