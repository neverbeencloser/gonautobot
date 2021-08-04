package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceCircuitimpactReadReader is a Reader for the PluginsCircuitMaintenanceCircuitimpactRead structure.
type PluginsCircuitMaintenanceCircuitimpactReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceCircuitimpactReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceCircuitimpactReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactReadOK creates a PluginsCircuitMaintenanceCircuitimpactReadOK with default headers values
func NewPluginsCircuitMaintenanceCircuitimpactReadOK() *PluginsCircuitMaintenanceCircuitimpactReadOK {
	return &PluginsCircuitMaintenanceCircuitimpactReadOK{}
}

/* PluginsCircuitMaintenanceCircuitimpactReadOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceCircuitimpactReadOK plugins circuit maintenance circuitimpact read o k
*/
type PluginsCircuitMaintenanceCircuitimpactReadOK struct {
	Payload *models.CircuitMaintenanceCircuitImpact
}

func (o *PluginsCircuitMaintenanceCircuitimpactReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/circuit-maintenance/circuitimpact/{id}/][%d] pluginsCircuitMaintenanceCircuitimpactReadOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceCircuitimpactReadOK) GetPayload() *models.CircuitMaintenanceCircuitImpact {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceCircuitimpactReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitMaintenanceCircuitImpact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
