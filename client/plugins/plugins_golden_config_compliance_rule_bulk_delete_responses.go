package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigComplianceRuleBulkDeleteReader is a Reader for the PluginsGoldenConfigComplianceRuleBulkDelete structure.
type PluginsGoldenConfigComplianceRuleBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceRuleBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigComplianceRuleBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkDeleteNoContent creates a PluginsGoldenConfigComplianceRuleBulkDeleteNoContent with default headers values
func NewPluginsGoldenConfigComplianceRuleBulkDeleteNoContent() *PluginsGoldenConfigComplianceRuleBulkDeleteNoContent {
	return &PluginsGoldenConfigComplianceRuleBulkDeleteNoContent{}
}

/* PluginsGoldenConfigComplianceRuleBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigComplianceRuleBulkDeleteNoContent plugins golden config compliance rule bulk delete no content
*/
type PluginsGoldenConfigComplianceRuleBulkDeleteNoContent struct {
}

func (o *PluginsGoldenConfigComplianceRuleBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/compliance-rule/][%d] pluginsGoldenConfigComplianceRuleBulkDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigComplianceRuleBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
