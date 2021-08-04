package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigConfigReplaceDeleteReader is a Reader for the PluginsGoldenConfigConfigReplaceDelete structure.
type PluginsGoldenConfigConfigReplaceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigReplaceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigConfigReplaceDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigReplaceDeleteNoContent creates a PluginsGoldenConfigConfigReplaceDeleteNoContent with default headers values
func NewPluginsGoldenConfigConfigReplaceDeleteNoContent() *PluginsGoldenConfigConfigReplaceDeleteNoContent {
	return &PluginsGoldenConfigConfigReplaceDeleteNoContent{}
}

/* PluginsGoldenConfigConfigReplaceDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigConfigReplaceDeleteNoContent plugins golden config config replace delete no content
*/
type PluginsGoldenConfigConfigReplaceDeleteNoContent struct {
}

func (o *PluginsGoldenConfigConfigReplaceDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/config-replace/{id}/][%d] pluginsGoldenConfigConfigReplaceDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigConfigReplaceDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
