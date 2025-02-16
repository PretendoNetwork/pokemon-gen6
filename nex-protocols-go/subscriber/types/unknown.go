// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Unknown is a type within the Shop protocol
type Unknown struct {
	types.Structure
	Unknown *types.QBuffer
}

// WriteTo writes the Unknown to the given writable
func (u *Unknown) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	u.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	u.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Unknown from the given readable
func (u *Unknown) ExtractFrom(readable types.Readable) error {
	var err error

	err = u.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Unknown header. %s", err.Error())
	}

	err = u.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Unknown.Unknown. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Unknown
func (u *Unknown) Copy() types.RVType {
	copied := NewUnknown()

	copied.StructureVersion = u.StructureVersion
	copied.Unknown = u.Unknown.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given Unknown contains the same data as the current Unknown
func (u *Unknown) Equals(o types.RVType) bool {
	if _, ok := o.(*Unknown); !ok {
		return false
	}

	other := o.(*Unknown)

	if u.StructureVersion != other.StructureVersion {
		return false
	}

	return u.Unknown.Equals(other.Unknown)
}

// String returns the string representation of the Unknown
func (u *Unknown) String() string {
	return u.FormatToString(0)
}

// FormatToString pretty-prints the Unknown using the provided indentation level
func (u *Unknown) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Unknown{\n")
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, u.Unknown))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewUnknown returns a new Unknown
func NewUnknown() *Unknown {
	u := &Unknown{
		Unknown: types.NewQBuffer(nil),
	}

	return u
}
