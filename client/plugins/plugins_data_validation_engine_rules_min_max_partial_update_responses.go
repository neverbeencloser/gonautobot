package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDataValidationEngineRulesMinMaxPartialUpdateReader is a Reader for the PluginsDataValidationEngineRulesMinMaxPartialUpdate structure.
type PluginsDataValidationEngineRulesMinMaxPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsDataValidationEngineRulesMinMaxPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesMinMaxPartialUpdateOK creates a PluginsDataValidationEngineRulesMinMaxPartialUpdateOK with default headers values
func NewPluginsDataValidationEngineRulesMinMaxPartialUpdateOK() *PluginsDataValidationEngineRulesMinMaxPartialUpdateOK {
	return &PluginsDataValidationEngineRulesMinMaxPartialUpdateOK{}
}

/* PluginsDataValidationEngineRulesMinMaxPartialUpdateOK describes a response with status code 200, with default header values.

PluginsDataValidationEngineRulesMinMaxPartialUpdateOK plugins data validation engine rules min max partial update o k
*/
type PluginsDataValidationEngineRulesMinMaxPartialUpdateOK struct {
	Payload *models.MinMaxValidationRule
}

func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/data-validation-engine/rules/min-max/{id}/][%d] pluginsDataValidationEngineRulesMinMaxPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateOK) GetPayload() *models.MinMaxValidationRule {
	return o.Payload
}

func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MinMaxValidationRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
