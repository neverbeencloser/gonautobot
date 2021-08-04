package users

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersGroupsDeleteReader is a Reader for the UsersGroupsDelete structure.
type UsersGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersGroupsDeleteNoContent creates a UsersGroupsDeleteNoContent with default headers values
func NewUsersGroupsDeleteNoContent() *UsersGroupsDeleteNoContent {
	return &UsersGroupsDeleteNoContent{}
}

/* UsersGroupsDeleteNoContent describes a response with status code 204, with default header values.

UsersGroupsDeleteNoContent users groups delete no content
*/
type UsersGroupsDeleteNoContent struct {
}

func (o *UsersGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/groups/{id}/][%d] usersGroupsDeleteNoContent ", 204)
}

func (o *UsersGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
