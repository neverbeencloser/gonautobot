package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigComplianceReadReader is a Reader for the PluginsGoldenConfigConfigComplianceRead structure.
type PluginsGoldenConfigConfigComplianceReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigComplianceReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigComplianceReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigComplianceReadOK creates a PluginsGoldenConfigConfigComplianceReadOK with default headers values
func NewPluginsGoldenConfigConfigComplianceReadOK() *PluginsGoldenConfigConfigComplianceReadOK {
	return &PluginsGoldenConfigConfigComplianceReadOK{}
}

/* PluginsGoldenConfigConfigComplianceReadOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigComplianceReadOK plugins golden config config compliance read o k
*/
type PluginsGoldenConfigConfigComplianceReadOK struct {
	Payload *models.ConfigCompliance
}

func (o *PluginsGoldenConfigConfigComplianceReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/golden-config/config-compliance/{id}/][%d] pluginsGoldenConfigConfigComplianceReadOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigComplianceReadOK) GetPayload() *models.ConfigCompliance {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigComplianceReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCompliance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
