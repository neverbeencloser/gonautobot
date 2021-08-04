package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateReader is a Reader for the PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdate structure.
type PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK creates a PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK with default headers values
func NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK() *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK{}
}

/* PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK plugins golden config golden config settings bulk partial update o k
*/
type PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK struct {
	Payload *models.GoldenConfigSetting
}

func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/golden-config-settings/][%d] pluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK) GetPayload() *models.GoldenConfigSetting {
	return o.Payload
}

func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoldenConfigSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
