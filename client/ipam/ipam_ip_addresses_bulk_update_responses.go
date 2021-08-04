package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamIPAddressesBulkUpdateReader is a Reader for the IpamIPAddressesBulkUpdate structure.
type IpamIPAddressesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPAddressesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamIPAddressesBulkUpdateOK creates a IpamIPAddressesBulkUpdateOK with default headers values
func NewIpamIPAddressesBulkUpdateOK() *IpamIPAddressesBulkUpdateOK {
	return &IpamIPAddressesBulkUpdateOK{}
}

/* IpamIPAddressesBulkUpdateOK describes a response with status code 200, with default header values.

IpamIPAddressesBulkUpdateOK ipam Ip addresses bulk update o k
*/
type IpamIPAddressesBulkUpdateOK struct {
	Payload *models.IPAddress
}

func (o *IpamIPAddressesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-addresses/][%d] ipamIpAddressesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamIPAddressesBulkUpdateOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
