package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigReplaceCreateReader is a Reader for the PluginsGoldenConfigConfigReplaceCreate structure.
type PluginsGoldenConfigConfigReplaceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigReplaceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsGoldenConfigConfigReplaceCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigReplaceCreateCreated creates a PluginsGoldenConfigConfigReplaceCreateCreated with default headers values
func NewPluginsGoldenConfigConfigReplaceCreateCreated() *PluginsGoldenConfigConfigReplaceCreateCreated {
	return &PluginsGoldenConfigConfigReplaceCreateCreated{}
}

/* PluginsGoldenConfigConfigReplaceCreateCreated describes a response with status code 201, with default header values.

PluginsGoldenConfigConfigReplaceCreateCreated plugins golden config config replace create created
*/
type PluginsGoldenConfigConfigReplaceCreateCreated struct {
	Payload *models.ConfigReplace
}

func (o *PluginsGoldenConfigConfigReplaceCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/golden-config/config-replace/][%d] pluginsGoldenConfigConfigReplaceCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsGoldenConfigConfigReplaceCreateCreated) GetPayload() *models.ConfigReplace {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigReplaceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigReplace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
