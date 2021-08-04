package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsDataValidationEngineRulesRegexDeleteReader is a Reader for the PluginsDataValidationEngineRulesRegexDelete structure.
type PluginsDataValidationEngineRulesRegexDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesRegexDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsDataValidationEngineRulesRegexDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesRegexDeleteNoContent creates a PluginsDataValidationEngineRulesRegexDeleteNoContent with default headers values
func NewPluginsDataValidationEngineRulesRegexDeleteNoContent() *PluginsDataValidationEngineRulesRegexDeleteNoContent {
	return &PluginsDataValidationEngineRulesRegexDeleteNoContent{}
}

/* PluginsDataValidationEngineRulesRegexDeleteNoContent describes a response with status code 204, with default header values.

PluginsDataValidationEngineRulesRegexDeleteNoContent plugins data validation engine rules regex delete no content
*/
type PluginsDataValidationEngineRulesRegexDeleteNoContent struct {
}

func (o *PluginsDataValidationEngineRulesRegexDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/data-validation-engine/rules/regex/{id}/][%d] pluginsDataValidationEngineRulesRegexDeleteNoContent ", 204)
}

func (o *PluginsDataValidationEngineRulesRegexDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
