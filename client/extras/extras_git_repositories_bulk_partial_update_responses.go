package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGitRepositoriesBulkPartialUpdateReader is a Reader for the ExtrasGitRepositoriesBulkPartialUpdate structure.
type ExtrasGitRepositoriesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGitRepositoriesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasGitRepositoriesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGitRepositoriesBulkPartialUpdateOK creates a ExtrasGitRepositoriesBulkPartialUpdateOK with default headers values
func NewExtrasGitRepositoriesBulkPartialUpdateOK() *ExtrasGitRepositoriesBulkPartialUpdateOK {
	return &ExtrasGitRepositoriesBulkPartialUpdateOK{}
}

/* ExtrasGitRepositoriesBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasGitRepositoriesBulkPartialUpdateOK extras git repositories bulk partial update o k
*/
type ExtrasGitRepositoriesBulkPartialUpdateOK struct {
	Payload *models.GitRepository
}

func (o *ExtrasGitRepositoriesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/git-repositories/][%d] extrasGitRepositoriesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasGitRepositoriesBulkPartialUpdateOK) GetPayload() *models.GitRepository {
	return o.Payload
}

func (o *ExtrasGitRepositoriesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
