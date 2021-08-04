package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigGoldenConfigBulkPartialUpdateReader is a Reader for the PluginsGoldenConfigGoldenConfigBulkPartialUpdate structure.
type PluginsGoldenConfigGoldenConfigBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateOK creates a PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK with default headers values
func NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateOK() *PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK {
	return &PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK{}
}

/* PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK plugins golden config golden config bulk partial update o k
*/
type PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK struct {
	Payload *models.GoldenConfig
}

func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/golden-config/][%d] pluginsGoldenConfigGoldenConfigBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK) GetPayload() *models.GoldenConfig {
	return o.Payload
}

func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoldenConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
