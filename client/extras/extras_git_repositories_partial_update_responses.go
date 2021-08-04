package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGitRepositoriesPartialUpdateReader is a Reader for the ExtrasGitRepositoriesPartialUpdate structure.
type ExtrasGitRepositoriesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGitRepositoriesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasGitRepositoriesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGitRepositoriesPartialUpdateOK creates a ExtrasGitRepositoriesPartialUpdateOK with default headers values
func NewExtrasGitRepositoriesPartialUpdateOK() *ExtrasGitRepositoriesPartialUpdateOK {
	return &ExtrasGitRepositoriesPartialUpdateOK{}
}

/* ExtrasGitRepositoriesPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasGitRepositoriesPartialUpdateOK extras git repositories partial update o k
*/
type ExtrasGitRepositoriesPartialUpdateOK struct {
	Payload *models.GitRepository
}

func (o *ExtrasGitRepositoriesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/git-repositories/{id}/][%d] extrasGitRepositoriesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasGitRepositoriesPartialUpdateOK) GetPayload() *models.GitRepository {
	return o.Payload
}

func (o *ExtrasGitRepositoriesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
