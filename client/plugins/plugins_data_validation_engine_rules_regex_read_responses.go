package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDataValidationEngineRulesRegexReadReader is a Reader for the PluginsDataValidationEngineRulesRegexRead structure.
type PluginsDataValidationEngineRulesRegexReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesRegexReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsDataValidationEngineRulesRegexReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesRegexReadOK creates a PluginsDataValidationEngineRulesRegexReadOK with default headers values
func NewPluginsDataValidationEngineRulesRegexReadOK() *PluginsDataValidationEngineRulesRegexReadOK {
	return &PluginsDataValidationEngineRulesRegexReadOK{}
}

/* PluginsDataValidationEngineRulesRegexReadOK describes a response with status code 200, with default header values.

PluginsDataValidationEngineRulesRegexReadOK plugins data validation engine rules regex read o k
*/
type PluginsDataValidationEngineRulesRegexReadOK struct {
	Payload *models.RegularExpressionValidationRule
}

func (o *PluginsDataValidationEngineRulesRegexReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/data-validation-engine/rules/regex/{id}/][%d] pluginsDataValidationEngineRulesRegexReadOK  %+v", 200, o.Payload)
}
func (o *PluginsDataValidationEngineRulesRegexReadOK) GetPayload() *models.RegularExpressionValidationRule {
	return o.Payload
}

func (o *PluginsDataValidationEngineRulesRegexReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularExpressionValidationRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
