// Package types implements all the types used by the ShopNintendoBadgeArcade protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ShopPostPlayLogParam is a type within the ShopNintendoBadgeArcade protocol
type ShopPostPlayLogParam struct {
	types.Structure
	Unknown1  *types.List[*types.PrimitiveU32]
	Timestamp *types.DateTime
	Unknown2  *types.String
}

// WriteTo writes the ShopPostPlayLogParam to the given writable
func (spplp *ShopPostPlayLogParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	spplp.Unknown1.WriteTo(contentWritable)
	spplp.Timestamp.WriteTo(contentWritable)
	spplp.Unknown2.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	spplp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ShopPostPlayLogParam from the given readable
func (spplp *ShopPostPlayLogParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = spplp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopPostPlayLogParam header. %s", err.Error())
	}

	err = spplp.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopPostPlayLogParam.Unknown1. %s", err.Error())
	}

	err = spplp.Timestamp.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopPostPlayLogParam.Timestamp. %s", err.Error())
	}

	err = spplp.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopPostPlayLogParam.Unknown2. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ShopPostPlayLogParam
func (spplp *ShopPostPlayLogParam) Copy() types.RVType {
	copied := NewShopPostPlayLogParam()

	copied.StructureVersion = spplp.StructureVersion
	copied.Unknown1 = spplp.Unknown1.Copy().(*types.List[*types.PrimitiveU32])
	copied.Timestamp = spplp.Timestamp.Copy().(*types.DateTime)
	copied.Unknown2 = spplp.Unknown2.Copy().(*types.String)

	return copied
}

// Equals checks if the given ShopPostPlayLogParam contains the same data as the current ShopPostPlayLogParam
func (spplp *ShopPostPlayLogParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ShopPostPlayLogParam); !ok {
		return false
	}

	other := o.(*ShopPostPlayLogParam)

	if spplp.StructureVersion != other.StructureVersion {
		return false
	}

	if !spplp.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !spplp.Timestamp.Equals(other.Timestamp) {
		return false
	}

	return spplp.Unknown2.Equals(other.Unknown2)
}

// String returns the string representation of the ShopPostPlayLogParam
func (spplp *ShopPostPlayLogParam) String() string {
	return spplp.FormatToString(0)
}

// FormatToString pretty-prints the ShopPostPlayLogParam using the provided indentation level
func (spplp *ShopPostPlayLogParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ShopPostPlayLogParam{\n")
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, spplp.Unknown1))
	b.WriteString(fmt.Sprintf("%sTimestamp: %s,\n", indentationValues, spplp.Timestamp.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, spplp.Unknown2))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewShopPostPlayLogParam returns a new ShopPostPlayLogParam
func NewShopPostPlayLogParam() *ShopPostPlayLogParam {
	spplp := &ShopPostPlayLogParam{
		Unknown1:  types.NewList[*types.PrimitiveU32](),
		Timestamp: types.NewDateTime(0),
		Unknown2:  types.NewString(""),
	}

	spplp.Unknown1.Type = types.NewPrimitiveU32(0)

	return spplp
}
