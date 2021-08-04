package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGraphqlQueriesUpdateReader is a Reader for the ExtrasGraphqlQueriesUpdate structure.
type ExtrasGraphqlQueriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGraphqlQueriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasGraphqlQueriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGraphqlQueriesUpdateOK creates a ExtrasGraphqlQueriesUpdateOK with default headers values
func NewExtrasGraphqlQueriesUpdateOK() *ExtrasGraphqlQueriesUpdateOK {
	return &ExtrasGraphqlQueriesUpdateOK{}
}

/* ExtrasGraphqlQueriesUpdateOK describes a response with status code 200, with default header values.

ExtrasGraphqlQueriesUpdateOK extras graphql queries update o k
*/
type ExtrasGraphqlQueriesUpdateOK struct {
	Payload *models.GraphQLQuery
}

func (o *ExtrasGraphqlQueriesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/graphql-queries/{id}/][%d] extrasGraphqlQueriesUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasGraphqlQueriesUpdateOK) GetPayload() *models.GraphQLQuery {
	return o.Payload
}

func (o *ExtrasGraphqlQueriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GraphQLQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
