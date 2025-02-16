// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MiiList is a type within the Friends3DS protocol
type MiiList struct {
	types.Structure
	*types.Data
	Unknown1    *types.String
	Unknown2    *types.PrimitiveBool
	Unknown3    *types.PrimitiveU8
	MiiDataList *types.List[*types.Buffer]
}

// WriteTo writes the MiiList to the given writable
func (ml *MiiList) WriteTo(writable types.Writable) {
	ml.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	ml.Unknown1.WriteTo(contentWritable)
	ml.Unknown2.WriteTo(contentWritable)
	ml.Unknown3.WriteTo(contentWritable)
	ml.MiiDataList.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ml.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MiiList from the given readable
func (ml *MiiList) ExtractFrom(readable types.Readable) error {
	var err error

	err = ml.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.Data. %s", err.Error())
	}

	err = ml.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList header. %s", err.Error())
	}

	err = ml.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.Unknown1. %s", err.Error())
	}

	err = ml.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.Unknown2. %s", err.Error())
	}

	err = ml.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.Unknown3. %s", err.Error())
	}

	err = ml.MiiDataList.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.MiiDataList. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MiiList
func (ml *MiiList) Copy() types.RVType {
	copied := NewMiiList()

	copied.StructureVersion = ml.StructureVersion
	copied.Data = ml.Data.Copy().(*types.Data)
	copied.Unknown1 = ml.Unknown1.Copy().(*types.String)
	copied.Unknown2 = ml.Unknown2.Copy().(*types.PrimitiveBool)
	copied.Unknown3 = ml.Unknown3.Copy().(*types.PrimitiveU8)
	copied.MiiDataList = ml.MiiDataList.Copy().(*types.List[*types.Buffer])

	return copied
}

// Equals checks if the given MiiList contains the same data as the current MiiList
func (ml *MiiList) Equals(o types.RVType) bool {
	if _, ok := o.(*MiiList); !ok {
		return false
	}

	other := o.(*MiiList)

	if ml.StructureVersion != other.StructureVersion {
		return false
	}

	if !ml.Data.Equals(other.Data) {
		return false
	}

	if !ml.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !ml.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !ml.Unknown3.Equals(other.Unknown3) {
		return false
	}

	return ml.MiiDataList.Equals(other.MiiDataList)
}

// String returns the string representation of the MiiList
func (ml *MiiList) String() string {
	return ml.FormatToString(0)
}

// FormatToString pretty-prints the MiiList using the provided indentation level
func (ml *MiiList) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MiiList{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, ml.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, ml.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, ml.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, ml.Unknown3))
	b.WriteString(fmt.Sprintf("%sMiiDataList: %s,\n", indentationValues, ml.MiiDataList))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMiiList returns a new MiiList
func NewMiiList() *MiiList {
	ml := &MiiList{
		Data:        types.NewData(),
		Unknown1:    types.NewString(""),
		Unknown2:    types.NewPrimitiveBool(false),
		Unknown3:    types.NewPrimitiveU8(0),
		MiiDataList: types.NewList[*types.Buffer](),
	}

	ml.MiiDataList.Type = types.NewBuffer(nil)

	return ml
}
