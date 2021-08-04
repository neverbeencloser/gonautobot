package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDataValidationEngineRulesMinMaxUpdateReader is a Reader for the PluginsDataValidationEngineRulesMinMaxUpdate structure.
type PluginsDataValidationEngineRulesMinMaxUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesMinMaxUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsDataValidationEngineRulesMinMaxUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesMinMaxUpdateOK creates a PluginsDataValidationEngineRulesMinMaxUpdateOK with default headers values
func NewPluginsDataValidationEngineRulesMinMaxUpdateOK() *PluginsDataValidationEngineRulesMinMaxUpdateOK {
	return &PluginsDataValidationEngineRulesMinMaxUpdateOK{}
}

/* PluginsDataValidationEngineRulesMinMaxUpdateOK describes a response with status code 200, with default header values.

PluginsDataValidationEngineRulesMinMaxUpdateOK plugins data validation engine rules min max update o k
*/
type PluginsDataValidationEngineRulesMinMaxUpdateOK struct {
	Payload *models.MinMaxValidationRule
}

func (o *PluginsDataValidationEngineRulesMinMaxUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/data-validation-engine/rules/min-max/{id}/][%d] pluginsDataValidationEngineRulesMinMaxUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsDataValidationEngineRulesMinMaxUpdateOK) GetPayload() *models.MinMaxValidationRule {
	return o.Payload
}

func (o *PluginsDataValidationEngineRulesMinMaxUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MinMaxValidationRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
