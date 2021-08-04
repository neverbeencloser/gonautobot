package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomLinksCreateReader is a Reader for the ExtrasCustomLinksCreate structure.
type ExtrasCustomLinksCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasCustomLinksCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomLinksCreateCreated creates a ExtrasCustomLinksCreateCreated with default headers values
func NewExtrasCustomLinksCreateCreated() *ExtrasCustomLinksCreateCreated {
	return &ExtrasCustomLinksCreateCreated{}
}

/* ExtrasCustomLinksCreateCreated describes a response with status code 201, with default header values.

ExtrasCustomLinksCreateCreated extras custom links create created
*/
type ExtrasCustomLinksCreateCreated struct {
	Payload *models.CustomLink
}

func (o *ExtrasCustomLinksCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/custom-links/][%d] extrasCustomLinksCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasCustomLinksCreateCreated) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
