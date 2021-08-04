package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersUsersReadReader is a Reader for the UsersUsersRead structure.
type UsersUsersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersReadOK creates a UsersUsersReadOK with default headers values
func NewUsersUsersReadOK() *UsersUsersReadOK {
	return &UsersUsersReadOK{}
}

/* UsersUsersReadOK describes a response with status code 200, with default header values.

UsersUsersReadOK users users read o k
*/
type UsersUsersReadOK struct {
	Payload *models.User
}

func (o *UsersUsersReadOK) Error() string {
	return fmt.Sprintf("[GET /users/users/{id}/][%d] usersUsersReadOK  %+v", 200, o.Payload)
}
func (o *UsersUsersReadOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
