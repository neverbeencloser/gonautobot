package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigGoldenConfigBulkDeleteReader is a Reader for the PluginsGoldenConfigGoldenConfigBulkDelete structure.
type PluginsGoldenConfigGoldenConfigBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigGoldenConfigBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkDeleteNoContent creates a PluginsGoldenConfigGoldenConfigBulkDeleteNoContent with default headers values
func NewPluginsGoldenConfigGoldenConfigBulkDeleteNoContent() *PluginsGoldenConfigGoldenConfigBulkDeleteNoContent {
	return &PluginsGoldenConfigGoldenConfigBulkDeleteNoContent{}
}

/* PluginsGoldenConfigGoldenConfigBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigGoldenConfigBulkDeleteNoContent plugins golden config golden config bulk delete no content
*/
type PluginsGoldenConfigGoldenConfigBulkDeleteNoContent struct {
}

func (o *PluginsGoldenConfigGoldenConfigBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/golden-config/][%d] pluginsGoldenConfigGoldenConfigBulkDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigGoldenConfigBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
