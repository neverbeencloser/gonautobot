package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGitRepositoriesUpdateReader is a Reader for the ExtrasGitRepositoriesUpdate structure.
type ExtrasGitRepositoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGitRepositoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasGitRepositoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGitRepositoriesUpdateOK creates a ExtrasGitRepositoriesUpdateOK with default headers values
func NewExtrasGitRepositoriesUpdateOK() *ExtrasGitRepositoriesUpdateOK {
	return &ExtrasGitRepositoriesUpdateOK{}
}

/* ExtrasGitRepositoriesUpdateOK describes a response with status code 200, with default header values.

ExtrasGitRepositoriesUpdateOK extras git repositories update o k
*/
type ExtrasGitRepositoriesUpdateOK struct {
	Payload *models.GitRepository
}

func (o *ExtrasGitRepositoriesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/git-repositories/{id}/][%d] extrasGitRepositoriesUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasGitRepositoriesUpdateOK) GetPayload() *models.GitRepository {
	return o.Payload
}

func (o *ExtrasGitRepositoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
