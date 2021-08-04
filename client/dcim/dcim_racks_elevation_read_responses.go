package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRacksElevationReadReader is a Reader for the DcimRacksElevationRead structure.
type DcimRacksElevationReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksElevationReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRacksElevationReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRacksElevationReadOK creates a DcimRacksElevationReadOK with default headers values
func NewDcimRacksElevationReadOK() *DcimRacksElevationReadOK {
	return &DcimRacksElevationReadOK{}
}

/* DcimRacksElevationReadOK describes a response with status code 200, with default header values.

DcimRacksElevationReadOK dcim racks elevation read o k
*/
type DcimRacksElevationReadOK struct {
	Payload []*models.RackUnit
}

func (o *DcimRacksElevationReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/elevation/][%d] dcimRacksElevationReadOK  %+v", 200, o.Payload)
}
func (o *DcimRacksElevationReadOK) GetPayload() []*models.RackUnit {
	return o.Payload
}

func (o *DcimRacksElevationReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
