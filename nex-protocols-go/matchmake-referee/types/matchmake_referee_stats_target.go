// Package types implements all the types used by the MatchmakeReferee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeRefereeStatsTarget is a type within the MatchmakeReferee protocol
type MatchmakeRefereeStatsTarget struct {
	types.Structure
	*types.Data
	PID      *types.PID
	Category *types.PrimitiveU32
}

// WriteTo writes the MatchmakeRefereeStatsTarget to the given writable
func (mrst *MatchmakeRefereeStatsTarget) WriteTo(writable types.Writable) {
	mrst.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	mrst.PID.WriteTo(contentWritable)
	mrst.Category.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mrst.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeStatsTarget from the given readable
func (mrst *MatchmakeRefereeStatsTarget) ExtractFrom(readable types.Readable) error {
	var err error

	err = mrst.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsTarget.Data. %s", err.Error())
	}

	err = mrst.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsTarget header. %s", err.Error())
	}

	err = mrst.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsTarget.PID. %s", err.Error())
	}

	err = mrst.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsTarget.Category. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStatsTarget
func (mrst *MatchmakeRefereeStatsTarget) Copy() types.RVType {
	copied := NewMatchmakeRefereeStatsTarget()

	copied.StructureVersion = mrst.StructureVersion
	copied.Data = mrst.Data.Copy().(*types.Data)
	copied.PID = mrst.PID.Copy().(*types.PID)
	copied.Category = mrst.Category.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given MatchmakeRefereeStatsTarget contains the same data as the current MatchmakeRefereeStatsTarget
func (mrst *MatchmakeRefereeStatsTarget) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereeStatsTarget); !ok {
		return false
	}

	other := o.(*MatchmakeRefereeStatsTarget)

	if mrst.StructureVersion != other.StructureVersion {
		return false
	}

	if !mrst.Data.Equals(other.Data) {
		return false
	}

	if !mrst.PID.Equals(other.PID) {
		return false
	}

	return mrst.Category.Equals(other.Category)
}

// String returns the string representation of the MatchmakeRefereeStatsTarget
func (mrst *MatchmakeRefereeStatsTarget) String() string {
	return mrst.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeRefereeStatsTarget using the provided indentation level
func (mrst *MatchmakeRefereeStatsTarget) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeStatsTarget{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, mrst.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, mrst.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, mrst.Category))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeStatsTarget returns a new MatchmakeRefereeStatsTarget
func NewMatchmakeRefereeStatsTarget() *MatchmakeRefereeStatsTarget {
	mrst := &MatchmakeRefereeStatsTarget{
		Data:     types.NewData(),
		PID:      types.NewPID(0),
		Category: types.NewPrimitiveU32(0),
	}

	return mrst
}
