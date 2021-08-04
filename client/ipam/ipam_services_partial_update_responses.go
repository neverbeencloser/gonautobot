package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamServicesPartialUpdateReader is a Reader for the IpamServicesPartialUpdate structure.
type IpamServicesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServicesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamServicesPartialUpdateOK creates a IpamServicesPartialUpdateOK with default headers values
func NewIpamServicesPartialUpdateOK() *IpamServicesPartialUpdateOK {
	return &IpamServicesPartialUpdateOK{}
}

/* IpamServicesPartialUpdateOK describes a response with status code 200, with default header values.

IpamServicesPartialUpdateOK ipam services partial update o k
*/
type IpamServicesPartialUpdateOK struct {
	Payload *models.Service
}

func (o *IpamServicesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/services/{id}/][%d] ipamServicesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamServicesPartialUpdateOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
