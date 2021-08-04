package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimSitesBulkPartialUpdateReader is a Reader for the DcimSitesBulkPartialUpdate structure.
type DcimSitesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimSitesBulkPartialUpdateOK creates a DcimSitesBulkPartialUpdateOK with default headers values
func NewDcimSitesBulkPartialUpdateOK() *DcimSitesBulkPartialUpdateOK {
	return &DcimSitesBulkPartialUpdateOK{}
}

/* DcimSitesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimSitesBulkPartialUpdateOK dcim sites bulk partial update o k
*/
type DcimSitesBulkPartialUpdateOK struct {
	Payload *models.Site
}

func (o *DcimSitesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/][%d] dcimSitesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimSitesBulkPartialUpdateOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
