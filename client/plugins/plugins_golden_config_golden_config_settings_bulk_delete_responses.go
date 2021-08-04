package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigGoldenConfigSettingsBulkDeleteReader is a Reader for the PluginsGoldenConfigGoldenConfigSettingsBulkDelete structure.
type PluginsGoldenConfigGoldenConfigSettingsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent creates a PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent with default headers values
func NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent() *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent{}
}

/* PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent plugins golden config golden config settings bulk delete no content
*/
type PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent struct {
}

func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/golden-config-settings/][%d] pluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
