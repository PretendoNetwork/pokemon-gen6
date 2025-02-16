// Package types implements all the types used by the MatchmakeExtension protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendUserParam is a type within the MatchmakeExtension protocol
type FriendUserParam struct {
	types.Structure
	Name *types.String
}

// WriteTo writes the FriendUserParam to the given writable
func (fup *FriendUserParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	fup.Name.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fup.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendUserParam from the given readable
func (fup *FriendUserParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = fup.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserParam header. %s", err.Error())
	}

	err = fup.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserParam.Name. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendUserParam
func (fup *FriendUserParam) Copy() types.RVType {
	copied := NewFriendUserParam()

	copied.StructureVersion = fup.StructureVersion
	copied.Name = fup.Name.Copy().(*types.String)

	return copied
}

// Equals checks if the given FriendUserParam contains the same data as the current FriendUserParam
func (fup *FriendUserParam) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendUserParam); !ok {
		return false
	}

	other := o.(*FriendUserParam)

	if fup.StructureVersion != other.StructureVersion {
		return false
	}

	return fup.Name.Equals(other.Name)
}

// String returns the string representation of the FriendUserParam
func (fup *FriendUserParam) String() string {
	return fup.FormatToString(0)
}

// FormatToString pretty-prints the FriendUserParam using the provided indentation level
func (fup *FriendUserParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendUserParam{\n")
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, fup.Name))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendUserParam returns a new FriendUserParam
func NewFriendUserParam() *FriendUserParam {
	fup := &FriendUserParam{
		Name: types.NewString(""),
	}

	return fup
}
