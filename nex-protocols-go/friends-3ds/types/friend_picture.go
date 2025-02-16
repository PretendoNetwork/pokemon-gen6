// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendPicture is a type within the Friends3DS protocol
type FriendPicture struct {
	types.Structure
	*types.Data
	Unknown1    *types.PrimitiveU32
	PictureData *types.Buffer
	Unknown2    *types.DateTime
}

// WriteTo writes the FriendPicture to the given writable
func (fp *FriendPicture) WriteTo(writable types.Writable) {
	fp.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	fp.Unknown1.WriteTo(contentWritable)
	fp.PictureData.WriteTo(contentWritable)
	fp.Unknown2.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendPicture from the given readable
func (fp *FriendPicture) ExtractFrom(readable types.Readable) error {
	var err error

	err = fp.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPicture.Data. %s", err.Error())
	}

	err = fp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPicture header. %s", err.Error())
	}

	err = fp.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPicture.Unknown1. %s", err.Error())
	}

	err = fp.PictureData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPicture.PictureData. %s", err.Error())
	}

	err = fp.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPicture.Unknown2. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendPicture
func (fp *FriendPicture) Copy() types.RVType {
	copied := NewFriendPicture()

	copied.StructureVersion = fp.StructureVersion
	copied.Data = fp.Data.Copy().(*types.Data)
	copied.Unknown1 = fp.Unknown1.Copy().(*types.PrimitiveU32)
	copied.PictureData = fp.PictureData.Copy().(*types.Buffer)
	copied.Unknown2 = fp.Unknown2.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given FriendPicture contains the same data as the current FriendPicture
func (fp *FriendPicture) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendPicture); !ok {
		return false
	}

	other := o.(*FriendPicture)

	if fp.StructureVersion != other.StructureVersion {
		return false
	}

	if !fp.Data.Equals(other.Data) {
		return false
	}

	if !fp.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !fp.PictureData.Equals(other.PictureData) {
		return false
	}

	return fp.Unknown2.Equals(other.Unknown2)
}

// String returns the string representation of the FriendPicture
func (fp *FriendPicture) String() string {
	return fp.FormatToString(0)
}

// FormatToString pretty-prints the FriendPicture using the provided indentation level
func (fp *FriendPicture) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendPicture{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, fp.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, fp.Unknown1))
	b.WriteString(fmt.Sprintf("%sPictureData: %s,\n", indentationValues, fp.PictureData))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, fp.Unknown2.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendPicture returns a new FriendPicture
func NewFriendPicture() *FriendPicture {
	fp := &FriendPicture{
		Data:        types.NewData(),
		Unknown1:    types.NewPrimitiveU32(0),
		PictureData: types.NewBuffer(nil),
		Unknown2:    types.NewDateTime(0),
	}

	return fp
}
