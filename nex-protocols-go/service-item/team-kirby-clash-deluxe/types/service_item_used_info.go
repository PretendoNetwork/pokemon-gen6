// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemUsedInfo is a type within the ServiceItem protocol
type ServiceItemUsedInfo struct {
	types.Structure
	AcquiredCount *types.PrimitiveU32
	UsedCount     *types.PrimitiveU32
}

// WriteTo writes the ServiceItemUsedInfo to the given writable
func (siui *ServiceItemUsedInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siui.AcquiredCount.WriteTo(contentWritable)
	siui.UsedCount.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siui.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemUsedInfo from the given readable
func (siui *ServiceItemUsedInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = siui.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUsedInfo header. %s", err.Error())
	}

	err = siui.AcquiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUsedInfo.AcquiredCount. %s", err.Error())
	}

	err = siui.UsedCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUsedInfo.UsedCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemUsedInfo
func (siui *ServiceItemUsedInfo) Copy() types.RVType {
	copied := NewServiceItemUsedInfo()

	copied.StructureVersion = siui.StructureVersion
	copied.AcquiredCount = siui.AcquiredCount.Copy().(*types.PrimitiveU32)
	copied.UsedCount = siui.UsedCount.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given ServiceItemUsedInfo contains the same data as the current ServiceItemUsedInfo
func (siui *ServiceItemUsedInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemUsedInfo); !ok {
		return false
	}

	other := o.(*ServiceItemUsedInfo)

	if siui.StructureVersion != other.StructureVersion {
		return false
	}

	if !siui.AcquiredCount.Equals(other.AcquiredCount) {
		return false
	}

	return siui.UsedCount.Equals(other.UsedCount)
}

// String returns the string representation of the ServiceItemUsedInfo
func (siui *ServiceItemUsedInfo) String() string {
	return siui.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemUsedInfo using the provided indentation level
func (siui *ServiceItemUsedInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemUsedInfo{\n")
	b.WriteString(fmt.Sprintf("%sAcquiredCount: %s,\n", indentationValues, siui.AcquiredCount))
	b.WriteString(fmt.Sprintf("%sUsedCount: %s,\n", indentationValues, siui.UsedCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemUsedInfo returns a new ServiceItemUsedInfo
func NewServiceItemUsedInfo() *ServiceItemUsedInfo {
	siui := &ServiceItemUsedInfo{
		AcquiredCount: types.NewPrimitiveU32(0),
		UsedCount:     types.NewPrimitiveU32(0),
	}

	return siui
}
