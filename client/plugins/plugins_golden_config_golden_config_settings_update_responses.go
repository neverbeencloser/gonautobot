package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigGoldenConfigSettingsUpdateReader is a Reader for the PluginsGoldenConfigGoldenConfigSettingsUpdate structure.
type PluginsGoldenConfigGoldenConfigSettingsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigGoldenConfigSettingsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsUpdateOK creates a PluginsGoldenConfigGoldenConfigSettingsUpdateOK with default headers values
func NewPluginsGoldenConfigGoldenConfigSettingsUpdateOK() *PluginsGoldenConfigGoldenConfigSettingsUpdateOK {
	return &PluginsGoldenConfigGoldenConfigSettingsUpdateOK{}
}

/* PluginsGoldenConfigGoldenConfigSettingsUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigGoldenConfigSettingsUpdateOK plugins golden config golden config settings update o k
*/
type PluginsGoldenConfigGoldenConfigSettingsUpdateOK struct {
	Payload *models.GoldenConfigSetting
}

func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/golden-config-settings/{id}/][%d] pluginsGoldenConfigGoldenConfigSettingsUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateOK) GetPayload() *models.GoldenConfigSetting {
	return o.Payload
}

func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoldenConfigSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
