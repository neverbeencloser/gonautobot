package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRegionsCreateReader is a Reader for the DcimRegionsCreate structure.
type DcimRegionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRegionsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRegionsCreateCreated creates a DcimRegionsCreateCreated with default headers values
func NewDcimRegionsCreateCreated() *DcimRegionsCreateCreated {
	return &DcimRegionsCreateCreated{}
}

/* DcimRegionsCreateCreated describes a response with status code 201, with default header values.

DcimRegionsCreateCreated dcim regions create created
*/
type DcimRegionsCreateCreated struct {
	Payload *models.Region
}

func (o *DcimRegionsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/regions/][%d] dcimRegionsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimRegionsCreateCreated) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
