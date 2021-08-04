package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersPermissionsPartialUpdateReader is a Reader for the UsersPermissionsPartialUpdate structure.
type UsersPermissionsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersPermissionsPartialUpdateOK creates a UsersPermissionsPartialUpdateOK with default headers values
func NewUsersPermissionsPartialUpdateOK() *UsersPermissionsPartialUpdateOK {
	return &UsersPermissionsPartialUpdateOK{}
}

/* UsersPermissionsPartialUpdateOK describes a response with status code 200, with default header values.

UsersPermissionsPartialUpdateOK users permissions partial update o k
*/
type UsersPermissionsPartialUpdateOK struct {
	Payload *models.ObjectPermission
}

func (o *UsersPermissionsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/permissions/{id}/][%d] usersPermissionsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *UsersPermissionsPartialUpdateOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
