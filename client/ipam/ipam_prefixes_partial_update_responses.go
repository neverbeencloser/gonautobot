package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamPrefixesPartialUpdateReader is a Reader for the IpamPrefixesPartialUpdate structure.
type IpamPrefixesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesPartialUpdateOK creates a IpamPrefixesPartialUpdateOK with default headers values
func NewIpamPrefixesPartialUpdateOK() *IpamPrefixesPartialUpdateOK {
	return &IpamPrefixesPartialUpdateOK{}
}

/* IpamPrefixesPartialUpdateOK describes a response with status code 200, with default header values.

IpamPrefixesPartialUpdateOK ipam prefixes partial update o k
*/
type IpamPrefixesPartialUpdateOK struct {
	Payload *models.Prefix
}

func (o *IpamPrefixesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/prefixes/{id}/][%d] ipamPrefixesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamPrefixesPartialUpdateOK) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
