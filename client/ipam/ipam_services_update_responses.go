package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamServicesUpdateReader is a Reader for the IpamServicesUpdate structure.
type IpamServicesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServicesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamServicesUpdateOK creates a IpamServicesUpdateOK with default headers values
func NewIpamServicesUpdateOK() *IpamServicesUpdateOK {
	return &IpamServicesUpdateOK{}
}

/* IpamServicesUpdateOK describes a response with status code 200, with default header values.

IpamServicesUpdateOK ipam services update o k
*/
type IpamServicesUpdateOK struct {
	Payload *models.Service
}

func (o *IpamServicesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/services/{id}/][%d] ipamServicesUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamServicesUpdateOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
