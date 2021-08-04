package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceCircuitimpactUpdateReader is a Reader for the PluginsCircuitMaintenanceCircuitimpactUpdate structure.
type PluginsCircuitMaintenanceCircuitimpactUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceCircuitimpactUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactUpdateOK creates a PluginsCircuitMaintenanceCircuitimpactUpdateOK with default headers values
func NewPluginsCircuitMaintenanceCircuitimpactUpdateOK() *PluginsCircuitMaintenanceCircuitimpactUpdateOK {
	return &PluginsCircuitMaintenanceCircuitimpactUpdateOK{}
}

/* PluginsCircuitMaintenanceCircuitimpactUpdateOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceCircuitimpactUpdateOK plugins circuit maintenance circuitimpact update o k
*/
type PluginsCircuitMaintenanceCircuitimpactUpdateOK struct {
	Payload *models.CircuitMaintenanceCircuitImpact
}

func (o *PluginsCircuitMaintenanceCircuitimpactUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/circuit-maintenance/circuitimpact/{id}/][%d] pluginsCircuitMaintenanceCircuitimpactUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateOK) GetPayload() *models.CircuitMaintenanceCircuitImpact {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceCircuitimpactUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitMaintenanceCircuitImpact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
