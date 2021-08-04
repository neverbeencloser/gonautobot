package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortsCreateReader is a Reader for the DcimFrontPortsCreate structure.
type DcimFrontPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimFrontPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortsCreateCreated creates a DcimFrontPortsCreateCreated with default headers values
func NewDcimFrontPortsCreateCreated() *DcimFrontPortsCreateCreated {
	return &DcimFrontPortsCreateCreated{}
}

/* DcimFrontPortsCreateCreated describes a response with status code 201, with default header values.

DcimFrontPortsCreateCreated dcim front ports create created
*/
type DcimFrontPortsCreateCreated struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/front-ports/][%d] dcimFrontPortsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimFrontPortsCreateCreated) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
