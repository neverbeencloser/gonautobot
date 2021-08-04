package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsCircuitMaintenanceCircuitimpactBulkDeleteReader is a Reader for the PluginsCircuitMaintenanceCircuitimpactBulkDelete structure.
type PluginsCircuitMaintenanceCircuitimpactBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent creates a PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent with default headers values
func NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent() *PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent {
	return &PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent{}
}

/* PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent plugins circuit maintenance circuitimpact bulk delete no content
*/
type PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent struct {
}

func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/circuit-maintenance/circuitimpact/][%d] pluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent ", 204)
}

func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
