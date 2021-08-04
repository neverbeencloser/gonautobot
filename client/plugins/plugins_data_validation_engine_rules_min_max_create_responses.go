package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDataValidationEngineRulesMinMaxCreateReader is a Reader for the PluginsDataValidationEngineRulesMinMaxCreate structure.
type PluginsDataValidationEngineRulesMinMaxCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesMinMaxCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsDataValidationEngineRulesMinMaxCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesMinMaxCreateCreated creates a PluginsDataValidationEngineRulesMinMaxCreateCreated with default headers values
func NewPluginsDataValidationEngineRulesMinMaxCreateCreated() *PluginsDataValidationEngineRulesMinMaxCreateCreated {
	return &PluginsDataValidationEngineRulesMinMaxCreateCreated{}
}

/* PluginsDataValidationEngineRulesMinMaxCreateCreated describes a response with status code 201, with default header values.

PluginsDataValidationEngineRulesMinMaxCreateCreated plugins data validation engine rules min max create created
*/
type PluginsDataValidationEngineRulesMinMaxCreateCreated struct {
	Payload *models.MinMaxValidationRule
}

func (o *PluginsDataValidationEngineRulesMinMaxCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/data-validation-engine/rules/min-max/][%d] pluginsDataValidationEngineRulesMinMaxCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsDataValidationEngineRulesMinMaxCreateCreated) GetPayload() *models.MinMaxValidationRule {
	return o.Payload
}

func (o *PluginsDataValidationEngineRulesMinMaxCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MinMaxValidationRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
