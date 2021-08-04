package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsCircuitMaintenanceNoteDeleteReader is a Reader for the PluginsCircuitMaintenanceNoteDelete structure.
type PluginsCircuitMaintenanceNoteDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceNoteDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsCircuitMaintenanceNoteDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceNoteDeleteNoContent creates a PluginsCircuitMaintenanceNoteDeleteNoContent with default headers values
func NewPluginsCircuitMaintenanceNoteDeleteNoContent() *PluginsCircuitMaintenanceNoteDeleteNoContent {
	return &PluginsCircuitMaintenanceNoteDeleteNoContent{}
}

/* PluginsCircuitMaintenanceNoteDeleteNoContent describes a response with status code 204, with default header values.

PluginsCircuitMaintenanceNoteDeleteNoContent plugins circuit maintenance note delete no content
*/
type PluginsCircuitMaintenanceNoteDeleteNoContent struct {
}

func (o *PluginsCircuitMaintenanceNoteDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/circuit-maintenance/note/{id}/][%d] pluginsCircuitMaintenanceNoteDeleteNoContent ", 204)
}

func (o *PluginsCircuitMaintenanceNoteDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
