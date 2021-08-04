package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRearPortTemplatesPartialUpdateReader is a Reader for the DcimRearPortTemplatesPartialUpdate structure.
type DcimRearPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortTemplatesPartialUpdateOK creates a DcimRearPortTemplatesPartialUpdateOK with default headers values
func NewDcimRearPortTemplatesPartialUpdateOK() *DcimRearPortTemplatesPartialUpdateOK {
	return &DcimRearPortTemplatesPartialUpdateOK{}
}

/* DcimRearPortTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimRearPortTemplatesPartialUpdateOK dcim rear port templates partial update o k
*/
type DcimRearPortTemplatesPartialUpdateOK struct {
	Payload *models.RearPortTemplate
}

func (o *DcimRearPortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRearPortTemplatesPartialUpdateOK) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
