package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClusterGroupsPartialUpdateReader is a Reader for the VirtualizationClusterGroupsPartialUpdate structure.
type VirtualizationClusterGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsPartialUpdateOK creates a VirtualizationClusterGroupsPartialUpdateOK with default headers values
func NewVirtualizationClusterGroupsPartialUpdateOK() *VirtualizationClusterGroupsPartialUpdateOK {
	return &VirtualizationClusterGroupsPartialUpdateOK{}
}

/* VirtualizationClusterGroupsPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsPartialUpdateOK virtualization cluster groups partial update o k
*/
type VirtualizationClusterGroupsPartialUpdateOK struct {
	Payload *models.ClusterGroup
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClusterGroupsPartialUpdateOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
