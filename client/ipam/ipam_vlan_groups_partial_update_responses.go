package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVlanGroupsPartialUpdateReader is a Reader for the IpamVlanGroupsPartialUpdate structure.
type IpamVlanGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlanGroupsPartialUpdateOK creates a IpamVlanGroupsPartialUpdateOK with default headers values
func NewIpamVlanGroupsPartialUpdateOK() *IpamVlanGroupsPartialUpdateOK {
	return &IpamVlanGroupsPartialUpdateOK{}
}

/* IpamVlanGroupsPartialUpdateOK describes a response with status code 200, with default header values.

IpamVlanGroupsPartialUpdateOK ipam vlan groups partial update o k
*/
type IpamVlanGroupsPartialUpdateOK struct {
	Payload *models.VLANGroup
}

func (o *IpamVlanGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlanGroupsPartialUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
