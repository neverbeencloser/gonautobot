package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersUsersBulkPartialUpdateReader is a Reader for the UsersUsersBulkPartialUpdate structure.
type UsersUsersBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersBulkPartialUpdateOK creates a UsersUsersBulkPartialUpdateOK with default headers values
func NewUsersUsersBulkPartialUpdateOK() *UsersUsersBulkPartialUpdateOK {
	return &UsersUsersBulkPartialUpdateOK{}
}

/* UsersUsersBulkPartialUpdateOK describes a response with status code 200, with default header values.

UsersUsersBulkPartialUpdateOK users users bulk partial update o k
*/
type UsersUsersBulkPartialUpdateOK struct {
	Payload *models.User
}

func (o *UsersUsersBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/users/][%d] usersUsersBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *UsersUsersBulkPartialUpdateOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
