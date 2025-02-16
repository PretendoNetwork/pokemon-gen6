// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MiiV2 is a type within the FriendsWiiU protocol
type MiiV2 struct {
	types.Structure
	*types.Data
	Name     *types.String
	Unknown1 *types.PrimitiveU8
	Unknown2 *types.PrimitiveU8
	MiiData  *types.Buffer
	Datetime *types.DateTime
}

// WriteTo writes the MiiV2 to the given writable
func (mv *MiiV2) WriteTo(writable types.Writable) {
	mv.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	mv.Name.WriteTo(contentWritable)
	mv.Unknown1.WriteTo(contentWritable)
	mv.Unknown2.WriteTo(contentWritable)
	mv.MiiData.WriteTo(contentWritable)
	mv.Datetime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MiiV2 from the given readable
func (mv *MiiV2) ExtractFrom(readable types.Readable) error {
	var err error

	err = mv.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Data. %s", err.Error())
	}

	err = mv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2 header. %s", err.Error())
	}

	err = mv.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Name. %s", err.Error())
	}

	err = mv.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Unknown1. %s", err.Error())
	}

	err = mv.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Unknown2. %s", err.Error())
	}

	err = mv.MiiData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.MiiData. %s", err.Error())
	}

	err = mv.Datetime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Datetime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MiiV2
func (mv *MiiV2) Copy() types.RVType {
	copied := NewMiiV2()

	copied.StructureVersion = mv.StructureVersion
	copied.Data = mv.Data.Copy().(*types.Data)
	copied.Name = mv.Name.Copy().(*types.String)
	copied.Unknown1 = mv.Unknown1.Copy().(*types.PrimitiveU8)
	copied.Unknown2 = mv.Unknown2.Copy().(*types.PrimitiveU8)
	copied.MiiData = mv.MiiData.Copy().(*types.Buffer)
	copied.Datetime = mv.Datetime.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given MiiV2 contains the same data as the current MiiV2
func (mv *MiiV2) Equals(o types.RVType) bool {
	if _, ok := o.(*MiiV2); !ok {
		return false
	}

	other := o.(*MiiV2)

	if mv.StructureVersion != other.StructureVersion {
		return false
	}

	if !mv.Data.Equals(other.Data) {
		return false
	}

	if !mv.Name.Equals(other.Name) {
		return false
	}

	if !mv.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !mv.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !mv.MiiData.Equals(other.MiiData) {
		return false
	}

	return mv.Datetime.Equals(other.Datetime)
}

// String returns the string representation of the MiiV2
func (mv *MiiV2) String() string {
	return mv.FormatToString(0)
}

// FormatToString pretty-prints the MiiV2 using the provided indentation level
func (mv *MiiV2) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MiiV2{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, mv.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, mv.Name))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, mv.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, mv.Unknown2))
	b.WriteString(fmt.Sprintf("%sMiiData: %s,\n", indentationValues, mv.MiiData))
	b.WriteString(fmt.Sprintf("%sDatetime: %s,\n", indentationValues, mv.Datetime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMiiV2 returns a new MiiV2
func NewMiiV2() *MiiV2 {
	mv := &MiiV2{
		Data:     types.NewData(),
		Name:     types.NewString(""),
		Unknown1: types.NewPrimitiveU8(0),
		Unknown2: types.NewPrimitiveU8(0),
		MiiData:  types.NewBuffer(nil),
		Datetime: types.NewDateTime(0),
	}

	return mv
}
