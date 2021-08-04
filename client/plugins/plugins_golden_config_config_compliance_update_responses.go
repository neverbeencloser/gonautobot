package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigComplianceUpdateReader is a Reader for the PluginsGoldenConfigConfigComplianceUpdate structure.
type PluginsGoldenConfigConfigComplianceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigComplianceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigComplianceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigComplianceUpdateOK creates a PluginsGoldenConfigConfigComplianceUpdateOK with default headers values
func NewPluginsGoldenConfigConfigComplianceUpdateOK() *PluginsGoldenConfigConfigComplianceUpdateOK {
	return &PluginsGoldenConfigConfigComplianceUpdateOK{}
}

/* PluginsGoldenConfigConfigComplianceUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigComplianceUpdateOK plugins golden config config compliance update o k
*/
type PluginsGoldenConfigConfigComplianceUpdateOK struct {
	Payload *models.ConfigCompliance
}

func (o *PluginsGoldenConfigConfigComplianceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/config-compliance/{id}/][%d] pluginsGoldenConfigConfigComplianceUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigComplianceUpdateOK) GetPayload() *models.ConfigCompliance {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigComplianceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCompliance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
