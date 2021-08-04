package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasGraphqlQueriesBulkPartialUpdateReader is a Reader for the ExtrasGraphqlQueriesBulkPartialUpdate structure.
type ExtrasGraphqlQueriesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGraphqlQueriesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasGraphqlQueriesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGraphqlQueriesBulkPartialUpdateOK creates a ExtrasGraphqlQueriesBulkPartialUpdateOK with default headers values
func NewExtrasGraphqlQueriesBulkPartialUpdateOK() *ExtrasGraphqlQueriesBulkPartialUpdateOK {
	return &ExtrasGraphqlQueriesBulkPartialUpdateOK{}
}

/* ExtrasGraphqlQueriesBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasGraphqlQueriesBulkPartialUpdateOK extras graphql queries bulk partial update o k
*/
type ExtrasGraphqlQueriesBulkPartialUpdateOK struct {
	Payload *models.GraphQLQuery
}

func (o *ExtrasGraphqlQueriesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/graphql-queries/][%d] extrasGraphqlQueriesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasGraphqlQueriesBulkPartialUpdateOK) GetPayload() *models.GraphQLQuery {
	return o.Payload
}

func (o *ExtrasGraphqlQueriesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GraphQLQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
