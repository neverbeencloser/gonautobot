package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasContentTypesReadReader is a Reader for the ExtrasContentTypesRead structure.
type ExtrasContentTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasContentTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasContentTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasContentTypesReadOK creates a ExtrasContentTypesReadOK with default headers values
func NewExtrasContentTypesReadOK() *ExtrasContentTypesReadOK {
	return &ExtrasContentTypesReadOK{}
}

/* ExtrasContentTypesReadOK describes a response with status code 200, with default header values.

ExtrasContentTypesReadOK extras content types read o k
*/
type ExtrasContentTypesReadOK struct {
	Payload *models.ContentType
}

func (o *ExtrasContentTypesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extrasContentTypesReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasContentTypesReadOK) GetPayload() *models.ContentType {
	return o.Payload
}

func (o *ExtrasContentTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
