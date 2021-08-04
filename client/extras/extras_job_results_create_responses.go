package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasJobResultsCreateReader is a Reader for the ExtrasJobResultsCreate structure.
type ExtrasJobResultsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJobResultsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasJobResultsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasJobResultsCreateCreated creates a ExtrasJobResultsCreateCreated with default headers values
func NewExtrasJobResultsCreateCreated() *ExtrasJobResultsCreateCreated {
	return &ExtrasJobResultsCreateCreated{}
}

/* ExtrasJobResultsCreateCreated describes a response with status code 201, with default header values.

ExtrasJobResultsCreateCreated extras job results create created
*/
type ExtrasJobResultsCreateCreated struct {
	Payload *models.JobResult
}

func (o *ExtrasJobResultsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/job-results/][%d] extrasJobResultsCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasJobResultsCreateCreated) GetPayload() *models.JobResult {
	return o.Payload
}

func (o *ExtrasJobResultsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
