package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateReader is a Reader for the PluginsDataValidationEngineRulesMinMaxBulkPartialUpdate structure.
type PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK creates a PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK with default headers values
func NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK() *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK {
	return &PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK{}
}

/* PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK plugins data validation engine rules min max bulk partial update o k
*/
type PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK struct {
	Payload *models.MinMaxValidationRule
}

func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/data-validation-engine/rules/min-max/][%d] pluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK) GetPayload() *models.MinMaxValidationRule {
	return o.Payload
}

func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MinMaxValidationRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
