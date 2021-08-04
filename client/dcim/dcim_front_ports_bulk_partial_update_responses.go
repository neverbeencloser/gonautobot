package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortsBulkPartialUpdateReader is a Reader for the DcimFrontPortsBulkPartialUpdate structure.
type DcimFrontPortsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortsBulkPartialUpdateOK creates a DcimFrontPortsBulkPartialUpdateOK with default headers values
func NewDcimFrontPortsBulkPartialUpdateOK() *DcimFrontPortsBulkPartialUpdateOK {
	return &DcimFrontPortsBulkPartialUpdateOK{}
}

/* DcimFrontPortsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortsBulkPartialUpdateOK dcim front ports bulk partial update o k
*/
type DcimFrontPortsBulkPartialUpdateOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/][%d] dcimFrontPortsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortsBulkPartialUpdateOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
