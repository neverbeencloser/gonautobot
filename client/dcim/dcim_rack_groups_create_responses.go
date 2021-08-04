package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackGroupsCreateReader is a Reader for the DcimRackGroupsCreate structure.
type DcimRackGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRackGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackGroupsCreateCreated creates a DcimRackGroupsCreateCreated with default headers values
func NewDcimRackGroupsCreateCreated() *DcimRackGroupsCreateCreated {
	return &DcimRackGroupsCreateCreated{}
}

/* DcimRackGroupsCreateCreated describes a response with status code 201, with default header values.

DcimRackGroupsCreateCreated dcim rack groups create created
*/
type DcimRackGroupsCreateCreated struct {
	Payload *models.RackGroup
}

func (o *DcimRackGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rack-groups/][%d] dcimRackGroupsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimRackGroupsCreateCreated) GetPayload() *models.RackGroup {
	return o.Payload
}

func (o *DcimRackGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
