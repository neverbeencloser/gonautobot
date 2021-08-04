package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRacksCreateReader is a Reader for the DcimRacksCreate structure.
type DcimRacksCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRacksCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRacksCreateCreated creates a DcimRacksCreateCreated with default headers values
func NewDcimRacksCreateCreated() *DcimRacksCreateCreated {
	return &DcimRacksCreateCreated{}
}

/* DcimRacksCreateCreated describes a response with status code 201, with default header values.

DcimRacksCreateCreated dcim racks create created
*/
type DcimRacksCreateCreated struct {
	Payload *models.Rack
}

func (o *DcimRacksCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/racks/][%d] dcimRacksCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimRacksCreateCreated) GetPayload() *models.Rack {
	return o.Payload
}

func (o *DcimRacksCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
