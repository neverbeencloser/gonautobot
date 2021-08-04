package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigRemoveCreateReader is a Reader for the PluginsGoldenConfigConfigRemoveCreate structure.
type PluginsGoldenConfigConfigRemoveCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigRemoveCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsGoldenConfigConfigRemoveCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigRemoveCreateCreated creates a PluginsGoldenConfigConfigRemoveCreateCreated with default headers values
func NewPluginsGoldenConfigConfigRemoveCreateCreated() *PluginsGoldenConfigConfigRemoveCreateCreated {
	return &PluginsGoldenConfigConfigRemoveCreateCreated{}
}

/* PluginsGoldenConfigConfigRemoveCreateCreated describes a response with status code 201, with default header values.

PluginsGoldenConfigConfigRemoveCreateCreated plugins golden config config remove create created
*/
type PluginsGoldenConfigConfigRemoveCreateCreated struct {
	Payload *models.ConfigRemove
}

func (o *PluginsGoldenConfigConfigRemoveCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/golden-config/config-remove/][%d] pluginsGoldenConfigConfigRemoveCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsGoldenConfigConfigRemoveCreateCreated) GetPayload() *models.ConfigRemove {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigRemoveCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigRemove)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
