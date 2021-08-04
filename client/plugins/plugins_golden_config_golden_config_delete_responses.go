package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigGoldenConfigDeleteReader is a Reader for the PluginsGoldenConfigGoldenConfigDelete structure.
type PluginsGoldenConfigGoldenConfigDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigGoldenConfigDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigDeleteNoContent creates a PluginsGoldenConfigGoldenConfigDeleteNoContent with default headers values
func NewPluginsGoldenConfigGoldenConfigDeleteNoContent() *PluginsGoldenConfigGoldenConfigDeleteNoContent {
	return &PluginsGoldenConfigGoldenConfigDeleteNoContent{}
}

/* PluginsGoldenConfigGoldenConfigDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigGoldenConfigDeleteNoContent plugins golden config golden config delete no content
*/
type PluginsGoldenConfigGoldenConfigDeleteNoContent struct {
}

func (o *PluginsGoldenConfigGoldenConfigDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/golden-config/{id}/][%d] pluginsGoldenConfigGoldenConfigDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigGoldenConfigDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
