// Package types implements all the types used by the MatchmakeExtension protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SimpleSearchDateTimeAttribute is a type within the MatchmakeExtension protocol
type SimpleSearchDateTimeAttribute struct {
	types.Structure
	Unknown   *types.PrimitiveU32
	Unknown2  *types.PrimitiveU32
	Unknown3  *types.PrimitiveU32
	Unknown4  *types.PrimitiveU32
	StartTime *types.DateTime
	EndTime   *types.DateTime
}

// WriteTo writes the SimpleSearchDateTimeAttribute to the given writable
func (ssdta *SimpleSearchDateTimeAttribute) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ssdta.Unknown.WriteTo(contentWritable)
	ssdta.Unknown2.WriteTo(contentWritable)
	ssdta.Unknown3.WriteTo(contentWritable)
	ssdta.Unknown4.WriteTo(contentWritable)
	ssdta.StartTime.WriteTo(contentWritable)
	ssdta.EndTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ssdta.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SimpleSearchDateTimeAttribute from the given readable
func (ssdta *SimpleSearchDateTimeAttribute) ExtractFrom(readable types.Readable) error {
	var err error

	err = ssdta.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute header. %s", err.Error())
	}

	err = ssdta.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown. %s", err.Error())
	}

	err = ssdta.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown2. %s", err.Error())
	}

	err = ssdta.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown3. %s", err.Error())
	}

	err = ssdta.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown4. %s", err.Error())
	}

	err = ssdta.StartTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.StartTime. %s", err.Error())
	}

	err = ssdta.EndTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.EndTime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SimpleSearchDateTimeAttribute
func (ssdta *SimpleSearchDateTimeAttribute) Copy() types.RVType {
	copied := NewSimpleSearchDateTimeAttribute()

	copied.StructureVersion = ssdta.StructureVersion
	copied.Unknown = ssdta.Unknown.Copy().(*types.PrimitiveU32)
	copied.Unknown2 = ssdta.Unknown2.Copy().(*types.PrimitiveU32)
	copied.Unknown3 = ssdta.Unknown3.Copy().(*types.PrimitiveU32)
	copied.Unknown4 = ssdta.Unknown4.Copy().(*types.PrimitiveU32)
	copied.StartTime = ssdta.StartTime.Copy().(*types.DateTime)
	copied.EndTime = ssdta.EndTime.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given SimpleSearchDateTimeAttribute contains the same data as the current SimpleSearchDateTimeAttribute
func (ssdta *SimpleSearchDateTimeAttribute) Equals(o types.RVType) bool {
	if _, ok := o.(*SimpleSearchDateTimeAttribute); !ok {
		return false
	}

	other := o.(*SimpleSearchDateTimeAttribute)

	if ssdta.StructureVersion != other.StructureVersion {
		return false
	}

	if !ssdta.Unknown.Equals(other.Unknown) {
		return false
	}

	if !ssdta.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !ssdta.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !ssdta.Unknown4.Equals(other.Unknown4) {
		return false
	}

	if !ssdta.StartTime.Equals(other.StartTime) {
		return false
	}

	return ssdta.EndTime.Equals(other.EndTime)
}

// String returns the string representation of the SimpleSearchDateTimeAttribute
func (ssdta *SimpleSearchDateTimeAttribute) String() string {
	return ssdta.FormatToString(0)
}

// FormatToString pretty-prints the SimpleSearchDateTimeAttribute using the provided indentation level
func (ssdta *SimpleSearchDateTimeAttribute) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimpleSearchDateTimeAttribute{\n")
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, ssdta.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, ssdta.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, ssdta.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %s,\n", indentationValues, ssdta.Unknown4))
	b.WriteString(fmt.Sprintf("%sStartTime: %s,\n", indentationValues, ssdta.StartTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sEndTime: %s,\n", indentationValues, ssdta.EndTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchDateTimeAttribute returns a new SimpleSearchDateTimeAttribute
func NewSimpleSearchDateTimeAttribute() *SimpleSearchDateTimeAttribute {
	ssdta := &SimpleSearchDateTimeAttribute{
		Unknown:   types.NewPrimitiveU32(0),
		Unknown2:  types.NewPrimitiveU32(0),
		Unknown3:  types.NewPrimitiveU32(0),
		Unknown4:  types.NewPrimitiveU32(0),
		StartTime: types.NewDateTime(0),
		EndTime:   types.NewDateTime(0),
	}

	return ssdta
}
