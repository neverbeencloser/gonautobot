package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamRolesBulkPartialUpdateReader is a Reader for the IpamRolesBulkPartialUpdate structure.
type IpamRolesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRolesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRolesBulkPartialUpdateOK creates a IpamRolesBulkPartialUpdateOK with default headers values
func NewIpamRolesBulkPartialUpdateOK() *IpamRolesBulkPartialUpdateOK {
	return &IpamRolesBulkPartialUpdateOK{}
}

/* IpamRolesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamRolesBulkPartialUpdateOK ipam roles bulk partial update o k
*/
type IpamRolesBulkPartialUpdateOK struct {
	Payload *models.Role
}

func (o *IpamRolesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/roles/][%d] ipamRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamRolesBulkPartialUpdateOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
