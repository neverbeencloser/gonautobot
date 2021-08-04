package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRegionsReadReader is a Reader for the DcimRegionsRead structure.
type DcimRegionsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRegionsReadOK creates a DcimRegionsReadOK with default headers values
func NewDcimRegionsReadOK() *DcimRegionsReadOK {
	return &DcimRegionsReadOK{}
}

/* DcimRegionsReadOK describes a response with status code 200, with default header values.

DcimRegionsReadOK dcim regions read o k
*/
type DcimRegionsReadOK struct {
	Payload *models.Region
}

func (o *DcimRegionsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/regions/{id}/][%d] dcimRegionsReadOK  %+v", 200, o.Payload)
}
func (o *DcimRegionsReadOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
