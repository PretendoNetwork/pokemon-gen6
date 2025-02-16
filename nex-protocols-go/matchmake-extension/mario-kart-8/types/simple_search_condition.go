// Package types implements all the types used by the MatchmakeExtension protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SimpleSearchCondition is a type within the MatchmakeExtension protocol
type SimpleSearchCondition struct {
	types.Structure
	Value              *types.PrimitiveU32
	ComparisonOperator *types.PrimitiveU32
}

// WriteTo writes the SimpleSearchCondition to the given writable
func (ssc *SimpleSearchCondition) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ssc.Value.WriteTo(contentWritable)
	ssc.ComparisonOperator.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ssc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SimpleSearchCondition from the given readable
func (ssc *SimpleSearchCondition) ExtractFrom(readable types.Readable) error {
	var err error

	err = ssc.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchCondition header. %s", err.Error())
	}

	err = ssc.Value.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchCondition.Value. %s", err.Error())
	}

	err = ssc.ComparisonOperator.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchCondition.ComparisonOperator. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SimpleSearchCondition
func (ssc *SimpleSearchCondition) Copy() types.RVType {
	copied := NewSimpleSearchCondition()

	copied.StructureVersion = ssc.StructureVersion
	copied.Value = ssc.Value.Copy().(*types.PrimitiveU32)
	copied.ComparisonOperator = ssc.ComparisonOperator.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given SimpleSearchCondition contains the same data as the current SimpleSearchCondition
func (ssc *SimpleSearchCondition) Equals(o types.RVType) bool {
	if _, ok := o.(*SimpleSearchCondition); !ok {
		return false
	}

	other := o.(*SimpleSearchCondition)

	if ssc.StructureVersion != other.StructureVersion {
		return false
	}

	if !ssc.Value.Equals(other.Value) {
		return false
	}

	return ssc.ComparisonOperator.Equals(other.ComparisonOperator)
}

// String returns the string representation of the SimpleSearchCondition
func (ssc *SimpleSearchCondition) String() string {
	return ssc.FormatToString(0)
}

// FormatToString pretty-prints the SimpleSearchCondition using the provided indentation level
func (ssc *SimpleSearchCondition) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimpleSearchCondition{\n")
	b.WriteString(fmt.Sprintf("%sValue: %s,\n", indentationValues, ssc.Value))
	b.WriteString(fmt.Sprintf("%sComparisonOperator: %s,\n", indentationValues, ssc.ComparisonOperator))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchCondition returns a new SimpleSearchCondition
func NewSimpleSearchCondition() *SimpleSearchCondition {
	ssc := &SimpleSearchCondition{
		Value:              types.NewPrimitiveU32(0),
		ComparisonOperator: types.NewPrimitiveU32(0),
	}

	return ssc
}
