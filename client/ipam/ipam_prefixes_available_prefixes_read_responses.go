package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamPrefixesAvailablePrefixesReadReader is a Reader for the IpamPrefixesAvailablePrefixesRead structure.
type IpamPrefixesAvailablePrefixesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailablePrefixesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesAvailablePrefixesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesAvailablePrefixesReadOK creates a IpamPrefixesAvailablePrefixesReadOK with default headers values
func NewIpamPrefixesAvailablePrefixesReadOK() *IpamPrefixesAvailablePrefixesReadOK {
	return &IpamPrefixesAvailablePrefixesReadOK{}
}

/* IpamPrefixesAvailablePrefixesReadOK describes a response with status code 200, with default header values.

IpamPrefixesAvailablePrefixesReadOK ipam prefixes available prefixes read o k
*/
type IpamPrefixesAvailablePrefixesReadOK struct {
	Payload []*models.AvailablePrefix
}

func (o *IpamPrefixesAvailablePrefixesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-prefixes/][%d] ipamPrefixesAvailablePrefixesReadOK  %+v", 200, o.Payload)
}
func (o *IpamPrefixesAvailablePrefixesReadOK) GetPayload() []*models.AvailablePrefix {
	return o.Payload
}

func (o *IpamPrefixesAvailablePrefixesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
