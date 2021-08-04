package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamServicesReadReader is a Reader for the IpamServicesRead structure.
type IpamServicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamServicesReadOK creates a IpamServicesReadOK with default headers values
func NewIpamServicesReadOK() *IpamServicesReadOK {
	return &IpamServicesReadOK{}
}

/* IpamServicesReadOK describes a response with status code 200, with default header values.

IpamServicesReadOK ipam services read o k
*/
type IpamServicesReadOK struct {
	Payload *models.Service
}

func (o *IpamServicesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/services/{id}/][%d] ipamServicesReadOK  %+v", 200, o.Payload)
}
func (o *IpamServicesReadOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
