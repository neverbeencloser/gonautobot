package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVrfsBulkUpdateReader is a Reader for the IpamVrfsBulkUpdate structure.
type IpamVrfsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVrfsBulkUpdateOK creates a IpamVrfsBulkUpdateOK with default headers values
func NewIpamVrfsBulkUpdateOK() *IpamVrfsBulkUpdateOK {
	return &IpamVrfsBulkUpdateOK{}
}

/* IpamVrfsBulkUpdateOK describes a response with status code 200, with default header values.

IpamVrfsBulkUpdateOK ipam vrfs bulk update o k
*/
type IpamVrfsBulkUpdateOK struct {
	Payload *models.VRF
}

func (o *IpamVrfsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vrfs/][%d] ipamVrfsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVrfsBulkUpdateOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
