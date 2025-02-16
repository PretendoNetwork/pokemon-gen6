// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SubscriberGetContentParam is a type within the Shop protocol
type SubscriberGetContentParam struct {
	types.Structure
	Unknown1 *types.String
	Unknown2 *types.PrimitiveU32
	Unknown3 *types.PrimitiveU32
	Unknown4 *types.PrimitiveU64
}

// WriteTo writes the SubscriberGetContentParam to the given writable
func (sgcp *SubscriberGetContentParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sgcp.Unknown1.WriteTo(contentWritable)
	sgcp.Unknown2.WriteTo(contentWritable)
	sgcp.Unknown3.WriteTo(contentWritable)
	sgcp.Unknown4.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sgcp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SubscriberGetContentParam from the given readable
func (sgcp *SubscriberGetContentParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sgcp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam header. %s", err.Error())
	}

	err = sgcp.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown1. %s", err.Error())
	}

	err = sgcp.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown2. %s", err.Error())
	}

	err = sgcp.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown3. %s", err.Error())
	}

	err = sgcp.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown4. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SubscriberGetContentParam
func (sgcp *SubscriberGetContentParam) Copy() types.RVType {
	copied := NewSubscriberGetContentParam()

	copied.StructureVersion = sgcp.StructureVersion
	copied.Unknown1 = sgcp.Unknown1.Copy().(*types.String)
	copied.Unknown2 = sgcp.Unknown2.Copy().(*types.PrimitiveU32)
	copied.Unknown3 = sgcp.Unknown3.Copy().(*types.PrimitiveU32)
	copied.Unknown4 = sgcp.Unknown4.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given SubscriberGetContentParam contains the same data as the current SubscriberGetContentParam
func (sgcp *SubscriberGetContentParam) Equals(o types.RVType) bool {
	if _, ok := o.(*SubscriberGetContentParam); !ok {
		return false
	}

	other := o.(*SubscriberGetContentParam)

	if sgcp.StructureVersion != other.StructureVersion {
		return false
	}

	if !sgcp.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !sgcp.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !sgcp.Unknown3.Equals(other.Unknown3) {
		return false
	}

	return sgcp.Unknown4.Equals(other.Unknown4)
}

// String returns the string representation of the SubscriberGetContentParam
func (sgcp *SubscriberGetContentParam) String() string {
	return sgcp.FormatToString(0)
}

// FormatToString pretty-prints the SubscriberGetContentParam using the provided indentation level
func (sgcp *SubscriberGetContentParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberGetContentParam{\n")
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, sgcp.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, sgcp.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, sgcp.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %s,\n", indentationValues, sgcp.Unknown4))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberGetContentParam returns a new SubscriberGetContentParam
func NewSubscriberGetContentParam() *SubscriberGetContentParam {
	sgcp := &SubscriberGetContentParam{
		Unknown1: types.NewString(""),
		Unknown2: types.NewPrimitiveU32(0),
		Unknown3: types.NewPrimitiveU32(0),
		Unknown4: types.NewPrimitiveU64(0),
	}

	return sgcp
}
