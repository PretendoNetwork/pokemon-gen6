// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemUserInfo is a type within the ServiceItem protocol
type ServiceItemUserInfo struct {
	types.Structure
	ApplicationBuffer *types.QBuffer
}

// WriteTo writes the ServiceItemUserInfo to the given writable
func (siui *ServiceItemUserInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siui.ApplicationBuffer.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siui.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemUserInfo from the given readable
func (siui *ServiceItemUserInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = siui.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUserInfo header. %s", err.Error())
	}

	err = siui.ApplicationBuffer.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUserInfo.ApplicationBuffer. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemUserInfo
func (siui *ServiceItemUserInfo) Copy() types.RVType {
	copied := NewServiceItemUserInfo()

	copied.StructureVersion = siui.StructureVersion
	copied.ApplicationBuffer = siui.ApplicationBuffer.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given ServiceItemUserInfo contains the same data as the current ServiceItemUserInfo
func (siui *ServiceItemUserInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemUserInfo); !ok {
		return false
	}

	other := o.(*ServiceItemUserInfo)

	if siui.StructureVersion != other.StructureVersion {
		return false
	}

	return siui.ApplicationBuffer.Equals(other.ApplicationBuffer)
}

// String returns the string representation of the ServiceItemUserInfo
func (siui *ServiceItemUserInfo) String() string {
	return siui.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemUserInfo using the provided indentation level
func (siui *ServiceItemUserInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemUserInfo{\n")
	b.WriteString(fmt.Sprintf("%sApplicationBuffer: %s,\n", indentationValues, siui.ApplicationBuffer))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemUserInfo returns a new ServiceItemUserInfo
func NewServiceItemUserInfo() *ServiceItemUserInfo {
	siui := &ServiceItemUserInfo{
		ApplicationBuffer: types.NewQBuffer(nil),
	}

	return siui
}
