package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamPrefixesCreateReader is a Reader for the IpamPrefixesCreate structure.
type IpamPrefixesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamPrefixesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesCreateCreated creates a IpamPrefixesCreateCreated with default headers values
func NewIpamPrefixesCreateCreated() *IpamPrefixesCreateCreated {
	return &IpamPrefixesCreateCreated{}
}

/* IpamPrefixesCreateCreated describes a response with status code 201, with default header values.

IpamPrefixesCreateCreated ipam prefixes create created
*/
type IpamPrefixesCreateCreated struct {
	Payload *models.Prefix
}

func (o *IpamPrefixesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/][%d] ipamPrefixesCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamPrefixesCreateCreated) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
