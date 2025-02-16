// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetServiceItemRightParam is a type within the ServiceItem protocol
type ServiceItemGetServiceItemRightParam struct {
	types.Structure
	ReferenceID *types.String
	DeviceID    *types.String
	UniqueID    *types.PrimitiveU32
	ItemGroup   *types.PrimitiveU8
	Platform    *types.PrimitiveU8 // * Revision 1
}

// WriteTo writes the ServiceItemGetServiceItemRightParam to the given writable
func (sigsirp *ServiceItemGetServiceItemRightParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sigsirp.ReferenceID.WriteTo(contentWritable)
	sigsirp.DeviceID.WriteTo(contentWritable)
	sigsirp.UniqueID.WriteTo(contentWritable)
	sigsirp.ItemGroup.WriteTo(contentWritable)

	if sigsirp.StructureVersion >= 1 {
		sigsirp.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	sigsirp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetServiceItemRightParam from the given readable
func (sigsirp *ServiceItemGetServiceItemRightParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sigsirp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam header. %s", err.Error())
	}

	err = sigsirp.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.ReferenceID. %s", err.Error())
	}

	err = sigsirp.DeviceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.DeviceID. %s", err.Error())
	}

	err = sigsirp.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.UniqueID. %s", err.Error())
	}

	err = sigsirp.ItemGroup.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.ItemGroup. %s", err.Error())
	}

	if sigsirp.StructureVersion >= 1 {
		err = sigsirp.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.Platform. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetServiceItemRightParam
func (sigsirp *ServiceItemGetServiceItemRightParam) Copy() types.RVType {
	copied := NewServiceItemGetServiceItemRightParam()

	copied.StructureVersion = sigsirp.StructureVersion
	copied.ReferenceID = sigsirp.ReferenceID.Copy().(*types.String)
	copied.DeviceID = sigsirp.DeviceID.Copy().(*types.String)
	copied.UniqueID = sigsirp.UniqueID.Copy().(*types.PrimitiveU32)
	copied.ItemGroup = sigsirp.ItemGroup.Copy().(*types.PrimitiveU8)
	copied.Platform = sigsirp.Platform.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given ServiceItemGetServiceItemRightParam contains the same data as the current ServiceItemGetServiceItemRightParam
func (sigsirp *ServiceItemGetServiceItemRightParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetServiceItemRightParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetServiceItemRightParam)

	if sigsirp.StructureVersion != other.StructureVersion {
		return false
	}

	if !sigsirp.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !sigsirp.DeviceID.Equals(other.DeviceID) {
		return false
	}

	if !sigsirp.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !sigsirp.ItemGroup.Equals(other.ItemGroup) {
		return false
	}

	return sigsirp.Platform.Equals(other.Platform)
}

// String returns the string representation of the ServiceItemGetServiceItemRightParam
func (sigsirp *ServiceItemGetServiceItemRightParam) String() string {
	return sigsirp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetServiceItemRightParam using the provided indentation level
func (sigsirp *ServiceItemGetServiceItemRightParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetServiceItemRightParam{\n")
	b.WriteString(fmt.Sprintf("%sReferenceID: %s,\n", indentationValues, sigsirp.ReferenceID))
	b.WriteString(fmt.Sprintf("%sDeviceID: %s,\n", indentationValues, sigsirp.DeviceID))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, sigsirp.UniqueID))
	b.WriteString(fmt.Sprintf("%sItemGroup: %s,\n", indentationValues, sigsirp.ItemGroup))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, sigsirp.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetServiceItemRightParam returns a new ServiceItemGetServiceItemRightParam
func NewServiceItemGetServiceItemRightParam() *ServiceItemGetServiceItemRightParam {
	sigsirp := &ServiceItemGetServiceItemRightParam{
		ReferenceID: types.NewString(""),
		DeviceID:    types.NewString(""),
		UniqueID:    types.NewPrimitiveU32(0),
		ItemGroup:   types.NewPrimitiveU8(0),
		Platform:    types.NewPrimitiveU8(0),
	}

	return sigsirp
}
