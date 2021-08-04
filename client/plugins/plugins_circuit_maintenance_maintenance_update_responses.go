package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceMaintenanceUpdateReader is a Reader for the PluginsCircuitMaintenanceMaintenanceUpdate structure.
type PluginsCircuitMaintenanceMaintenanceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceMaintenanceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceMaintenanceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceMaintenanceUpdateOK creates a PluginsCircuitMaintenanceMaintenanceUpdateOK with default headers values
func NewPluginsCircuitMaintenanceMaintenanceUpdateOK() *PluginsCircuitMaintenanceMaintenanceUpdateOK {
	return &PluginsCircuitMaintenanceMaintenanceUpdateOK{}
}

/* PluginsCircuitMaintenanceMaintenanceUpdateOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceMaintenanceUpdateOK plugins circuit maintenance maintenance update o k
*/
type PluginsCircuitMaintenanceMaintenanceUpdateOK struct {
	Payload *models.CircuitMaintenance
}

func (o *PluginsCircuitMaintenanceMaintenanceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/circuit-maintenance/maintenance/{id}/][%d] pluginsCircuitMaintenanceMaintenanceUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceMaintenanceUpdateOK) GetPayload() *models.CircuitMaintenance {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceMaintenanceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitMaintenance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
