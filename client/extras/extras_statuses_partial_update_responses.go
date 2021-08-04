package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasStatusesPartialUpdateReader is a Reader for the ExtrasStatusesPartialUpdate structure.
type ExtrasStatusesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasStatusesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasStatusesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasStatusesPartialUpdateOK creates a ExtrasStatusesPartialUpdateOK with default headers values
func NewExtrasStatusesPartialUpdateOK() *ExtrasStatusesPartialUpdateOK {
	return &ExtrasStatusesPartialUpdateOK{}
}

/* ExtrasStatusesPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasStatusesPartialUpdateOK extras statuses partial update o k
*/
type ExtrasStatusesPartialUpdateOK struct {
	Payload *models.Status
}

func (o *ExtrasStatusesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/statuses/{id}/][%d] extrasStatusesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasStatusesPartialUpdateOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *ExtrasStatusesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
