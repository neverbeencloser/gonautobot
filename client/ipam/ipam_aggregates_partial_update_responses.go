package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamAggregatesPartialUpdateReader is a Reader for the IpamAggregatesPartialUpdate structure.
type IpamAggregatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAggregatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamAggregatesPartialUpdateOK creates a IpamAggregatesPartialUpdateOK with default headers values
func NewIpamAggregatesPartialUpdateOK() *IpamAggregatesPartialUpdateOK {
	return &IpamAggregatesPartialUpdateOK{}
}

/* IpamAggregatesPartialUpdateOK describes a response with status code 200, with default header values.

IpamAggregatesPartialUpdateOK ipam aggregates partial update o k
*/
type IpamAggregatesPartialUpdateOK struct {
	Payload *models.Aggregate
}

func (o *IpamAggregatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/{id}/][%d] ipamAggregatesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamAggregatesPartialUpdateOK) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *IpamAggregatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
