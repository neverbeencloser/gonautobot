package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGitRepositoriesBulkUpdateReader is a Reader for the ExtrasGitRepositoriesBulkUpdate structure.
type ExtrasGitRepositoriesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGitRepositoriesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasGitRepositoriesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGitRepositoriesBulkUpdateOK creates a ExtrasGitRepositoriesBulkUpdateOK with default headers values
func NewExtrasGitRepositoriesBulkUpdateOK() *ExtrasGitRepositoriesBulkUpdateOK {
	return &ExtrasGitRepositoriesBulkUpdateOK{}
}

/* ExtrasGitRepositoriesBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasGitRepositoriesBulkUpdateOK extras git repositories bulk update o k
*/
type ExtrasGitRepositoriesBulkUpdateOK struct {
	Payload *models.GitRepository
}

func (o *ExtrasGitRepositoriesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/git-repositories/][%d] extrasGitRepositoriesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasGitRepositoriesBulkUpdateOK) GetPayload() *models.GitRepository {
	return o.Payload
}

func (o *ExtrasGitRepositoriesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
