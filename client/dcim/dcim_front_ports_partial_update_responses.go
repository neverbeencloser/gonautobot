package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortsPartialUpdateReader is a Reader for the DcimFrontPortsPartialUpdate structure.
type DcimFrontPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortsPartialUpdateOK creates a DcimFrontPortsPartialUpdateOK with default headers values
func NewDcimFrontPortsPartialUpdateOK() *DcimFrontPortsPartialUpdateOK {
	return &DcimFrontPortsPartialUpdateOK{}
}

/* DcimFrontPortsPartialUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortsPartialUpdateOK dcim front ports partial update o k
*/
type DcimFrontPortsPartialUpdateOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/{id}/][%d] dcimFrontPortsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortsPartialUpdateOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
