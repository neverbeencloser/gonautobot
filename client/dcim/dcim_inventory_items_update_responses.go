package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInventoryItemsUpdateReader is a Reader for the DcimInventoryItemsUpdate structure.
type DcimInventoryItemsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInventoryItemsUpdateOK creates a DcimInventoryItemsUpdateOK with default headers values
func NewDcimInventoryItemsUpdateOK() *DcimInventoryItemsUpdateOK {
	return &DcimInventoryItemsUpdateOK{}
}

/* DcimInventoryItemsUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemsUpdateOK dcim inventory items update o k
*/
type DcimInventoryItemsUpdateOK struct {
	Payload *models.InventoryItem
}

func (o *DcimInventoryItemsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-items/{id}/][%d] dcimInventoryItemsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemsUpdateOK) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
