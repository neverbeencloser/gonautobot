package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasTagsBulkPartialUpdateReader is a Reader for the ExtrasTagsBulkPartialUpdate structure.
type ExtrasTagsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasTagsBulkPartialUpdateOK creates a ExtrasTagsBulkPartialUpdateOK with default headers values
func NewExtrasTagsBulkPartialUpdateOK() *ExtrasTagsBulkPartialUpdateOK {
	return &ExtrasTagsBulkPartialUpdateOK{}
}

/* ExtrasTagsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasTagsBulkPartialUpdateOK extras tags bulk partial update o k
*/
type ExtrasTagsBulkPartialUpdateOK struct {
	Payload *models.Tag
}

func (o *ExtrasTagsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/tags/][%d] extrasTagsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasTagsBulkPartialUpdateOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
