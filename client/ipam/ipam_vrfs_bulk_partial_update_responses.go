package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVrfsBulkPartialUpdateReader is a Reader for the IpamVrfsBulkPartialUpdate structure.
type IpamVrfsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVrfsBulkPartialUpdateOK creates a IpamVrfsBulkPartialUpdateOK with default headers values
func NewIpamVrfsBulkPartialUpdateOK() *IpamVrfsBulkPartialUpdateOK {
	return &IpamVrfsBulkPartialUpdateOK{}
}

/* IpamVrfsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamVrfsBulkPartialUpdateOK ipam vrfs bulk partial update o k
*/
type IpamVrfsBulkPartialUpdateOK struct {
	Payload *models.VRF
}

func (o *IpamVrfsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/][%d] ipamVrfsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVrfsBulkPartialUpdateOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
