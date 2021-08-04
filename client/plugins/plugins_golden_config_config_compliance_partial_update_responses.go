package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigCompliancePartialUpdateReader is a Reader for the PluginsGoldenConfigConfigCompliancePartialUpdate structure.
type PluginsGoldenConfigConfigCompliancePartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigCompliancePartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigCompliancePartialUpdateOK creates a PluginsGoldenConfigConfigCompliancePartialUpdateOK with default headers values
func NewPluginsGoldenConfigConfigCompliancePartialUpdateOK() *PluginsGoldenConfigConfigCompliancePartialUpdateOK {
	return &PluginsGoldenConfigConfigCompliancePartialUpdateOK{}
}

/* PluginsGoldenConfigConfigCompliancePartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigCompliancePartialUpdateOK plugins golden config config compliance partial update o k
*/
type PluginsGoldenConfigConfigCompliancePartialUpdateOK struct {
	Payload *models.ConfigCompliance
}

func (o *PluginsGoldenConfigConfigCompliancePartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/config-compliance/{id}/][%d] pluginsGoldenConfigConfigCompliancePartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateOK) GetPayload() *models.ConfigCompliance {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigCompliancePartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCompliance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
