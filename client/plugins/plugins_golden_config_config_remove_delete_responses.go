package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigConfigRemoveDeleteReader is a Reader for the PluginsGoldenConfigConfigRemoveDelete structure.
type PluginsGoldenConfigConfigRemoveDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigRemoveDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigConfigRemoveDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigRemoveDeleteNoContent creates a PluginsGoldenConfigConfigRemoveDeleteNoContent with default headers values
func NewPluginsGoldenConfigConfigRemoveDeleteNoContent() *PluginsGoldenConfigConfigRemoveDeleteNoContent {
	return &PluginsGoldenConfigConfigRemoveDeleteNoContent{}
}

/* PluginsGoldenConfigConfigRemoveDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigConfigRemoveDeleteNoContent plugins golden config config remove delete no content
*/
type PluginsGoldenConfigConfigRemoveDeleteNoContent struct {
}

func (o *PluginsGoldenConfigConfigRemoveDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/config-remove/{id}/][%d] pluginsGoldenConfigConfigRemoveDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigConfigRemoveDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
