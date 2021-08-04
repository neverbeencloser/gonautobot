package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigGoldenConfigBulkUpdateReader is a Reader for the PluginsGoldenConfigGoldenConfigBulkUpdate structure.
type PluginsGoldenConfigGoldenConfigBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigGoldenConfigBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkUpdateOK creates a PluginsGoldenConfigGoldenConfigBulkUpdateOK with default headers values
func NewPluginsGoldenConfigGoldenConfigBulkUpdateOK() *PluginsGoldenConfigGoldenConfigBulkUpdateOK {
	return &PluginsGoldenConfigGoldenConfigBulkUpdateOK{}
}

/* PluginsGoldenConfigGoldenConfigBulkUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigGoldenConfigBulkUpdateOK plugins golden config golden config bulk update o k
*/
type PluginsGoldenConfigGoldenConfigBulkUpdateOK struct {
	Payload *models.GoldenConfig
}

func (o *PluginsGoldenConfigGoldenConfigBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/golden-config/][%d] pluginsGoldenConfigGoldenConfigBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateOK) GetPayload() *models.GoldenConfig {
	return o.Payload
}

func (o *PluginsGoldenConfigGoldenConfigBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoldenConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
