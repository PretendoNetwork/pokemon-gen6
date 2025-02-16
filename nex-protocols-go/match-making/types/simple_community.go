// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SimpleCommunity is a type within the Matchmaking protocol
type SimpleCommunity struct {
	types.Structure
	GatheringID           *types.PrimitiveU32
	MatchmakeSessionCount *types.PrimitiveU32
}

// WriteTo writes the SimpleCommunity to the given writable
func (sc *SimpleCommunity) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sc.GatheringID.WriteTo(contentWritable)
	sc.MatchmakeSessionCount.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SimpleCommunity from the given readable
func (sc *SimpleCommunity) ExtractFrom(readable types.Readable) error {
	var err error

	err = sc.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleCommunity header. %s", err.Error())
	}

	err = sc.GatheringID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleCommunity.GatheringID. %s", err.Error())
	}

	err = sc.MatchmakeSessionCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleCommunity.MatchmakeSessionCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SimpleCommunity
func (sc *SimpleCommunity) Copy() types.RVType {
	copied := NewSimpleCommunity()

	copied.StructureVersion = sc.StructureVersion
	copied.GatheringID = sc.GatheringID.Copy().(*types.PrimitiveU32)
	copied.MatchmakeSessionCount = sc.MatchmakeSessionCount.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given SimpleCommunity contains the same data as the current SimpleCommunity
func (sc *SimpleCommunity) Equals(o types.RVType) bool {
	if _, ok := o.(*SimpleCommunity); !ok {
		return false
	}

	other := o.(*SimpleCommunity)

	if sc.StructureVersion != other.StructureVersion {
		return false
	}

	if !sc.GatheringID.Equals(other.GatheringID) {
		return false
	}

	return sc.MatchmakeSessionCount.Equals(other.MatchmakeSessionCount)
}

// String returns the string representation of the SimpleCommunity
func (sc *SimpleCommunity) String() string {
	return sc.FormatToString(0)
}

// FormatToString pretty-prints the SimpleCommunity using the provided indentation level
func (sc *SimpleCommunity) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimpleCommunity{\n")
	b.WriteString(fmt.Sprintf("%sGatheringID: %s,\n", indentationValues, sc.GatheringID))
	b.WriteString(fmt.Sprintf("%sMatchmakeSessionCount: %s,\n", indentationValues, sc.MatchmakeSessionCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleCommunity returns a new SimpleCommunity
func NewSimpleCommunity() *SimpleCommunity {
	sc := &SimpleCommunity{
		GatheringID:           types.NewPrimitiveU32(0),
		MatchmakeSessionCount: types.NewPrimitiveU32(0),
	}

	return sc
}
