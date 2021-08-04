package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasStatusesBulkUpdateReader is a Reader for the ExtrasStatusesBulkUpdate structure.
type ExtrasStatusesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasStatusesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasStatusesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasStatusesBulkUpdateOK creates a ExtrasStatusesBulkUpdateOK with default headers values
func NewExtrasStatusesBulkUpdateOK() *ExtrasStatusesBulkUpdateOK {
	return &ExtrasStatusesBulkUpdateOK{}
}

/* ExtrasStatusesBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasStatusesBulkUpdateOK extras statuses bulk update o k
*/
type ExtrasStatusesBulkUpdateOK struct {
	Payload *models.Status
}

func (o *ExtrasStatusesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/statuses/][%d] extrasStatusesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasStatusesBulkUpdateOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *ExtrasStatusesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
