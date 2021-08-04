package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceMaintenanceBulkUpdateReader is a Reader for the PluginsCircuitMaintenanceMaintenanceBulkUpdate structure.
type PluginsCircuitMaintenanceMaintenanceBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceMaintenanceBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkUpdateOK creates a PluginsCircuitMaintenanceMaintenanceBulkUpdateOK with default headers values
func NewPluginsCircuitMaintenanceMaintenanceBulkUpdateOK() *PluginsCircuitMaintenanceMaintenanceBulkUpdateOK {
	return &PluginsCircuitMaintenanceMaintenanceBulkUpdateOK{}
}

/* PluginsCircuitMaintenanceMaintenanceBulkUpdateOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceMaintenanceBulkUpdateOK plugins circuit maintenance maintenance bulk update o k
*/
type PluginsCircuitMaintenanceMaintenanceBulkUpdateOK struct {
	Payload *models.CircuitMaintenance
}

func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/circuit-maintenance/maintenance/][%d] pluginsCircuitMaintenanceMaintenanceBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateOK) GetPayload() *models.CircuitMaintenance {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitMaintenance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
