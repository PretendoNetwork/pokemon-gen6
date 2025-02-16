// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PlayingSession is a type within the Matchmaking protocol
type PlayingSession struct {
	types.Structure
	PrincipalID *types.PID
	Gathering   *types.AnyDataHolder
}

// WriteTo writes the PlayingSession to the given writable
func (ps *PlayingSession) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ps.PrincipalID.WriteTo(contentWritable)
	ps.Gathering.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ps.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PlayingSession from the given readable
func (ps *PlayingSession) ExtractFrom(readable types.Readable) error {
	var err error

	err = ps.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PlayingSession header. %s", err.Error())
	}

	err = ps.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PlayingSession.PrincipalID. %s", err.Error())
	}

	err = ps.Gathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PlayingSession.Gathering. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PlayingSession
func (ps *PlayingSession) Copy() types.RVType {
	copied := NewPlayingSession()

	copied.StructureVersion = ps.StructureVersion
	copied.PrincipalID = ps.PrincipalID.Copy().(*types.PID)
	copied.Gathering = ps.Gathering.Copy().(*types.AnyDataHolder)

	return copied
}

// Equals checks if the given PlayingSession contains the same data as the current PlayingSession
func (ps *PlayingSession) Equals(o types.RVType) bool {
	if _, ok := o.(*PlayingSession); !ok {
		return false
	}

	other := o.(*PlayingSession)

	if ps.StructureVersion != other.StructureVersion {
		return false
	}

	if !ps.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	return ps.Gathering.Equals(other.Gathering)
}

// String returns the string representation of the PlayingSession
func (ps *PlayingSession) String() string {
	return ps.FormatToString(0)
}

// FormatToString pretty-prints the PlayingSession using the provided indentation level
func (ps *PlayingSession) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PlayingSession{\n")
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, ps.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sGathering: %s,\n", indentationValues, ps.Gathering.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPlayingSession returns a new PlayingSession
func NewPlayingSession() *PlayingSession {
	ps := &PlayingSession{
		PrincipalID: types.NewPID(0),
		Gathering:   types.NewAnyDataHolder(),
	}

	return ps
}
