package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasTagsUpdateReader is a Reader for the ExtrasTagsUpdate structure.
type ExtrasTagsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasTagsUpdateOK creates a ExtrasTagsUpdateOK with default headers values
func NewExtrasTagsUpdateOK() *ExtrasTagsUpdateOK {
	return &ExtrasTagsUpdateOK{}
}

/* ExtrasTagsUpdateOK describes a response with status code 200, with default header values.

ExtrasTagsUpdateOK extras tags update o k
*/
type ExtrasTagsUpdateOK struct {
	Payload *models.Tag
}

func (o *ExtrasTagsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/{id}/][%d] extrasTagsUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasTagsUpdateOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
