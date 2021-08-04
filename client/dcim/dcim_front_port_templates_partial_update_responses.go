package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortTemplatesPartialUpdateReader is a Reader for the DcimFrontPortTemplatesPartialUpdate structure.
type DcimFrontPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortTemplatesPartialUpdateOK creates a DcimFrontPortTemplatesPartialUpdateOK with default headers values
func NewDcimFrontPortTemplatesPartialUpdateOK() *DcimFrontPortTemplatesPartialUpdateOK {
	return &DcimFrontPortTemplatesPartialUpdateOK{}
}

/* DcimFrontPortTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesPartialUpdateOK dcim front port templates partial update o k
*/
type DcimFrontPortTemplatesPartialUpdateOK struct {
	Payload *models.FrontPortTemplate
}

func (o *DcimFrontPortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortTemplatesPartialUpdateOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
