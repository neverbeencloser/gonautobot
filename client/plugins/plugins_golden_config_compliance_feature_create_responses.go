package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigComplianceFeatureCreateReader is a Reader for the PluginsGoldenConfigComplianceFeatureCreate structure.
type PluginsGoldenConfigComplianceFeatureCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceFeatureCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsGoldenConfigComplianceFeatureCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceFeatureCreateCreated creates a PluginsGoldenConfigComplianceFeatureCreateCreated with default headers values
func NewPluginsGoldenConfigComplianceFeatureCreateCreated() *PluginsGoldenConfigComplianceFeatureCreateCreated {
	return &PluginsGoldenConfigComplianceFeatureCreateCreated{}
}

/* PluginsGoldenConfigComplianceFeatureCreateCreated describes a response with status code 201, with default header values.

PluginsGoldenConfigComplianceFeatureCreateCreated plugins golden config compliance feature create created
*/
type PluginsGoldenConfigComplianceFeatureCreateCreated struct {
	Payload *models.ComplianceFeature
}

func (o *PluginsGoldenConfigComplianceFeatureCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/golden-config/compliance-feature/][%d] pluginsGoldenConfigComplianceFeatureCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsGoldenConfigComplianceFeatureCreateCreated) GetPayload() *models.ComplianceFeature {
	return o.Payload
}

func (o *PluginsGoldenConfigComplianceFeatureCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComplianceFeature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
