// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SubscriberPostContentParam is a type within the Shop protocol
type SubscriberPostContentParam struct {
	types.Structure
	Unknown1 *types.List[*types.String]
	Unknown2 *types.String
	Unknown3 *types.QBuffer
}

// WriteTo writes the SubscriberPostContentParam to the given writable
func (spcp *SubscriberPostContentParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	spcp.Unknown1.WriteTo(contentWritable)
	spcp.Unknown2.WriteTo(contentWritable)
	spcp.Unknown3.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	spcp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SubscriberPostContentParam from the given readable
func (spcp *SubscriberPostContentParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = spcp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam header. %s", err.Error())
	}

	err = spcp.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Unknown1. %s", err.Error())
	}

	err = spcp.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Unknown2. %s", err.Error())
	}

	err = spcp.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Unknown3. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SubscriberPostContentParam
func (spcp *SubscriberPostContentParam) Copy() types.RVType {
	copied := NewSubscriberPostContentParam()

	copied.StructureVersion = spcp.StructureVersion
	copied.Unknown1 = spcp.Unknown1.Copy().(*types.List[*types.String])
	copied.Unknown2 = spcp.Unknown2.Copy().(*types.String)
	copied.Unknown3 = spcp.Unknown3.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given SubscriberPostContentParam contains the same data as the current SubscriberPostContentParam
func (spcp *SubscriberPostContentParam) Equals(o types.RVType) bool {
	if _, ok := o.(*SubscriberPostContentParam); !ok {
		return false
	}

	other := o.(*SubscriberPostContentParam)

	if spcp.StructureVersion != other.StructureVersion {
		return false
	}

	if !spcp.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !spcp.Unknown2.Equals(other.Unknown2) {
		return false
	}

	return spcp.Unknown3.Equals(other.Unknown3)
}

// String returns the string representation of the SubscriberPostContentParam
func (spcp *SubscriberPostContentParam) String() string {
	return spcp.FormatToString(0)
}

// FormatToString pretty-prints the SubscriberPostContentParam using the provided indentation level
func (spcp *SubscriberPostContentParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberPostContentParam{\n")
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, spcp.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, spcp.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, spcp.Unknown3))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberPostContentParam returns a new SubscriberPostContentParam
func NewSubscriberPostContentParam() *SubscriberPostContentParam {
	spcp := &SubscriberPostContentParam{
		Unknown1: types.NewList[*types.String](),
		Unknown2: types.NewString(""),
		Unknown3: types.NewQBuffer(nil),
	}

	spcp.Unknown1.Type = types.NewString("")

	return spcp
}
