// Package types implements all the types used by the MatchmakeReferee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeRefereePersonalRoundResult is a type within the MatchmakeReferee protocol
type MatchmakeRefereePersonalRoundResult struct {
	types.Structure
	*types.Data
	PID                     *types.PID
	PersonalRoundResultFlag *types.PrimitiveU32
	RoundWinLoss            *types.PrimitiveU32
	RatingValueChange       *types.PrimitiveS32
	Buffer                  *types.QBuffer
}

// WriteTo writes the MatchmakeRefereePersonalRoundResult to the given writable
func (mrprr *MatchmakeRefereePersonalRoundResult) WriteTo(writable types.Writable) {
	mrprr.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	mrprr.PID.WriteTo(contentWritable)
	mrprr.PersonalRoundResultFlag.WriteTo(contentWritable)
	mrprr.RoundWinLoss.WriteTo(contentWritable)
	mrprr.RatingValueChange.WriteTo(contentWritable)
	mrprr.Buffer.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mrprr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereePersonalRoundResult from the given readable
func (mrprr *MatchmakeRefereePersonalRoundResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = mrprr.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.Data. %s", err.Error())
	}

	err = mrprr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult header. %s", err.Error())
	}

	err = mrprr.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.PID. %s", err.Error())
	}

	err = mrprr.PersonalRoundResultFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.PersonalRoundResultFlag. %s", err.Error())
	}

	err = mrprr.RoundWinLoss.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.RoundWinLoss. %s", err.Error())
	}

	err = mrprr.RatingValueChange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.RatingValueChange. %s", err.Error())
	}

	err = mrprr.Buffer.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.Buffer. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereePersonalRoundResult
func (mrprr *MatchmakeRefereePersonalRoundResult) Copy() types.RVType {
	copied := NewMatchmakeRefereePersonalRoundResult()

	copied.StructureVersion = mrprr.StructureVersion
	copied.Data = mrprr.Data.Copy().(*types.Data)
	copied.PID = mrprr.PID.Copy().(*types.PID)
	copied.PersonalRoundResultFlag = mrprr.PersonalRoundResultFlag.Copy().(*types.PrimitiveU32)
	copied.RoundWinLoss = mrprr.RoundWinLoss.Copy().(*types.PrimitiveU32)
	copied.RatingValueChange = mrprr.RatingValueChange.Copy().(*types.PrimitiveS32)
	copied.Buffer = mrprr.Buffer.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given MatchmakeRefereePersonalRoundResult contains the same data as the current MatchmakeRefereePersonalRoundResult
func (mrprr *MatchmakeRefereePersonalRoundResult) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereePersonalRoundResult); !ok {
		return false
	}

	other := o.(*MatchmakeRefereePersonalRoundResult)

	if mrprr.StructureVersion != other.StructureVersion {
		return false
	}

	if !mrprr.Data.Equals(other.Data) {
		return false
	}

	if !mrprr.PID.Equals(other.PID) {
		return false
	}

	if !mrprr.PersonalRoundResultFlag.Equals(other.PersonalRoundResultFlag) {
		return false
	}

	if !mrprr.RoundWinLoss.Equals(other.RoundWinLoss) {
		return false
	}

	if !mrprr.RatingValueChange.Equals(other.RatingValueChange) {
		return false
	}

	return mrprr.Buffer.Equals(other.Buffer)
}

// String returns the string representation of the MatchmakeRefereePersonalRoundResult
func (mrprr *MatchmakeRefereePersonalRoundResult) String() string {
	return mrprr.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeRefereePersonalRoundResult using the provided indentation level
func (mrprr *MatchmakeRefereePersonalRoundResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereePersonalRoundResult{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, mrprr.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, mrprr.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPersonalRoundResultFlag: %s,\n", indentationValues, mrprr.PersonalRoundResultFlag))
	b.WriteString(fmt.Sprintf("%sRoundWinLoss: %s,\n", indentationValues, mrprr.RoundWinLoss))
	b.WriteString(fmt.Sprintf("%sRatingValueChange: %s,\n", indentationValues, mrprr.RatingValueChange))
	b.WriteString(fmt.Sprintf("%sBuffer: %s,\n", indentationValues, mrprr.Buffer))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereePersonalRoundResult returns a new MatchmakeRefereePersonalRoundResult
func NewMatchmakeRefereePersonalRoundResult() *MatchmakeRefereePersonalRoundResult {
	mrprr := &MatchmakeRefereePersonalRoundResult{
		Data:                    types.NewData(),
		PID:                     types.NewPID(0),
		PersonalRoundResultFlag: types.NewPrimitiveU32(0),
		RoundWinLoss:            types.NewPrimitiveU32(0),
		RatingValueChange:       types.NewPrimitiveS32(0),
		Buffer:                  types.NewQBuffer(nil),
	}

	return mrprr
}
