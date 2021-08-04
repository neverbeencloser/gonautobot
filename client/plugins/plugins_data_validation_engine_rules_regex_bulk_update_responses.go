package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDataValidationEngineRulesRegexBulkUpdateReader is a Reader for the PluginsDataValidationEngineRulesRegexBulkUpdate structure.
type PluginsDataValidationEngineRulesRegexBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsDataValidationEngineRulesRegexBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkUpdateOK creates a PluginsDataValidationEngineRulesRegexBulkUpdateOK with default headers values
func NewPluginsDataValidationEngineRulesRegexBulkUpdateOK() *PluginsDataValidationEngineRulesRegexBulkUpdateOK {
	return &PluginsDataValidationEngineRulesRegexBulkUpdateOK{}
}

/* PluginsDataValidationEngineRulesRegexBulkUpdateOK describes a response with status code 200, with default header values.

PluginsDataValidationEngineRulesRegexBulkUpdateOK plugins data validation engine rules regex bulk update o k
*/
type PluginsDataValidationEngineRulesRegexBulkUpdateOK struct {
	Payload *models.RegularExpressionValidationRule
}

func (o *PluginsDataValidationEngineRulesRegexBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/data-validation-engine/rules/regex/][%d] pluginsDataValidationEngineRulesRegexBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateOK) GetPayload() *models.RegularExpressionValidationRule {
	return o.Payload
}

func (o *PluginsDataValidationEngineRulesRegexBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularExpressionValidationRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
