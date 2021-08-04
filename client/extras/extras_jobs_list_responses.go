package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasJobsListReader is a Reader for the ExtrasJobsList structure.
type ExtrasJobsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJobsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJobsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasJobsListOK creates a ExtrasJobsListOK with default headers values
func NewExtrasJobsListOK() *ExtrasJobsListOK {
	return &ExtrasJobsListOK{}
}

/* ExtrasJobsListOK describes a response with status code 200, with default header values.

ExtrasJobsListOK extras jobs list o k
*/
type ExtrasJobsListOK struct {
}

func (o *ExtrasJobsListOK) Error() string {
	return fmt.Sprintf("[GET /extras/jobs/][%d] extrasJobsListOK ", 200)
}

func (o *ExtrasJobsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
