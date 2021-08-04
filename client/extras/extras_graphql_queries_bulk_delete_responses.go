package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasGraphqlQueriesBulkDeleteReader is a Reader for the ExtrasGraphqlQueriesBulkDelete structure.
type ExtrasGraphqlQueriesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGraphqlQueriesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasGraphqlQueriesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGraphqlQueriesBulkDeleteNoContent creates a ExtrasGraphqlQueriesBulkDeleteNoContent with default headers values
func NewExtrasGraphqlQueriesBulkDeleteNoContent() *ExtrasGraphqlQueriesBulkDeleteNoContent {
	return &ExtrasGraphqlQueriesBulkDeleteNoContent{}
}

/* ExtrasGraphqlQueriesBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasGraphqlQueriesBulkDeleteNoContent extras graphql queries bulk delete no content
*/
type ExtrasGraphqlQueriesBulkDeleteNoContent struct {
}

func (o *ExtrasGraphqlQueriesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/graphql-queries/][%d] extrasGraphqlQueriesBulkDeleteNoContent ", 204)
}

func (o *ExtrasGraphqlQueriesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
