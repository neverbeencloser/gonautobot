package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimCablesBulkPartialUpdateReader is a Reader for the DcimCablesBulkPartialUpdate structure.
type DcimCablesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimCablesBulkPartialUpdateOK creates a DcimCablesBulkPartialUpdateOK with default headers values
func NewDcimCablesBulkPartialUpdateOK() *DcimCablesBulkPartialUpdateOK {
	return &DcimCablesBulkPartialUpdateOK{}
}

/* DcimCablesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimCablesBulkPartialUpdateOK dcim cables bulk partial update o k
*/
type DcimCablesBulkPartialUpdateOK struct {
	Payload *models.Cable
}

func (o *DcimCablesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/cables/][%d] dcimCablesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimCablesBulkPartialUpdateOK) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
