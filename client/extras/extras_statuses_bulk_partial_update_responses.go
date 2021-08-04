package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasStatusesBulkPartialUpdateReader is a Reader for the ExtrasStatusesBulkPartialUpdate structure.
type ExtrasStatusesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasStatusesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasStatusesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasStatusesBulkPartialUpdateOK creates a ExtrasStatusesBulkPartialUpdateOK with default headers values
func NewExtrasStatusesBulkPartialUpdateOK() *ExtrasStatusesBulkPartialUpdateOK {
	return &ExtrasStatusesBulkPartialUpdateOK{}
}

/* ExtrasStatusesBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasStatusesBulkPartialUpdateOK extras statuses bulk partial update o k
*/
type ExtrasStatusesBulkPartialUpdateOK struct {
	Payload *models.Status
}

func (o *ExtrasStatusesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/statuses/][%d] extrasStatusesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasStatusesBulkPartialUpdateOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *ExtrasStatusesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
