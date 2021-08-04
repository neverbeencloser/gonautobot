package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfacesTraceReader is a Reader for the DcimInterfacesTrace structure.
type DcimInterfacesTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfacesTraceOK creates a DcimInterfacesTraceOK with default headers values
func NewDcimInterfacesTraceOK() *DcimInterfacesTraceOK {
	return &DcimInterfacesTraceOK{}
}

/* DcimInterfacesTraceOK describes a response with status code 200, with default header values.

DcimInterfacesTraceOK dcim interfaces trace o k
*/
type DcimInterfacesTraceOK struct {
	Payload *models.Interface
}

func (o *DcimInterfacesTraceOK) Error() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/trace/][%d] dcimInterfacesTraceOK  %+v", 200, o.Payload)
}
func (o *DcimInterfacesTraceOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
