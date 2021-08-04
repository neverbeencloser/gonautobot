package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsCircuitMaintenanceMaintenanceBulkDeleteReader is a Reader for the PluginsCircuitMaintenanceMaintenanceBulkDelete structure.
type PluginsCircuitMaintenanceMaintenanceBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent creates a PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent with default headers values
func NewPluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent() *PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent {
	return &PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent{}
}

/* PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent plugins circuit maintenance maintenance bulk delete no content
*/
type PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent struct {
}

func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/circuit-maintenance/maintenance/][%d] pluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent ", 204)
}

func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
