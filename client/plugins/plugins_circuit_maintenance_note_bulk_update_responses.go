package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceNoteBulkUpdateReader is a Reader for the PluginsCircuitMaintenanceNoteBulkUpdate structure.
type PluginsCircuitMaintenanceNoteBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceNoteBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceNoteBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceNoteBulkUpdateOK creates a PluginsCircuitMaintenanceNoteBulkUpdateOK with default headers values
func NewPluginsCircuitMaintenanceNoteBulkUpdateOK() *PluginsCircuitMaintenanceNoteBulkUpdateOK {
	return &PluginsCircuitMaintenanceNoteBulkUpdateOK{}
}

/* PluginsCircuitMaintenanceNoteBulkUpdateOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceNoteBulkUpdateOK plugins circuit maintenance note bulk update o k
*/
type PluginsCircuitMaintenanceNoteBulkUpdateOK struct {
	Payload *models.Note
}

func (o *PluginsCircuitMaintenanceNoteBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/circuit-maintenance/note/][%d] pluginsCircuitMaintenanceNoteBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceNoteBulkUpdateOK) GetPayload() *models.Note {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceNoteBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Note)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
