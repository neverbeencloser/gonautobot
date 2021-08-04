package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigRemovePartialUpdateReader is a Reader for the PluginsGoldenConfigConfigRemovePartialUpdate structure.
type PluginsGoldenConfigConfigRemovePartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigRemovePartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigRemovePartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigRemovePartialUpdateOK creates a PluginsGoldenConfigConfigRemovePartialUpdateOK with default headers values
func NewPluginsGoldenConfigConfigRemovePartialUpdateOK() *PluginsGoldenConfigConfigRemovePartialUpdateOK {
	return &PluginsGoldenConfigConfigRemovePartialUpdateOK{}
}

/* PluginsGoldenConfigConfigRemovePartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigRemovePartialUpdateOK plugins golden config config remove partial update o k
*/
type PluginsGoldenConfigConfigRemovePartialUpdateOK struct {
	Payload *models.ConfigRemove
}

func (o *PluginsGoldenConfigConfigRemovePartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/config-remove/{id}/][%d] pluginsGoldenConfigConfigRemovePartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigRemovePartialUpdateOK) GetPayload() *models.ConfigRemove {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigRemovePartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigRemove)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
