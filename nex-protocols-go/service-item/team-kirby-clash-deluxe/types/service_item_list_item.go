// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemListItem is a type within the ServiceItem protocol
type ServiceItemListItem struct {
	types.Structure
	ItemCode            *types.String
	RegularPrice        *ServiceItemAmount
	TaxExcluded         *types.PrimitiveBool
	InitialPurchaseOnly *types.PrimitiveBool
	Limitation          *ServiceItemLimitation
	Attributes          *types.List[*ServiceItemAttribute]
}

// WriteTo writes the ServiceItemListItem to the given writable
func (sili *ServiceItemListItem) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sili.ItemCode.WriteTo(contentWritable)
	sili.RegularPrice.WriteTo(contentWritable)
	sili.TaxExcluded.WriteTo(contentWritable)
	sili.InitialPurchaseOnly.WriteTo(contentWritable)
	sili.Limitation.WriteTo(contentWritable)
	sili.Attributes.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sili.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemListItem from the given readable
func (sili *ServiceItemListItem) ExtractFrom(readable types.Readable) error {
	var err error

	err = sili.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem header. %s", err.Error())
	}

	err = sili.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.ItemCode. %s", err.Error())
	}

	err = sili.RegularPrice.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.RegularPrice. %s", err.Error())
	}

	err = sili.TaxExcluded.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.TaxExcluded. %s", err.Error())
	}

	err = sili.InitialPurchaseOnly.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.InitialPurchaseOnly. %s", err.Error())
	}

	err = sili.Limitation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.Limitation. %s", err.Error())
	}

	err = sili.Attributes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.Attributes. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemListItem
func (sili *ServiceItemListItem) Copy() types.RVType {
	copied := NewServiceItemListItem()

	copied.StructureVersion = sili.StructureVersion
	copied.ItemCode = sili.ItemCode.Copy().(*types.String)
	copied.RegularPrice = sili.RegularPrice.Copy().(*ServiceItemAmount)
	copied.TaxExcluded = sili.TaxExcluded.Copy().(*types.PrimitiveBool)
	copied.InitialPurchaseOnly = sili.InitialPurchaseOnly.Copy().(*types.PrimitiveBool)
	copied.Limitation = sili.Limitation.Copy().(*ServiceItemLimitation)
	copied.Attributes = sili.Attributes.Copy().(*types.List[*ServiceItemAttribute])

	return copied
}

// Equals checks if the given ServiceItemListItem contains the same data as the current ServiceItemListItem
func (sili *ServiceItemListItem) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemListItem); !ok {
		return false
	}

	other := o.(*ServiceItemListItem)

	if sili.StructureVersion != other.StructureVersion {
		return false
	}

	if !sili.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !sili.RegularPrice.Equals(other.RegularPrice) {
		return false
	}

	if !sili.TaxExcluded.Equals(other.TaxExcluded) {
		return false
	}

	if !sili.InitialPurchaseOnly.Equals(other.InitialPurchaseOnly) {
		return false
	}

	if !sili.Limitation.Equals(other.Limitation) {
		return false
	}

	return sili.Attributes.Equals(other.Attributes)
}

// String returns the string representation of the ServiceItemListItem
func (sili *ServiceItemListItem) String() string {
	return sili.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemListItem using the provided indentation level
func (sili *ServiceItemListItem) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemListItem{\n")
	b.WriteString(fmt.Sprintf("%sItemCode: %s,\n", indentationValues, sili.ItemCode))
	b.WriteString(fmt.Sprintf("%sRegularPrice: %s,\n", indentationValues, sili.RegularPrice.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sTaxExcluded: %s,\n", indentationValues, sili.TaxExcluded))
	b.WriteString(fmt.Sprintf("%sInitialPurchaseOnly: %s,\n", indentationValues, sili.InitialPurchaseOnly))
	b.WriteString(fmt.Sprintf("%sLimitation: %s,\n", indentationValues, sili.Limitation.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sAttributes: %s,\n", indentationValues, sili.Attributes))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemListItem returns a new ServiceItemListItem
func NewServiceItemListItem() *ServiceItemListItem {
	sili := &ServiceItemListItem{
		ItemCode:            types.NewString(""),
		RegularPrice:        NewServiceItemAmount(),
		TaxExcluded:         types.NewPrimitiveBool(false),
		InitialPurchaseOnly: types.NewPrimitiveBool(false),
		Limitation:          NewServiceItemLimitation(),
		Attributes:          types.NewList[*ServiceItemAttribute](),
	}

	sili.Attributes.Type = NewServiceItemAttribute()

	return sili
}
