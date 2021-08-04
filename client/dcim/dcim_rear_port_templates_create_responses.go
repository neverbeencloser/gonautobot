package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRearPortTemplatesCreateReader is a Reader for the DcimRearPortTemplatesCreate structure.
type DcimRearPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRearPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortTemplatesCreateCreated creates a DcimRearPortTemplatesCreateCreated with default headers values
func NewDcimRearPortTemplatesCreateCreated() *DcimRearPortTemplatesCreateCreated {
	return &DcimRearPortTemplatesCreateCreated{}
}

/* DcimRearPortTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimRearPortTemplatesCreateCreated dcim rear port templates create created
*/
type DcimRearPortTemplatesCreateCreated struct {
	Payload *models.RearPortTemplate
}

func (o *DcimRearPortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rear-port-templates/][%d] dcimRearPortTemplatesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimRearPortTemplatesCreateCreated) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
