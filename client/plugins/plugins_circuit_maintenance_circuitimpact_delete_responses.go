package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsCircuitMaintenanceCircuitimpactDeleteReader is a Reader for the PluginsCircuitMaintenanceCircuitimpactDelete structure.
type PluginsCircuitMaintenanceCircuitimpactDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceCircuitimpactDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsCircuitMaintenanceCircuitimpactDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactDeleteNoContent creates a PluginsCircuitMaintenanceCircuitimpactDeleteNoContent with default headers values
func NewPluginsCircuitMaintenanceCircuitimpactDeleteNoContent() *PluginsCircuitMaintenanceCircuitimpactDeleteNoContent {
	return &PluginsCircuitMaintenanceCircuitimpactDeleteNoContent{}
}

/* PluginsCircuitMaintenanceCircuitimpactDeleteNoContent describes a response with status code 204, with default header values.

PluginsCircuitMaintenanceCircuitimpactDeleteNoContent plugins circuit maintenance circuitimpact delete no content
*/
type PluginsCircuitMaintenanceCircuitimpactDeleteNoContent struct {
}

func (o *PluginsCircuitMaintenanceCircuitimpactDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/circuit-maintenance/circuitimpact/{id}/][%d] pluginsCircuitMaintenanceCircuitimpactDeleteNoContent ", 204)
}

func (o *PluginsCircuitMaintenanceCircuitimpactDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
