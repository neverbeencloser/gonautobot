package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackGroupsReadReader is a Reader for the DcimRackGroupsRead structure.
type DcimRackGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackGroupsReadOK creates a DcimRackGroupsReadOK with default headers values
func NewDcimRackGroupsReadOK() *DcimRackGroupsReadOK {
	return &DcimRackGroupsReadOK{}
}

/* DcimRackGroupsReadOK describes a response with status code 200, with default header values.

DcimRackGroupsReadOK dcim rack groups read o k
*/
type DcimRackGroupsReadOK struct {
	Payload *models.RackGroup
}

func (o *DcimRackGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-groups/{id}/][%d] dcimRackGroupsReadOK  %+v", 200, o.Payload)
}
func (o *DcimRackGroupsReadOK) GetPayload() *models.RackGroup {
	return o.Payload
}

func (o *DcimRackGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
