package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfaceTemplatesCreateReader is a Reader for the DcimInterfaceTemplatesCreate structure.
type DcimInterfaceTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimInterfaceTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesCreateCreated creates a DcimInterfaceTemplatesCreateCreated with default headers values
func NewDcimInterfaceTemplatesCreateCreated() *DcimInterfaceTemplatesCreateCreated {
	return &DcimInterfaceTemplatesCreateCreated{}
}

/* DcimInterfaceTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimInterfaceTemplatesCreateCreated dcim interface templates create created
*/
type DcimInterfaceTemplatesCreateCreated struct {
	Payload *models.InterfaceTemplate
}

func (o *DcimInterfaceTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/interface-templates/][%d] dcimInterfaceTemplatesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimInterfaceTemplatesCreateCreated) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
