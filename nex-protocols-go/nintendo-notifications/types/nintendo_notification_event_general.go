// Package types implements all the types used by the NintendoNotifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NintendoNotificationEventGeneral is a type within the NintendoNotifications protocol
type NintendoNotificationEventGeneral struct {
	types.Structure
	U32Param  *types.PrimitiveU32
	U64Param1 *types.PrimitiveU64
	U64Param2 *types.PrimitiveU64
	StrParam  *types.String
}

// WriteTo writes the NintendoNotificationEventGeneral to the given writable
func (nneg *NintendoNotificationEventGeneral) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nneg.U32Param.WriteTo(contentWritable)
	nneg.U64Param1.WriteTo(contentWritable)
	nneg.U64Param2.WriteTo(contentWritable)
	nneg.StrParam.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nneg.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoNotificationEventGeneral from the given readable
func (nneg *NintendoNotificationEventGeneral) ExtractFrom(readable types.Readable) error {
	var err error

	err = nneg.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral header. %s", err.Error())
	}

	err = nneg.U32Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral.U32Param. %s", err.Error())
	}

	err = nneg.U64Param1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral.U64Param1. %s", err.Error())
	}

	err = nneg.U64Param2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral.U64Param2. %s", err.Error())
	}

	err = nneg.StrParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral.StrParam. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoNotificationEventGeneral
func (nneg *NintendoNotificationEventGeneral) Copy() types.RVType {
	copied := NewNintendoNotificationEventGeneral()

	copied.StructureVersion = nneg.StructureVersion
	copied.U32Param = nneg.U32Param.Copy().(*types.PrimitiveU32)
	copied.U64Param1 = nneg.U64Param1.Copy().(*types.PrimitiveU64)
	copied.U64Param2 = nneg.U64Param2.Copy().(*types.PrimitiveU64)
	copied.StrParam = nneg.StrParam.Copy().(*types.String)

	return copied
}

// Equals checks if the given NintendoNotificationEventGeneral contains the same data as the current NintendoNotificationEventGeneral
func (nneg *NintendoNotificationEventGeneral) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoNotificationEventGeneral); !ok {
		return false
	}

	other := o.(*NintendoNotificationEventGeneral)

	if nneg.StructureVersion != other.StructureVersion {
		return false
	}

	if !nneg.U32Param.Equals(other.U32Param) {
		return false
	}

	if !nneg.U64Param1.Equals(other.U64Param1) {
		return false
	}

	if !nneg.U64Param2.Equals(other.U64Param2) {
		return false
	}

	return nneg.StrParam.Equals(other.StrParam)
}

// String returns the string representation of the NintendoNotificationEventGeneral
func (nneg *NintendoNotificationEventGeneral) String() string {
	return nneg.FormatToString(0)
}

// FormatToString pretty-prints the NintendoNotificationEventGeneral using the provided indentation level
func (nneg *NintendoNotificationEventGeneral) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoNotificationEventGeneral{\n")
	b.WriteString(fmt.Sprintf("%sU32Param: %s,\n", indentationValues, nneg.U32Param))
	b.WriteString(fmt.Sprintf("%sU64Param1: %s,\n", indentationValues, nneg.U64Param1))
	b.WriteString(fmt.Sprintf("%sU64Param2: %s,\n", indentationValues, nneg.U64Param2))
	b.WriteString(fmt.Sprintf("%sStrParam: %s,\n", indentationValues, nneg.StrParam))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoNotificationEventGeneral returns a new NintendoNotificationEventGeneral
func NewNintendoNotificationEventGeneral() *NintendoNotificationEventGeneral {
	nneg := &NintendoNotificationEventGeneral{
		U32Param:  types.NewPrimitiveU32(0),
		U64Param1: types.NewPrimitiveU64(0),
		U64Param2: types.NewPrimitiveU64(0),
		StrParam:  types.NewString(""),
	}

	return nneg
}
