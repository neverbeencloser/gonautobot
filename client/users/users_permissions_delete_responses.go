package users

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersPermissionsDeleteReader is a Reader for the UsersPermissionsDelete structure.
type UsersPermissionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersPermissionsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersPermissionsDeleteNoContent creates a UsersPermissionsDeleteNoContent with default headers values
func NewUsersPermissionsDeleteNoContent() *UsersPermissionsDeleteNoContent {
	return &UsersPermissionsDeleteNoContent{}
}

/* UsersPermissionsDeleteNoContent describes a response with status code 204, with default header values.

UsersPermissionsDeleteNoContent users permissions delete no content
*/
type UsersPermissionsDeleteNoContent struct {
}

func (o *UsersPermissionsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/permissions/{id}/][%d] usersPermissionsDeleteNoContent ", 204)
}

func (o *UsersPermissionsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
