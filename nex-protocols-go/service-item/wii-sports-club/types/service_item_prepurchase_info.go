// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemPrepurchaseInfo is a type within the ServiceItem protocol
type ServiceItemPrepurchaseInfo struct {
	types.Structure
	ItemCode       *types.String
	PriceID        *types.String
	RegularPrice   *ServiceItemAmount
	IsTaxAvailable *types.PrimitiveBool
	TaxAmount      *ServiceItemAmount
	TotalAmount    *ServiceItemAmount
	CurrentBalance *ServiceItemAmount
	PostBalance    *ServiceItemAmount
}

// WriteTo writes the ServiceItemPrepurchaseInfo to the given writable
func (sipi *ServiceItemPrepurchaseInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sipi.ItemCode.WriteTo(contentWritable)
	sipi.PriceID.WriteTo(contentWritable)
	sipi.RegularPrice.WriteTo(contentWritable)
	sipi.IsTaxAvailable.WriteTo(contentWritable)
	sipi.TaxAmount.WriteTo(contentWritable)
	sipi.TotalAmount.WriteTo(contentWritable)
	sipi.CurrentBalance.WriteTo(contentWritable)
	sipi.PostBalance.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sipi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemPrepurchaseInfo from the given readable
func (sipi *ServiceItemPrepurchaseInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = sipi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo header. %s", err.Error())
	}

	err = sipi.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.ItemCode. %s", err.Error())
	}

	err = sipi.PriceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.PriceID. %s", err.Error())
	}

	err = sipi.RegularPrice.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.RegularPrice. %s", err.Error())
	}

	err = sipi.IsTaxAvailable.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.IsTaxAvailable. %s", err.Error())
	}

	err = sipi.TaxAmount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.TaxAmount. %s", err.Error())
	}

	err = sipi.TotalAmount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.TotalAmount. %s", err.Error())
	}

	err = sipi.CurrentBalance.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.CurrentBalance. %s", err.Error())
	}

	err = sipi.PostBalance.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.PostBalance. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemPrepurchaseInfo
func (sipi *ServiceItemPrepurchaseInfo) Copy() types.RVType {
	copied := NewServiceItemPrepurchaseInfo()

	copied.StructureVersion = sipi.StructureVersion
	copied.ItemCode = sipi.ItemCode.Copy().(*types.String)
	copied.PriceID = sipi.PriceID.Copy().(*types.String)
	copied.RegularPrice = sipi.RegularPrice.Copy().(*ServiceItemAmount)
	copied.IsTaxAvailable = sipi.IsTaxAvailable.Copy().(*types.PrimitiveBool)
	copied.TaxAmount = sipi.TaxAmount.Copy().(*ServiceItemAmount)
	copied.TotalAmount = sipi.TotalAmount.Copy().(*ServiceItemAmount)
	copied.CurrentBalance = sipi.CurrentBalance.Copy().(*ServiceItemAmount)
	copied.PostBalance = sipi.PostBalance.Copy().(*ServiceItemAmount)

	return copied
}

// Equals checks if the given ServiceItemPrepurchaseInfo contains the same data as the current ServiceItemPrepurchaseInfo
func (sipi *ServiceItemPrepurchaseInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPrepurchaseInfo); !ok {
		return false
	}

	other := o.(*ServiceItemPrepurchaseInfo)

	if sipi.StructureVersion != other.StructureVersion {
		return false
	}

	if !sipi.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !sipi.PriceID.Equals(other.PriceID) {
		return false
	}

	if !sipi.RegularPrice.Equals(other.RegularPrice) {
		return false
	}

	if !sipi.IsTaxAvailable.Equals(other.IsTaxAvailable) {
		return false
	}

	if !sipi.TaxAmount.Equals(other.TaxAmount) {
		return false
	}

	if !sipi.TotalAmount.Equals(other.TotalAmount) {
		return false
	}

	if !sipi.CurrentBalance.Equals(other.CurrentBalance) {
		return false
	}

	return sipi.PostBalance.Equals(other.PostBalance)
}

// String returns the string representation of the ServiceItemPrepurchaseInfo
func (sipi *ServiceItemPrepurchaseInfo) String() string {
	return sipi.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemPrepurchaseInfo using the provided indentation level
func (sipi *ServiceItemPrepurchaseInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPrepurchaseInfo{\n")
	b.WriteString(fmt.Sprintf("%sItemCode: %s,\n", indentationValues, sipi.ItemCode))
	b.WriteString(fmt.Sprintf("%sPriceID: %s,\n", indentationValues, sipi.PriceID))
	b.WriteString(fmt.Sprintf("%sRegularPrice: %s,\n", indentationValues, sipi.RegularPrice.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sIsTaxAvailable: %s,\n", indentationValues, sipi.IsTaxAvailable))
	b.WriteString(fmt.Sprintf("%sTaxAmount: %s,\n", indentationValues, sipi.TaxAmount.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sTotalAmount: %s,\n", indentationValues, sipi.TotalAmount.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCurrentBalance: %s,\n", indentationValues, sipi.CurrentBalance.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPostBalance: %s,\n", indentationValues, sipi.PostBalance.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPrepurchaseInfo returns a new ServiceItemPrepurchaseInfo
func NewServiceItemPrepurchaseInfo() *ServiceItemPrepurchaseInfo {
	sipi := &ServiceItemPrepurchaseInfo{
		ItemCode:       types.NewString(""),
		PriceID:        types.NewString(""),
		RegularPrice:   NewServiceItemAmount(),
		IsTaxAvailable: types.NewPrimitiveBool(false),
		TaxAmount:      NewServiceItemAmount(),
		TotalAmount:    NewServiceItemAmount(),
		CurrentBalance: NewServiceItemAmount(),
		PostBalance:    NewServiceItemAmount(),
	}

	return sipi
}
