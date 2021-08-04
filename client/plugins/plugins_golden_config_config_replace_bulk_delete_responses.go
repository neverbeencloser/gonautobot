package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigConfigReplaceBulkDeleteReader is a Reader for the PluginsGoldenConfigConfigReplaceBulkDelete structure.
type PluginsGoldenConfigConfigReplaceBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigReplaceBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigConfigReplaceBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkDeleteNoContent creates a PluginsGoldenConfigConfigReplaceBulkDeleteNoContent with default headers values
func NewPluginsGoldenConfigConfigReplaceBulkDeleteNoContent() *PluginsGoldenConfigConfigReplaceBulkDeleteNoContent {
	return &PluginsGoldenConfigConfigReplaceBulkDeleteNoContent{}
}

/* PluginsGoldenConfigConfigReplaceBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigConfigReplaceBulkDeleteNoContent plugins golden config config replace bulk delete no content
*/
type PluginsGoldenConfigConfigReplaceBulkDeleteNoContent struct {
}

func (o *PluginsGoldenConfigConfigReplaceBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/config-replace/][%d] pluginsGoldenConfigConfigReplaceBulkDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigConfigReplaceBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
