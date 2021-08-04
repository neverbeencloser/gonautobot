package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClusterGroupsReadReader is a Reader for the VirtualizationClusterGroupsRead structure.
type VirtualizationClusterGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsReadOK creates a VirtualizationClusterGroupsReadOK with default headers values
func NewVirtualizationClusterGroupsReadOK() *VirtualizationClusterGroupsReadOK {
	return &VirtualizationClusterGroupsReadOK{}
}

/* VirtualizationClusterGroupsReadOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsReadOK virtualization cluster groups read o k
*/
type VirtualizationClusterGroupsReadOK struct {
	Payload *models.ClusterGroup
}

func (o *VirtualizationClusterGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsReadOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClusterGroupsReadOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
