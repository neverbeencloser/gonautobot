package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsDataValidationEngineRulesMinMaxDeleteReader is a Reader for the PluginsDataValidationEngineRulesMinMaxDelete structure.
type PluginsDataValidationEngineRulesMinMaxDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesMinMaxDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsDataValidationEngineRulesMinMaxDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesMinMaxDeleteNoContent creates a PluginsDataValidationEngineRulesMinMaxDeleteNoContent with default headers values
func NewPluginsDataValidationEngineRulesMinMaxDeleteNoContent() *PluginsDataValidationEngineRulesMinMaxDeleteNoContent {
	return &PluginsDataValidationEngineRulesMinMaxDeleteNoContent{}
}

/* PluginsDataValidationEngineRulesMinMaxDeleteNoContent describes a response with status code 204, with default header values.

PluginsDataValidationEngineRulesMinMaxDeleteNoContent plugins data validation engine rules min max delete no content
*/
type PluginsDataValidationEngineRulesMinMaxDeleteNoContent struct {
}

func (o *PluginsDataValidationEngineRulesMinMaxDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/data-validation-engine/rules/min-max/{id}/][%d] pluginsDataValidationEngineRulesMinMaxDeleteNoContent ", 204)
}

func (o *PluginsDataValidationEngineRulesMinMaxDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
