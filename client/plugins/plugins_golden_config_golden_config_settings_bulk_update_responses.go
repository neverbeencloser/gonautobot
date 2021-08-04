package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigGoldenConfigSettingsBulkUpdateReader is a Reader for the PluginsGoldenConfigGoldenConfigSettingsBulkUpdate structure.
type PluginsGoldenConfigGoldenConfigSettingsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK creates a PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK with default headers values
func NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK() *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK{}
}

/* PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK plugins golden config golden config settings bulk update o k
*/
type PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK struct {
	Payload *models.GoldenConfigSetting
}

func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/golden-config-settings/][%d] pluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK) GetPayload() *models.GoldenConfigSetting {
	return o.Payload
}

func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoldenConfigSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
