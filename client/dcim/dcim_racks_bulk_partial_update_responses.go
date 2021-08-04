package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRacksBulkPartialUpdateReader is a Reader for the DcimRacksBulkPartialUpdate structure.
type DcimRacksBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRacksBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRacksBulkPartialUpdateOK creates a DcimRacksBulkPartialUpdateOK with default headers values
func NewDcimRacksBulkPartialUpdateOK() *DcimRacksBulkPartialUpdateOK {
	return &DcimRacksBulkPartialUpdateOK{}
}

/* DcimRacksBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRacksBulkPartialUpdateOK dcim racks bulk partial update o k
*/
type DcimRacksBulkPartialUpdateOK struct {
	Payload *models.Rack
}

func (o *DcimRacksBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/racks/][%d] dcimRacksBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRacksBulkPartialUpdateOK) GetPayload() *models.Rack {
	return o.Payload
}

func (o *DcimRacksBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
