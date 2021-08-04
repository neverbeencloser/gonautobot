package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsDeviceOnboardingOnboardingDeleteReader is a Reader for the PluginsDeviceOnboardingOnboardingDelete structure.
type PluginsDeviceOnboardingOnboardingDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDeviceOnboardingOnboardingDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsDeviceOnboardingOnboardingDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDeviceOnboardingOnboardingDeleteNoContent creates a PluginsDeviceOnboardingOnboardingDeleteNoContent with default headers values
func NewPluginsDeviceOnboardingOnboardingDeleteNoContent() *PluginsDeviceOnboardingOnboardingDeleteNoContent {
	return &PluginsDeviceOnboardingOnboardingDeleteNoContent{}
}

/* PluginsDeviceOnboardingOnboardingDeleteNoContent describes a response with status code 204, with default header values.

PluginsDeviceOnboardingOnboardingDeleteNoContent plugins device onboarding onboarding delete no content
*/
type PluginsDeviceOnboardingOnboardingDeleteNoContent struct {
}

func (o *PluginsDeviceOnboardingOnboardingDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/device-onboarding/onboarding/{id}/][%d] pluginsDeviceOnboardingOnboardingDeleteNoContent ", 204)
}

func (o *PluginsDeviceOnboardingOnboardingDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
