// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PersistentGathering is a type within the Matchmaking protocol
type PersistentGathering struct {
	types.Structure
	*Gathering
	CommunityType          *types.PrimitiveU32
	Password               *types.String
	Attribs                *types.List[*types.PrimitiveU32]
	ApplicationBuffer      *types.Buffer
	ParticipationStartDate *types.DateTime
	ParticipationEndDate   *types.DateTime
	MatchmakeSessionCount  *types.PrimitiveU32
	ParticipationCount     *types.PrimitiveU32
}

// WriteTo writes the PersistentGathering to the given writable
func (pg *PersistentGathering) WriteTo(writable types.Writable) {
	pg.Gathering.WriteTo(writable)

	contentWritable := writable.CopyNew()

	pg.CommunityType.WriteTo(contentWritable)
	pg.Password.WriteTo(contentWritable)
	pg.Attribs.WriteTo(contentWritable)
	pg.ApplicationBuffer.WriteTo(contentWritable)
	pg.ParticipationStartDate.WriteTo(contentWritable)
	pg.ParticipationEndDate.WriteTo(contentWritable)
	pg.MatchmakeSessionCount.WriteTo(contentWritable)
	pg.ParticipationCount.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pg.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PersistentGathering from the given readable
func (pg *PersistentGathering) ExtractFrom(readable types.Readable) error {
	var err error

	err = pg.Gathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.Gathering. %s", err.Error())
	}

	err = pg.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering header. %s", err.Error())
	}

	err = pg.CommunityType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.CommunityType. %s", err.Error())
	}

	err = pg.Password.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.Password. %s", err.Error())
	}

	err = pg.Attribs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.Attribs. %s", err.Error())
	}

	err = pg.ApplicationBuffer.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.ApplicationBuffer. %s", err.Error())
	}

	err = pg.ParticipationStartDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.ParticipationStartDate. %s", err.Error())
	}

	err = pg.ParticipationEndDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.ParticipationEndDate. %s", err.Error())
	}

	err = pg.MatchmakeSessionCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.MatchmakeSessionCount. %s", err.Error())
	}

	err = pg.ParticipationCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentGathering.ParticipationCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PersistentGathering
func (pg *PersistentGathering) Copy() types.RVType {
	copied := NewPersistentGathering()

	copied.StructureVersion = pg.StructureVersion
	copied.Gathering = pg.Gathering.Copy().(*Gathering)
	copied.CommunityType = pg.CommunityType.Copy().(*types.PrimitiveU32)
	copied.Password = pg.Password.Copy().(*types.String)
	copied.Attribs = pg.Attribs.Copy().(*types.List[*types.PrimitiveU32])
	copied.ApplicationBuffer = pg.ApplicationBuffer.Copy().(*types.Buffer)
	copied.ParticipationStartDate = pg.ParticipationStartDate.Copy().(*types.DateTime)
	copied.ParticipationEndDate = pg.ParticipationEndDate.Copy().(*types.DateTime)
	copied.MatchmakeSessionCount = pg.MatchmakeSessionCount.Copy().(*types.PrimitiveU32)
	copied.ParticipationCount = pg.ParticipationCount.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given PersistentGathering contains the same data as the current PersistentGathering
func (pg *PersistentGathering) Equals(o types.RVType) bool {
	if _, ok := o.(*PersistentGathering); !ok {
		return false
	}

	other := o.(*PersistentGathering)

	if pg.StructureVersion != other.StructureVersion {
		return false
	}

	if !pg.Gathering.Equals(other.Gathering) {
		return false
	}

	if !pg.CommunityType.Equals(other.CommunityType) {
		return false
	}

	if !pg.Password.Equals(other.Password) {
		return false
	}

	if !pg.Attribs.Equals(other.Attribs) {
		return false
	}

	if !pg.ApplicationBuffer.Equals(other.ApplicationBuffer) {
		return false
	}

	if !pg.ParticipationStartDate.Equals(other.ParticipationStartDate) {
		return false
	}

	if !pg.ParticipationEndDate.Equals(other.ParticipationEndDate) {
		return false
	}

	if !pg.MatchmakeSessionCount.Equals(other.MatchmakeSessionCount) {
		return false
	}

	return pg.ParticipationCount.Equals(other.ParticipationCount)
}

// String returns the string representation of the PersistentGathering
func (pg *PersistentGathering) String() string {
	return pg.FormatToString(0)
}

// FormatToString pretty-prints the PersistentGathering using the provided indentation level
func (pg *PersistentGathering) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PersistentGathering{\n")
	b.WriteString(fmt.Sprintf("%sGathering (parent): %s,\n", indentationValues, pg.Gathering.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCommunityType: %s,\n", indentationValues, pg.CommunityType))
	b.WriteString(fmt.Sprintf("%sPassword: %s,\n", indentationValues, pg.Password))
	b.WriteString(fmt.Sprintf("%sAttribs: %s,\n", indentationValues, pg.Attribs))
	b.WriteString(fmt.Sprintf("%sApplicationBuffer: %s,\n", indentationValues, pg.ApplicationBuffer))
	b.WriteString(fmt.Sprintf("%sParticipationStartDate: %s,\n", indentationValues, pg.ParticipationStartDate.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sParticipationEndDate: %s,\n", indentationValues, pg.ParticipationEndDate.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMatchmakeSessionCount: %s,\n", indentationValues, pg.MatchmakeSessionCount))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %s,\n", indentationValues, pg.ParticipationCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPersistentGathering returns a new PersistentGathering
func NewPersistentGathering() *PersistentGathering {
	pg := &PersistentGathering{
		Gathering:              NewGathering(),
		CommunityType:          types.NewPrimitiveU32(0),
		Password:               types.NewString(""),
		Attribs:                types.NewList[*types.PrimitiveU32](),
		ApplicationBuffer:      types.NewBuffer(nil),
		ParticipationStartDate: types.NewDateTime(0),
		ParticipationEndDate:   types.NewDateTime(0),
		MatchmakeSessionCount:  types.NewPrimitiveU32(0),
		ParticipationCount:     types.NewPrimitiveU32(0),
	}

	pg.Attribs.Type = types.NewPrimitiveU32(0)

	return pg
}
