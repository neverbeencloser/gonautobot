package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamPrefixesBulkUpdateReader is a Reader for the IpamPrefixesBulkUpdate structure.
type IpamPrefixesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesBulkUpdateOK creates a IpamPrefixesBulkUpdateOK with default headers values
func NewIpamPrefixesBulkUpdateOK() *IpamPrefixesBulkUpdateOK {
	return &IpamPrefixesBulkUpdateOK{}
}

/* IpamPrefixesBulkUpdateOK describes a response with status code 200, with default header values.

IpamPrefixesBulkUpdateOK ipam prefixes bulk update o k
*/
type IpamPrefixesBulkUpdateOK struct {
	Payload *models.Prefix
}

func (o *IpamPrefixesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/prefixes/][%d] ipamPrefixesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamPrefixesBulkUpdateOK) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
