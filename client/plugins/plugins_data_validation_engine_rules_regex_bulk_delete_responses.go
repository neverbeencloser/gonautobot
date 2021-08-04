package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsDataValidationEngineRulesRegexBulkDeleteReader is a Reader for the PluginsDataValidationEngineRulesRegexBulkDelete structure.
type PluginsDataValidationEngineRulesRegexBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesRegexBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsDataValidationEngineRulesRegexBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkDeleteNoContent creates a PluginsDataValidationEngineRulesRegexBulkDeleteNoContent with default headers values
func NewPluginsDataValidationEngineRulesRegexBulkDeleteNoContent() *PluginsDataValidationEngineRulesRegexBulkDeleteNoContent {
	return &PluginsDataValidationEngineRulesRegexBulkDeleteNoContent{}
}

/* PluginsDataValidationEngineRulesRegexBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsDataValidationEngineRulesRegexBulkDeleteNoContent plugins data validation engine rules regex bulk delete no content
*/
type PluginsDataValidationEngineRulesRegexBulkDeleteNoContent struct {
}

func (o *PluginsDataValidationEngineRulesRegexBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/data-validation-engine/rules/regex/][%d] pluginsDataValidationEngineRulesRegexBulkDeleteNoContent ", 204)
}

func (o *PluginsDataValidationEngineRulesRegexBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
