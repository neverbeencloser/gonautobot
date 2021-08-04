package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigComplianceCreateReader is a Reader for the PluginsGoldenConfigConfigComplianceCreate structure.
type PluginsGoldenConfigConfigComplianceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigComplianceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsGoldenConfigConfigComplianceCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigComplianceCreateCreated creates a PluginsGoldenConfigConfigComplianceCreateCreated with default headers values
func NewPluginsGoldenConfigConfigComplianceCreateCreated() *PluginsGoldenConfigConfigComplianceCreateCreated {
	return &PluginsGoldenConfigConfigComplianceCreateCreated{}
}

/* PluginsGoldenConfigConfigComplianceCreateCreated describes a response with status code 201, with default header values.

PluginsGoldenConfigConfigComplianceCreateCreated plugins golden config config compliance create created
*/
type PluginsGoldenConfigConfigComplianceCreateCreated struct {
	Payload *models.ConfigCompliance
}

func (o *PluginsGoldenConfigConfigComplianceCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/golden-config/config-compliance/][%d] pluginsGoldenConfigConfigComplianceCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsGoldenConfigConfigComplianceCreateCreated) GetPayload() *models.ConfigCompliance {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigComplianceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCompliance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
