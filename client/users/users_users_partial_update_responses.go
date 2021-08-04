package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersUsersPartialUpdateReader is a Reader for the UsersUsersPartialUpdate structure.
type UsersUsersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersPartialUpdateOK creates a UsersUsersPartialUpdateOK with default headers values
func NewUsersUsersPartialUpdateOK() *UsersUsersPartialUpdateOK {
	return &UsersUsersPartialUpdateOK{}
}

/* UsersUsersPartialUpdateOK describes a response with status code 200, with default header values.

UsersUsersPartialUpdateOK users users partial update o k
*/
type UsersUsersPartialUpdateOK struct {
	Payload *models.User
}

func (o *UsersUsersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/users/{id}/][%d] usersUsersPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *UsersUsersPartialUpdateOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
