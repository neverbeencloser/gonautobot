package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigComplianceRuleUpdateReader is a Reader for the PluginsGoldenConfigComplianceRuleUpdate structure.
type PluginsGoldenConfigComplianceRuleUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceRuleUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigComplianceRuleUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceRuleUpdateOK creates a PluginsGoldenConfigComplianceRuleUpdateOK with default headers values
func NewPluginsGoldenConfigComplianceRuleUpdateOK() *PluginsGoldenConfigComplianceRuleUpdateOK {
	return &PluginsGoldenConfigComplianceRuleUpdateOK{}
}

/* PluginsGoldenConfigComplianceRuleUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigComplianceRuleUpdateOK plugins golden config compliance rule update o k
*/
type PluginsGoldenConfigComplianceRuleUpdateOK struct {
	Payload *models.ComplianceRule
}

func (o *PluginsGoldenConfigComplianceRuleUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/compliance-rule/{id}/][%d] pluginsGoldenConfigComplianceRuleUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigComplianceRuleUpdateOK) GetPayload() *models.ComplianceRule {
	return o.Payload
}

func (o *PluginsGoldenConfigComplianceRuleUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComplianceRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
