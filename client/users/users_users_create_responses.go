package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersUsersCreateReader is a Reader for the UsersUsersCreate structure.
type UsersUsersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersUsersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersCreateCreated creates a UsersUsersCreateCreated with default headers values
func NewUsersUsersCreateCreated() *UsersUsersCreateCreated {
	return &UsersUsersCreateCreated{}
}

/* UsersUsersCreateCreated describes a response with status code 201, with default header values.

UsersUsersCreateCreated users users create created
*/
type UsersUsersCreateCreated struct {
	Payload *models.User
}

func (o *UsersUsersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /users/users/][%d] usersUsersCreateCreated  %+v", 201, o.Payload)
}
func (o *UsersUsersCreateCreated) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
