package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceNotePartialUpdateReader is a Reader for the PluginsCircuitMaintenanceNotePartialUpdate structure.
type PluginsCircuitMaintenanceNotePartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceNotePartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceNotePartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceNotePartialUpdateOK creates a PluginsCircuitMaintenanceNotePartialUpdateOK with default headers values
func NewPluginsCircuitMaintenanceNotePartialUpdateOK() *PluginsCircuitMaintenanceNotePartialUpdateOK {
	return &PluginsCircuitMaintenanceNotePartialUpdateOK{}
}

/* PluginsCircuitMaintenanceNotePartialUpdateOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceNotePartialUpdateOK plugins circuit maintenance note partial update o k
*/
type PluginsCircuitMaintenanceNotePartialUpdateOK struct {
	Payload *models.Note
}

func (o *PluginsCircuitMaintenanceNotePartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/circuit-maintenance/note/{id}/][%d] pluginsCircuitMaintenanceNotePartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceNotePartialUpdateOK) GetPayload() *models.Note {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceNotePartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Note)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
