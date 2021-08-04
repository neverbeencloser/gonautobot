package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPlatformsBulkUpdateReader is a Reader for the DcimPlatformsBulkUpdate structure.
type DcimPlatformsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPlatformsBulkUpdateOK creates a DcimPlatformsBulkUpdateOK with default headers values
func NewDcimPlatformsBulkUpdateOK() *DcimPlatformsBulkUpdateOK {
	return &DcimPlatformsBulkUpdateOK{}
}

/* DcimPlatformsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPlatformsBulkUpdateOK dcim platforms bulk update o k
*/
type DcimPlatformsBulkUpdateOK struct {
	Payload *models.Platform
}

func (o *DcimPlatformsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/platforms/][%d] dcimPlatformsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPlatformsBulkUpdateOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
