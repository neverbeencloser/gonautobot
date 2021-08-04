package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVrfsPartialUpdateReader is a Reader for the IpamVrfsPartialUpdate structure.
type IpamVrfsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVrfsPartialUpdateOK creates a IpamVrfsPartialUpdateOK with default headers values
func NewIpamVrfsPartialUpdateOK() *IpamVrfsPartialUpdateOK {
	return &IpamVrfsPartialUpdateOK{}
}

/* IpamVrfsPartialUpdateOK describes a response with status code 200, with default header values.

IpamVrfsPartialUpdateOK ipam vrfs partial update o k
*/
type IpamVrfsPartialUpdateOK struct {
	Payload *models.VRF
}

func (o *IpamVrfsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/{id}/][%d] ipamVrfsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVrfsPartialUpdateOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
