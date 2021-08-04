package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGraphqlQueriesRunReader is a Reader for the ExtrasGraphqlQueriesRun structure.
type ExtrasGraphqlQueriesRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGraphqlQueriesRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasGraphqlQueriesRunCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGraphqlQueriesRunCreated creates a ExtrasGraphqlQueriesRunCreated with default headers values
func NewExtrasGraphqlQueriesRunCreated() *ExtrasGraphqlQueriesRunCreated {
	return &ExtrasGraphqlQueriesRunCreated{}
}

/* ExtrasGraphqlQueriesRunCreated describes a response with status code 201, with default header values.

ExtrasGraphqlQueriesRunCreated extras graphql queries run created
*/
type ExtrasGraphqlQueriesRunCreated struct {
	Payload *models.GraphQLQuery
}

func (o *ExtrasGraphqlQueriesRunCreated) Error() string {
	return fmt.Sprintf("[POST /extras/graphql-queries/{id}/run/][%d] extrasGraphqlQueriesRunCreated  %+v", 201, o.Payload)
}
func (o *ExtrasGraphqlQueriesRunCreated) GetPayload() *models.GraphQLQuery {
	return o.Payload
}

func (o *ExtrasGraphqlQueriesRunCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GraphQLQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
