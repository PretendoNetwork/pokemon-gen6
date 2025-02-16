// Package types implements all the types used by the MatchmakeReferee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeRefereeStatsInitParam is a type within the MatchmakeReferee protocol
type MatchmakeRefereeStatsInitParam struct {
	types.Structure
	*types.Data
	Category           *types.PrimitiveU32
	InitialRatingValue *types.PrimitiveU32
}

// WriteTo writes the MatchmakeRefereeStatsInitParam to the given writable
func (mrsip *MatchmakeRefereeStatsInitParam) WriteTo(writable types.Writable) {
	mrsip.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	mrsip.Category.WriteTo(contentWritable)
	mrsip.InitialRatingValue.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mrsip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeStatsInitParam from the given readable
func (mrsip *MatchmakeRefereeStatsInitParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = mrsip.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsInitParam.Data. %s", err.Error())
	}

	err = mrsip.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsInitParam header. %s", err.Error())
	}

	err = mrsip.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsInitParam.Category. %s", err.Error())
	}

	err = mrsip.InitialRatingValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsInitParam.InitialRatingValue. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStatsInitParam
func (mrsip *MatchmakeRefereeStatsInitParam) Copy() types.RVType {
	copied := NewMatchmakeRefereeStatsInitParam()

	copied.StructureVersion = mrsip.StructureVersion
	copied.Data = mrsip.Data.Copy().(*types.Data)
	copied.Category = mrsip.Category.Copy().(*types.PrimitiveU32)
	copied.InitialRatingValue = mrsip.InitialRatingValue.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given MatchmakeRefereeStatsInitParam contains the same data as the current MatchmakeRefereeStatsInitParam
func (mrsip *MatchmakeRefereeStatsInitParam) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereeStatsInitParam); !ok {
		return false
	}

	other := o.(*MatchmakeRefereeStatsInitParam)

	if mrsip.StructureVersion != other.StructureVersion {
		return false
	}

	if !mrsip.Data.Equals(other.Data) {
		return false
	}

	if !mrsip.Category.Equals(other.Category) {
		return false
	}

	return mrsip.InitialRatingValue.Equals(other.InitialRatingValue)
}

// String returns the string representation of the MatchmakeRefereeStatsInitParam
func (mrsip *MatchmakeRefereeStatsInitParam) String() string {
	return mrsip.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeRefereeStatsInitParam using the provided indentation level
func (mrsip *MatchmakeRefereeStatsInitParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeStatsInitParam{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, mrsip.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, mrsip.Category))
	b.WriteString(fmt.Sprintf("%sInitialRatingValue: %s,\n", indentationValues, mrsip.InitialRatingValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeStatsInitParam returns a new MatchmakeRefereeStatsInitParam
func NewMatchmakeRefereeStatsInitParam() *MatchmakeRefereeStatsInitParam {
	mrsip := &MatchmakeRefereeStatsInitParam{
		Data:               types.NewData(),
		Category:           types.NewPrimitiveU32(0),
		InitialRatingValue: types.NewPrimitiveU32(0),
	}

	return mrsip
}
