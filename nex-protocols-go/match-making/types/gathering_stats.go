// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GatheringStats is a type within the Matchmaking protocol
type GatheringStats struct {
	types.Structure
	PIDParticipant *types.PID
	UIFlags        *types.PrimitiveU32
	LstValues      *types.List[*types.PrimitiveF32]
}

// WriteTo writes the GatheringStats to the given writable
func (gs *GatheringStats) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gs.PIDParticipant.WriteTo(contentWritable)
	gs.UIFlags.WriteTo(contentWritable)
	gs.LstValues.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GatheringStats from the given readable
func (gs *GatheringStats) ExtractFrom(readable types.Readable) error {
	var err error

	err = gs.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringStats header. %s", err.Error())
	}

	err = gs.PIDParticipant.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringStats.PIDParticipant. %s", err.Error())
	}

	err = gs.UIFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringStats.UIFlags. %s", err.Error())
	}

	err = gs.LstValues.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringStats.LstValues. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GatheringStats
func (gs *GatheringStats) Copy() types.RVType {
	copied := NewGatheringStats()

	copied.StructureVersion = gs.StructureVersion
	copied.PIDParticipant = gs.PIDParticipant.Copy().(*types.PID)
	copied.UIFlags = gs.UIFlags.Copy().(*types.PrimitiveU32)
	copied.LstValues = gs.LstValues.Copy().(*types.List[*types.PrimitiveF32])

	return copied
}

// Equals checks if the given GatheringStats contains the same data as the current GatheringStats
func (gs *GatheringStats) Equals(o types.RVType) bool {
	if _, ok := o.(*GatheringStats); !ok {
		return false
	}

	other := o.(*GatheringStats)

	if gs.StructureVersion != other.StructureVersion {
		return false
	}

	if !gs.PIDParticipant.Equals(other.PIDParticipant) {
		return false
	}

	if !gs.UIFlags.Equals(other.UIFlags) {
		return false
	}

	return gs.LstValues.Equals(other.LstValues)
}

// String returns the string representation of the GatheringStats
func (gs *GatheringStats) String() string {
	return gs.FormatToString(0)
}

// FormatToString pretty-prints the GatheringStats using the provided indentation level
func (gs *GatheringStats) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GatheringStats{\n")
	b.WriteString(fmt.Sprintf("%sPIDParticipant: %s,\n", indentationValues, gs.PIDParticipant.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUIFlags: %s,\n", indentationValues, gs.UIFlags))
	b.WriteString(fmt.Sprintf("%sLstValues: %s,\n", indentationValues, gs.LstValues))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGatheringStats returns a new GatheringStats
func NewGatheringStats() *GatheringStats {
	gs := &GatheringStats{
		PIDParticipant: types.NewPID(0),
		UIFlags:        types.NewPrimitiveU32(0),
		LstValues:      types.NewList[*types.PrimitiveF32](),
	}

	gs.LstValues.Type = types.NewPrimitiveF32(0)

	return gs
}
