package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigComplianceFeatureDeleteReader is a Reader for the PluginsGoldenConfigComplianceFeatureDelete structure.
type PluginsGoldenConfigComplianceFeatureDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceFeatureDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigComplianceFeatureDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceFeatureDeleteNoContent creates a PluginsGoldenConfigComplianceFeatureDeleteNoContent with default headers values
func NewPluginsGoldenConfigComplianceFeatureDeleteNoContent() *PluginsGoldenConfigComplianceFeatureDeleteNoContent {
	return &PluginsGoldenConfigComplianceFeatureDeleteNoContent{}
}

/* PluginsGoldenConfigComplianceFeatureDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigComplianceFeatureDeleteNoContent plugins golden config compliance feature delete no content
*/
type PluginsGoldenConfigComplianceFeatureDeleteNoContent struct {
}

func (o *PluginsGoldenConfigComplianceFeatureDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/compliance-feature/{id}/][%d] pluginsGoldenConfigComplianceFeatureDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigComplianceFeatureDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
