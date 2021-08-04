package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasJobsReadReader is a Reader for the ExtrasJobsRead structure.
type ExtrasJobsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJobsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJobsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasJobsReadOK creates a ExtrasJobsReadOK with default headers values
func NewExtrasJobsReadOK() *ExtrasJobsReadOK {
	return &ExtrasJobsReadOK{}
}

/* ExtrasJobsReadOK describes a response with status code 200, with default header values.

ExtrasJobsReadOK extras jobs read o k
*/
type ExtrasJobsReadOK struct {
}

func (o *ExtrasJobsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/jobs/{class_path}/][%d] extrasJobsReadOK ", 200)
}

func (o *ExtrasJobsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
