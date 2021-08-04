package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigRemoveUpdateReader is a Reader for the PluginsGoldenConfigConfigRemoveUpdate structure.
type PluginsGoldenConfigConfigRemoveUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigRemoveUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigRemoveUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigRemoveUpdateOK creates a PluginsGoldenConfigConfigRemoveUpdateOK with default headers values
func NewPluginsGoldenConfigConfigRemoveUpdateOK() *PluginsGoldenConfigConfigRemoveUpdateOK {
	return &PluginsGoldenConfigConfigRemoveUpdateOK{}
}

/* PluginsGoldenConfigConfigRemoveUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigRemoveUpdateOK plugins golden config config remove update o k
*/
type PluginsGoldenConfigConfigRemoveUpdateOK struct {
	Payload *models.ConfigRemove
}

func (o *PluginsGoldenConfigConfigRemoveUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/config-remove/{id}/][%d] pluginsGoldenConfigConfigRemoveUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigRemoveUpdateOK) GetPayload() *models.ConfigRemove {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigRemoveUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigRemove)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
