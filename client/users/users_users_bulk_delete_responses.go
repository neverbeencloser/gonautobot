package users

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersUsersBulkDeleteReader is a Reader for the UsersUsersBulkDelete structure.
type UsersUsersBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersUsersBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersBulkDeleteNoContent creates a UsersUsersBulkDeleteNoContent with default headers values
func NewUsersUsersBulkDeleteNoContent() *UsersUsersBulkDeleteNoContent {
	return &UsersUsersBulkDeleteNoContent{}
}

/* UsersUsersBulkDeleteNoContent describes a response with status code 204, with default header values.

UsersUsersBulkDeleteNoContent users users bulk delete no content
*/
type UsersUsersBulkDeleteNoContent struct {
}

func (o *UsersUsersBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/users/][%d] usersUsersBulkDeleteNoContent ", 204)
}

func (o *UsersUsersBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
