package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDataValidationEngineRulesRegexPartialUpdateReader is a Reader for the PluginsDataValidationEngineRulesRegexPartialUpdate structure.
type PluginsDataValidationEngineRulesRegexPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsDataValidationEngineRulesRegexPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesRegexPartialUpdateOK creates a PluginsDataValidationEngineRulesRegexPartialUpdateOK with default headers values
func NewPluginsDataValidationEngineRulesRegexPartialUpdateOK() *PluginsDataValidationEngineRulesRegexPartialUpdateOK {
	return &PluginsDataValidationEngineRulesRegexPartialUpdateOK{}
}

/* PluginsDataValidationEngineRulesRegexPartialUpdateOK describes a response with status code 200, with default header values.

PluginsDataValidationEngineRulesRegexPartialUpdateOK plugins data validation engine rules regex partial update o k
*/
type PluginsDataValidationEngineRulesRegexPartialUpdateOK struct {
	Payload *models.RegularExpressionValidationRule
}

func (o *PluginsDataValidationEngineRulesRegexPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/data-validation-engine/rules/regex/{id}/][%d] pluginsDataValidationEngineRulesRegexPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateOK) GetPayload() *models.RegularExpressionValidationRule {
	return o.Payload
}

func (o *PluginsDataValidationEngineRulesRegexPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularExpressionValidationRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
