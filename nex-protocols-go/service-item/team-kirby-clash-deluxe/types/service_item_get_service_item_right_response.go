// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetServiceItemRightResponse is a type within the ServiceItem protocol
type ServiceItemGetServiceItemRightResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullableRightInfos *types.List[*ServiceItemRightInfos]
}

// WriteTo writes the ServiceItemGetServiceItemRightResponse to the given writable
func (sigsirr *ServiceItemGetServiceItemRightResponse) WriteTo(writable types.Writable) {
	sigsirr.ServiceItemEShopResponse.WriteTo(writable)

	contentWritable := writable.CopyNew()

	sigsirr.NullableRightInfos.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sigsirr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetServiceItemRightResponse from the given readable
func (sigsirr *ServiceItemGetServiceItemRightResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = sigsirr.ServiceItemEShopResponse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightResponse.ServiceItemEShopResponse. %s", err.Error())
	}

	err = sigsirr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightResponse header. %s", err.Error())
	}

	err = sigsirr.NullableRightInfos.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightResponse.NullableRightInfos. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetServiceItemRightResponse
func (sigsirr *ServiceItemGetServiceItemRightResponse) Copy() types.RVType {
	copied := NewServiceItemGetServiceItemRightResponse()

	copied.StructureVersion = sigsirr.StructureVersion
	copied.ServiceItemEShopResponse = sigsirr.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.NullableRightInfos = sigsirr.NullableRightInfos.Copy().(*types.List[*ServiceItemRightInfos])

	return copied
}

// Equals checks if the given ServiceItemGetServiceItemRightResponse contains the same data as the current ServiceItemGetServiceItemRightResponse
func (sigsirr *ServiceItemGetServiceItemRightResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetServiceItemRightResponse); !ok {
		return false
	}

	other := o.(*ServiceItemGetServiceItemRightResponse)

	if sigsirr.StructureVersion != other.StructureVersion {
		return false
	}

	if !sigsirr.ServiceItemEShopResponse.Equals(other.ServiceItemEShopResponse) {
		return false
	}

	return sigsirr.NullableRightInfos.Equals(other.NullableRightInfos)
}

// String returns the string representation of the ServiceItemGetServiceItemRightResponse
func (sigsirr *ServiceItemGetServiceItemRightResponse) String() string {
	return sigsirr.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetServiceItemRightResponse using the provided indentation level
func (sigsirr *ServiceItemGetServiceItemRightResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetServiceItemRightResponse{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemEShopResponse (parent): %s,\n", indentationValues, sigsirr.ServiceItemEShopResponse.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNullableRightInfos: %s,\n", indentationValues, sigsirr.NullableRightInfos))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetServiceItemRightResponse returns a new ServiceItemGetServiceItemRightResponse
func NewServiceItemGetServiceItemRightResponse() *ServiceItemGetServiceItemRightResponse {
	sigsirr := &ServiceItemGetServiceItemRightResponse{
		ServiceItemEShopResponse: NewServiceItemEShopResponse(),
		NullableRightInfos:       types.NewList[*ServiceItemRightInfos](),
	}

	sigsirr.NullableRightInfos.Type = NewServiceItemRightInfos()

	return sigsirr
}
