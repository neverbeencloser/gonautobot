package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimCablesBulkUpdateReader is a Reader for the DcimCablesBulkUpdate structure.
type DcimCablesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimCablesBulkUpdateOK creates a DcimCablesBulkUpdateOK with default headers values
func NewDcimCablesBulkUpdateOK() *DcimCablesBulkUpdateOK {
	return &DcimCablesBulkUpdateOK{}
}

/* DcimCablesBulkUpdateOK describes a response with status code 200, with default header values.

DcimCablesBulkUpdateOK dcim cables bulk update o k
*/
type DcimCablesBulkUpdateOK struct {
	Payload *models.Cable
}

func (o *DcimCablesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/cables/][%d] dcimCablesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimCablesBulkUpdateOK) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
