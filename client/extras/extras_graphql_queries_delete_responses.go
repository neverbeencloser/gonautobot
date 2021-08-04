package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasGraphqlQueriesDeleteReader is a Reader for the ExtrasGraphqlQueriesDelete structure.
type ExtrasGraphqlQueriesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGraphqlQueriesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasGraphqlQueriesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGraphqlQueriesDeleteNoContent creates a ExtrasGraphqlQueriesDeleteNoContent with default headers values
func NewExtrasGraphqlQueriesDeleteNoContent() *ExtrasGraphqlQueriesDeleteNoContent {
	return &ExtrasGraphqlQueriesDeleteNoContent{}
}

/* ExtrasGraphqlQueriesDeleteNoContent describes a response with status code 204, with default header values.

ExtrasGraphqlQueriesDeleteNoContent extras graphql queries delete no content
*/
type ExtrasGraphqlQueriesDeleteNoContent struct {
}

func (o *ExtrasGraphqlQueriesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/graphql-queries/{id}/][%d] extrasGraphqlQueriesDeleteNoContent ", 204)
}

func (o *ExtrasGraphqlQueriesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
