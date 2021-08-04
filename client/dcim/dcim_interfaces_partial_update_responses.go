package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfacesPartialUpdateReader is a Reader for the DcimInterfacesPartialUpdate structure.
type DcimInterfacesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfacesPartialUpdateOK creates a DcimInterfacesPartialUpdateOK with default headers values
func NewDcimInterfacesPartialUpdateOK() *DcimInterfacesPartialUpdateOK {
	return &DcimInterfacesPartialUpdateOK{}
}

/* DcimInterfacesPartialUpdateOK describes a response with status code 200, with default header values.

DcimInterfacesPartialUpdateOK dcim interfaces partial update o k
*/
type DcimInterfacesPartialUpdateOK struct {
	Payload *models.Interface
}

func (o *DcimInterfacesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/{id}/][%d] dcimInterfacesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInterfacesPartialUpdateOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
