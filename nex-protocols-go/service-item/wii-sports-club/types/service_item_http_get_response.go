// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemHTTPGetResponse is a type within the ServiceItem protocol
type ServiceItemHTTPGetResponse struct {
	types.Structure
	Response *types.QBuffer
}

// WriteTo writes the ServiceItemHTTPGetResponse to the given writable
func (sihttpgr *ServiceItemHTTPGetResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sihttpgr.Response.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sihttpgr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemHTTPGetResponse from the given readable
func (sihttpgr *ServiceItemHTTPGetResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = sihttpgr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemHTTPGetResponse header. %s", err.Error())
	}

	err = sihttpgr.Response.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemHTTPGetResponse.Response. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemHTTPGetResponse
func (sihttpgr *ServiceItemHTTPGetResponse) Copy() types.RVType {
	copied := NewServiceItemHTTPGetResponse()

	copied.StructureVersion = sihttpgr.StructureVersion
	copied.Response = sihttpgr.Response.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given ServiceItemHTTPGetResponse contains the same data as the current ServiceItemHTTPGetResponse
func (sihttpgr *ServiceItemHTTPGetResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemHTTPGetResponse); !ok {
		return false
	}

	other := o.(*ServiceItemHTTPGetResponse)

	if sihttpgr.StructureVersion != other.StructureVersion {
		return false
	}

	return sihttpgr.Response.Equals(other.Response)
}

// String returns the string representation of the ServiceItemHTTPGetResponse
func (sihttpgr *ServiceItemHTTPGetResponse) String() string {
	return sihttpgr.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemHTTPGetResponse using the provided indentation level
func (sihttpgr *ServiceItemHTTPGetResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemHTTPGetResponse{\n")
	b.WriteString(fmt.Sprintf("%sResponse: %s,\n", indentationValues, sihttpgr.Response))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemHTTPGetResponse returns a new ServiceItemHTTPGetResponse
func NewServiceItemHTTPGetResponse() *ServiceItemHTTPGetResponse {
	sihttpgr := &ServiceItemHTTPGetResponse{
		Response: types.NewQBuffer(nil),
	}

	return sihttpgr
}
