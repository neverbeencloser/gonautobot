package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPanelsBulkUpdateReader is a Reader for the DcimPowerPanelsBulkUpdate structure.
type DcimPowerPanelsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPanelsBulkUpdateOK creates a DcimPowerPanelsBulkUpdateOK with default headers values
func NewDcimPowerPanelsBulkUpdateOK() *DcimPowerPanelsBulkUpdateOK {
	return &DcimPowerPanelsBulkUpdateOK{}
}

/* DcimPowerPanelsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerPanelsBulkUpdateOK dcim power panels bulk update o k
*/
type DcimPowerPanelsBulkUpdateOK struct {
	Payload *models.PowerPanel
}

func (o *DcimPowerPanelsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-panels/][%d] dcimPowerPanelsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPanelsBulkUpdateOK) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
