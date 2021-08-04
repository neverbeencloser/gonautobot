package users

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersUsersDeleteReader is a Reader for the UsersUsersDelete structure.
type UsersUsersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersUsersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersDeleteNoContent creates a UsersUsersDeleteNoContent with default headers values
func NewUsersUsersDeleteNoContent() *UsersUsersDeleteNoContent {
	return &UsersUsersDeleteNoContent{}
}

/* UsersUsersDeleteNoContent describes a response with status code 204, with default header values.

UsersUsersDeleteNoContent users users delete no content
*/
type UsersUsersDeleteNoContent struct {
}

func (o *UsersUsersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/users/{id}/][%d] usersUsersDeleteNoContent ", 204)
}

func (o *UsersUsersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
