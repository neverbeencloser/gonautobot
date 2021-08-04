package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVrfsUpdateReader is a Reader for the IpamVrfsUpdate structure.
type IpamVrfsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVrfsUpdateOK creates a IpamVrfsUpdateOK with default headers values
func NewIpamVrfsUpdateOK() *IpamVrfsUpdateOK {
	return &IpamVrfsUpdateOK{}
}

/* IpamVrfsUpdateOK describes a response with status code 200, with default header values.

IpamVrfsUpdateOK ipam vrfs update o k
*/
type IpamVrfsUpdateOK struct {
	Payload *models.VRF
}

func (o *IpamVrfsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vrfs/{id}/][%d] ipamVrfsUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVrfsUpdateOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
