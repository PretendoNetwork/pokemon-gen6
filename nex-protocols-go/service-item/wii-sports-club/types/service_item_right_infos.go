// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemRightInfos is a type within the ServiceItem protocol
type ServiceItemRightInfos struct {
	types.Structure
	RightInfos *types.List[*ServiceItemRightInfo]
}

// WriteTo writes the ServiceItemRightInfos to the given writable
func (siri *ServiceItemRightInfos) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siri.RightInfos.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siri.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemRightInfos from the given readable
func (siri *ServiceItemRightInfos) ExtractFrom(readable types.Readable) error {
	var err error

	err = siri.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos header. %s", err.Error())
	}

	err = siri.RightInfos.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.RightInfos. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemRightInfos
func (siri *ServiceItemRightInfos) Copy() types.RVType {
	copied := NewServiceItemRightInfos()

	copied.StructureVersion = siri.StructureVersion
	copied.RightInfos = siri.RightInfos.Copy().(*types.List[*ServiceItemRightInfo])

	return copied
}

// Equals checks if the given ServiceItemRightInfos contains the same data as the current ServiceItemRightInfos
func (siri *ServiceItemRightInfos) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightInfos); !ok {
		return false
	}

	other := o.(*ServiceItemRightInfos)

	if siri.StructureVersion != other.StructureVersion {
		return false
	}

	return siri.RightInfos.Equals(other.RightInfos)
}

// String returns the string representation of the ServiceItemRightInfos
func (siri *ServiceItemRightInfos) String() string {
	return siri.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemRightInfos using the provided indentation level
func (siri *ServiceItemRightInfos) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightInfos{\n")
	b.WriteString(fmt.Sprintf("%sRightInfos: %s,\n", indentationValues, siri.RightInfos))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightInfos returns a new ServiceItemRightInfos
func NewServiceItemRightInfos() *ServiceItemRightInfos {
	siri := &ServiceItemRightInfos{
		RightInfos: types.NewList[*ServiceItemRightInfo](),
	}

	siri.RightInfos.Type = NewServiceItemRightInfo()

	return siri
}
