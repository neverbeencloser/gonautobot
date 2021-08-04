package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDataValidationEngineRulesRegexBulkPartialUpdateReader is a Reader for the PluginsDataValidationEngineRulesRegexBulkPartialUpdate structure.
type PluginsDataValidationEngineRulesRegexBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateOK creates a PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK with default headers values
func NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateOK() *PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK {
	return &PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK{}
}

/* PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK plugins data validation engine rules regex bulk partial update o k
*/
type PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK struct {
	Payload *models.RegularExpressionValidationRule
}

func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/data-validation-engine/rules/regex/][%d] pluginsDataValidationEngineRulesRegexBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK) GetPayload() *models.RegularExpressionValidationRule {
	return o.Payload
}

func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularExpressionValidationRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
