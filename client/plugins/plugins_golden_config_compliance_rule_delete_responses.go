package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigComplianceRuleDeleteReader is a Reader for the PluginsGoldenConfigComplianceRuleDelete structure.
type PluginsGoldenConfigComplianceRuleDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceRuleDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigComplianceRuleDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceRuleDeleteNoContent creates a PluginsGoldenConfigComplianceRuleDeleteNoContent with default headers values
func NewPluginsGoldenConfigComplianceRuleDeleteNoContent() *PluginsGoldenConfigComplianceRuleDeleteNoContent {
	return &PluginsGoldenConfigComplianceRuleDeleteNoContent{}
}

/* PluginsGoldenConfigComplianceRuleDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigComplianceRuleDeleteNoContent plugins golden config compliance rule delete no content
*/
type PluginsGoldenConfigComplianceRuleDeleteNoContent struct {
}

func (o *PluginsGoldenConfigComplianceRuleDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/compliance-rule/{id}/][%d] pluginsGoldenConfigComplianceRuleDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigComplianceRuleDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
