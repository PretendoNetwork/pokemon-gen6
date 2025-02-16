// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemListServiceItemResponse is a type within the ServiceItem protocol
type ServiceItemListServiceItemResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullableCatalog *types.List[*ServiceItemCatalog]
}

// WriteTo writes the ServiceItemListServiceItemResponse to the given writable
func (silsir *ServiceItemListServiceItemResponse) WriteTo(writable types.Writable) {
	silsir.ServiceItemEShopResponse.WriteTo(writable)

	contentWritable := writable.CopyNew()

	silsir.NullableCatalog.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	silsir.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemListServiceItemResponse from the given readable
func (silsir *ServiceItemListServiceItemResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = silsir.ServiceItemEShopResponse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemResponse.ServiceItemEShopResponse. %s", err.Error())
	}

	err = silsir.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemResponse header. %s", err.Error())
	}

	err = silsir.NullableCatalog.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemResponse.NullableCatalog. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemListServiceItemResponse
func (silsir *ServiceItemListServiceItemResponse) Copy() types.RVType {
	copied := NewServiceItemListServiceItemResponse()

	copied.StructureVersion = silsir.StructureVersion
	copied.ServiceItemEShopResponse = silsir.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.NullableCatalog = silsir.NullableCatalog.Copy().(*types.List[*ServiceItemCatalog])

	return copied
}

// Equals checks if the given ServiceItemListServiceItemResponse contains the same data as the current ServiceItemListServiceItemResponse
func (silsir *ServiceItemListServiceItemResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemListServiceItemResponse); !ok {
		return false
	}

	other := o.(*ServiceItemListServiceItemResponse)

	if silsir.StructureVersion != other.StructureVersion {
		return false
	}

	if !silsir.ServiceItemEShopResponse.Equals(other.ServiceItemEShopResponse) {
		return false
	}

	return silsir.NullableCatalog.Equals(other.NullableCatalog)
}

// String returns the string representation of the ServiceItemListServiceItemResponse
func (silsir *ServiceItemListServiceItemResponse) String() string {
	return silsir.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemListServiceItemResponse using the provided indentation level
func (silsir *ServiceItemListServiceItemResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemListServiceItemResponse{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemEShopResponse (parent): %s,\n", indentationValues, silsir.ServiceItemEShopResponse.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNullableCatalog: %s,\n", indentationValues, silsir.NullableCatalog))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemListServiceItemResponse returns a new ServiceItemListServiceItemResponse
func NewServiceItemListServiceItemResponse() *ServiceItemListServiceItemResponse {
	silsir := &ServiceItemListServiceItemResponse{
		ServiceItemEShopResponse: NewServiceItemEShopResponse(),
		NullableCatalog:          types.NewList[*ServiceItemCatalog](),
	}

	silsir.NullableCatalog.Type = NewServiceItemCatalog()

	return silsir
}
