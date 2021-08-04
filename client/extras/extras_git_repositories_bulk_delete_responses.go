package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasGitRepositoriesBulkDeleteReader is a Reader for the ExtrasGitRepositoriesBulkDelete structure.
type ExtrasGitRepositoriesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGitRepositoriesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasGitRepositoriesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGitRepositoriesBulkDeleteNoContent creates a ExtrasGitRepositoriesBulkDeleteNoContent with default headers values
func NewExtrasGitRepositoriesBulkDeleteNoContent() *ExtrasGitRepositoriesBulkDeleteNoContent {
	return &ExtrasGitRepositoriesBulkDeleteNoContent{}
}

/* ExtrasGitRepositoriesBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasGitRepositoriesBulkDeleteNoContent extras git repositories bulk delete no content
*/
type ExtrasGitRepositoriesBulkDeleteNoContent struct {
}

func (o *ExtrasGitRepositoriesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/git-repositories/][%d] extrasGitRepositoriesBulkDeleteNoContent ", 204)
}

func (o *ExtrasGitRepositoriesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
