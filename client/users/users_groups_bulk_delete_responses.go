package users

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersGroupsBulkDeleteReader is a Reader for the UsersGroupsBulkDelete structure.
type UsersGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersGroupsBulkDeleteNoContent creates a UsersGroupsBulkDeleteNoContent with default headers values
func NewUsersGroupsBulkDeleteNoContent() *UsersGroupsBulkDeleteNoContent {
	return &UsersGroupsBulkDeleteNoContent{}
}

/* UsersGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

UsersGroupsBulkDeleteNoContent users groups bulk delete no content
*/
type UsersGroupsBulkDeleteNoContent struct {
}

func (o *UsersGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/groups/][%d] usersGroupsBulkDeleteNoContent ", 204)
}

func (o *UsersGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
