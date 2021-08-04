package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasGitRepositoriesDeleteReader is a Reader for the ExtrasGitRepositoriesDelete structure.
type ExtrasGitRepositoriesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGitRepositoriesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasGitRepositoriesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGitRepositoriesDeleteNoContent creates a ExtrasGitRepositoriesDeleteNoContent with default headers values
func NewExtrasGitRepositoriesDeleteNoContent() *ExtrasGitRepositoriesDeleteNoContent {
	return &ExtrasGitRepositoriesDeleteNoContent{}
}

/* ExtrasGitRepositoriesDeleteNoContent describes a response with status code 204, with default header values.

ExtrasGitRepositoriesDeleteNoContent extras git repositories delete no content
*/
type ExtrasGitRepositoriesDeleteNoContent struct {
}

func (o *ExtrasGitRepositoriesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/git-repositories/{id}/][%d] extrasGitRepositoriesDeleteNoContent ", 204)
}

func (o *ExtrasGitRepositoriesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
