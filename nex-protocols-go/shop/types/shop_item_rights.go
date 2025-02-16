// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ShopItemRights is a type within the Shop protocol
type ShopItemRights struct {
	types.Structure
	ReferenceID *types.QBuffer
	ItemType    *types.PrimitiveS8
	Attribute   *types.PrimitiveU32
}

// WriteTo writes the ShopItemRights to the given writable
func (sir *ShopItemRights) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sir.ReferenceID.WriteTo(contentWritable)
	sir.ItemType.WriteTo(contentWritable)
	sir.Attribute.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sir.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ShopItemRights from the given readable
func (sir *ShopItemRights) ExtractFrom(readable types.Readable) error {
	var err error

	err = sir.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItemRights header. %s", err.Error())
	}

	err = sir.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItemRights.ReferenceID. %s", err.Error())
	}

	err = sir.ItemType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItemRights.ItemType. %s", err.Error())
	}

	err = sir.Attribute.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItemRights.Attribute. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ShopItemRights
func (sir *ShopItemRights) Copy() types.RVType {
	copied := NewShopItemRights()

	copied.StructureVersion = sir.StructureVersion
	copied.ReferenceID = sir.ReferenceID.Copy().(*types.QBuffer)
	copied.ItemType = sir.ItemType.Copy().(*types.PrimitiveS8)
	copied.Attribute = sir.Attribute.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given ShopItemRights contains the same data as the current ShopItemRights
func (sir *ShopItemRights) Equals(o types.RVType) bool {
	if _, ok := o.(*ShopItemRights); !ok {
		return false
	}

	other := o.(*ShopItemRights)

	if sir.StructureVersion != other.StructureVersion {
		return false
	}

	if !sir.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !sir.ItemType.Equals(other.ItemType) {
		return false
	}

	return sir.Attribute.Equals(other.Attribute)
}

// String returns the string representation of the ShopItemRights
func (sir *ShopItemRights) String() string {
	return sir.FormatToString(0)
}

// FormatToString pretty-prints the ShopItemRights using the provided indentation level
func (sir *ShopItemRights) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ShopItemRights{\n")
	b.WriteString(fmt.Sprintf("%sReferenceID: %s,\n", indentationValues, sir.ReferenceID))
	b.WriteString(fmt.Sprintf("%sItemType: %s,\n", indentationValues, sir.ItemType))
	b.WriteString(fmt.Sprintf("%sAttribute: %s,\n", indentationValues, sir.Attribute))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewShopItemRights returns a new ShopItemRights
func NewShopItemRights() *ShopItemRights {
	sir := &ShopItemRights{
		ReferenceID: types.NewQBuffer(nil),
		ItemType:    types.NewPrimitiveS8(0),
		Attribute:   types.NewPrimitiveU32(0),
	}

	return sir
}
