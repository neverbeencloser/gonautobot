package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortTemplatesCreateReader is a Reader for the DcimFrontPortTemplatesCreate structure.
type DcimFrontPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimFrontPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortTemplatesCreateCreated creates a DcimFrontPortTemplatesCreateCreated with default headers values
func NewDcimFrontPortTemplatesCreateCreated() *DcimFrontPortTemplatesCreateCreated {
	return &DcimFrontPortTemplatesCreateCreated{}
}

/* DcimFrontPortTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimFrontPortTemplatesCreateCreated dcim front port templates create created
*/
type DcimFrontPortTemplatesCreateCreated struct {
	Payload *models.FrontPortTemplate
}

func (o *DcimFrontPortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/front-port-templates/][%d] dcimFrontPortTemplatesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimFrontPortTemplatesCreateCreated) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
