// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Invitation is a type within the Matchmaking protocol
type Invitation struct {
	types.Structure
	IDGathering *types.PrimitiveU32
	IDGuest     *types.PrimitiveU32
	StrMessage  *types.String
}

// WriteTo writes the Invitation to the given writable
func (i *Invitation) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	i.IDGathering.WriteTo(contentWritable)
	i.IDGuest.WriteTo(contentWritable)
	i.StrMessage.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	i.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Invitation from the given readable
func (i *Invitation) ExtractFrom(readable types.Readable) error {
	var err error

	err = i.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Invitation header. %s", err.Error())
	}

	err = i.IDGathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Invitation.IDGathering. %s", err.Error())
	}

	err = i.IDGuest.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Invitation.IDGuest. %s", err.Error())
	}

	err = i.StrMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Invitation.StrMessage. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Invitation
func (i *Invitation) Copy() types.RVType {
	copied := NewInvitation()

	copied.StructureVersion = i.StructureVersion
	copied.IDGathering = i.IDGathering.Copy().(*types.PrimitiveU32)
	copied.IDGuest = i.IDGuest.Copy().(*types.PrimitiveU32)
	copied.StrMessage = i.StrMessage.Copy().(*types.String)

	return copied
}

// Equals checks if the given Invitation contains the same data as the current Invitation
func (i *Invitation) Equals(o types.RVType) bool {
	if _, ok := o.(*Invitation); !ok {
		return false
	}

	other := o.(*Invitation)

	if i.StructureVersion != other.StructureVersion {
		return false
	}

	if !i.IDGathering.Equals(other.IDGathering) {
		return false
	}

	if !i.IDGuest.Equals(other.IDGuest) {
		return false
	}

	return i.StrMessage.Equals(other.StrMessage)
}

// String returns the string representation of the Invitation
func (i *Invitation) String() string {
	return i.FormatToString(0)
}

// FormatToString pretty-prints the Invitation using the provided indentation level
func (i *Invitation) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Invitation{\n")
	b.WriteString(fmt.Sprintf("%sIDGathering: %s,\n", indentationValues, i.IDGathering))
	b.WriteString(fmt.Sprintf("%sIDGuest: %s,\n", indentationValues, i.IDGuest))
	b.WriteString(fmt.Sprintf("%sStrMessage: %s,\n", indentationValues, i.StrMessage))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewInvitation returns a new Invitation
func NewInvitation() *Invitation {
	i := &Invitation{
		IDGathering: types.NewPrimitiveU32(0),
		IDGuest:     types.NewPrimitiveU32(0),
		StrMessage:  types.NewString(""),
	}

	return i
}
