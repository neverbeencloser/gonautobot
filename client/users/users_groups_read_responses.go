package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersGroupsReadReader is a Reader for the UsersGroupsRead structure.
type UsersGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersGroupsReadOK creates a UsersGroupsReadOK with default headers values
func NewUsersGroupsReadOK() *UsersGroupsReadOK {
	return &UsersGroupsReadOK{}
}

/* UsersGroupsReadOK describes a response with status code 200, with default header values.

UsersGroupsReadOK users groups read o k
*/
type UsersGroupsReadOK struct {
	Payload *models.Group
}

func (o *UsersGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /users/groups/{id}/][%d] usersGroupsReadOK  %+v", 200, o.Payload)
}
func (o *UsersGroupsReadOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
