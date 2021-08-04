package circuits

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CircuitsProvidersBulkDeleteReader is a Reader for the CircuitsProvidersBulkDelete structure.
type CircuitsProvidersBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsProvidersBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsProvidersBulkDeleteNoContent creates a CircuitsProvidersBulkDeleteNoContent with default headers values
func NewCircuitsProvidersBulkDeleteNoContent() *CircuitsProvidersBulkDeleteNoContent {
	return &CircuitsProvidersBulkDeleteNoContent{}
}

/* CircuitsProvidersBulkDeleteNoContent describes a response with status code 204, with default header values.

CircuitsProvidersBulkDeleteNoContent circuits providers bulk delete no content
*/
type CircuitsProvidersBulkDeleteNoContent struct {
}

func (o *CircuitsProvidersBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/providers/][%d] circuitsProvidersBulkDeleteNoContent ", 204)
}

func (o *CircuitsProvidersBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
