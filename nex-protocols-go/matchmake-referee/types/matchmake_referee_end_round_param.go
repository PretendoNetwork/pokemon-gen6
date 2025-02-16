// Package types implements all the types used by the MatchmakeReferee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeRefereeEndRoundParam is a type within the MatchmakeReferee protocol
type MatchmakeRefereeEndRoundParam struct {
	types.Structure
	*types.Data
	RoundID              *types.PrimitiveU64
	PersonalRoundResults *types.List[*MatchmakeRefereePersonalRoundResult]
}

// WriteTo writes the MatchmakeRefereeEndRoundParam to the given writable
func (mrerp *MatchmakeRefereeEndRoundParam) WriteTo(writable types.Writable) {
	mrerp.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	mrerp.RoundID.WriteTo(contentWritable)
	mrerp.PersonalRoundResults.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mrerp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeEndRoundParam from the given readable
func (mrerp *MatchmakeRefereeEndRoundParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = mrerp.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeEndRoundParam.Data. %s", err.Error())
	}

	err = mrerp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeEndRoundParam header. %s", err.Error())
	}

	err = mrerp.RoundID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeEndRoundParam.RoundID. %s", err.Error())
	}

	err = mrerp.PersonalRoundResults.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeEndRoundParam.PersonalRoundResults. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeEndRoundParam
func (mrerp *MatchmakeRefereeEndRoundParam) Copy() types.RVType {
	copied := NewMatchmakeRefereeEndRoundParam()

	copied.StructureVersion = mrerp.StructureVersion
	copied.Data = mrerp.Data.Copy().(*types.Data)
	copied.RoundID = mrerp.RoundID.Copy().(*types.PrimitiveU64)
	copied.PersonalRoundResults = mrerp.PersonalRoundResults.Copy().(*types.List[*MatchmakeRefereePersonalRoundResult])

	return copied
}

// Equals checks if the given MatchmakeRefereeEndRoundParam contains the same data as the current MatchmakeRefereeEndRoundParam
func (mrerp *MatchmakeRefereeEndRoundParam) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereeEndRoundParam); !ok {
		return false
	}

	other := o.(*MatchmakeRefereeEndRoundParam)

	if mrerp.StructureVersion != other.StructureVersion {
		return false
	}

	if !mrerp.Data.Equals(other.Data) {
		return false
	}

	if !mrerp.RoundID.Equals(other.RoundID) {
		return false
	}

	return mrerp.PersonalRoundResults.Equals(other.PersonalRoundResults)
}

// String returns the string representation of the MatchmakeRefereeEndRoundParam
func (mrerp *MatchmakeRefereeEndRoundParam) String() string {
	return mrerp.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeRefereeEndRoundParam using the provided indentation level
func (mrerp *MatchmakeRefereeEndRoundParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeEndRoundParam{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, mrerp.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRoundID: %s,\n", indentationValues, mrerp.RoundID))
	b.WriteString(fmt.Sprintf("%sPersonalRoundResults: %s,\n", indentationValues, mrerp.PersonalRoundResults))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeEndRoundParam returns a new MatchmakeRefereeEndRoundParam
func NewMatchmakeRefereeEndRoundParam() *MatchmakeRefereeEndRoundParam {
	mrerp := &MatchmakeRefereeEndRoundParam{
		Data:                 types.NewData(),
		RoundID:              types.NewPrimitiveU64(0),
		PersonalRoundResults: types.NewList[*MatchmakeRefereePersonalRoundResult](),
	}

	mrerp.PersonalRoundResults.Type = NewMatchmakeRefereePersonalRoundResult()

	return mrerp
}
