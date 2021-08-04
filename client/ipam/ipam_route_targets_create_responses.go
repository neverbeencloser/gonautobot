package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamRouteTargetsCreateReader is a Reader for the IpamRouteTargetsCreate structure.
type IpamRouteTargetsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamRouteTargetsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRouteTargetsCreateCreated creates a IpamRouteTargetsCreateCreated with default headers values
func NewIpamRouteTargetsCreateCreated() *IpamRouteTargetsCreateCreated {
	return &IpamRouteTargetsCreateCreated{}
}

/* IpamRouteTargetsCreateCreated describes a response with status code 201, with default header values.

IpamRouteTargetsCreateCreated ipam route targets create created
*/
type IpamRouteTargetsCreateCreated struct {
	Payload *models.RouteTarget
}

func (o *IpamRouteTargetsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/route-targets/][%d] ipamRouteTargetsCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamRouteTargetsCreateCreated) GetPayload() *models.RouteTarget {
	return o.Payload
}

func (o *IpamRouteTargetsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
