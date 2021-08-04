package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigRemoveBulkUpdateReader is a Reader for the PluginsGoldenConfigConfigRemoveBulkUpdate structure.
type PluginsGoldenConfigConfigRemoveBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigRemoveBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigRemoveBulkUpdateOK creates a PluginsGoldenConfigConfigRemoveBulkUpdateOK with default headers values
func NewPluginsGoldenConfigConfigRemoveBulkUpdateOK() *PluginsGoldenConfigConfigRemoveBulkUpdateOK {
	return &PluginsGoldenConfigConfigRemoveBulkUpdateOK{}
}

/* PluginsGoldenConfigConfigRemoveBulkUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigRemoveBulkUpdateOK plugins golden config config remove bulk update o k
*/
type PluginsGoldenConfigConfigRemoveBulkUpdateOK struct {
	Payload *models.ConfigRemove
}

func (o *PluginsGoldenConfigConfigRemoveBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/config-remove/][%d] pluginsGoldenConfigConfigRemoveBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateOK) GetPayload() *models.ConfigRemove {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigRemoveBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigRemove)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
