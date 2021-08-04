package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamIPAddressesBulkPartialUpdateReader is a Reader for the IpamIPAddressesBulkPartialUpdate structure.
type IpamIPAddressesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPAddressesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamIPAddressesBulkPartialUpdateOK creates a IpamIPAddressesBulkPartialUpdateOK with default headers values
func NewIpamIPAddressesBulkPartialUpdateOK() *IpamIPAddressesBulkPartialUpdateOK {
	return &IpamIPAddressesBulkPartialUpdateOK{}
}

/* IpamIPAddressesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamIPAddressesBulkPartialUpdateOK ipam Ip addresses bulk partial update o k
*/
type IpamIPAddressesBulkPartialUpdateOK struct {
	Payload *models.IPAddress
}

func (o *IpamIPAddressesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/][%d] ipamIpAddressesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamIPAddressesBulkPartialUpdateOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
