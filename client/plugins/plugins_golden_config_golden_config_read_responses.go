package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigGoldenConfigReadReader is a Reader for the PluginsGoldenConfigGoldenConfigRead structure.
type PluginsGoldenConfigGoldenConfigReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigGoldenConfigReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigGoldenConfigReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigGoldenConfigReadOK creates a PluginsGoldenConfigGoldenConfigReadOK with default headers values
func NewPluginsGoldenConfigGoldenConfigReadOK() *PluginsGoldenConfigGoldenConfigReadOK {
	return &PluginsGoldenConfigGoldenConfigReadOK{}
}

/* PluginsGoldenConfigGoldenConfigReadOK describes a response with status code 200, with default header values.

PluginsGoldenConfigGoldenConfigReadOK plugins golden config golden config read o k
*/
type PluginsGoldenConfigGoldenConfigReadOK struct {
	Payload *models.GoldenConfig
}

func (o *PluginsGoldenConfigGoldenConfigReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/golden-config/golden-config/{id}/][%d] pluginsGoldenConfigGoldenConfigReadOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigGoldenConfigReadOK) GetPayload() *models.GoldenConfig {
	return o.Payload
}

func (o *PluginsGoldenConfigGoldenConfigReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoldenConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
