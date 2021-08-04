package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsGoldenConfigSotaggReadReader is a Reader for the PluginsGoldenConfigSotaggRead structure.
type PluginsGoldenConfigSotaggReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigSotaggReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigSotaggReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigSotaggReadOK creates a PluginsGoldenConfigSotaggReadOK with default headers values
func NewPluginsGoldenConfigSotaggReadOK() *PluginsGoldenConfigSotaggReadOK {
	return &PluginsGoldenConfigSotaggReadOK{}
}

/* PluginsGoldenConfigSotaggReadOK describes a response with status code 200, with default header values.

PluginsGoldenConfigSotaggReadOK plugins golden config sotagg read o k
*/
type PluginsGoldenConfigSotaggReadOK struct {
}

func (o *PluginsGoldenConfigSotaggReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/golden-config/sotagg/{id}/][%d] pluginsGoldenConfigSotaggReadOK ", 200)
}

func (o *PluginsGoldenConfigSotaggReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
