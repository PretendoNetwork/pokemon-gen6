// Package types implements all the types used by the MatchmakeExtension protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SimpleSearchParam is a type within the MatchmakeExtension protocol
type SimpleSearchParam struct {
	types.Structure
	Unknown     *types.PrimitiveU32
	Unknown2    *types.PID
	Conditions  *types.List[*SimpleSearchCondition]
	Unknown3    *types.String
	ResultRange *types.ResultRange
	Unknown4    *types.DateTime
}

// WriteTo writes the SimpleSearchParam to the given writable
func (ssp *SimpleSearchParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ssp.Unknown.WriteTo(contentWritable)
	ssp.Unknown2.WriteTo(contentWritable)
	ssp.Conditions.WriteTo(contentWritable)
	ssp.Unknown3.WriteTo(contentWritable)
	ssp.ResultRange.WriteTo(contentWritable)
	ssp.Unknown4.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ssp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SimpleSearchParam from the given readable
func (ssp *SimpleSearchParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = ssp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam header. %s", err.Error())
	}

	err = ssp.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown. %s", err.Error())
	}

	err = ssp.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown2. %s", err.Error())
	}

	err = ssp.Conditions.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Conditions. %s", err.Error())
	}

	err = ssp.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown3. %s", err.Error())
	}

	err = ssp.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.ResultRange. %s", err.Error())
	}

	err = ssp.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown4. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SimpleSearchParam
func (ssp *SimpleSearchParam) Copy() types.RVType {
	copied := NewSimpleSearchParam()

	copied.StructureVersion = ssp.StructureVersion
	copied.Unknown = ssp.Unknown.Copy().(*types.PrimitiveU32)
	copied.Unknown2 = ssp.Unknown2.Copy().(*types.PID)
	copied.Conditions = ssp.Conditions.Copy().(*types.List[*SimpleSearchCondition])
	copied.Unknown3 = ssp.Unknown3.Copy().(*types.String)
	copied.ResultRange = ssp.ResultRange.Copy().(*types.ResultRange)
	copied.Unknown4 = ssp.Unknown4.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given SimpleSearchParam contains the same data as the current SimpleSearchParam
func (ssp *SimpleSearchParam) Equals(o types.RVType) bool {
	if _, ok := o.(*SimpleSearchParam); !ok {
		return false
	}

	other := o.(*SimpleSearchParam)

	if ssp.StructureVersion != other.StructureVersion {
		return false
	}

	if !ssp.Unknown.Equals(other.Unknown) {
		return false
	}

	if !ssp.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !ssp.Conditions.Equals(other.Conditions) {
		return false
	}

	if !ssp.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !ssp.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return ssp.Unknown4.Equals(other.Unknown4)
}

// String returns the string representation of the SimpleSearchParam
func (ssp *SimpleSearchParam) String() string {
	return ssp.FormatToString(0)
}

// FormatToString pretty-prints the SimpleSearchParam using the provided indentation level
func (ssp *SimpleSearchParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimpleSearchParam{\n")
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, ssp.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, ssp.Unknown2.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sConditions: %s,\n", indentationValues, ssp.Conditions))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, ssp.Unknown3))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, ssp.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown4: %s,\n", indentationValues, ssp.Unknown4.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchParam returns a new SimpleSearchParam
func NewSimpleSearchParam() *SimpleSearchParam {
	ssp := &SimpleSearchParam{
		Unknown:     types.NewPrimitiveU32(0),
		Unknown2:    types.NewPID(0),
		Conditions:  types.NewList[*SimpleSearchCondition](),
		Unknown3:    types.NewString(""),
		ResultRange: types.NewResultRange(),
		Unknown4:    types.NewDateTime(0),
	}

	ssp.Conditions.Type = NewSimpleSearchCondition()

	return ssp
}
