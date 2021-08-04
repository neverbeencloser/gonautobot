package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimManufacturersBulkPartialUpdateReader is a Reader for the DcimManufacturersBulkPartialUpdate structure.
type DcimManufacturersBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimManufacturersBulkPartialUpdateOK creates a DcimManufacturersBulkPartialUpdateOK with default headers values
func NewDcimManufacturersBulkPartialUpdateOK() *DcimManufacturersBulkPartialUpdateOK {
	return &DcimManufacturersBulkPartialUpdateOK{}
}

/* DcimManufacturersBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimManufacturersBulkPartialUpdateOK dcim manufacturers bulk partial update o k
*/
type DcimManufacturersBulkPartialUpdateOK struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/][%d] dcimManufacturersBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimManufacturersBulkPartialUpdateOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
