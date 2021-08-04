package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRegionsUpdateReader is a Reader for the DcimRegionsUpdate structure.
type DcimRegionsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRegionsUpdateOK creates a DcimRegionsUpdateOK with default headers values
func NewDcimRegionsUpdateOK() *DcimRegionsUpdateOK {
	return &DcimRegionsUpdateOK{}
}

/* DcimRegionsUpdateOK describes a response with status code 200, with default header values.

DcimRegionsUpdateOK dcim regions update o k
*/
type DcimRegionsUpdateOK struct {
	Payload *models.Region
}

func (o *DcimRegionsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcimRegionsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRegionsUpdateOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
