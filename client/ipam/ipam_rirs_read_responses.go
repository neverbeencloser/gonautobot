package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamRirsReadReader is a Reader for the IpamRirsRead structure.
type IpamRirsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRirsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRirsReadOK creates a IpamRirsReadOK with default headers values
func NewIpamRirsReadOK() *IpamRirsReadOK {
	return &IpamRirsReadOK{}
}

/* IpamRirsReadOK describes a response with status code 200, with default header values.

IpamRirsReadOK ipam rirs read o k
*/
type IpamRirsReadOK struct {
	Payload *models.RIR
}

func (o *IpamRirsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/rirs/{id}/][%d] ipamRirsReadOK  %+v", 200, o.Payload)
}
func (o *IpamRirsReadOK) GetPayload() *models.RIR {
	return o.Payload
}

func (o *IpamRirsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RIR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
