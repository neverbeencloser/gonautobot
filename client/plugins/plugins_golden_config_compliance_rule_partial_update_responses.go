package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigComplianceRulePartialUpdateReader is a Reader for the PluginsGoldenConfigComplianceRulePartialUpdate structure.
type PluginsGoldenConfigComplianceRulePartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceRulePartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigComplianceRulePartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceRulePartialUpdateOK creates a PluginsGoldenConfigComplianceRulePartialUpdateOK with default headers values
func NewPluginsGoldenConfigComplianceRulePartialUpdateOK() *PluginsGoldenConfigComplianceRulePartialUpdateOK {
	return &PluginsGoldenConfigComplianceRulePartialUpdateOK{}
}

/* PluginsGoldenConfigComplianceRulePartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigComplianceRulePartialUpdateOK plugins golden config compliance rule partial update o k
*/
type PluginsGoldenConfigComplianceRulePartialUpdateOK struct {
	Payload *models.ComplianceRule
}

func (o *PluginsGoldenConfigComplianceRulePartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/compliance-rule/{id}/][%d] pluginsGoldenConfigComplianceRulePartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigComplianceRulePartialUpdateOK) GetPayload() *models.ComplianceRule {
	return o.Payload
}

func (o *PluginsGoldenConfigComplianceRulePartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComplianceRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
