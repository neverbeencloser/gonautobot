package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamAggregatesBulkUpdateReader is a Reader for the IpamAggregatesBulkUpdate structure.
type IpamAggregatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAggregatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamAggregatesBulkUpdateOK creates a IpamAggregatesBulkUpdateOK with default headers values
func NewIpamAggregatesBulkUpdateOK() *IpamAggregatesBulkUpdateOK {
	return &IpamAggregatesBulkUpdateOK{}
}

/* IpamAggregatesBulkUpdateOK describes a response with status code 200, with default header values.

IpamAggregatesBulkUpdateOK ipam aggregates bulk update o k
*/
type IpamAggregatesBulkUpdateOK struct {
	Payload *models.Aggregate
}

func (o *IpamAggregatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/aggregates/][%d] ipamAggregatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamAggregatesBulkUpdateOK) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *IpamAggregatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
