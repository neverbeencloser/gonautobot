package status

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StatusListReader is a Reader for the StatusList structure.
type StatusListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStatusListOK creates a StatusListOK with default headers values
func NewStatusListOK() *StatusListOK {
	return &StatusListOK{}
}

/* StatusListOK describes a response with status code 200, with default header values.

StatusListOK status list o k
*/
type StatusListOK struct {
}

func (o *StatusListOK) Error() string {
	return fmt.Sprintf("[GET /status/][%d] statusListOK ", 200)
}

func (o *StatusListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
