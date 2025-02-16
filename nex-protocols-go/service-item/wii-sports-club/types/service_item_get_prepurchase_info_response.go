// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetPrepurchaseInfoResponse is a type within the ServiceItem protocol
type ServiceItemGetPrepurchaseInfoResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullablePrepurchaseInfo *types.List[*ServiceItemPrepurchaseInfo]
}

// WriteTo writes the ServiceItemGetPrepurchaseInfoResponse to the given writable
func (sigpir *ServiceItemGetPrepurchaseInfoResponse) WriteTo(writable types.Writable) {
	sigpir.ServiceItemEShopResponse.WriteTo(writable)

	contentWritable := writable.CopyNew()

	sigpir.NullablePrepurchaseInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sigpir.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetPrepurchaseInfoResponse from the given readable
func (sigpir *ServiceItemGetPrepurchaseInfoResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = sigpir.ServiceItemEShopResponse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoResponse.ServiceItemEShopResponse. %s", err.Error())
	}

	err = sigpir.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoResponse header. %s", err.Error())
	}

	err = sigpir.NullablePrepurchaseInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetPrepurchaseInfoResponse
func (sigpir *ServiceItemGetPrepurchaseInfoResponse) Copy() types.RVType {
	copied := NewServiceItemGetPrepurchaseInfoResponse()

	copied.StructureVersion = sigpir.StructureVersion
	copied.ServiceItemEShopResponse = sigpir.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.NullablePrepurchaseInfo = sigpir.NullablePrepurchaseInfo.Copy().(*types.List[*ServiceItemPrepurchaseInfo])

	return copied
}

// Equals checks if the given ServiceItemGetPrepurchaseInfoResponse contains the same data as the current ServiceItemGetPrepurchaseInfoResponse
func (sigpir *ServiceItemGetPrepurchaseInfoResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetPrepurchaseInfoResponse); !ok {
		return false
	}

	other := o.(*ServiceItemGetPrepurchaseInfoResponse)

	if sigpir.StructureVersion != other.StructureVersion {
		return false
	}

	if !sigpir.ServiceItemEShopResponse.Equals(other.ServiceItemEShopResponse) {
		return false
	}

	return sigpir.NullablePrepurchaseInfo.Equals(other.NullablePrepurchaseInfo)
}

// String returns the string representation of the ServiceItemGetPrepurchaseInfoResponse
func (sigpir *ServiceItemGetPrepurchaseInfoResponse) String() string {
	return sigpir.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetPrepurchaseInfoResponse using the provided indentation level
func (sigpir *ServiceItemGetPrepurchaseInfoResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetPrepurchaseInfoResponse{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemEShopResponse (parent): %s,\n", indentationValues, sigpir.ServiceItemEShopResponse.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNullablePrepurchaseInfo: %s,\n", indentationValues, sigpir.NullablePrepurchaseInfo))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetPrepurchaseInfoResponse returns a new ServiceItemGetPrepurchaseInfoResponse
func NewServiceItemGetPrepurchaseInfoResponse() *ServiceItemGetPrepurchaseInfoResponse {
	sigpir := &ServiceItemGetPrepurchaseInfoResponse{
		ServiceItemEShopResponse: NewServiceItemEShopResponse(),
		NullablePrepurchaseInfo:  types.NewList[*ServiceItemPrepurchaseInfo](),
	}

	sigpir.NullablePrepurchaseInfo.Type = NewServiceItemPrepurchaseInfo()

	return sigpir
}
