// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetBalanceResponse is a type within the ServiceItem protocol
type ServiceItemGetBalanceResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullableBalance *types.List[*ServiceItemAmount]
}

// WriteTo writes the ServiceItemGetBalanceResponse to the given writable
func (sigbr *ServiceItemGetBalanceResponse) WriteTo(writable types.Writable) {
	sigbr.ServiceItemEShopResponse.WriteTo(writable)

	contentWritable := writable.CopyNew()

	sigbr.NullableBalance.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sigbr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetBalanceResponse from the given readable
func (sigbr *ServiceItemGetBalanceResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = sigbr.ServiceItemEShopResponse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceResponse.ServiceItemEShopResponse. %s", err.Error())
	}

	err = sigbr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceResponse header. %s", err.Error())
	}

	err = sigbr.NullableBalance.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceResponse.NullableBalance. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetBalanceResponse
func (sigbr *ServiceItemGetBalanceResponse) Copy() types.RVType {
	copied := NewServiceItemGetBalanceResponse()

	copied.StructureVersion = sigbr.StructureVersion
	copied.ServiceItemEShopResponse = sigbr.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.NullableBalance = sigbr.NullableBalance.Copy().(*types.List[*ServiceItemAmount])

	return copied
}

// Equals checks if the given ServiceItemGetBalanceResponse contains the same data as the current ServiceItemGetBalanceResponse
func (sigbr *ServiceItemGetBalanceResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetBalanceResponse); !ok {
		return false
	}

	other := o.(*ServiceItemGetBalanceResponse)

	if sigbr.StructureVersion != other.StructureVersion {
		return false
	}

	if !sigbr.ServiceItemEShopResponse.Equals(other.ServiceItemEShopResponse) {
		return false
	}

	return sigbr.NullableBalance.Equals(other.NullableBalance)
}

// String returns the string representation of the ServiceItemGetBalanceResponse
func (sigbr *ServiceItemGetBalanceResponse) String() string {
	return sigbr.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetBalanceResponse using the provided indentation level
func (sigbr *ServiceItemGetBalanceResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetBalanceResponse{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemEShopResponse (parent): %s,\n", indentationValues, sigbr.ServiceItemEShopResponse.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNullableBalance: %s,\n", indentationValues, sigbr.NullableBalance))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetBalanceResponse returns a new ServiceItemGetBalanceResponse
func NewServiceItemGetBalanceResponse() *ServiceItemGetBalanceResponse {
	sigbr := &ServiceItemGetBalanceResponse{
		ServiceItemEShopResponse: NewServiceItemEShopResponse(),
		NullableBalance:          types.NewList[*ServiceItemAmount](),
	}

	sigbr.NullableBalance.Type = NewServiceItemAmount()

	return sigbr
}
