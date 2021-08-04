package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigComplianceFeatureBulkPartialUpdateReader is a Reader for the PluginsGoldenConfigComplianceFeatureBulkPartialUpdate structure.
type PluginsGoldenConfigComplianceFeatureBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK creates a PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK with default headers values
func NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK() *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK {
	return &PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK{}
}

/* PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK plugins golden config compliance feature bulk partial update o k
*/
type PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK struct {
	Payload *models.ComplianceFeature
}

func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/compliance-feature/][%d] pluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK) GetPayload() *models.ComplianceFeature {
	return o.Payload
}

func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComplianceFeature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
