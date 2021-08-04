package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClustersBulkUpdateReader is a Reader for the VirtualizationClustersBulkUpdate structure.
type VirtualizationClustersBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClustersBulkUpdateOK creates a VirtualizationClustersBulkUpdateOK with default headers values
func NewVirtualizationClustersBulkUpdateOK() *VirtualizationClustersBulkUpdateOK {
	return &VirtualizationClustersBulkUpdateOK{}
}

/* VirtualizationClustersBulkUpdateOK describes a response with status code 200, with default header values.

VirtualizationClustersBulkUpdateOK virtualization clusters bulk update o k
*/
type VirtualizationClustersBulkUpdateOK struct {
	Payload *models.Cluster
}

func (o *VirtualizationClustersBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualizationClustersBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClustersBulkUpdateOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
