package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceMaintenanceReadReader is a Reader for the PluginsCircuitMaintenanceMaintenanceRead structure.
type PluginsCircuitMaintenanceMaintenanceReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceMaintenanceReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceMaintenanceReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceMaintenanceReadOK creates a PluginsCircuitMaintenanceMaintenanceReadOK with default headers values
func NewPluginsCircuitMaintenanceMaintenanceReadOK() *PluginsCircuitMaintenanceMaintenanceReadOK {
	return &PluginsCircuitMaintenanceMaintenanceReadOK{}
}

/* PluginsCircuitMaintenanceMaintenanceReadOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceMaintenanceReadOK plugins circuit maintenance maintenance read o k
*/
type PluginsCircuitMaintenanceMaintenanceReadOK struct {
	Payload *models.CircuitMaintenance
}

func (o *PluginsCircuitMaintenanceMaintenanceReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/circuit-maintenance/maintenance/{id}/][%d] pluginsCircuitMaintenanceMaintenanceReadOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceMaintenanceReadOK) GetPayload() *models.CircuitMaintenance {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceMaintenanceReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitMaintenance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
