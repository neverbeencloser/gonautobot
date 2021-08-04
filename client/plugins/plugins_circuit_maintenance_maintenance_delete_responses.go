package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsCircuitMaintenanceMaintenanceDeleteReader is a Reader for the PluginsCircuitMaintenanceMaintenanceDelete structure.
type PluginsCircuitMaintenanceMaintenanceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceMaintenanceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsCircuitMaintenanceMaintenanceDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceMaintenanceDeleteNoContent creates a PluginsCircuitMaintenanceMaintenanceDeleteNoContent with default headers values
func NewPluginsCircuitMaintenanceMaintenanceDeleteNoContent() *PluginsCircuitMaintenanceMaintenanceDeleteNoContent {
	return &PluginsCircuitMaintenanceMaintenanceDeleteNoContent{}
}

/* PluginsCircuitMaintenanceMaintenanceDeleteNoContent describes a response with status code 204, with default header values.

PluginsCircuitMaintenanceMaintenanceDeleteNoContent plugins circuit maintenance maintenance delete no content
*/
type PluginsCircuitMaintenanceMaintenanceDeleteNoContent struct {
}

func (o *PluginsCircuitMaintenanceMaintenanceDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/circuit-maintenance/maintenance/{id}/][%d] pluginsCircuitMaintenanceMaintenanceDeleteNoContent ", 204)
}

func (o *PluginsCircuitMaintenanceMaintenanceDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
