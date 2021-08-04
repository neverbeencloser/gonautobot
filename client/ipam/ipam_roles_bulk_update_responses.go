package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamRolesBulkUpdateReader is a Reader for the IpamRolesBulkUpdate structure.
type IpamRolesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRolesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRolesBulkUpdateOK creates a IpamRolesBulkUpdateOK with default headers values
func NewIpamRolesBulkUpdateOK() *IpamRolesBulkUpdateOK {
	return &IpamRolesBulkUpdateOK{}
}

/* IpamRolesBulkUpdateOK describes a response with status code 200, with default header values.

IpamRolesBulkUpdateOK ipam roles bulk update o k
*/
type IpamRolesBulkUpdateOK struct {
	Payload *models.Role
}

func (o *IpamRolesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/roles/][%d] ipamRolesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamRolesBulkUpdateOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
