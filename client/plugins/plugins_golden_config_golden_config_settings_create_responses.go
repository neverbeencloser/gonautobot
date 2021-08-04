package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigGoldenConfigSettingsCreateReader is a Reader for the PluginsGoldenConfigGoldenConfigSettingsCreate structure.
type PluginsGoldenConfigGoldenConfigSettingsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsGoldenConfigGoldenConfigSettingsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsCreateCreated creates a PluginsGoldenConfigGoldenConfigSettingsCreateCreated with default headers values
func NewPluginsGoldenConfigGoldenConfigSettingsCreateCreated() *PluginsGoldenConfigGoldenConfigSettingsCreateCreated {
	return &PluginsGoldenConfigGoldenConfigSettingsCreateCreated{}
}

/* PluginsGoldenConfigGoldenConfigSettingsCreateCreated describes a response with status code 201, with default header values.

PluginsGoldenConfigGoldenConfigSettingsCreateCreated plugins golden config golden config settings create created
*/
type PluginsGoldenConfigGoldenConfigSettingsCreateCreated struct {
	Payload *models.GoldenConfigSetting
}

func (o *PluginsGoldenConfigGoldenConfigSettingsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/golden-config/golden-config-settings/][%d] pluginsGoldenConfigGoldenConfigSettingsCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateCreated) GetPayload() *models.GoldenConfigSetting {
	return o.Payload
}

func (o *PluginsGoldenConfigGoldenConfigSettingsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoldenConfigSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
