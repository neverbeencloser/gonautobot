package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInventoryItemsBulkUpdateReader is a Reader for the DcimInventoryItemsBulkUpdate structure.
type DcimInventoryItemsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInventoryItemsBulkUpdateOK creates a DcimInventoryItemsBulkUpdateOK with default headers values
func NewDcimInventoryItemsBulkUpdateOK() *DcimInventoryItemsBulkUpdateOK {
	return &DcimInventoryItemsBulkUpdateOK{}
}

/* DcimInventoryItemsBulkUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemsBulkUpdateOK dcim inventory items bulk update o k
*/
type DcimInventoryItemsBulkUpdateOK struct {
	Payload *models.InventoryItem
}

func (o *DcimInventoryItemsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-items/][%d] dcimInventoryItemsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemsBulkUpdateOK) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
