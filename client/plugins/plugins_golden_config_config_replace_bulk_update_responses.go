package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigReplaceBulkUpdateReader is a Reader for the PluginsGoldenConfigConfigReplaceBulkUpdate structure.
type PluginsGoldenConfigConfigReplaceBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigReplaceBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkUpdateOK creates a PluginsGoldenConfigConfigReplaceBulkUpdateOK with default headers values
func NewPluginsGoldenConfigConfigReplaceBulkUpdateOK() *PluginsGoldenConfigConfigReplaceBulkUpdateOK {
	return &PluginsGoldenConfigConfigReplaceBulkUpdateOK{}
}

/* PluginsGoldenConfigConfigReplaceBulkUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigReplaceBulkUpdateOK plugins golden config config replace bulk update o k
*/
type PluginsGoldenConfigConfigReplaceBulkUpdateOK struct {
	Payload *models.ConfigReplace
}

func (o *PluginsGoldenConfigConfigReplaceBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/config-replace/][%d] pluginsGoldenConfigConfigReplaceBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateOK) GetPayload() *models.ConfigReplace {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigReplaceBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigReplace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
