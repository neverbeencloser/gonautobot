package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRegionsBulkUpdateReader is a Reader for the DcimRegionsBulkUpdate structure.
type DcimRegionsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRegionsBulkUpdateOK creates a DcimRegionsBulkUpdateOK with default headers values
func NewDcimRegionsBulkUpdateOK() *DcimRegionsBulkUpdateOK {
	return &DcimRegionsBulkUpdateOK{}
}

/* DcimRegionsBulkUpdateOK describes a response with status code 200, with default header values.

DcimRegionsBulkUpdateOK dcim regions bulk update o k
*/
type DcimRegionsBulkUpdateOK struct {
	Payload *models.Region
}

func (o *DcimRegionsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/regions/][%d] dcimRegionsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRegionsBulkUpdateOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
