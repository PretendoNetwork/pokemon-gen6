// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendMiiList is a type within the Friends3DS protocol
type FriendMiiList struct {
	types.Structure
	*types.Data
	Unknown1 *types.PrimitiveU32
	MiiList  *MiiList
	Unknown2 *types.DateTime
}

// WriteTo writes the FriendMiiList to the given writable
func (fml *FriendMiiList) WriteTo(writable types.Writable) {
	fml.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	fml.Unknown1.WriteTo(contentWritable)
	fml.MiiList.WriteTo(contentWritable)
	fml.Unknown2.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fml.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendMiiList from the given readable
func (fml *FriendMiiList) ExtractFrom(readable types.Readable) error {
	var err error

	err = fml.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendMiiList.Data. %s", err.Error())
	}

	err = fml.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendMiiList header. %s", err.Error())
	}

	err = fml.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendMiiList.Unknown1. %s", err.Error())
	}

	err = fml.MiiList.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendMiiList.MiiList. %s", err.Error())
	}

	err = fml.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendMiiList.Unknown2. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendMiiList
func (fml *FriendMiiList) Copy() types.RVType {
	copied := NewFriendMiiList()

	copied.StructureVersion = fml.StructureVersion
	copied.Data = fml.Data.Copy().(*types.Data)
	copied.Unknown1 = fml.Unknown1.Copy().(*types.PrimitiveU32)
	copied.MiiList = fml.MiiList.Copy().(*MiiList)
	copied.Unknown2 = fml.Unknown2.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given FriendMiiList contains the same data as the current FriendMiiList
func (fml *FriendMiiList) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendMiiList); !ok {
		return false
	}

	other := o.(*FriendMiiList)

	if fml.StructureVersion != other.StructureVersion {
		return false
	}

	if !fml.Data.Equals(other.Data) {
		return false
	}

	if !fml.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !fml.MiiList.Equals(other.MiiList) {
		return false
	}

	return fml.Unknown2.Equals(other.Unknown2)
}

// String returns the string representation of the FriendMiiList
func (fml *FriendMiiList) String() string {
	return fml.FormatToString(0)
}

// FormatToString pretty-prints the FriendMiiList using the provided indentation level
func (fml *FriendMiiList) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendMiiList{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, fml.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, fml.Unknown1))
	b.WriteString(fmt.Sprintf("%sMiiList: %s,\n", indentationValues, fml.MiiList.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, fml.Unknown2.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendMiiList returns a new FriendMiiList
func NewFriendMiiList() *FriendMiiList {
	fml := &FriendMiiList{
		Data:     types.NewData(),
		Unknown1: types.NewPrimitiveU32(0),
		MiiList:  NewMiiList(),
		Unknown2: types.NewDateTime(0),
	}

	return fml
}
