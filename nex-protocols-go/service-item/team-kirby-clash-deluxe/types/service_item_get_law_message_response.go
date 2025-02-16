// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetLawMessageResponse is a type within the ServiceItem protocol
type ServiceItemGetLawMessageResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullableLawMessage *types.List[*ServiceItemLawMessage]
}

// WriteTo writes the ServiceItemGetLawMessageResponse to the given writable
func (siglmr *ServiceItemGetLawMessageResponse) WriteTo(writable types.Writable) {
	siglmr.ServiceItemEShopResponse.WriteTo(writable)

	contentWritable := writable.CopyNew()

	siglmr.NullableLawMessage.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siglmr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetLawMessageResponse from the given readable
func (siglmr *ServiceItemGetLawMessageResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = siglmr.ServiceItemEShopResponse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageResponse.ServiceItemEShopResponse. %s", err.Error())
	}

	err = siglmr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageResponse header. %s", err.Error())
	}

	err = siglmr.NullableLawMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageResponse.NullableLawMessage. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetLawMessageResponse
func (siglmr *ServiceItemGetLawMessageResponse) Copy() types.RVType {
	copied := NewServiceItemGetLawMessageResponse()

	copied.StructureVersion = siglmr.StructureVersion
	copied.ServiceItemEShopResponse = siglmr.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.NullableLawMessage = siglmr.NullableLawMessage.Copy().(*types.List[*ServiceItemLawMessage])

	return copied
}

// Equals checks if the given ServiceItemGetLawMessageResponse contains the same data as the current ServiceItemGetLawMessageResponse
func (siglmr *ServiceItemGetLawMessageResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetLawMessageResponse); !ok {
		return false
	}

	other := o.(*ServiceItemGetLawMessageResponse)

	if siglmr.StructureVersion != other.StructureVersion {
		return false
	}

	if !siglmr.ServiceItemEShopResponse.Equals(other.ServiceItemEShopResponse) {
		return false
	}

	return siglmr.NullableLawMessage.Equals(other.NullableLawMessage)
}

// String returns the string representation of the ServiceItemGetLawMessageResponse
func (siglmr *ServiceItemGetLawMessageResponse) String() string {
	return siglmr.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetLawMessageResponse using the provided indentation level
func (siglmr *ServiceItemGetLawMessageResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetLawMessageResponse{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemEShopResponse (parent): %s,\n", indentationValues, siglmr.ServiceItemEShopResponse.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNullableLawMessage: %s,\n", indentationValues, siglmr.NullableLawMessage))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetLawMessageResponse returns a new ServiceItemGetLawMessageResponse
func NewServiceItemGetLawMessageResponse() *ServiceItemGetLawMessageResponse {
	siglmr := &ServiceItemGetLawMessageResponse{
		ServiceItemEShopResponse: NewServiceItemEShopResponse(),
		NullableLawMessage:       types.NewList[*ServiceItemLawMessage](),
	}

	siglmr.NullableLawMessage.Type = NewServiceItemLawMessage()

	return siglmr
}
