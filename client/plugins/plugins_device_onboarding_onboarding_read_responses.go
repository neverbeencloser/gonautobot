package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDeviceOnboardingOnboardingReadReader is a Reader for the PluginsDeviceOnboardingOnboardingRead structure.
type PluginsDeviceOnboardingOnboardingReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDeviceOnboardingOnboardingReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsDeviceOnboardingOnboardingReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDeviceOnboardingOnboardingReadOK creates a PluginsDeviceOnboardingOnboardingReadOK with default headers values
func NewPluginsDeviceOnboardingOnboardingReadOK() *PluginsDeviceOnboardingOnboardingReadOK {
	return &PluginsDeviceOnboardingOnboardingReadOK{}
}

/* PluginsDeviceOnboardingOnboardingReadOK describes a response with status code 200, with default header values.

PluginsDeviceOnboardingOnboardingReadOK plugins device onboarding onboarding read o k
*/
type PluginsDeviceOnboardingOnboardingReadOK struct {
	Payload *models.OnboardingTask
}

func (o *PluginsDeviceOnboardingOnboardingReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/device-onboarding/onboarding/{id}/][%d] pluginsDeviceOnboardingOnboardingReadOK  %+v", 200, o.Payload)
}
func (o *PluginsDeviceOnboardingOnboardingReadOK) GetPayload() *models.OnboardingTask {
	return o.Payload
}

func (o *PluginsDeviceOnboardingOnboardingReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OnboardingTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
