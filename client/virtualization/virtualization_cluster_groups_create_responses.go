package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClusterGroupsCreateReader is a Reader for the VirtualizationClusterGroupsCreate structure.
type VirtualizationClusterGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVirtualizationClusterGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsCreateCreated creates a VirtualizationClusterGroupsCreateCreated with default headers values
func NewVirtualizationClusterGroupsCreateCreated() *VirtualizationClusterGroupsCreateCreated {
	return &VirtualizationClusterGroupsCreateCreated{}
}

/* VirtualizationClusterGroupsCreateCreated describes a response with status code 201, with default header values.

VirtualizationClusterGroupsCreateCreated virtualization cluster groups create created
*/
type VirtualizationClusterGroupsCreateCreated struct {
	Payload *models.ClusterGroup
}

func (o *VirtualizationClusterGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /virtualization/cluster-groups/][%d] virtualizationClusterGroupsCreateCreated  %+v", 201, o.Payload)
}
func (o *VirtualizationClusterGroupsCreateCreated) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
