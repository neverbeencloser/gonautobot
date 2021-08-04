package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigGoldenConfigUpdateReader is a Reader for the PluginsGoldenConfigGoldenConfigUpdate structure.
type PluginsGoldenConfigGoldenConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigGoldenConfigUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigUpdateOK creates a PluginsGoldenConfigGoldenConfigUpdateOK with default headers values
func NewPluginsGoldenConfigGoldenConfigUpdateOK() *PluginsGoldenConfigGoldenConfigUpdateOK {
	return &PluginsGoldenConfigGoldenConfigUpdateOK{}
}

/* PluginsGoldenConfigGoldenConfigUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigGoldenConfigUpdateOK plugins golden config golden config update o k
*/
type PluginsGoldenConfigGoldenConfigUpdateOK struct {
	Payload *models.GoldenConfig
}

func (o *PluginsGoldenConfigGoldenConfigUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/golden-config/{id}/][%d] pluginsGoldenConfigGoldenConfigUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigGoldenConfigUpdateOK) GetPayload() *models.GoldenConfig {
	return o.Payload
}

func (o *PluginsGoldenConfigGoldenConfigUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoldenConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
