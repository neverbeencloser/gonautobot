package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasJobResultsBulkUpdateReader is a Reader for the ExtrasJobResultsBulkUpdate structure.
type ExtrasJobResultsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJobResultsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJobResultsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasJobResultsBulkUpdateOK creates a ExtrasJobResultsBulkUpdateOK with default headers values
func NewExtrasJobResultsBulkUpdateOK() *ExtrasJobResultsBulkUpdateOK {
	return &ExtrasJobResultsBulkUpdateOK{}
}

/* ExtrasJobResultsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasJobResultsBulkUpdateOK extras job results bulk update o k
*/
type ExtrasJobResultsBulkUpdateOK struct {
	Payload *models.JobResult
}

func (o *ExtrasJobResultsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/job-results/][%d] extrasJobResultsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasJobResultsBulkUpdateOK) GetPayload() *models.JobResult {
	return o.Payload
}

func (o *ExtrasJobResultsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
