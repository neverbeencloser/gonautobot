package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfacesUpdateReader is a Reader for the DcimInterfacesUpdate structure.
type DcimInterfacesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfacesUpdateOK creates a DcimInterfacesUpdateOK with default headers values
func NewDcimInterfacesUpdateOK() *DcimInterfacesUpdateOK {
	return &DcimInterfacesUpdateOK{}
}

/* DcimInterfacesUpdateOK describes a response with status code 200, with default header values.

DcimInterfacesUpdateOK dcim interfaces update o k
*/
type DcimInterfacesUpdateOK struct {
	Payload *models.Interface
}

func (o *DcimInterfacesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/interfaces/{id}/][%d] dcimInterfacesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInterfacesUpdateOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
