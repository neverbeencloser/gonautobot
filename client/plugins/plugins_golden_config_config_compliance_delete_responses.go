package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigConfigComplianceDeleteReader is a Reader for the PluginsGoldenConfigConfigComplianceDelete structure.
type PluginsGoldenConfigConfigComplianceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigComplianceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigConfigComplianceDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigComplianceDeleteNoContent creates a PluginsGoldenConfigConfigComplianceDeleteNoContent with default headers values
func NewPluginsGoldenConfigConfigComplianceDeleteNoContent() *PluginsGoldenConfigConfigComplianceDeleteNoContent {
	return &PluginsGoldenConfigConfigComplianceDeleteNoContent{}
}

/* PluginsGoldenConfigConfigComplianceDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigConfigComplianceDeleteNoContent plugins golden config config compliance delete no content
*/
type PluginsGoldenConfigConfigComplianceDeleteNoContent struct {
}

func (o *PluginsGoldenConfigConfigComplianceDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/config-compliance/{id}/][%d] pluginsGoldenConfigConfigComplianceDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigConfigComplianceDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
