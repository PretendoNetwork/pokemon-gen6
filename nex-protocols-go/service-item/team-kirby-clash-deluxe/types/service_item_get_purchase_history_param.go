// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetPurchaseHistoryParam is a type within the ServiceItem protocol
type ServiceItemGetPurchaseHistoryParam struct {
	types.Structure
	Language *types.String
	Offset   *types.PrimitiveU32
	Size     *types.PrimitiveU32
	UniqueID *types.PrimitiveU32
	Platform *types.PrimitiveU8
}

// WriteTo writes the ServiceItemGetPurchaseHistoryParam to the given writable
func (sigphp *ServiceItemGetPurchaseHistoryParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sigphp.Language.WriteTo(contentWritable)
	sigphp.Offset.WriteTo(contentWritable)
	sigphp.Size.WriteTo(contentWritable)
	sigphp.UniqueID.WriteTo(contentWritable)
	sigphp.Platform.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sigphp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetPurchaseHistoryParam from the given readable
func (sigphp *ServiceItemGetPurchaseHistoryParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sigphp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam header. %s", err.Error())
	}

	err = sigphp.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Language. %s", err.Error())
	}

	err = sigphp.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Offset. %s", err.Error())
	}

	err = sigphp.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Size. %s", err.Error())
	}

	err = sigphp.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.UniqueID. %s", err.Error())
	}

	err = sigphp.Platform.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Platform. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetPurchaseHistoryParam
func (sigphp *ServiceItemGetPurchaseHistoryParam) Copy() types.RVType {
	copied := NewServiceItemGetPurchaseHistoryParam()

	copied.StructureVersion = sigphp.StructureVersion
	copied.Language = sigphp.Language.Copy().(*types.String)
	copied.Offset = sigphp.Offset.Copy().(*types.PrimitiveU32)
	copied.Size = sigphp.Size.Copy().(*types.PrimitiveU32)
	copied.UniqueID = sigphp.UniqueID.Copy().(*types.PrimitiveU32)
	copied.Platform = sigphp.Platform.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given ServiceItemGetPurchaseHistoryParam contains the same data as the current ServiceItemGetPurchaseHistoryParam
func (sigphp *ServiceItemGetPurchaseHistoryParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetPurchaseHistoryParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetPurchaseHistoryParam)

	if sigphp.StructureVersion != other.StructureVersion {
		return false
	}

	if !sigphp.Language.Equals(other.Language) {
		return false
	}

	if !sigphp.Offset.Equals(other.Offset) {
		return false
	}

	if !sigphp.Size.Equals(other.Size) {
		return false
	}

	if !sigphp.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return sigphp.Platform.Equals(other.Platform)
}

// String returns the string representation of the ServiceItemGetPurchaseHistoryParam
func (sigphp *ServiceItemGetPurchaseHistoryParam) String() string {
	return sigphp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetPurchaseHistoryParam using the provided indentation level
func (sigphp *ServiceItemGetPurchaseHistoryParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetPurchaseHistoryParam{\n")
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, sigphp.Language))
	b.WriteString(fmt.Sprintf("%sOffset: %s,\n", indentationValues, sigphp.Offset))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, sigphp.Size))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, sigphp.UniqueID))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, sigphp.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetPurchaseHistoryParam returns a new ServiceItemGetPurchaseHistoryParam
func NewServiceItemGetPurchaseHistoryParam() *ServiceItemGetPurchaseHistoryParam {
	sigphp := &ServiceItemGetPurchaseHistoryParam{
		Language: types.NewString(""),
		Offset:   types.NewPrimitiveU32(0),
		Size:     types.NewPrimitiveU32(0),
		UniqueID: types.NewPrimitiveU32(0),
		Platform: types.NewPrimitiveU8(0),
	}

	return sigphp
}
