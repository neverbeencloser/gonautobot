package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPanelsBulkPartialUpdateReader is a Reader for the DcimPowerPanelsBulkPartialUpdate structure.
type DcimPowerPanelsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPanelsBulkPartialUpdateOK creates a DcimPowerPanelsBulkPartialUpdateOK with default headers values
func NewDcimPowerPanelsBulkPartialUpdateOK() *DcimPowerPanelsBulkPartialUpdateOK {
	return &DcimPowerPanelsBulkPartialUpdateOK{}
}

/* DcimPowerPanelsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPanelsBulkPartialUpdateOK dcim power panels bulk partial update o k
*/
type DcimPowerPanelsBulkPartialUpdateOK struct {
	Payload *models.PowerPanel
}

func (o *DcimPowerPanelsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/][%d] dcimPowerPanelsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPanelsBulkPartialUpdateOK) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
