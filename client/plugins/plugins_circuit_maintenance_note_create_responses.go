package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceNoteCreateReader is a Reader for the PluginsCircuitMaintenanceNoteCreate structure.
type PluginsCircuitMaintenanceNoteCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceNoteCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsCircuitMaintenanceNoteCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceNoteCreateCreated creates a PluginsCircuitMaintenanceNoteCreateCreated with default headers values
func NewPluginsCircuitMaintenanceNoteCreateCreated() *PluginsCircuitMaintenanceNoteCreateCreated {
	return &PluginsCircuitMaintenanceNoteCreateCreated{}
}

/* PluginsCircuitMaintenanceNoteCreateCreated describes a response with status code 201, with default header values.

PluginsCircuitMaintenanceNoteCreateCreated plugins circuit maintenance note create created
*/
type PluginsCircuitMaintenanceNoteCreateCreated struct {
	Payload *models.Note
}

func (o *PluginsCircuitMaintenanceNoteCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/circuit-maintenance/note/][%d] pluginsCircuitMaintenanceNoteCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsCircuitMaintenanceNoteCreateCreated) GetPayload() *models.Note {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceNoteCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Note)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
