package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigConfigComplianceBulkDeleteReader is a Reader for the PluginsGoldenConfigConfigComplianceBulkDelete structure.
type PluginsGoldenConfigConfigComplianceBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigComplianceBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsGoldenConfigConfigComplianceBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigComplianceBulkDeleteNoContent creates a PluginsGoldenConfigConfigComplianceBulkDeleteNoContent with default headers values
func NewPluginsGoldenConfigConfigComplianceBulkDeleteNoContent() *PluginsGoldenConfigConfigComplianceBulkDeleteNoContent {
	return &PluginsGoldenConfigConfigComplianceBulkDeleteNoContent{}
}

/* PluginsGoldenConfigConfigComplianceBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsGoldenConfigConfigComplianceBulkDeleteNoContent plugins golden config config compliance bulk delete no content
*/
type PluginsGoldenConfigConfigComplianceBulkDeleteNoContent struct {
}

func (o *PluginsGoldenConfigConfigComplianceBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/golden-config/config-compliance/][%d] pluginsGoldenConfigConfigComplianceBulkDeleteNoContent ", 204)
}

func (o *PluginsGoldenConfigConfigComplianceBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
