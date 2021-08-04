package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamRolesReadReader is a Reader for the IpamRolesRead structure.
type IpamRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRolesReadOK creates a IpamRolesReadOK with default headers values
func NewIpamRolesReadOK() *IpamRolesReadOK {
	return &IpamRolesReadOK{}
}

/* IpamRolesReadOK describes a response with status code 200, with default header values.

IpamRolesReadOK ipam roles read o k
*/
type IpamRolesReadOK struct {
	Payload *models.Role
}

func (o *IpamRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipamRolesReadOK  %+v", 200, o.Payload)
}
func (o *IpamRolesReadOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
