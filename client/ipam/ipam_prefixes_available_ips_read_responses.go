package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamPrefixesAvailableIpsReadReader is a Reader for the IpamPrefixesAvailableIpsRead structure.
type IpamPrefixesAvailableIpsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailableIpsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesAvailableIpsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesAvailableIpsReadOK creates a IpamPrefixesAvailableIpsReadOK with default headers values
func NewIpamPrefixesAvailableIpsReadOK() *IpamPrefixesAvailableIpsReadOK {
	return &IpamPrefixesAvailableIpsReadOK{}
}

/* IpamPrefixesAvailableIpsReadOK describes a response with status code 200, with default header values.

IpamPrefixesAvailableIpsReadOK ipam prefixes available ips read o k
*/
type IpamPrefixesAvailableIpsReadOK struct {
	Payload []*models.AvailableIP
}

func (o *IpamPrefixesAvailableIpsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsReadOK  %+v", 200, o.Payload)
}
func (o *IpamPrefixesAvailableIpsReadOK) GetPayload() []*models.AvailableIP {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
