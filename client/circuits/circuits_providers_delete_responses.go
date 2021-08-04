package circuits

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CircuitsProvidersDeleteReader is a Reader for the CircuitsProvidersDelete structure.
type CircuitsProvidersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsProvidersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsProvidersDeleteNoContent creates a CircuitsProvidersDeleteNoContent with default headers values
func NewCircuitsProvidersDeleteNoContent() *CircuitsProvidersDeleteNoContent {
	return &CircuitsProvidersDeleteNoContent{}
}

/* CircuitsProvidersDeleteNoContent describes a response with status code 204, with default header values.

CircuitsProvidersDeleteNoContent circuits providers delete no content
*/
type CircuitsProvidersDeleteNoContent struct {
}

func (o *CircuitsProvidersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/providers/{id}/][%d] circuitsProvidersDeleteNoContent ", 204)
}

func (o *CircuitsProvidersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
