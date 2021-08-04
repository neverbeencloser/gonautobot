package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceNoteReadReader is a Reader for the PluginsCircuitMaintenanceNoteRead structure.
type PluginsCircuitMaintenanceNoteReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceNoteReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceNoteReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceNoteReadOK creates a PluginsCircuitMaintenanceNoteReadOK with default headers values
func NewPluginsCircuitMaintenanceNoteReadOK() *PluginsCircuitMaintenanceNoteReadOK {
	return &PluginsCircuitMaintenanceNoteReadOK{}
}

/* PluginsCircuitMaintenanceNoteReadOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceNoteReadOK plugins circuit maintenance note read o k
*/
type PluginsCircuitMaintenanceNoteReadOK struct {
	Payload *models.Note
}

func (o *PluginsCircuitMaintenanceNoteReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/circuit-maintenance/note/{id}/][%d] pluginsCircuitMaintenanceNoteReadOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceNoteReadOK) GetPayload() *models.Note {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceNoteReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Note)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
