package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigComplianceFeatureBulkDeleteReader is a Reader for the PluginsGoldenConfigComplianceFeatureBulkDelete structure.
type PluginsGoldenConfigComplianceFeatureBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigComplianceFeatureBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkDeleteNoContent creates a PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent with default headers values
func NewPluginsGoldenConfigComplianceFeatureBulkDeleteNoContent() *PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent {
	return &PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent{}
}

/* PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent plugins golden config compliance feature bulk delete no content
*/
type PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent struct {
}

func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/compliance-feature/][%d] pluginsGoldenConfigComplianceFeatureBulkDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
