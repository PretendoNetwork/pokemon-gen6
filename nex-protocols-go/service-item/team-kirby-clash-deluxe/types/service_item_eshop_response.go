// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemEShopResponse is a type within the ServiceItem protocol
type ServiceItemEShopResponse struct {
	types.Structure
	HTTPStatus    *types.PrimitiveU32
	ErrorCode     *types.PrimitiveU32
	CorrelationID *types.String
}

// WriteTo writes the ServiceItemEShopResponse to the given writable
func (siesr *ServiceItemEShopResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siesr.HTTPStatus.WriteTo(contentWritable)
	siesr.ErrorCode.WriteTo(contentWritable)
	siesr.CorrelationID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siesr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemEShopResponse from the given readable
func (siesr *ServiceItemEShopResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = siesr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEShopResponse header. %s", err.Error())
	}

	err = siesr.HTTPStatus.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEShopResponse.HTTPStatus. %s", err.Error())
	}

	err = siesr.ErrorCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEShopResponse.ErrorCode. %s", err.Error())
	}

	err = siesr.CorrelationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEShopResponse.CorrelationID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemEShopResponse
func (siesr *ServiceItemEShopResponse) Copy() types.RVType {
	copied := NewServiceItemEShopResponse()

	copied.StructureVersion = siesr.StructureVersion
	copied.HTTPStatus = siesr.HTTPStatus.Copy().(*types.PrimitiveU32)
	copied.ErrorCode = siesr.ErrorCode.Copy().(*types.PrimitiveU32)
	copied.CorrelationID = siesr.CorrelationID.Copy().(*types.String)

	return copied
}

// Equals checks if the given ServiceItemEShopResponse contains the same data as the current ServiceItemEShopResponse
func (siesr *ServiceItemEShopResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemEShopResponse); !ok {
		return false
	}

	other := o.(*ServiceItemEShopResponse)

	if siesr.StructureVersion != other.StructureVersion {
		return false
	}

	if !siesr.HTTPStatus.Equals(other.HTTPStatus) {
		return false
	}

	if !siesr.ErrorCode.Equals(other.ErrorCode) {
		return false
	}

	return siesr.CorrelationID.Equals(other.CorrelationID)
}

// String returns the string representation of the ServiceItemEShopResponse
func (siesr *ServiceItemEShopResponse) String() string {
	return siesr.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemEShopResponse using the provided indentation level
func (siesr *ServiceItemEShopResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemEShopResponse{\n")
	b.WriteString(fmt.Sprintf("%sHTTPStatus: %s,\n", indentationValues, siesr.HTTPStatus))
	b.WriteString(fmt.Sprintf("%sErrorCode: %s,\n", indentationValues, siesr.ErrorCode))
	b.WriteString(fmt.Sprintf("%sCorrelationID: %s,\n", indentationValues, siesr.CorrelationID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemEShopResponse returns a new ServiceItemEShopResponse
func NewServiceItemEShopResponse() *ServiceItemEShopResponse {
	siesr := &ServiceItemEShopResponse{
		HTTPStatus:    types.NewPrimitiveU32(0),
		ErrorCode:     types.NewPrimitiveU32(0),
		CorrelationID: types.NewString(""),
	}

	return siesr
}
