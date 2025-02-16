// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemCatalog is a type within the ServiceItem protocol
type ServiceItemCatalog struct {
	types.Structure
	TotalSize *types.PrimitiveU32
	Offset    *types.PrimitiveU32
	ListItems *types.List[*ServiceItemListItem]
}

// WriteTo writes the ServiceItemCatalog to the given writable
func (sic *ServiceItemCatalog) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sic.TotalSize.WriteTo(contentWritable)
	sic.Offset.WriteTo(contentWritable)
	sic.ListItems.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sic.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemCatalog from the given readable
func (sic *ServiceItemCatalog) ExtractFrom(readable types.Readable) error {
	var err error

	err = sic.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog header. %s", err.Error())
	}

	err = sic.TotalSize.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.TotalSize. %s", err.Error())
	}

	err = sic.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.Offset. %s", err.Error())
	}

	err = sic.ListItems.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.ListItems. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemCatalog
func (sic *ServiceItemCatalog) Copy() types.RVType {
	copied := NewServiceItemCatalog()

	copied.StructureVersion = sic.StructureVersion
	copied.TotalSize = sic.TotalSize.Copy().(*types.PrimitiveU32)
	copied.Offset = sic.Offset.Copy().(*types.PrimitiveU32)
	copied.ListItems = sic.ListItems.Copy().(*types.List[*ServiceItemListItem])

	return copied
}

// Equals checks if the given ServiceItemCatalog contains the same data as the current ServiceItemCatalog
func (sic *ServiceItemCatalog) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemCatalog); !ok {
		return false
	}

	other := o.(*ServiceItemCatalog)

	if sic.StructureVersion != other.StructureVersion {
		return false
	}

	if !sic.TotalSize.Equals(other.TotalSize) {
		return false
	}

	if !sic.Offset.Equals(other.Offset) {
		return false
	}

	return sic.ListItems.Equals(other.ListItems)
}

// String returns the string representation of the ServiceItemCatalog
func (sic *ServiceItemCatalog) String() string {
	return sic.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemCatalog using the provided indentation level
func (sic *ServiceItemCatalog) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemCatalog{\n")
	b.WriteString(fmt.Sprintf("%sTotalSize: %s,\n", indentationValues, sic.TotalSize))
	b.WriteString(fmt.Sprintf("%sOffset: %s,\n", indentationValues, sic.Offset))
	b.WriteString(fmt.Sprintf("%sListItems: %s,\n", indentationValues, sic.ListItems))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemCatalog returns a new ServiceItemCatalog
func NewServiceItemCatalog() *ServiceItemCatalog {
	sic := &ServiceItemCatalog{
		TotalSize: types.NewPrimitiveU32(0),
		Offset:    types.NewPrimitiveU32(0),
		ListItems: types.NewList[*ServiceItemListItem](),
	}

	sic.ListItems.Type = NewServiceItemListItem()

	return sic
}
