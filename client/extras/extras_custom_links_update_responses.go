package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomLinksUpdateReader is a Reader for the ExtrasCustomLinksUpdate structure.
type ExtrasCustomLinksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomLinksUpdateOK creates a ExtrasCustomLinksUpdateOK with default headers values
func NewExtrasCustomLinksUpdateOK() *ExtrasCustomLinksUpdateOK {
	return &ExtrasCustomLinksUpdateOK{}
}

/* ExtrasCustomLinksUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomLinksUpdateOK extras custom links update o k
*/
type ExtrasCustomLinksUpdateOK struct {
	Payload *models.CustomLink
}

func (o *ExtrasCustomLinksUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-links/{id}/][%d] extrasCustomLinksUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomLinksUpdateOK) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
