package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimManufacturersBulkUpdateReader is a Reader for the DcimManufacturersBulkUpdate structure.
type DcimManufacturersBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimManufacturersBulkUpdateOK creates a DcimManufacturersBulkUpdateOK with default headers values
func NewDcimManufacturersBulkUpdateOK() *DcimManufacturersBulkUpdateOK {
	return &DcimManufacturersBulkUpdateOK{}
}

/* DcimManufacturersBulkUpdateOK describes a response with status code 200, with default header values.

DcimManufacturersBulkUpdateOK dcim manufacturers bulk update o k
*/
type DcimManufacturersBulkUpdateOK struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/manufacturers/][%d] dcimManufacturersBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimManufacturersBulkUpdateOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
