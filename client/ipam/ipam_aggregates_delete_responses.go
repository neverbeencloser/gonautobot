package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamAggregatesDeleteReader is a Reader for the IpamAggregatesDelete structure.
type IpamAggregatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamAggregatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamAggregatesDeleteNoContent creates a IpamAggregatesDeleteNoContent with default headers values
func NewIpamAggregatesDeleteNoContent() *IpamAggregatesDeleteNoContent {
	return &IpamAggregatesDeleteNoContent{}
}

/* IpamAggregatesDeleteNoContent describes a response with status code 204, with default header values.

IpamAggregatesDeleteNoContent ipam aggregates delete no content
*/
type IpamAggregatesDeleteNoContent struct {
}

func (o *IpamAggregatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/aggregates/{id}/][%d] ipamAggregatesDeleteNoContent ", 204)
}

func (o *IpamAggregatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
