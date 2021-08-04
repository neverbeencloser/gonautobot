package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGraphqlQueriesCreateReader is a Reader for the ExtrasGraphqlQueriesCreate structure.
type ExtrasGraphqlQueriesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGraphqlQueriesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasGraphqlQueriesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGraphqlQueriesCreateCreated creates a ExtrasGraphqlQueriesCreateCreated with default headers values
func NewExtrasGraphqlQueriesCreateCreated() *ExtrasGraphqlQueriesCreateCreated {
	return &ExtrasGraphqlQueriesCreateCreated{}
}

/* ExtrasGraphqlQueriesCreateCreated describes a response with status code 201, with default header values.

ExtrasGraphqlQueriesCreateCreated extras graphql queries create created
*/
type ExtrasGraphqlQueriesCreateCreated struct {
	Payload *models.GraphQLQuery
}

func (o *ExtrasGraphqlQueriesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/graphql-queries/][%d] extrasGraphqlQueriesCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasGraphqlQueriesCreateCreated) GetPayload() *models.GraphQLQuery {
	return o.Payload
}

func (o *ExtrasGraphqlQueriesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GraphQLQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
