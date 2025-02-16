// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetPurchaseHistoryResponse is a type within the ServiceItem protocol
type ServiceItemGetPurchaseHistoryResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullablePurchaseHistory *types.List[*ServiceItemPurchaseHistory]
}

// WriteTo writes the ServiceItemGetPurchaseHistoryResponse to the given writable
func (sigphr *ServiceItemGetPurchaseHistoryResponse) WriteTo(writable types.Writable) {
	sigphr.ServiceItemEShopResponse.WriteTo(writable)

	contentWritable := writable.CopyNew()

	sigphr.NullablePurchaseHistory.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sigphr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetPurchaseHistoryResponse from the given readable
func (sigphr *ServiceItemGetPurchaseHistoryResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = sigphr.ServiceItemEShopResponse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryResponse.ServiceItemEShopResponse. %s", err.Error())
	}

	err = sigphr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryResponse header. %s", err.Error())
	}

	err = sigphr.NullablePurchaseHistory.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryResponse.NullablePurchaseHistory. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetPurchaseHistoryResponse
func (sigphr *ServiceItemGetPurchaseHistoryResponse) Copy() types.RVType {
	copied := NewServiceItemGetPurchaseHistoryResponse()

	copied.StructureVersion = sigphr.StructureVersion
	copied.ServiceItemEShopResponse = sigphr.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.NullablePurchaseHistory = sigphr.NullablePurchaseHistory.Copy().(*types.List[*ServiceItemPurchaseHistory])

	return copied
}

// Equals checks if the given ServiceItemGetPurchaseHistoryResponse contains the same data as the current ServiceItemGetPurchaseHistoryResponse
func (sigphr *ServiceItemGetPurchaseHistoryResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetPurchaseHistoryResponse); !ok {
		return false
	}

	other := o.(*ServiceItemGetPurchaseHistoryResponse)

	if sigphr.StructureVersion != other.StructureVersion {
		return false
	}

	if !sigphr.ServiceItemEShopResponse.Equals(other.ServiceItemEShopResponse) {
		return false
	}

	return sigphr.NullablePurchaseHistory.Equals(other.NullablePurchaseHistory)
}

// String returns the string representation of the ServiceItemGetPurchaseHistoryResponse
func (sigphr *ServiceItemGetPurchaseHistoryResponse) String() string {
	return sigphr.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetPurchaseHistoryResponse using the provided indentation level
func (sigphr *ServiceItemGetPurchaseHistoryResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetPurchaseHistoryResponse{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemEShopResponse (parent): %s,\n", indentationValues, sigphr.ServiceItemEShopResponse.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNullablePurchaseHistory: %s,\n", indentationValues, sigphr.NullablePurchaseHistory))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetPurchaseHistoryResponse returns a new ServiceItemGetPurchaseHistoryResponse
func NewServiceItemGetPurchaseHistoryResponse() *ServiceItemGetPurchaseHistoryResponse {
	sigphr := &ServiceItemGetPurchaseHistoryResponse{
		ServiceItemEShopResponse: NewServiceItemEShopResponse(),
		NullablePurchaseHistory:  types.NewList[*ServiceItemPurchaseHistory](),
	}

	sigphr.NullablePurchaseHistory.Type = NewServiceItemPurchaseHistory()

	return sigphr
}
