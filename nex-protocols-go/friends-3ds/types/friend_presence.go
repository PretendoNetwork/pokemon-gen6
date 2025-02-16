// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendPresence is a type within the Friends3DS protocol
type FriendPresence struct {
	types.Structure
	*types.Data
	PID      *types.PID
	Presence *NintendoPresence
}

// WriteTo writes the FriendPresence to the given writable
func (fp *FriendPresence) WriteTo(writable types.Writable) {
	fp.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	fp.PID.WriteTo(contentWritable)
	fp.Presence.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendPresence from the given readable
func (fp *FriendPresence) ExtractFrom(readable types.Readable) error {
	var err error

	err = fp.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPresence.Data. %s", err.Error())
	}

	err = fp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPresence header. %s", err.Error())
	}

	err = fp.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPresence.PID. %s", err.Error())
	}

	err = fp.Presence.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPresence.Presence. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendPresence
func (fp *FriendPresence) Copy() types.RVType {
	copied := NewFriendPresence()

	copied.StructureVersion = fp.StructureVersion
	copied.Data = fp.Data.Copy().(*types.Data)
	copied.PID = fp.PID.Copy().(*types.PID)
	copied.Presence = fp.Presence.Copy().(*NintendoPresence)

	return copied
}

// Equals checks if the given FriendPresence contains the same data as the current FriendPresence
func (fp *FriendPresence) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendPresence); !ok {
		return false
	}

	other := o.(*FriendPresence)

	if fp.StructureVersion != other.StructureVersion {
		return false
	}

	if !fp.Data.Equals(other.Data) {
		return false
	}

	if !fp.PID.Equals(other.PID) {
		return false
	}

	return fp.Presence.Equals(other.Presence)
}

// String returns the string representation of the FriendPresence
func (fp *FriendPresence) String() string {
	return fp.FormatToString(0)
}

// FormatToString pretty-prints the FriendPresence using the provided indentation level
func (fp *FriendPresence) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendPresence{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, fp.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, fp.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPresence: %s,\n", indentationValues, fp.Presence.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendPresence returns a new FriendPresence
func NewFriendPresence() *FriendPresence {
	fp := &FriendPresence{
		Data:     types.NewData(),
		PID:      types.NewPID(0),
		Presence: NewNintendoPresence(),
	}

	return fp
}
