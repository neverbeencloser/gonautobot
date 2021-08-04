package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomLinksBulkUpdateReader is a Reader for the ExtrasCustomLinksBulkUpdate structure.
type ExtrasCustomLinksBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomLinksBulkUpdateOK creates a ExtrasCustomLinksBulkUpdateOK with default headers values
func NewExtrasCustomLinksBulkUpdateOK() *ExtrasCustomLinksBulkUpdateOK {
	return &ExtrasCustomLinksBulkUpdateOK{}
}

/* ExtrasCustomLinksBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomLinksBulkUpdateOK extras custom links bulk update o k
*/
type ExtrasCustomLinksBulkUpdateOK struct {
	Payload *models.CustomLink
}

func (o *ExtrasCustomLinksBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-links/][%d] extrasCustomLinksBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomLinksBulkUpdateOK) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
