package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceCircuitimpactPartialUpdateReader is a Reader for the PluginsCircuitMaintenanceCircuitimpactPartialUpdate structure.
type PluginsCircuitMaintenanceCircuitimpactPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateOK creates a PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK with default headers values
func NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateOK() *PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK {
	return &PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK{}
}

/* PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK plugins circuit maintenance circuitimpact partial update o k
*/
type PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK struct {
	Payload *models.CircuitMaintenanceCircuitImpact
}

func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/circuit-maintenance/circuitimpact/{id}/][%d] pluginsCircuitMaintenanceCircuitimpactPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK) GetPayload() *models.CircuitMaintenanceCircuitImpact {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitMaintenanceCircuitImpact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
