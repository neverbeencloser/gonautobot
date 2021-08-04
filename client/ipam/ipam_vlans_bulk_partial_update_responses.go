package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVlansBulkPartialUpdateReader is a Reader for the IpamVlansBulkPartialUpdate structure.
type IpamVlansBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlansBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlansBulkPartialUpdateOK creates a IpamVlansBulkPartialUpdateOK with default headers values
func NewIpamVlansBulkPartialUpdateOK() *IpamVlansBulkPartialUpdateOK {
	return &IpamVlansBulkPartialUpdateOK{}
}

/* IpamVlansBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamVlansBulkPartialUpdateOK ipam vlans bulk partial update o k
*/
type IpamVlansBulkPartialUpdateOK struct {
	Payload *models.VLAN
}

func (o *IpamVlansBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlans/][%d] ipamVlansBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlansBulkPartialUpdateOK) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
