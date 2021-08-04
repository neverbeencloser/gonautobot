package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGitRepositoriesCreateReader is a Reader for the ExtrasGitRepositoriesCreate structure.
type ExtrasGitRepositoriesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGitRepositoriesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasGitRepositoriesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGitRepositoriesCreateCreated creates a ExtrasGitRepositoriesCreateCreated with default headers values
func NewExtrasGitRepositoriesCreateCreated() *ExtrasGitRepositoriesCreateCreated {
	return &ExtrasGitRepositoriesCreateCreated{}
}

/* ExtrasGitRepositoriesCreateCreated describes a response with status code 201, with default header values.

ExtrasGitRepositoriesCreateCreated extras git repositories create created
*/
type ExtrasGitRepositoriesCreateCreated struct {
	Payload *models.GitRepository
}

func (o *ExtrasGitRepositoriesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/git-repositories/][%d] extrasGitRepositoriesCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasGitRepositoriesCreateCreated) GetPayload() *models.GitRepository {
	return o.Payload
}

func (o *ExtrasGitRepositoriesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
