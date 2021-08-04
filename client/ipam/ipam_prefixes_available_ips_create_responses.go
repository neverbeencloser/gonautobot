package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamPrefixesAvailableIpsCreateReader is a Reader for the IpamPrefixesAvailableIpsCreate structure.
type IpamPrefixesAvailableIpsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailableIpsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamPrefixesAvailableIpsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesAvailableIpsCreateCreated creates a IpamPrefixesAvailableIpsCreateCreated with default headers values
func NewIpamPrefixesAvailableIpsCreateCreated() *IpamPrefixesAvailableIpsCreateCreated {
	return &IpamPrefixesAvailableIpsCreateCreated{}
}

/* IpamPrefixesAvailableIpsCreateCreated describes a response with status code 201, with default header values.

IpamPrefixesAvailableIpsCreateCreated ipam prefixes available ips create created
*/
type IpamPrefixesAvailableIpsCreateCreated struct {
	Payload []*models.AvailableIP
}

func (o *IpamPrefixesAvailableIpsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamPrefixesAvailableIpsCreateCreated) GetPayload() []*models.AvailableIP {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
