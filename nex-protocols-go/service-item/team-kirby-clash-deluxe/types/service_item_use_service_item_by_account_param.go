// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemUseServiceItemByAccountParam is a type within the ServiceItem protocol
type ServiceItemUseServiceItemByAccountParam struct {
	types.Structure
	ReferenceIDForUse         *types.String
	ReferenceIDForRightBinary *types.String
	UseType                   *types.PrimitiveU8
	UseNumber                 *types.PrimitiveU8
	RightBinary               *types.QBuffer
	LogMessage                *types.String
	UniqueID                  *types.PrimitiveU32
	Platform                  *types.PrimitiveU8
}

// WriteTo writes the ServiceItemUseServiceItemByAccountParam to the given writable
func (siusibap *ServiceItemUseServiceItemByAccountParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siusibap.ReferenceIDForUse.WriteTo(contentWritable)
	siusibap.ReferenceIDForRightBinary.WriteTo(contentWritable)
	siusibap.UseType.WriteTo(contentWritable)
	siusibap.UseNumber.WriteTo(contentWritable)
	siusibap.RightBinary.WriteTo(contentWritable)
	siusibap.LogMessage.WriteTo(contentWritable)
	siusibap.UniqueID.WriteTo(contentWritable)
	siusibap.Platform.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siusibap.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemUseServiceItemByAccountParam from the given readable
func (siusibap *ServiceItemUseServiceItemByAccountParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = siusibap.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam header. %s", err.Error())
	}

	err = siusibap.ReferenceIDForUse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.ReferenceIDForUse. %s", err.Error())
	}

	err = siusibap.ReferenceIDForRightBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.ReferenceIDForRightBinary. %s", err.Error())
	}

	err = siusibap.UseType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.UseType. %s", err.Error())
	}

	err = siusibap.UseNumber.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.UseNumber. %s", err.Error())
	}

	err = siusibap.RightBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.RightBinary. %s", err.Error())
	}

	err = siusibap.LogMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.LogMessage. %s", err.Error())
	}

	err = siusibap.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.UniqueID. %s", err.Error())
	}

	err = siusibap.Platform.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemByAccountParam.Platform. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemUseServiceItemByAccountParam
func (siusibap *ServiceItemUseServiceItemByAccountParam) Copy() types.RVType {
	copied := NewServiceItemUseServiceItemByAccountParam()

	copied.StructureVersion = siusibap.StructureVersion
	copied.ReferenceIDForUse = siusibap.ReferenceIDForUse.Copy().(*types.String)
	copied.ReferenceIDForRightBinary = siusibap.ReferenceIDForRightBinary.Copy().(*types.String)
	copied.UseType = siusibap.UseType.Copy().(*types.PrimitiveU8)
	copied.UseNumber = siusibap.UseNumber.Copy().(*types.PrimitiveU8)
	copied.RightBinary = siusibap.RightBinary.Copy().(*types.QBuffer)
	copied.LogMessage = siusibap.LogMessage.Copy().(*types.String)
	copied.UniqueID = siusibap.UniqueID.Copy().(*types.PrimitiveU32)
	copied.Platform = siusibap.Platform.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given ServiceItemUseServiceItemByAccountParam contains the same data as the current ServiceItemUseServiceItemByAccountParam
func (siusibap *ServiceItemUseServiceItemByAccountParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemUseServiceItemByAccountParam); !ok {
		return false
	}

	other := o.(*ServiceItemUseServiceItemByAccountParam)

	if siusibap.StructureVersion != other.StructureVersion {
		return false
	}

	if !siusibap.ReferenceIDForUse.Equals(other.ReferenceIDForUse) {
		return false
	}

	if !siusibap.ReferenceIDForRightBinary.Equals(other.ReferenceIDForRightBinary) {
		return false
	}

	if !siusibap.UseType.Equals(other.UseType) {
		return false
	}

	if !siusibap.UseNumber.Equals(other.UseNumber) {
		return false
	}

	if !siusibap.RightBinary.Equals(other.RightBinary) {
		return false
	}

	if !siusibap.LogMessage.Equals(other.LogMessage) {
		return false
	}

	if !siusibap.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return siusibap.Platform.Equals(other.Platform)
}

// String returns the string representation of the ServiceItemUseServiceItemByAccountParam
func (siusibap *ServiceItemUseServiceItemByAccountParam) String() string {
	return siusibap.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemUseServiceItemByAccountParam using the provided indentation level
func (siusibap *ServiceItemUseServiceItemByAccountParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemUseServiceItemByAccountParam{\n")
	b.WriteString(fmt.Sprintf("%sReferenceIDForUse: %s,\n", indentationValues, siusibap.ReferenceIDForUse))
	b.WriteString(fmt.Sprintf("%sReferenceIDForRightBinary: %s,\n", indentationValues, siusibap.ReferenceIDForRightBinary))
	b.WriteString(fmt.Sprintf("%sUseType: %s,\n", indentationValues, siusibap.UseType))
	b.WriteString(fmt.Sprintf("%sUseNumber: %s,\n", indentationValues, siusibap.UseNumber))
	b.WriteString(fmt.Sprintf("%sRightBinary: %s,\n", indentationValues, siusibap.RightBinary))
	b.WriteString(fmt.Sprintf("%sLogMessage: %s,\n", indentationValues, siusibap.LogMessage))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, siusibap.UniqueID))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, siusibap.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemUseServiceItemByAccountParam returns a new ServiceItemUseServiceItemByAccountParam
func NewServiceItemUseServiceItemByAccountParam() *ServiceItemUseServiceItemByAccountParam {
	siusibap := &ServiceItemUseServiceItemByAccountParam{
		ReferenceIDForUse:         types.NewString(""),
		ReferenceIDForRightBinary: types.NewString(""),
		UseType:                   types.NewPrimitiveU8(0),
		UseNumber:                 types.NewPrimitiveU8(0),
		RightBinary:               types.NewQBuffer(nil),
		LogMessage:                types.NewString(""),
		UniqueID:                  types.NewPrimitiveU32(0),
		Platform:                  types.NewPrimitiveU8(0),
	}

	return siusibap
}
