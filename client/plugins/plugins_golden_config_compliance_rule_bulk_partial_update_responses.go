package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigComplianceRuleBulkPartialUpdateReader is a Reader for the PluginsGoldenConfigComplianceRuleBulkPartialUpdate structure.
type PluginsGoldenConfigComplianceRuleBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateOK creates a PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK with default headers values
func NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateOK() *PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK {
	return &PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK{}
}

/* PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK plugins golden config compliance rule bulk partial update o k
*/
type PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK struct {
	Payload *models.ComplianceRule
}

func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/compliance-rule/][%d] pluginsGoldenConfigComplianceRuleBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK) GetPayload() *models.ComplianceRule {
	return o.Payload
}

func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComplianceRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
