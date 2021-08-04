package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVlansBulkUpdateReader is a Reader for the IpamVlansBulkUpdate structure.
type IpamVlansBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlansBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlansBulkUpdateOK creates a IpamVlansBulkUpdateOK with default headers values
func NewIpamVlansBulkUpdateOK() *IpamVlansBulkUpdateOK {
	return &IpamVlansBulkUpdateOK{}
}

/* IpamVlansBulkUpdateOK describes a response with status code 200, with default header values.

IpamVlansBulkUpdateOK ipam vlans bulk update o k
*/
type IpamVlansBulkUpdateOK struct {
	Payload *models.VLAN
}

func (o *IpamVlansBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlans/][%d] ipamVlansBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlansBulkUpdateOK) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
