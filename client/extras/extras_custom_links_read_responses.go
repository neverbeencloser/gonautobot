package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomLinksReadReader is a Reader for the ExtrasCustomLinksRead structure.
type ExtrasCustomLinksReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomLinksReadOK creates a ExtrasCustomLinksReadOK with default headers values
func NewExtrasCustomLinksReadOK() *ExtrasCustomLinksReadOK {
	return &ExtrasCustomLinksReadOK{}
}

/* ExtrasCustomLinksReadOK describes a response with status code 200, with default header values.

ExtrasCustomLinksReadOK extras custom links read o k
*/
type ExtrasCustomLinksReadOK struct {
	Payload *models.CustomLink
}

func (o *ExtrasCustomLinksReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/custom-links/{id}/][%d] extrasCustomLinksReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomLinksReadOK) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
