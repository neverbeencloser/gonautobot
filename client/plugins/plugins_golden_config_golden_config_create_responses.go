package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigGoldenConfigCreateReader is a Reader for the PluginsGoldenConfigGoldenConfigCreate structure.
type PluginsGoldenConfigGoldenConfigCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPluginsGoldenConfigGoldenConfigCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigCreateCreated creates a PluginsGoldenConfigGoldenConfigCreateCreated with default headers values
func NewPluginsGoldenConfigGoldenConfigCreateCreated() *PluginsGoldenConfigGoldenConfigCreateCreated {
	return &PluginsGoldenConfigGoldenConfigCreateCreated{}
}

/* PluginsGoldenConfigGoldenConfigCreateCreated describes a response with status code 201, with default header values.

PluginsGoldenConfigGoldenConfigCreateCreated plugins golden config golden config create created
*/
type PluginsGoldenConfigGoldenConfigCreateCreated struct {
	Payload *models.GoldenConfig
}

func (o *PluginsGoldenConfigGoldenConfigCreateCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/golden-config/golden-config/][%d] pluginsGoldenConfigGoldenConfigCreateCreated  %+v", 201, o.Payload)
}
func (o *PluginsGoldenConfigGoldenConfigCreateCreated) GetPayload() *models.GoldenConfig {
	return o.Payload
}

func (o *PluginsGoldenConfigGoldenConfigCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoldenConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
