package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomLinksPartialUpdateReader is a Reader for the ExtrasCustomLinksPartialUpdate structure.
type ExtrasCustomLinksPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomLinksPartialUpdateOK creates a ExtrasCustomLinksPartialUpdateOK with default headers values
func NewExtrasCustomLinksPartialUpdateOK() *ExtrasCustomLinksPartialUpdateOK {
	return &ExtrasCustomLinksPartialUpdateOK{}
}

/* ExtrasCustomLinksPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomLinksPartialUpdateOK extras custom links partial update o k
*/
type ExtrasCustomLinksPartialUpdateOK struct {
	Payload *models.CustomLink
}

func (o *ExtrasCustomLinksPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-links/{id}/][%d] extrasCustomLinksPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomLinksPartialUpdateOK) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
