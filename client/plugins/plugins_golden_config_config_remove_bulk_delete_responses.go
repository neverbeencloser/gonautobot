package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigConfigRemoveBulkDeleteReader is a Reader for the PluginsGoldenConfigConfigRemoveBulkDelete structure.
type PluginsGoldenConfigConfigRemoveBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigRemoveBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigConfigRemoveBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigRemoveBulkDeleteNoContent creates a PluginsGoldenConfigConfigRemoveBulkDeleteNoContent with default headers values
func NewPluginsGoldenConfigConfigRemoveBulkDeleteNoContent() *PluginsGoldenConfigConfigRemoveBulkDeleteNoContent {
	return &PluginsGoldenConfigConfigRemoveBulkDeleteNoContent{}
}

/* PluginsGoldenConfigConfigRemoveBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigConfigRemoveBulkDeleteNoContent plugins golden config config remove bulk delete no content
*/
type PluginsGoldenConfigConfigRemoveBulkDeleteNoContent struct {
}

func (o *PluginsGoldenConfigConfigRemoveBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/config-remove/][%d] pluginsGoldenConfigConfigRemoveBulkDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigConfigRemoveBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
