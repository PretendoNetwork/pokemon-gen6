// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendMii is a type within the Friends3DS protocol
type FriendMii struct {
	types.Structure
	*types.Data
	PID        *types.PID
	Mii        *Mii
	ModifiedAt *types.DateTime
}

// WriteTo writes the FriendMii to the given writable
func (fm *FriendMii) WriteTo(writable types.Writable) {
	fm.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	fm.PID.WriteTo(contentWritable)
	fm.Mii.WriteTo(contentWritable)
	fm.ModifiedAt.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fm.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendMii from the given readable
func (fm *FriendMii) ExtractFrom(readable types.Readable) error {
	var err error

	err = fm.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendMii.Data. %s", err.Error())
	}

	err = fm.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendMii header. %s", err.Error())
	}

	err = fm.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendMii.PID. %s", err.Error())
	}

	err = fm.Mii.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendMii.Mii. %s", err.Error())
	}

	err = fm.ModifiedAt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendMii.ModifiedAt. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendMii
func (fm *FriendMii) Copy() types.RVType {
	copied := NewFriendMii()

	copied.StructureVersion = fm.StructureVersion
	copied.Data = fm.Data.Copy().(*types.Data)
	copied.PID = fm.PID.Copy().(*types.PID)
	copied.Mii = fm.Mii.Copy().(*Mii)
	copied.ModifiedAt = fm.ModifiedAt.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given FriendMii contains the same data as the current FriendMii
func (fm *FriendMii) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendMii); !ok {
		return false
	}

	other := o.(*FriendMii)

	if fm.StructureVersion != other.StructureVersion {
		return false
	}

	if !fm.Data.Equals(other.Data) {
		return false
	}

	if !fm.PID.Equals(other.PID) {
		return false
	}

	if !fm.Mii.Equals(other.Mii) {
		return false
	}

	return fm.ModifiedAt.Equals(other.ModifiedAt)
}

// String returns the string representation of the FriendMii
func (fm *FriendMii) String() string {
	return fm.FormatToString(0)
}

// FormatToString pretty-prints the FriendMii using the provided indentation level
func (fm *FriendMii) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendMii{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, fm.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, fm.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMii: %s,\n", indentationValues, fm.Mii.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sModifiedAt: %s,\n", indentationValues, fm.ModifiedAt.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendMii returns a new FriendMii
func NewFriendMii() *FriendMii {
	fm := &FriendMii{
		Data:       types.NewData(),
		PID:        types.NewPID(0),
		Mii:        NewMii(),
		ModifiedAt: types.NewDateTime(0),
	}

	return fm
}
