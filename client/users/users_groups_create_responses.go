package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersGroupsCreateReader is a Reader for the UsersGroupsCreate structure.
type UsersGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersGroupsCreateCreated creates a UsersGroupsCreateCreated with default headers values
func NewUsersGroupsCreateCreated() *UsersGroupsCreateCreated {
	return &UsersGroupsCreateCreated{}
}

/* UsersGroupsCreateCreated describes a response with status code 201, with default header values.

UsersGroupsCreateCreated users groups create created
*/
type UsersGroupsCreateCreated struct {
	Payload *models.Group
}

func (o *UsersGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /users/groups/][%d] usersGroupsCreateCreated  %+v", 201, o.Payload)
}
func (o *UsersGroupsCreateCreated) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
