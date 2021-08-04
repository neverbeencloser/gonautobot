package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortsReadReader is a Reader for the DcimFrontPortsRead structure.
type DcimFrontPortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortsReadOK creates a DcimFrontPortsReadOK with default headers values
func NewDcimFrontPortsReadOK() *DcimFrontPortsReadOK {
	return &DcimFrontPortsReadOK{}
}

/* DcimFrontPortsReadOK describes a response with status code 200, with default header values.

DcimFrontPortsReadOK dcim front ports read o k
*/
type DcimFrontPortsReadOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/][%d] dcimFrontPortsReadOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortsReadOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
