package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPlatformsBulkPartialUpdateReader is a Reader for the DcimPlatformsBulkPartialUpdate structure.
type DcimPlatformsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPlatformsBulkPartialUpdateOK creates a DcimPlatformsBulkPartialUpdateOK with default headers values
func NewDcimPlatformsBulkPartialUpdateOK() *DcimPlatformsBulkPartialUpdateOK {
	return &DcimPlatformsBulkPartialUpdateOK{}
}

/* DcimPlatformsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPlatformsBulkPartialUpdateOK dcim platforms bulk partial update o k
*/
type DcimPlatformsBulkPartialUpdateOK struct {
	Payload *models.Platform
}

func (o *DcimPlatformsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/][%d] dcimPlatformsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPlatformsBulkPartialUpdateOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
