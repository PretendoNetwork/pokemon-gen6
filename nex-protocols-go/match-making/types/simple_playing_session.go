// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SimplePlayingSession is a type within the Matchmaking protocol
type SimplePlayingSession struct {
	types.Structure
	PrincipalID *types.PID
	GatheringID *types.PrimitiveU32
	GameMode    *types.PrimitiveU32
	Attribute0  *types.PrimitiveU32
}

// WriteTo writes the SimplePlayingSession to the given writable
func (sps *SimplePlayingSession) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sps.PrincipalID.WriteTo(contentWritable)
	sps.GatheringID.WriteTo(contentWritable)
	sps.GameMode.WriteTo(contentWritable)
	sps.Attribute0.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sps.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SimplePlayingSession from the given readable
func (sps *SimplePlayingSession) ExtractFrom(readable types.Readable) error {
	var err error

	err = sps.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession header. %s", err.Error())
	}

	err = sps.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.PrincipalID. %s", err.Error())
	}

	err = sps.GatheringID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.GatheringID. %s", err.Error())
	}

	err = sps.GameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.GameMode. %s", err.Error())
	}

	err = sps.Attribute0.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.Attribute0. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SimplePlayingSession
func (sps *SimplePlayingSession) Copy() types.RVType {
	copied := NewSimplePlayingSession()

	copied.StructureVersion = sps.StructureVersion
	copied.PrincipalID = sps.PrincipalID.Copy().(*types.PID)
	copied.GatheringID = sps.GatheringID.Copy().(*types.PrimitiveU32)
	copied.GameMode = sps.GameMode.Copy().(*types.PrimitiveU32)
	copied.Attribute0 = sps.Attribute0.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given SimplePlayingSession contains the same data as the current SimplePlayingSession
func (sps *SimplePlayingSession) Equals(o types.RVType) bool {
	if _, ok := o.(*SimplePlayingSession); !ok {
		return false
	}

	other := o.(*SimplePlayingSession)

	if sps.StructureVersion != other.StructureVersion {
		return false
	}

	if !sps.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !sps.GatheringID.Equals(other.GatheringID) {
		return false
	}

	if !sps.GameMode.Equals(other.GameMode) {
		return false
	}

	return sps.Attribute0.Equals(other.Attribute0)
}

// String returns the string representation of the SimplePlayingSession
func (sps *SimplePlayingSession) String() string {
	return sps.FormatToString(0)
}

// FormatToString pretty-prints the SimplePlayingSession using the provided indentation level
func (sps *SimplePlayingSession) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimplePlayingSession{\n")
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, sps.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sGatheringID: %s,\n", indentationValues, sps.GatheringID))
	b.WriteString(fmt.Sprintf("%sGameMode: %s,\n", indentationValues, sps.GameMode))
	b.WriteString(fmt.Sprintf("%sAttribute0: %s,\n", indentationValues, sps.Attribute0))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimplePlayingSession returns a new SimplePlayingSession
func NewSimplePlayingSession() *SimplePlayingSession {
	sps := &SimplePlayingSession{
		PrincipalID: types.NewPID(0),
		GatheringID: types.NewPrimitiveU32(0),
		GameMode:    types.NewPrimitiveU32(0),
		Attribute0:  types.NewPrimitiveU32(0),
	}

	return sps
}
