package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasStatusesUpdateReader is a Reader for the ExtrasStatusesUpdate structure.
type ExtrasStatusesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasStatusesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasStatusesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasStatusesUpdateOK creates a ExtrasStatusesUpdateOK with default headers values
func NewExtrasStatusesUpdateOK() *ExtrasStatusesUpdateOK {
	return &ExtrasStatusesUpdateOK{}
}

/* ExtrasStatusesUpdateOK describes a response with status code 200, with default header values.

ExtrasStatusesUpdateOK extras statuses update o k
*/
type ExtrasStatusesUpdateOK struct {
	Payload *models.Status
}

func (o *ExtrasStatusesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/statuses/{id}/][%d] extrasStatusesUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasStatusesUpdateOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *ExtrasStatusesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
