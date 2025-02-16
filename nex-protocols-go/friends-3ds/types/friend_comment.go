// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendComment is a type within the Friends3DS protocol
type FriendComment struct {
	types.Structure
	*types.Data
	PID        *types.PID
	Comment    *types.String
	ModifiedAt *types.DateTime
}

// WriteTo writes the FriendComment to the given writable
func (fc *FriendComment) WriteTo(writable types.Writable) {
	fc.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	fc.PID.WriteTo(contentWritable)
	fc.Comment.WriteTo(contentWritable)
	fc.ModifiedAt.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendComment from the given readable
func (fc *FriendComment) ExtractFrom(readable types.Readable) error {
	var err error

	err = fc.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendComment.Data. %s", err.Error())
	}

	err = fc.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendComment header. %s", err.Error())
	}

	err = fc.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendComment.PID. %s", err.Error())
	}

	err = fc.Comment.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendComment.Comment. %s", err.Error())
	}

	err = fc.ModifiedAt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendComment.ModifiedAt. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendComment
func (fc *FriendComment) Copy() types.RVType {
	copied := NewFriendComment()

	copied.StructureVersion = fc.StructureVersion
	copied.Data = fc.Data.Copy().(*types.Data)
	copied.PID = fc.PID.Copy().(*types.PID)
	copied.Comment = fc.Comment.Copy().(*types.String)
	copied.ModifiedAt = fc.ModifiedAt.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given FriendComment contains the same data as the current FriendComment
func (fc *FriendComment) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendComment); !ok {
		return false
	}

	other := o.(*FriendComment)

	if fc.StructureVersion != other.StructureVersion {
		return false
	}

	if !fc.Data.Equals(other.Data) {
		return false
	}

	if !fc.PID.Equals(other.PID) {
		return false
	}

	if !fc.Comment.Equals(other.Comment) {
		return false
	}

	return fc.ModifiedAt.Equals(other.ModifiedAt)
}

// String returns the string representation of the FriendComment
func (fc *FriendComment) String() string {
	return fc.FormatToString(0)
}

// FormatToString pretty-prints the FriendComment using the provided indentation level
func (fc *FriendComment) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendComment{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, fc.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, fc.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sComment: %s,\n", indentationValues, fc.Comment))
	b.WriteString(fmt.Sprintf("%sModifiedAt: %s,\n", indentationValues, fc.ModifiedAt.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendComment returns a new FriendComment
func NewFriendComment() *FriendComment {
	fc := &FriendComment{
		Data:       types.NewData(),
		PID:        types.NewPID(0),
		Comment:    types.NewString(""),
		ModifiedAt: types.NewDateTime(0),
	}

	return fc
}
