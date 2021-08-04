package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceMaintenanceCreateReader is a Reader for the PluginsCircuitMaintenanceMaintenanceCreate structure.
type PluginsCircuitMaintenanceMaintenanceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceMaintenanceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsCircuitMaintenanceMaintenanceCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceMaintenanceCreateCreated creates a PluginsCircuitMaintenanceMaintenanceCreateCreated with default headers values
func NewPluginsCircuitMaintenanceMaintenanceCreateCreated() *PluginsCircuitMaintenanceMaintenanceCreateCreated {
	return &PluginsCircuitMaintenanceMaintenanceCreateCreated{}
}

/* PluginsCircuitMaintenanceMaintenanceCreateCreated describes a response with status code 201, with default header values.

PluginsCircuitMaintenanceMaintenanceCreateCreated plugins circuit maintenance maintenance create created
*/
type PluginsCircuitMaintenanceMaintenanceCreateCreated struct {
	Payload *models.CircuitMaintenance
}

func (o *PluginsCircuitMaintenanceMaintenanceCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/circuit-maintenance/maintenance/][%d] pluginsCircuitMaintenanceMaintenanceCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsCircuitMaintenanceMaintenanceCreateCreated) GetPayload() *models.CircuitMaintenance {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceMaintenanceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitMaintenance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
