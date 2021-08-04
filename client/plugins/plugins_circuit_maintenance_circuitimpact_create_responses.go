package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceCircuitimpactCreateReader is a Reader for the PluginsCircuitMaintenanceCircuitimpactCreate structure.
type PluginsCircuitMaintenanceCircuitimpactCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceCircuitimpactCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsCircuitMaintenanceCircuitimpactCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactCreateCreated creates a PluginsCircuitMaintenanceCircuitimpactCreateCreated with default headers values
func NewPluginsCircuitMaintenanceCircuitimpactCreateCreated() *PluginsCircuitMaintenanceCircuitimpactCreateCreated {
	return &PluginsCircuitMaintenanceCircuitimpactCreateCreated{}
}

/* PluginsCircuitMaintenanceCircuitimpactCreateCreated describes a response with status code 201, with default header values.

PluginsCircuitMaintenanceCircuitimpactCreateCreated plugins circuit maintenance circuitimpact create created
*/
type PluginsCircuitMaintenanceCircuitimpactCreateCreated struct {
	Payload *models.CircuitMaintenanceCircuitImpact
}

func (o *PluginsCircuitMaintenanceCircuitimpactCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/circuit-maintenance/circuitimpact/][%d] pluginsCircuitMaintenanceCircuitimpactCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsCircuitMaintenanceCircuitimpactCreateCreated) GetPayload() *models.CircuitMaintenanceCircuitImpact {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceCircuitimpactCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitMaintenanceCircuitImpact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
