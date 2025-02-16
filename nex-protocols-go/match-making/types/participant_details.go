// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ParticipantDetails is a type within the Matchmaking protocol
type ParticipantDetails struct {
	types.Structure
	IDParticipant  *types.PID
	StrName        *types.String
	StrMessage     *types.String
	UIParticipants *types.PrimitiveU16
}

// WriteTo writes the ParticipantDetails to the given writable
func (pd *ParticipantDetails) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	pd.IDParticipant.WriteTo(contentWritable)
	pd.StrName.WriteTo(contentWritable)
	pd.StrMessage.WriteTo(contentWritable)
	pd.UIParticipants.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ParticipantDetails from the given readable
func (pd *ParticipantDetails) ExtractFrom(readable types.Readable) error {
	var err error

	err = pd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails header. %s", err.Error())
	}

	err = pd.IDParticipant.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.IDParticipant. %s", err.Error())
	}

	err = pd.StrName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.StrName. %s", err.Error())
	}

	err = pd.StrMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.StrMessage. %s", err.Error())
	}

	err = pd.UIParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.UIParticipants. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ParticipantDetails
func (pd *ParticipantDetails) Copy() types.RVType {
	copied := NewParticipantDetails()

	copied.StructureVersion = pd.StructureVersion
	copied.IDParticipant = pd.IDParticipant.Copy().(*types.PID)
	copied.StrName = pd.StrName.Copy().(*types.String)
	copied.StrMessage = pd.StrMessage.Copy().(*types.String)
	copied.UIParticipants = pd.UIParticipants.Copy().(*types.PrimitiveU16)

	return copied
}

// Equals checks if the given ParticipantDetails contains the same data as the current ParticipantDetails
func (pd *ParticipantDetails) Equals(o types.RVType) bool {
	if _, ok := o.(*ParticipantDetails); !ok {
		return false
	}

	other := o.(*ParticipantDetails)

	if pd.StructureVersion != other.StructureVersion {
		return false
	}

	if !pd.IDParticipant.Equals(other.IDParticipant) {
		return false
	}

	if !pd.StrName.Equals(other.StrName) {
		return false
	}

	if !pd.StrMessage.Equals(other.StrMessage) {
		return false
	}

	return pd.UIParticipants.Equals(other.UIParticipants)
}

// String returns the string representation of the ParticipantDetails
func (pd *ParticipantDetails) String() string {
	return pd.FormatToString(0)
}

// FormatToString pretty-prints the ParticipantDetails using the provided indentation level
func (pd *ParticipantDetails) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ParticipantDetails{\n")
	b.WriteString(fmt.Sprintf("%sIDParticipant: %s,\n", indentationValues, pd.IDParticipant.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrName: %s,\n", indentationValues, pd.StrName))
	b.WriteString(fmt.Sprintf("%sStrMessage: %s,\n", indentationValues, pd.StrMessage))
	b.WriteString(fmt.Sprintf("%sUIParticipants: %s,\n", indentationValues, pd.UIParticipants))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewParticipantDetails returns a new ParticipantDetails
func NewParticipantDetails() *ParticipantDetails {
	pd := &ParticipantDetails{
		IDParticipant:  types.NewPID(0),
		StrName:        types.NewString(""),
		StrMessage:     types.NewString(""),
		UIParticipants: types.NewPrimitiveU16(0),
	}

	return pd
}
