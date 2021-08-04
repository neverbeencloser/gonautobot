package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPlatformsPartialUpdateReader is a Reader for the DcimPlatformsPartialUpdate structure.
type DcimPlatformsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPlatformsPartialUpdateOK creates a DcimPlatformsPartialUpdateOK with default headers values
func NewDcimPlatformsPartialUpdateOK() *DcimPlatformsPartialUpdateOK {
	return &DcimPlatformsPartialUpdateOK{}
}

/* DcimPlatformsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPlatformsPartialUpdateOK dcim platforms partial update o k
*/
type DcimPlatformsPartialUpdateOK struct {
	Payload *models.Platform
}

func (o *DcimPlatformsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/{id}/][%d] dcimPlatformsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPlatformsPartialUpdateOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
