package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimManufacturersUpdateReader is a Reader for the DcimManufacturersUpdate structure.
type DcimManufacturersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimManufacturersUpdateOK creates a DcimManufacturersUpdateOK with default headers values
func NewDcimManufacturersUpdateOK() *DcimManufacturersUpdateOK {
	return &DcimManufacturersUpdateOK{}
}

/* DcimManufacturersUpdateOK describes a response with status code 200, with default header values.

DcimManufacturersUpdateOK dcim manufacturers update o k
*/
type DcimManufacturersUpdateOK struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/manufacturers/{id}/][%d] dcimManufacturersUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimManufacturersUpdateOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
