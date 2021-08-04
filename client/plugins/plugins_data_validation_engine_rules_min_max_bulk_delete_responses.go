package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsDataValidationEngineRulesMinMaxBulkDeleteReader is a Reader for the PluginsDataValidationEngineRulesMinMaxBulkDelete structure.
type PluginsDataValidationEngineRulesMinMaxBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent creates a PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent with default headers values
func NewPluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent() *PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent {
	return &PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent{}
}

/* PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent plugins data validation engine rules min max bulk delete no content
*/
type PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent struct {
}

func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/data-validation-engine/rules/min-max/][%d] pluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent ", 204)
}

func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
