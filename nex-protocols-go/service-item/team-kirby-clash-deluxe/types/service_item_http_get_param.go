// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemHTTPGetParam is a type within the ServiceItem protocol
type ServiceItemHTTPGetParam struct {
	types.Structure
	URL *types.String
}

// WriteTo writes the ServiceItemHTTPGetParam to the given writable
func (sihttpgp *ServiceItemHTTPGetParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sihttpgp.URL.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sihttpgp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemHTTPGetParam from the given readable
func (sihttpgp *ServiceItemHTTPGetParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sihttpgp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemHTTPGetParam header. %s", err.Error())
	}

	err = sihttpgp.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemHTTPGetParam.URL. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemHTTPGetParam
func (sihttpgp *ServiceItemHTTPGetParam) Copy() types.RVType {
	copied := NewServiceItemHTTPGetParam()

	copied.StructureVersion = sihttpgp.StructureVersion
	copied.URL = sihttpgp.URL.Copy().(*types.String)

	return copied
}

// Equals checks if the given ServiceItemHTTPGetParam contains the same data as the current ServiceItemHTTPGetParam
func (sihttpgp *ServiceItemHTTPGetParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemHTTPGetParam); !ok {
		return false
	}

	other := o.(*ServiceItemHTTPGetParam)

	if sihttpgp.StructureVersion != other.StructureVersion {
		return false
	}

	return sihttpgp.URL.Equals(other.URL)
}

// String returns the string representation of the ServiceItemHTTPGetParam
func (sihttpgp *ServiceItemHTTPGetParam) String() string {
	return sihttpgp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemHTTPGetParam using the provided indentation level
func (sihttpgp *ServiceItemHTTPGetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemHTTPGetParam{\n")
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, sihttpgp.URL))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemHTTPGetParam returns a new ServiceItemHTTPGetParam
func NewServiceItemHTTPGetParam() *ServiceItemHTTPGetParam {
	sihttpgp := &ServiceItemHTTPGetParam{
		URL: types.NewString(""),
	}

	return sihttpgp
}
