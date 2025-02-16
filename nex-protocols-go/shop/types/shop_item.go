// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ShopItem is a type within the Shop protocol
type ShopItem struct {
	types.Structure
	ItemID      *types.PrimitiveU32
	ReferenceID *types.QBuffer
	ServiceName *types.String
	ItemCode    *types.String
}

// WriteTo writes the ShopItem to the given writable
func (si *ShopItem) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	si.ItemID.WriteTo(contentWritable)
	si.ReferenceID.WriteTo(contentWritable)
	si.ServiceName.WriteTo(contentWritable)
	si.ItemCode.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	si.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ShopItem from the given readable
func (si *ShopItem) ExtractFrom(readable types.Readable) error {
	var err error

	err = si.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem header. %s", err.Error())
	}

	err = si.ItemID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ItemID. %s", err.Error())
	}

	err = si.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ReferenceID. %s", err.Error())
	}

	err = si.ServiceName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ServiceName. %s", err.Error())
	}

	err = si.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ItemCode. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ShopItem
func (si *ShopItem) Copy() types.RVType {
	copied := NewShopItem()

	copied.StructureVersion = si.StructureVersion
	copied.ItemID = si.ItemID.Copy().(*types.PrimitiveU32)
	copied.ReferenceID = si.ReferenceID.Copy().(*types.QBuffer)
	copied.ServiceName = si.ServiceName.Copy().(*types.String)
	copied.ItemCode = si.ItemCode.Copy().(*types.String)

	return copied
}

// Equals checks if the given ShopItem contains the same data as the current ShopItem
func (si *ShopItem) Equals(o types.RVType) bool {
	if _, ok := o.(*ShopItem); !ok {
		return false
	}

	other := o.(*ShopItem)

	if si.StructureVersion != other.StructureVersion {
		return false
	}

	if !si.ItemID.Equals(other.ItemID) {
		return false
	}

	if !si.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !si.ServiceName.Equals(other.ServiceName) {
		return false
	}

	return si.ItemCode.Equals(other.ItemCode)
}

// String returns the string representation of the ShopItem
func (si *ShopItem) String() string {
	return si.FormatToString(0)
}

// FormatToString pretty-prints the ShopItem using the provided indentation level
func (si *ShopItem) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ShopItem{\n")
	b.WriteString(fmt.Sprintf("%sItemID: %s,\n", indentationValues, si.ItemID))
	b.WriteString(fmt.Sprintf("%sReferenceID: %s,\n", indentationValues, si.ReferenceID))
	b.WriteString(fmt.Sprintf("%sServiceName: %s,\n", indentationValues, si.ServiceName))
	b.WriteString(fmt.Sprintf("%sItemCode: %s,\n", indentationValues, si.ItemCode))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewShopItem returns a new ShopItem
func NewShopItem() *ShopItem {
	si := &ShopItem{
		ItemID:      types.NewPrimitiveU32(0),
		ReferenceID: types.NewQBuffer(nil),
		ServiceName: types.NewString(""),
		ItemCode:    types.NewString(""),
	}

	return si
}
