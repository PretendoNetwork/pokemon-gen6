// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetLawMessageParam is a type within the ServiceItem protocol
type ServiceItemGetLawMessageParam struct {
	types.Structure
	Language *types.String
	UniqueID *types.PrimitiveU32
	Platform *types.PrimitiveU8 // * Revision 1
}

// WriteTo writes the ServiceItemGetLawMessageParam to the given writable
func (siglmp *ServiceItemGetLawMessageParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siglmp.Language.WriteTo(contentWritable)
	siglmp.UniqueID.WriteTo(contentWritable)

	if siglmp.StructureVersion >= 1 {
		siglmp.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	siglmp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetLawMessageParam from the given readable
func (siglmp *ServiceItemGetLawMessageParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = siglmp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageParam header. %s", err.Error())
	}

	err = siglmp.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageParam.Language. %s", err.Error())
	}

	err = siglmp.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageParam.UniqueID. %s", err.Error())
	}

	if siglmp.StructureVersion >= 1 {
		err = siglmp.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetLawMessageParam.Platform. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetLawMessageParam
func (siglmp *ServiceItemGetLawMessageParam) Copy() types.RVType {
	copied := NewServiceItemGetLawMessageParam()

	copied.StructureVersion = siglmp.StructureVersion
	copied.Language = siglmp.Language.Copy().(*types.String)
	copied.UniqueID = siglmp.UniqueID.Copy().(*types.PrimitiveU32)
	copied.Platform = siglmp.Platform.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given ServiceItemGetLawMessageParam contains the same data as the current ServiceItemGetLawMessageParam
func (siglmp *ServiceItemGetLawMessageParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetLawMessageParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetLawMessageParam)

	if siglmp.StructureVersion != other.StructureVersion {
		return false
	}

	if !siglmp.Language.Equals(other.Language) {
		return false
	}

	if !siglmp.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return siglmp.Platform.Equals(other.Platform)
}

// String returns the string representation of the ServiceItemGetLawMessageParam
func (siglmp *ServiceItemGetLawMessageParam) String() string {
	return siglmp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetLawMessageParam using the provided indentation level
func (siglmp *ServiceItemGetLawMessageParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetLawMessageParam{\n")
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, siglmp.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, siglmp.UniqueID))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, siglmp.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetLawMessageParam returns a new ServiceItemGetLawMessageParam
func NewServiceItemGetLawMessageParam() *ServiceItemGetLawMessageParam {
	siglmp := &ServiceItemGetLawMessageParam{
		Language: types.NewString(""),
		UniqueID: types.NewPrimitiveU32(0),
		Platform: types.NewPrimitiveU8(0),
	}

	return siglmp
}
