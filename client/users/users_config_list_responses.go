package users

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersConfigListReader is a Reader for the UsersConfigList structure.
type UsersConfigListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersConfigListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersConfigListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersConfigListOK creates a UsersConfigListOK with default headers values
func NewUsersConfigListOK() *UsersConfigListOK {
	return &UsersConfigListOK{}
}

/* UsersConfigListOK describes a response with status code 200, with default header values.

UsersConfigListOK users config list o k
*/
type UsersConfigListOK struct {
}

func (o *UsersConfigListOK) Error() string {
	return fmt.Sprintf("[GET /users/config/][%d] usersConfigListOK ", 200)
}

func (o *UsersConfigListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
