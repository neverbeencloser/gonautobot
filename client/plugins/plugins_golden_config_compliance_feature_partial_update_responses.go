package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigComplianceFeaturePartialUpdateReader is a Reader for the PluginsGoldenConfigComplianceFeaturePartialUpdate structure.
type PluginsGoldenConfigComplianceFeaturePartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigComplianceFeaturePartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigComplianceFeaturePartialUpdateOK creates a PluginsGoldenConfigComplianceFeaturePartialUpdateOK with default headers values
func NewPluginsGoldenConfigComplianceFeaturePartialUpdateOK() *PluginsGoldenConfigComplianceFeaturePartialUpdateOK {
	return &PluginsGoldenConfigComplianceFeaturePartialUpdateOK{}
}

/* PluginsGoldenConfigComplianceFeaturePartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigComplianceFeaturePartialUpdateOK plugins golden config compliance feature partial update o k
*/
type PluginsGoldenConfigComplianceFeaturePartialUpdateOK struct {
	Payload *models.ComplianceFeature
}

func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/compliance-feature/{id}/][%d] pluginsGoldenConfigComplianceFeaturePartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateOK) GetPayload() *models.ComplianceFeature {
	return o.Payload
}

func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComplianceFeature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
