package users

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersPermissionsBulkDeleteReader is a Reader for the UsersPermissionsBulkDelete structure.
type UsersPermissionsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersPermissionsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersPermissionsBulkDeleteNoContent creates a UsersPermissionsBulkDeleteNoContent with default headers values
func NewUsersPermissionsBulkDeleteNoContent() *UsersPermissionsBulkDeleteNoContent {
	return &UsersPermissionsBulkDeleteNoContent{}
}

/* UsersPermissionsBulkDeleteNoContent describes a response with status code 204, with default header values.

UsersPermissionsBulkDeleteNoContent users permissions bulk delete no content
*/
type UsersPermissionsBulkDeleteNoContent struct {
}

func (o *UsersPermissionsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/permissions/][%d] usersPermissionsBulkDeleteNoContent ", 204)
}

func (o *UsersPermissionsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
