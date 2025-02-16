// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DeletionEntry is a type within the Matchmaking protocol
type DeletionEntry struct {
	types.Structure
	IDGathering *types.PrimitiveU32
	PID         *types.PID
	UIReason    *types.PrimitiveU32
}

// WriteTo writes the DeletionEntry to the given writable
func (de *DeletionEntry) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	de.IDGathering.WriteTo(contentWritable)
	de.PID.WriteTo(contentWritable)
	de.UIReason.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	de.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DeletionEntry from the given readable
func (de *DeletionEntry) ExtractFrom(readable types.Readable) error {
	var err error

	err = de.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DeletionEntry header. %s", err.Error())
	}

	err = de.IDGathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DeletionEntry.IDGathering. %s", err.Error())
	}

	err = de.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DeletionEntry.PID. %s", err.Error())
	}

	err = de.UIReason.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DeletionEntry.UIReason. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DeletionEntry
func (de *DeletionEntry) Copy() types.RVType {
	copied := NewDeletionEntry()

	copied.StructureVersion = de.StructureVersion
	copied.IDGathering = de.IDGathering.Copy().(*types.PrimitiveU32)
	copied.PID = de.PID.Copy().(*types.PID)
	copied.UIReason = de.UIReason.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given DeletionEntry contains the same data as the current DeletionEntry
func (de *DeletionEntry) Equals(o types.RVType) bool {
	if _, ok := o.(*DeletionEntry); !ok {
		return false
	}

	other := o.(*DeletionEntry)

	if de.StructureVersion != other.StructureVersion {
		return false
	}

	if !de.IDGathering.Equals(other.IDGathering) {
		return false
	}

	if !de.PID.Equals(other.PID) {
		return false
	}

	return de.UIReason.Equals(other.UIReason)
}

// String returns the string representation of the DeletionEntry
func (de *DeletionEntry) String() string {
	return de.FormatToString(0)
}

// FormatToString pretty-prints the DeletionEntry using the provided indentation level
func (de *DeletionEntry) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DeletionEntry{\n")
	b.WriteString(fmt.Sprintf("%sIDGathering: %s,\n", indentationValues, de.IDGathering))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, de.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUIReason: %s,\n", indentationValues, de.UIReason))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDeletionEntry returns a new DeletionEntry
func NewDeletionEntry() *DeletionEntry {
	de := &DeletionEntry{
		IDGathering: types.NewPrimitiveU32(0),
		PID:         types.NewPID(0),
		UIReason:    types.NewPrimitiveU32(0),
	}

	return de
}
