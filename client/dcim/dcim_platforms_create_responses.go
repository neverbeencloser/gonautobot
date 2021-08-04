package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPlatformsCreateReader is a Reader for the DcimPlatformsCreate structure.
type DcimPlatformsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPlatformsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPlatformsCreateCreated creates a DcimPlatformsCreateCreated with default headers values
func NewDcimPlatformsCreateCreated() *DcimPlatformsCreateCreated {
	return &DcimPlatformsCreateCreated{}
}

/* DcimPlatformsCreateCreated describes a response with status code 201, with default header values.

DcimPlatformsCreateCreated dcim platforms create created
*/
type DcimPlatformsCreateCreated struct {
	Payload *models.Platform
}

func (o *DcimPlatformsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/platforms/][%d] dcimPlatformsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimPlatformsCreateCreated) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
