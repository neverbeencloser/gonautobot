package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateReader is a Reader for the PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdate structure.
type PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK creates a PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK with default headers values
func NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK() *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK {
	return &PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK{}
}

/* PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK plugins circuit maintenance circuitimpact bulk partial update o k
*/
type PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK struct {
	Payload *models.CircuitMaintenanceCircuitImpact
}

func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/circuit-maintenance/circuitimpact/][%d] pluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK) GetPayload() *models.CircuitMaintenanceCircuitImpact {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitMaintenanceCircuitImpact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
