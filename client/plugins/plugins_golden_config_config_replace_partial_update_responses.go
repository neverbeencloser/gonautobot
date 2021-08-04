package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigReplacePartialUpdateReader is a Reader for the PluginsGoldenConfigConfigReplacePartialUpdate structure.
type PluginsGoldenConfigConfigReplacePartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigReplacePartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigReplacePartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigReplacePartialUpdateOK creates a PluginsGoldenConfigConfigReplacePartialUpdateOK with default headers values
func NewPluginsGoldenConfigConfigReplacePartialUpdateOK() *PluginsGoldenConfigConfigReplacePartialUpdateOK {
	return &PluginsGoldenConfigConfigReplacePartialUpdateOK{}
}

/* PluginsGoldenConfigConfigReplacePartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigReplacePartialUpdateOK plugins golden config config replace partial update o k
*/
type PluginsGoldenConfigConfigReplacePartialUpdateOK struct {
	Payload *models.ConfigReplace
}

func (o *PluginsGoldenConfigConfigReplacePartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/config-replace/{id}/][%d] pluginsGoldenConfigConfigReplacePartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigReplacePartialUpdateOK) GetPayload() *models.ConfigReplace {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigReplacePartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigReplace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
