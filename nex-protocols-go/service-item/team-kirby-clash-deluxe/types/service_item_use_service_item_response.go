// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemUseServiceItemResponse is a type within the ServiceItem protocol
type ServiceItemUseServiceItemResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullableUsedInfo *types.List[*ServiceItemUsedInfo]
}

// WriteTo writes the ServiceItemUseServiceItemResponse to the given writable
func (siusir *ServiceItemUseServiceItemResponse) WriteTo(writable types.Writable) {
	siusir.ServiceItemEShopResponse.WriteTo(writable)

	contentWritable := writable.CopyNew()

	siusir.NullableUsedInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siusir.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemUseServiceItemResponse from the given readable
func (siusir *ServiceItemUseServiceItemResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = siusir.ServiceItemEShopResponse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemResponse.ServiceItemEShopResponse. %s", err.Error())
	}

	err = siusir.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemResponse header. %s", err.Error())
	}

	err = siusir.NullableUsedInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemResponse.NullableUsedInfo. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemUseServiceItemResponse
func (siusir *ServiceItemUseServiceItemResponse) Copy() types.RVType {
	copied := NewServiceItemUseServiceItemResponse()

	copied.StructureVersion = siusir.StructureVersion
	copied.ServiceItemEShopResponse = siusir.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.NullableUsedInfo = siusir.NullableUsedInfo.Copy().(*types.List[*ServiceItemUsedInfo])

	return copied
}

// Equals checks if the given ServiceItemUseServiceItemResponse contains the same data as the current ServiceItemUseServiceItemResponse
func (siusir *ServiceItemUseServiceItemResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemUseServiceItemResponse); !ok {
		return false
	}

	other := o.(*ServiceItemUseServiceItemResponse)

	if siusir.StructureVersion != other.StructureVersion {
		return false
	}

	if !siusir.ServiceItemEShopResponse.Equals(other.ServiceItemEShopResponse) {
		return false
	}

	return siusir.NullableUsedInfo.Equals(other.NullableUsedInfo)
}

// String returns the string representation of the ServiceItemUseServiceItemResponse
func (siusir *ServiceItemUseServiceItemResponse) String() string {
	return siusir.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemUseServiceItemResponse using the provided indentation level
func (siusir *ServiceItemUseServiceItemResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemUseServiceItemResponse{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemEShopResponse (parent): %s,\n", indentationValues, siusir.ServiceItemEShopResponse.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNullableUsedInfo: %s,\n", indentationValues, siusir.NullableUsedInfo))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemUseServiceItemResponse returns a new ServiceItemUseServiceItemResponse
func NewServiceItemUseServiceItemResponse() *ServiceItemUseServiceItemResponse {
	siusir := &ServiceItemUseServiceItemResponse{
		ServiceItemEShopResponse: NewServiceItemEShopResponse(),
		NullableUsedInfo:         types.NewList[*ServiceItemUsedInfo](),
	}

	siusir.NullableUsedInfo.Type = NewServiceItemUsedInfo()

	return siusir
}
