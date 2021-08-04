package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGitRepositoriesReadReader is a Reader for the ExtrasGitRepositoriesRead structure.
type ExtrasGitRepositoriesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGitRepositoriesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasGitRepositoriesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGitRepositoriesReadOK creates a ExtrasGitRepositoriesReadOK with default headers values
func NewExtrasGitRepositoriesReadOK() *ExtrasGitRepositoriesReadOK {
	return &ExtrasGitRepositoriesReadOK{}
}

/* ExtrasGitRepositoriesReadOK describes a response with status code 200, with default header values.

ExtrasGitRepositoriesReadOK extras git repositories read o k
*/
type ExtrasGitRepositoriesReadOK struct {
	Payload *models.GitRepository
}

func (o *ExtrasGitRepositoriesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/git-repositories/{id}/][%d] extrasGitRepositoriesReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasGitRepositoriesReadOK) GetPayload() *models.GitRepository {
	return o.Payload
}

func (o *ExtrasGitRepositoriesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
