package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimCablesReadReader is a Reader for the DcimCablesRead structure.
type DcimCablesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimCablesReadOK creates a DcimCablesReadOK with default headers values
func NewDcimCablesReadOK() *DcimCablesReadOK {
	return &DcimCablesReadOK{}
}

/* DcimCablesReadOK describes a response with status code 200, with default header values.

DcimCablesReadOK dcim cables read o k
*/
type DcimCablesReadOK struct {
	Payload *models.Cable
}

func (o *DcimCablesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/cables/{id}/][%d] dcimCablesReadOK  %+v", 200, o.Payload)
}
func (o *DcimCablesReadOK) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
