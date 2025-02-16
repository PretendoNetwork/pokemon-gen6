// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Comment is a type within the FriendsWiiU protocol
type Comment struct {
	types.Structure
	*types.Data
	Unknown     *types.PrimitiveU8
	Contents    *types.String
	LastChanged *types.DateTime
}

// WriteTo writes the Comment to the given writable
func (c *Comment) WriteTo(writable types.Writable) {
	c.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	c.Unknown.WriteTo(contentWritable)
	c.Contents.WriteTo(contentWritable)
	c.LastChanged.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	c.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Comment from the given readable
func (c *Comment) ExtractFrom(readable types.Readable) error {
	var err error

	err = c.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.Data. %s", err.Error())
	}

	err = c.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Comment header. %s", err.Error())
	}

	err = c.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.Unknown. %s", err.Error())
	}

	err = c.Contents.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.Contents. %s", err.Error())
	}

	err = c.LastChanged.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.LastChanged. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Comment
func (c *Comment) Copy() types.RVType {
	copied := NewComment()

	copied.StructureVersion = c.StructureVersion
	copied.Data = c.Data.Copy().(*types.Data)
	copied.Unknown = c.Unknown.Copy().(*types.PrimitiveU8)
	copied.Contents = c.Contents.Copy().(*types.String)
	copied.LastChanged = c.LastChanged.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given Comment contains the same data as the current Comment
func (c *Comment) Equals(o types.RVType) bool {
	if _, ok := o.(*Comment); !ok {
		return false
	}

	other := o.(*Comment)

	if c.StructureVersion != other.StructureVersion {
		return false
	}

	if !c.Data.Equals(other.Data) {
		return false
	}

	if !c.Unknown.Equals(other.Unknown) {
		return false
	}

	if !c.Contents.Equals(other.Contents) {
		return false
	}

	return c.LastChanged.Equals(other.LastChanged)
}

// String returns the string representation of the Comment
func (c *Comment) String() string {
	return c.FormatToString(0)
}

// FormatToString pretty-prints the Comment using the provided indentation level
func (c *Comment) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Comment{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, c.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, c.Unknown))
	b.WriteString(fmt.Sprintf("%sContents: %s,\n", indentationValues, c.Contents))
	b.WriteString(fmt.Sprintf("%sLastChanged: %s,\n", indentationValues, c.LastChanged.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewComment returns a new Comment
func NewComment() *Comment {
	c := &Comment{
		Data:        types.NewData(),
		Unknown:     types.NewPrimitiveU8(0),
		Contents:    types.NewString(""),
		LastChanged: types.NewDateTime(0),
	}

	return c
}
