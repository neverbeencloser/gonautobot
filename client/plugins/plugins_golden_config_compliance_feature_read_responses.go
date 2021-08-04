package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigComplianceFeatureReadReader is a Reader for the PluginsGoldenConfigComplianceFeatureRead structure.
type PluginsGoldenConfigComplianceFeatureReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceFeatureReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigComplianceFeatureReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceFeatureReadOK creates a PluginsGoldenConfigComplianceFeatureReadOK with default headers values
func NewPluginsGoldenConfigComplianceFeatureReadOK() *PluginsGoldenConfigComplianceFeatureReadOK {
	return &PluginsGoldenConfigComplianceFeatureReadOK{}
}

/* PluginsGoldenConfigComplianceFeatureReadOK describes a response with status code 200, with default header values.

PluginsGoldenConfigComplianceFeatureReadOK plugins golden config compliance feature read o k
*/
type PluginsGoldenConfigComplianceFeatureReadOK struct {
	Payload *models.ComplianceFeature
}

func (o *PluginsGoldenConfigComplianceFeatureReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/golden-config/compliance-feature/{id}/][%d] pluginsGoldenConfigComplianceFeatureReadOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigComplianceFeatureReadOK) GetPayload() *models.ComplianceFeature {
	return o.Payload
}

func (o *PluginsGoldenConfigComplianceFeatureReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComplianceFeature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
