package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasImageAttachmentsBulkPartialUpdateReader is a Reader for the ExtrasImageAttachmentsBulkPartialUpdate structure.
type ExtrasImageAttachmentsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasImageAttachmentsBulkPartialUpdateOK creates a ExtrasImageAttachmentsBulkPartialUpdateOK with default headers values
func NewExtrasImageAttachmentsBulkPartialUpdateOK() *ExtrasImageAttachmentsBulkPartialUpdateOK {
	return &ExtrasImageAttachmentsBulkPartialUpdateOK{}
}

/* ExtrasImageAttachmentsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsBulkPartialUpdateOK extras image attachments bulk partial update o k
*/
type ExtrasImageAttachmentsBulkPartialUpdateOK struct {
	Payload *models.ImageAttachment
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/][%d] extrasImageAttachmentsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
