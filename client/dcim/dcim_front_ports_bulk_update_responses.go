package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortsBulkUpdateReader is a Reader for the DcimFrontPortsBulkUpdate structure.
type DcimFrontPortsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortsBulkUpdateOK creates a DcimFrontPortsBulkUpdateOK with default headers values
func NewDcimFrontPortsBulkUpdateOK() *DcimFrontPortsBulkUpdateOK {
	return &DcimFrontPortsBulkUpdateOK{}
}

/* DcimFrontPortsBulkUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortsBulkUpdateOK dcim front ports bulk update o k
*/
type DcimFrontPortsBulkUpdateOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/][%d] dcimFrontPortsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortsBulkUpdateOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
