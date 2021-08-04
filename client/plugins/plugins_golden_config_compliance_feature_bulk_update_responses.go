package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigComplianceFeatureBulkUpdateReader is a Reader for the PluginsGoldenConfigComplianceFeatureBulkUpdate structure.
type PluginsGoldenConfigComplianceFeatureBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigComplianceFeatureBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkUpdateOK creates a PluginsGoldenConfigComplianceFeatureBulkUpdateOK with default headers values
func NewPluginsGoldenConfigComplianceFeatureBulkUpdateOK() *PluginsGoldenConfigComplianceFeatureBulkUpdateOK {
	return &PluginsGoldenConfigComplianceFeatureBulkUpdateOK{}
}

/* PluginsGoldenConfigComplianceFeatureBulkUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigComplianceFeatureBulkUpdateOK plugins golden config compliance feature bulk update o k
*/
type PluginsGoldenConfigComplianceFeatureBulkUpdateOK struct {
	Payload *models.ComplianceFeature
}

func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/compliance-feature/][%d] pluginsGoldenConfigComplianceFeatureBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateOK) GetPayload() *models.ComplianceFeature {
	return o.Payload
}

func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComplianceFeature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
