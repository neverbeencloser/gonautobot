package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersGroupsBulkUpdateReader is a Reader for the UsersGroupsBulkUpdate structure.
type UsersGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersGroupsBulkUpdateOK creates a UsersGroupsBulkUpdateOK with default headers values
func NewUsersGroupsBulkUpdateOK() *UsersGroupsBulkUpdateOK {
	return &UsersGroupsBulkUpdateOK{}
}

/* UsersGroupsBulkUpdateOK describes a response with status code 200, with default header values.

UsersGroupsBulkUpdateOK users groups bulk update o k
*/
type UsersGroupsBulkUpdateOK struct {
	Payload *models.Group
}

func (o *UsersGroupsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/groups/][%d] usersGroupsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *UsersGroupsBulkUpdateOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
