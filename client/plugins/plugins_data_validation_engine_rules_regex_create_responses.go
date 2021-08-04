package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDataValidationEngineRulesRegexCreateReader is a Reader for the PluginsDataValidationEngineRulesRegexCreate structure.
type PluginsDataValidationEngineRulesRegexCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesRegexCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsDataValidationEngineRulesRegexCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesRegexCreateCreated creates a PluginsDataValidationEngineRulesRegexCreateCreated with default headers values
func NewPluginsDataValidationEngineRulesRegexCreateCreated() *PluginsDataValidationEngineRulesRegexCreateCreated {
	return &PluginsDataValidationEngineRulesRegexCreateCreated{}
}

/* PluginsDataValidationEngineRulesRegexCreateCreated describes a response with status code 201, with default header values.

PluginsDataValidationEngineRulesRegexCreateCreated plugins data validation engine rules regex create created
*/
type PluginsDataValidationEngineRulesRegexCreateCreated struct {
	Payload *models.RegularExpressionValidationRule
}

func (o *PluginsDataValidationEngineRulesRegexCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/data-validation-engine/rules/regex/][%d] pluginsDataValidationEngineRulesRegexCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsDataValidationEngineRulesRegexCreateCreated) GetPayload() *models.RegularExpressionValidationRule {
	return o.Payload
}

func (o *PluginsDataValidationEngineRulesRegexCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularExpressionValidationRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
