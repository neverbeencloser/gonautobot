package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigComplianceBulkPartialUpdateReader is a Reader for the PluginsGoldenConfigConfigComplianceBulkPartialUpdate structure.
type PluginsGoldenConfigConfigComplianceBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateOK creates a PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK with default headers values
func NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateOK() *PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK {
	return &PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK{}
}

/* PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK plugins golden config config compliance bulk partial update o k
*/
type PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK struct {
	Payload *models.ConfigCompliance
}

func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/config-compliance/][%d] pluginsGoldenConfigConfigComplianceBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK) GetPayload() *models.ConfigCompliance {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCompliance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
