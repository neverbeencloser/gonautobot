package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVlanGroupsBulkPartialUpdateReader is a Reader for the IpamVlanGroupsBulkPartialUpdate structure.
type IpamVlanGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlanGroupsBulkPartialUpdateOK creates a IpamVlanGroupsBulkPartialUpdateOK with default headers values
func NewIpamVlanGroupsBulkPartialUpdateOK() *IpamVlanGroupsBulkPartialUpdateOK {
	return &IpamVlanGroupsBulkPartialUpdateOK{}
}

/* IpamVlanGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamVlanGroupsBulkPartialUpdateOK ipam vlan groups bulk partial update o k
*/
type IpamVlanGroupsBulkPartialUpdateOK struct {
	Payload *models.VLANGroup
}

func (o *IpamVlanGroupsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlan-groups/][%d] ipamVlanGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlanGroupsBulkPartialUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
