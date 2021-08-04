package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClusterGroupsBulkPartialUpdateReader is a Reader for the VirtualizationClusterGroupsBulkPartialUpdate structure.
type VirtualizationClusterGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsBulkPartialUpdateOK creates a VirtualizationClusterGroupsBulkPartialUpdateOK with default headers values
func NewVirtualizationClusterGroupsBulkPartialUpdateOK() *VirtualizationClusterGroupsBulkPartialUpdateOK {
	return &VirtualizationClusterGroupsBulkPartialUpdateOK{}
}

/* VirtualizationClusterGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsBulkPartialUpdateOK virtualization cluster groups bulk partial update o k
*/
type VirtualizationClusterGroupsBulkPartialUpdateOK struct {
	Payload *models.ClusterGroup
}

func (o *VirtualizationClusterGroupsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/cluster-groups/][%d] virtualizationClusterGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClusterGroupsBulkPartialUpdateOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
