package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimCablesCreateReader is a Reader for the DcimCablesCreate structure.
type DcimCablesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimCablesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimCablesCreateCreated creates a DcimCablesCreateCreated with default headers values
func NewDcimCablesCreateCreated() *DcimCablesCreateCreated {
	return &DcimCablesCreateCreated{}
}

/* DcimCablesCreateCreated describes a response with status code 201, with default header values.

DcimCablesCreateCreated dcim cables create created
*/
type DcimCablesCreateCreated struct {
	Payload *models.Cable
}

func (o *DcimCablesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/cables/][%d] dcimCablesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimCablesCreateCreated) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
