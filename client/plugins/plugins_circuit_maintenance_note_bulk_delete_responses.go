package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsCircuitMaintenanceNoteBulkDeleteReader is a Reader for the PluginsCircuitMaintenanceNoteBulkDelete structure.
type PluginsCircuitMaintenanceNoteBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceNoteBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsCircuitMaintenanceNoteBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceNoteBulkDeleteNoContent creates a PluginsCircuitMaintenanceNoteBulkDeleteNoContent with default headers values
func NewPluginsCircuitMaintenanceNoteBulkDeleteNoContent() *PluginsCircuitMaintenanceNoteBulkDeleteNoContent {
	return &PluginsCircuitMaintenanceNoteBulkDeleteNoContent{}
}

/* PluginsCircuitMaintenanceNoteBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsCircuitMaintenanceNoteBulkDeleteNoContent plugins circuit maintenance note bulk delete no content
*/
type PluginsCircuitMaintenanceNoteBulkDeleteNoContent struct {
}

func (o *PluginsCircuitMaintenanceNoteBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/circuit-maintenance/note/][%d] pluginsCircuitMaintenanceNoteBulkDeleteNoContent ", 204)
}

func (o *PluginsCircuitMaintenanceNoteBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
