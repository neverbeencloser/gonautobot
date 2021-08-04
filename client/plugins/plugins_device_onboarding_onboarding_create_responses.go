package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsDeviceOnboardingOnboardingCreateReader is a Reader for the PluginsDeviceOnboardingOnboardingCreate structure.
type PluginsDeviceOnboardingOnboardingCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDeviceOnboardingOnboardingCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsDeviceOnboardingOnboardingCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDeviceOnboardingOnboardingCreateCreated creates a PluginsDeviceOnboardingOnboardingCreateCreated with default headers values
func NewPluginsDeviceOnboardingOnboardingCreateCreated() *PluginsDeviceOnboardingOnboardingCreateCreated {
	return &PluginsDeviceOnboardingOnboardingCreateCreated{}
}

/* PluginsDeviceOnboardingOnboardingCreateCreated describes a response with status code 201, with default header values.

PluginsDeviceOnboardingOnboardingCreateCreated plugins device onboarding onboarding create created
*/
type PluginsDeviceOnboardingOnboardingCreateCreated struct {
	Payload *models.OnboardingTask
}

func (o *PluginsDeviceOnboardingOnboardingCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/device-onboarding/onboarding/][%d] pluginsDeviceOnboardingOnboardingCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsDeviceOnboardingOnboardingCreateCreated) GetPayload() *models.OnboardingTask {
	return o.Payload
}

func (o *PluginsDeviceOnboardingOnboardingCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OnboardingTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
