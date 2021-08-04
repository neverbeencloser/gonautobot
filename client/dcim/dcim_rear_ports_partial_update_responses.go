package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRearPortsPartialUpdateReader is a Reader for the DcimRearPortsPartialUpdate structure.
type DcimRearPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortsPartialUpdateOK creates a DcimRearPortsPartialUpdateOK with default headers values
func NewDcimRearPortsPartialUpdateOK() *DcimRearPortsPartialUpdateOK {
	return &DcimRearPortsPartialUpdateOK{}
}

/* DcimRearPortsPartialUpdateOK describes a response with status code 200, with default header values.

DcimRearPortsPartialUpdateOK dcim rear ports partial update o k
*/
type DcimRearPortsPartialUpdateOK struct {
	Payload *models.RearPort
}

func (o *DcimRearPortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rear-ports/{id}/][%d] dcimRearPortsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRearPortsPartialUpdateOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
