package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceNoteBulkPartialUpdateReader is a Reader for the PluginsCircuitMaintenanceNoteBulkPartialUpdate structure.
type PluginsCircuitMaintenanceNoteBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceNoteBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceNoteBulkPartialUpdateOK creates a PluginsCircuitMaintenanceNoteBulkPartialUpdateOK with default headers values
func NewPluginsCircuitMaintenanceNoteBulkPartialUpdateOK() *PluginsCircuitMaintenanceNoteBulkPartialUpdateOK {
	return &PluginsCircuitMaintenanceNoteBulkPartialUpdateOK{}
}

/* PluginsCircuitMaintenanceNoteBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceNoteBulkPartialUpdateOK plugins circuit maintenance note bulk partial update o k
*/
type PluginsCircuitMaintenanceNoteBulkPartialUpdateOK struct {
	Payload *models.Note
}

func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/circuit-maintenance/note/][%d] pluginsCircuitMaintenanceNoteBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateOK) GetPayload() *models.Note {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Note)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
