package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersPermissionsBulkUpdateReader is a Reader for the UsersPermissionsBulkUpdate structure.
type UsersPermissionsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersPermissionsBulkUpdateOK creates a UsersPermissionsBulkUpdateOK with default headers values
func NewUsersPermissionsBulkUpdateOK() *UsersPermissionsBulkUpdateOK {
	return &UsersPermissionsBulkUpdateOK{}
}

/* UsersPermissionsBulkUpdateOK describes a response with status code 200, with default header values.

UsersPermissionsBulkUpdateOK users permissions bulk update o k
*/
type UsersPermissionsBulkUpdateOK struct {
	Payload *models.ObjectPermission
}

func (o *UsersPermissionsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/permissions/][%d] usersPermissionsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *UsersPermissionsBulkUpdateOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
