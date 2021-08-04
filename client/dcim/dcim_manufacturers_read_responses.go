package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimManufacturersReadReader is a Reader for the DcimManufacturersRead structure.
type DcimManufacturersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimManufacturersReadOK creates a DcimManufacturersReadOK with default headers values
func NewDcimManufacturersReadOK() *DcimManufacturersReadOK {
	return &DcimManufacturersReadOK{}
}

/* DcimManufacturersReadOK describes a response with status code 200, with default header values.

DcimManufacturersReadOK dcim manufacturers read o k
*/
type DcimManufacturersReadOK struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/{id}/][%d] dcimManufacturersReadOK  %+v", 200, o.Payload)
}
func (o *DcimManufacturersReadOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
