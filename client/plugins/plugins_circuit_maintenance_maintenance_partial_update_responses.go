package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceMaintenancePartialUpdateReader is a Reader for the PluginsCircuitMaintenanceMaintenancePartialUpdate structure.
type PluginsCircuitMaintenanceMaintenancePartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceMaintenancePartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceMaintenancePartialUpdateOK creates a PluginsCircuitMaintenanceMaintenancePartialUpdateOK with default headers values
func NewPluginsCircuitMaintenanceMaintenancePartialUpdateOK() *PluginsCircuitMaintenanceMaintenancePartialUpdateOK {
	return &PluginsCircuitMaintenanceMaintenancePartialUpdateOK{}
}

/* PluginsCircuitMaintenanceMaintenancePartialUpdateOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceMaintenancePartialUpdateOK plugins circuit maintenance maintenance partial update o k
*/
type PluginsCircuitMaintenanceMaintenancePartialUpdateOK struct {
	Payload *models.CircuitMaintenance
}

func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/circuit-maintenance/maintenance/{id}/][%d] pluginsCircuitMaintenanceMaintenancePartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateOK) GetPayload() *models.CircuitMaintenance {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitMaintenance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
