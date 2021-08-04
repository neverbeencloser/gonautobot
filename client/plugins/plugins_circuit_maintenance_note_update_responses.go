package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceNoteUpdateReader is a Reader for the PluginsCircuitMaintenanceNoteUpdate structure.
type PluginsCircuitMaintenanceNoteUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceNoteUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceNoteUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceNoteUpdateOK creates a PluginsCircuitMaintenanceNoteUpdateOK with default headers values
func NewPluginsCircuitMaintenanceNoteUpdateOK() *PluginsCircuitMaintenanceNoteUpdateOK {
	return &PluginsCircuitMaintenanceNoteUpdateOK{}
}

/* PluginsCircuitMaintenanceNoteUpdateOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceNoteUpdateOK plugins circuit maintenance note update o k
*/
type PluginsCircuitMaintenanceNoteUpdateOK struct {
	Payload *models.Note
}

func (o *PluginsCircuitMaintenanceNoteUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/circuit-maintenance/note/{id}/][%d] pluginsCircuitMaintenanceNoteUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceNoteUpdateOK) GetPayload() *models.Note {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceNoteUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Note)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
