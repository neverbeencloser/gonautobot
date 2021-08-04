package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersUsersUpdateReader is a Reader for the UsersUsersUpdate structure.
type UsersUsersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersUpdateOK creates a UsersUsersUpdateOK with default headers values
func NewUsersUsersUpdateOK() *UsersUsersUpdateOK {
	return &UsersUsersUpdateOK{}
}

/* UsersUsersUpdateOK describes a response with status code 200, with default header values.

UsersUsersUpdateOK users users update o k
*/
type UsersUsersUpdateOK struct {
	Payload *models.User
}

func (o *UsersUsersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/users/{id}/][%d] usersUsersUpdateOK  %+v", 200, o.Payload)
}
func (o *UsersUsersUpdateOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
