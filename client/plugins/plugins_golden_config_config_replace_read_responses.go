package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigReplaceReadReader is a Reader for the PluginsGoldenConfigConfigReplaceRead structure.
type PluginsGoldenConfigConfigReplaceReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigReplaceReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigReplaceReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigReplaceReadOK creates a PluginsGoldenConfigConfigReplaceReadOK with default headers values
func NewPluginsGoldenConfigConfigReplaceReadOK() *PluginsGoldenConfigConfigReplaceReadOK {
	return &PluginsGoldenConfigConfigReplaceReadOK{}
}

/* PluginsGoldenConfigConfigReplaceReadOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigReplaceReadOK plugins golden config config replace read o k
*/
type PluginsGoldenConfigConfigReplaceReadOK struct {
	Payload *models.ConfigReplace
}

func (o *PluginsGoldenConfigConfigReplaceReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/golden-config/config-replace/{id}/][%d] pluginsGoldenConfigConfigReplaceReadOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigReplaceReadOK) GetPayload() *models.ConfigReplace {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigReplaceReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigReplace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
