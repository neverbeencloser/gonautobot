package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigReplaceBulkPartialUpdateReader is a Reader for the PluginsGoldenConfigConfigReplaceBulkPartialUpdate structure.
type PluginsGoldenConfigConfigReplaceBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateOK creates a PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK with default headers values
func NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateOK() *PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK {
	return &PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK{}
}

/* PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK plugins golden config config replace bulk partial update o k
*/
type PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK struct {
	Payload *models.ConfigReplace
}

func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/config-replace/][%d] pluginsGoldenConfigConfigReplaceBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK) GetPayload() *models.ConfigReplace {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigReplace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
