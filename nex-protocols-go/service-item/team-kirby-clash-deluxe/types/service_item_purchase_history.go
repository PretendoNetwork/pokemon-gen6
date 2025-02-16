// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemPurchaseHistory is a type within the ServiceItem protocol
type ServiceItemPurchaseHistory struct {
	types.Structure
	TotalSize    *types.PrimitiveU32
	Offset       *types.PrimitiveU32
	Transactions *types.List[*ServiceItemTransaction]
}

// WriteTo writes the ServiceItemPurchaseHistory to the given writable
func (siph *ServiceItemPurchaseHistory) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siph.TotalSize.WriteTo(contentWritable)
	siph.Offset.WriteTo(contentWritable)
	siph.Transactions.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siph.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemPurchaseHistory from the given readable
func (siph *ServiceItemPurchaseHistory) ExtractFrom(readable types.Readable) error {
	var err error

	err = siph.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseHistory header. %s", err.Error())
	}

	err = siph.TotalSize.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseHistory.TotalSize. %s", err.Error())
	}

	err = siph.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseHistory.Offset. %s", err.Error())
	}

	err = siph.Transactions.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseHistory.Transactions. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemPurchaseHistory
func (siph *ServiceItemPurchaseHistory) Copy() types.RVType {
	copied := NewServiceItemPurchaseHistory()

	copied.StructureVersion = siph.StructureVersion
	copied.TotalSize = siph.TotalSize.Copy().(*types.PrimitiveU32)
	copied.Offset = siph.Offset.Copy().(*types.PrimitiveU32)
	copied.Transactions = siph.Transactions.Copy().(*types.List[*ServiceItemTransaction])

	return copied
}

// Equals checks if the given ServiceItemPurchaseHistory contains the same data as the current ServiceItemPurchaseHistory
func (siph *ServiceItemPurchaseHistory) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPurchaseHistory); !ok {
		return false
	}

	other := o.(*ServiceItemPurchaseHistory)

	if siph.StructureVersion != other.StructureVersion {
		return false
	}

	if !siph.TotalSize.Equals(other.TotalSize) {
		return false
	}

	if !siph.Offset.Equals(other.Offset) {
		return false
	}

	return siph.Transactions.Equals(other.Transactions)
}

// String returns the string representation of the ServiceItemPurchaseHistory
func (siph *ServiceItemPurchaseHistory) String() string {
	return siph.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemPurchaseHistory using the provided indentation level
func (siph *ServiceItemPurchaseHistory) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseHistory{\n")
	b.WriteString(fmt.Sprintf("%sTotalSize: %s,\n", indentationValues, siph.TotalSize))
	b.WriteString(fmt.Sprintf("%sOffset: %s,\n", indentationValues, siph.Offset))
	b.WriteString(fmt.Sprintf("%sTransactions: %s,\n", indentationValues, siph.Transactions))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPurchaseHistory returns a new ServiceItemPurchaseHistory
func NewServiceItemPurchaseHistory() *ServiceItemPurchaseHistory {
	siph := &ServiceItemPurchaseHistory{
		TotalSize:    types.NewPrimitiveU32(0),
		Offset:       types.NewPrimitiveU32(0),
		Transactions: types.NewList[*ServiceItemTransaction](),
	}

	siph.Transactions.Type = NewServiceItemTransaction()

	return siph
}
