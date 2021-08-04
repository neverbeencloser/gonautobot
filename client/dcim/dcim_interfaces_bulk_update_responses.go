package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfacesBulkUpdateReader is a Reader for the DcimInterfacesBulkUpdate structure.
type DcimInterfacesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfacesBulkUpdateOK creates a DcimInterfacesBulkUpdateOK with default headers values
func NewDcimInterfacesBulkUpdateOK() *DcimInterfacesBulkUpdateOK {
	return &DcimInterfacesBulkUpdateOK{}
}

/* DcimInterfacesBulkUpdateOK describes a response with status code 200, with default header values.

DcimInterfacesBulkUpdateOK dcim interfaces bulk update o k
*/
type DcimInterfacesBulkUpdateOK struct {
	Payload *models.Interface
}

func (o *DcimInterfacesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/interfaces/][%d] dcimInterfacesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInterfacesBulkUpdateOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
