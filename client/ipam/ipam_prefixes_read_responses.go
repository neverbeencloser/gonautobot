package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamPrefixesReadReader is a Reader for the IpamPrefixesRead structure.
type IpamPrefixesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesReadOK creates a IpamPrefixesReadOK with default headers values
func NewIpamPrefixesReadOK() *IpamPrefixesReadOK {
	return &IpamPrefixesReadOK{}
}

/* IpamPrefixesReadOK describes a response with status code 200, with default header values.

IpamPrefixesReadOK ipam prefixes read o k
*/
type IpamPrefixesReadOK struct {
	Payload *models.Prefix
}

func (o *IpamPrefixesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/][%d] ipamPrefixesReadOK  %+v", 200, o.Payload)
}
func (o *IpamPrefixesReadOK) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
