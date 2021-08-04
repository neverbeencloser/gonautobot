package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfacesCreateReader is a Reader for the DcimInterfacesCreate structure.
type DcimInterfacesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimInterfacesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfacesCreateCreated creates a DcimInterfacesCreateCreated with default headers values
func NewDcimInterfacesCreateCreated() *DcimInterfacesCreateCreated {
	return &DcimInterfacesCreateCreated{}
}

/* DcimInterfacesCreateCreated describes a response with status code 201, with default header values.

DcimInterfacesCreateCreated dcim interfaces create created
*/
type DcimInterfacesCreateCreated struct {
	Payload *models.Interface
}

func (o *DcimInterfacesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/interfaces/][%d] dcimInterfacesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimInterfacesCreateCreated) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
