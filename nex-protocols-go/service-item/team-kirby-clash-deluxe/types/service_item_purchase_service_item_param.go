// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemPurchaseServiceItemParam is a type within the ServiceItem protocol
type ServiceItemPurchaseServiceItemParam struct {
	types.Structure
	ItemCode       *types.String
	PriceID        *types.String
	ReferenceID    *types.String
	Balance        *types.String
	ItemName       *types.String
	EcServiceToken *types.String
	Language       *types.String
	UniqueID       *types.PrimitiveU32
	Platform       *types.PrimitiveU8 // * Revision 1
}

// WriteTo writes the ServiceItemPurchaseServiceItemParam to the given writable
func (sipsip *ServiceItemPurchaseServiceItemParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sipsip.ItemCode.WriteTo(contentWritable)
	sipsip.PriceID.WriteTo(contentWritable)
	sipsip.ReferenceID.WriteTo(contentWritable)
	sipsip.Balance.WriteTo(contentWritable)
	sipsip.ItemName.WriteTo(contentWritable)
	sipsip.EcServiceToken.WriteTo(contentWritable)
	sipsip.Language.WriteTo(contentWritable)
	sipsip.UniqueID.WriteTo(contentWritable)

	if sipsip.StructureVersion >= 1 {
		sipsip.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	sipsip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemPurchaseServiceItemParam from the given readable
func (sipsip *ServiceItemPurchaseServiceItemParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sipsip.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam header. %s", err.Error())
	}

	err = sipsip.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.ItemCode. %s", err.Error())
	}

	err = sipsip.PriceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.PriceID. %s", err.Error())
	}

	err = sipsip.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.ReferenceID. %s", err.Error())
	}

	err = sipsip.Balance.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.Balance. %s", err.Error())
	}

	err = sipsip.ItemName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.ItemName. %s", err.Error())
	}

	err = sipsip.EcServiceToken.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.EcServiceToken. %s", err.Error())
	}

	err = sipsip.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.Language. %s", err.Error())
	}

	err = sipsip.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.UniqueID. %s", err.Error())
	}

	if sipsip.StructureVersion >= 1 {
		err = sipsip.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.Platform. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemPurchaseServiceItemParam
func (sipsip *ServiceItemPurchaseServiceItemParam) Copy() types.RVType {
	copied := NewServiceItemPurchaseServiceItemParam()

	copied.StructureVersion = sipsip.StructureVersion
	copied.ItemCode = sipsip.ItemCode.Copy().(*types.String)
	copied.PriceID = sipsip.PriceID.Copy().(*types.String)
	copied.ReferenceID = sipsip.ReferenceID.Copy().(*types.String)
	copied.Balance = sipsip.Balance.Copy().(*types.String)
	copied.ItemName = sipsip.ItemName.Copy().(*types.String)
	copied.EcServiceToken = sipsip.EcServiceToken.Copy().(*types.String)
	copied.Language = sipsip.Language.Copy().(*types.String)
	copied.UniqueID = sipsip.UniqueID.Copy().(*types.PrimitiveU32)
	copied.Platform = sipsip.Platform.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given ServiceItemPurchaseServiceItemParam contains the same data as the current ServiceItemPurchaseServiceItemParam
func (sipsip *ServiceItemPurchaseServiceItemParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPurchaseServiceItemParam); !ok {
		return false
	}

	other := o.(*ServiceItemPurchaseServiceItemParam)

	if sipsip.StructureVersion != other.StructureVersion {
		return false
	}

	if !sipsip.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !sipsip.PriceID.Equals(other.PriceID) {
		return false
	}

	if !sipsip.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !sipsip.Balance.Equals(other.Balance) {
		return false
	}

	if !sipsip.ItemName.Equals(other.ItemName) {
		return false
	}

	if !sipsip.EcServiceToken.Equals(other.EcServiceToken) {
		return false
	}

	if !sipsip.Language.Equals(other.Language) {
		return false
	}

	if !sipsip.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return sipsip.Platform.Equals(other.Platform)
}

// String returns the string representation of the ServiceItemPurchaseServiceItemParam
func (sipsip *ServiceItemPurchaseServiceItemParam) String() string {
	return sipsip.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemPurchaseServiceItemParam using the provided indentation level
func (sipsip *ServiceItemPurchaseServiceItemParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseServiceItemParam{\n")
	b.WriteString(fmt.Sprintf("%sItemCode: %s,\n", indentationValues, sipsip.ItemCode))
	b.WriteString(fmt.Sprintf("%sPriceID: %s,\n", indentationValues, sipsip.PriceID))
	b.WriteString(fmt.Sprintf("%sReferenceID: %s,\n", indentationValues, sipsip.ReferenceID))
	b.WriteString(fmt.Sprintf("%sBalance: %s,\n", indentationValues, sipsip.Balance))
	b.WriteString(fmt.Sprintf("%sItemName: %s,\n", indentationValues, sipsip.ItemName))
	b.WriteString(fmt.Sprintf("%sEcServiceToken: %s,\n", indentationValues, sipsip.EcServiceToken))
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, sipsip.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, sipsip.UniqueID))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, sipsip.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPurchaseServiceItemParam returns a new ServiceItemPurchaseServiceItemParam
func NewServiceItemPurchaseServiceItemParam() *ServiceItemPurchaseServiceItemParam {
	sipsip := &ServiceItemPurchaseServiceItemParam{
		ItemCode:       types.NewString(""),
		PriceID:        types.NewString(""),
		ReferenceID:    types.NewString(""),
		Balance:        types.NewString(""),
		ItemName:       types.NewString(""),
		EcServiceToken: types.NewString(""),
		Language:       types.NewString(""),
		UniqueID:       types.NewPrimitiveU32(0),
		Platform:       types.NewPrimitiveU8(0),
	}

	return sipsip
}
