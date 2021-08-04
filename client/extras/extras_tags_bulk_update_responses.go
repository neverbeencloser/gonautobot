package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasTagsBulkUpdateReader is a Reader for the ExtrasTagsBulkUpdate structure.
type ExtrasTagsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasTagsBulkUpdateOK creates a ExtrasTagsBulkUpdateOK with default headers values
func NewExtrasTagsBulkUpdateOK() *ExtrasTagsBulkUpdateOK {
	return &ExtrasTagsBulkUpdateOK{}
}

/* ExtrasTagsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasTagsBulkUpdateOK extras tags bulk update o k
*/
type ExtrasTagsBulkUpdateOK struct {
	Payload *models.Tag
}

func (o *ExtrasTagsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extrasTagsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasTagsBulkUpdateOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
