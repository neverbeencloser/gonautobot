package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimSitesBulkUpdateReader is a Reader for the DcimSitesBulkUpdate structure.
type DcimSitesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimSitesBulkUpdateOK creates a DcimSitesBulkUpdateOK with default headers values
func NewDcimSitesBulkUpdateOK() *DcimSitesBulkUpdateOK {
	return &DcimSitesBulkUpdateOK{}
}

/* DcimSitesBulkUpdateOK describes a response with status code 200, with default header values.

DcimSitesBulkUpdateOK dcim sites bulk update o k
*/
type DcimSitesBulkUpdateOK struct {
	Payload *models.Site
}

func (o *DcimSitesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/sites/][%d] dcimSitesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimSitesBulkUpdateOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
