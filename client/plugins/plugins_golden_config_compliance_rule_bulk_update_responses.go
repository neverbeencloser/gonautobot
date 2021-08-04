package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigComplianceRuleBulkUpdateReader is a Reader for the PluginsGoldenConfigComplianceRuleBulkUpdate structure.
type PluginsGoldenConfigComplianceRuleBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigComplianceRuleBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkUpdateOK creates a PluginsGoldenConfigComplianceRuleBulkUpdateOK with default headers values
func NewPluginsGoldenConfigComplianceRuleBulkUpdateOK() *PluginsGoldenConfigComplianceRuleBulkUpdateOK {
	return &PluginsGoldenConfigComplianceRuleBulkUpdateOK{}
}

/* PluginsGoldenConfigComplianceRuleBulkUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigComplianceRuleBulkUpdateOK plugins golden config compliance rule bulk update o k
*/
type PluginsGoldenConfigComplianceRuleBulkUpdateOK struct {
	Payload *models.ComplianceRule
}

func (o *PluginsGoldenConfigComplianceRuleBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/compliance-rule/][%d] pluginsGoldenConfigComplianceRuleBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateOK) GetPayload() *models.ComplianceRule {
	return o.Payload
}

func (o *PluginsGoldenConfigComplianceRuleBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComplianceRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
