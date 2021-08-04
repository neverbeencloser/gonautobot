package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasImageAttachmentsUpdateReader is a Reader for the ExtrasImageAttachmentsUpdate structure.
type ExtrasImageAttachmentsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasImageAttachmentsUpdateOK creates a ExtrasImageAttachmentsUpdateOK with default headers values
func NewExtrasImageAttachmentsUpdateOK() *ExtrasImageAttachmentsUpdateOK {
	return &ExtrasImageAttachmentsUpdateOK{}
}

/* ExtrasImageAttachmentsUpdateOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsUpdateOK extras image attachments update o k
*/
type ExtrasImageAttachmentsUpdateOK struct {
	Payload *models.ImageAttachment
}

func (o *ExtrasImageAttachmentsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/{id}/][%d] extrasImageAttachmentsUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasImageAttachmentsUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
