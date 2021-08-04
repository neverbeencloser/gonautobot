package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGitRepositoriesSyncReader is a Reader for the ExtrasGitRepositoriesSync structure.
type ExtrasGitRepositoriesSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGitRepositoriesSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasGitRepositoriesSyncCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGitRepositoriesSyncCreated creates a ExtrasGitRepositoriesSyncCreated with default headers values
func NewExtrasGitRepositoriesSyncCreated() *ExtrasGitRepositoriesSyncCreated {
	return &ExtrasGitRepositoriesSyncCreated{}
}

/* ExtrasGitRepositoriesSyncCreated describes a response with status code 201, with default header values.

ExtrasGitRepositoriesSyncCreated extras git repositories sync created
*/
type ExtrasGitRepositoriesSyncCreated struct {
	Payload *models.GitRepository
}

func (o *ExtrasGitRepositoriesSyncCreated) Error() string {
	return fmt.Sprintf("[POST /extras/git-repositories/{id}/sync/][%d] extrasGitRepositoriesSyncCreated  %+v", 201, o.Payload)
}
func (o *ExtrasGitRepositoriesSyncCreated) GetPayload() *models.GitRepository {
	return o.Payload
}

func (o *ExtrasGitRepositoriesSyncCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
