// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RankingCachedResult is a type within the Ranking protocol
type RankingCachedResult struct {
	types.Structure
	*RankingResult
	CreatedTime *types.DateTime
	ExpiredTime *types.DateTime
	MaxLength   *types.PrimitiveU8
}

// WriteTo writes the RankingCachedResult to the given writable
func (rcr *RankingCachedResult) WriteTo(writable types.Writable) {
	rcr.RankingResult.WriteTo(writable)

	contentWritable := writable.CopyNew()

	rcr.CreatedTime.WriteTo(contentWritable)
	rcr.ExpiredTime.WriteTo(contentWritable)
	rcr.MaxLength.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rcr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingCachedResult from the given readable
func (rcr *RankingCachedResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = rcr.RankingResult.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult.RankingResult. %s", err.Error())
	}

	err = rcr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult header. %s", err.Error())
	}

	err = rcr.CreatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult.CreatedTime. %s", err.Error())
	}

	err = rcr.ExpiredTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult.ExpiredTime. %s", err.Error())
	}

	err = rcr.MaxLength.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult.MaxLength. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingCachedResult
func (rcr *RankingCachedResult) Copy() types.RVType {
	copied := NewRankingCachedResult()

	copied.StructureVersion = rcr.StructureVersion
	copied.RankingResult = rcr.RankingResult.Copy().(*RankingResult)
	copied.CreatedTime = rcr.CreatedTime.Copy().(*types.DateTime)
	copied.ExpiredTime = rcr.ExpiredTime.Copy().(*types.DateTime)
	copied.MaxLength = rcr.MaxLength.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given RankingCachedResult contains the same data as the current RankingCachedResult
func (rcr *RankingCachedResult) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingCachedResult); !ok {
		return false
	}

	other := o.(*RankingCachedResult)

	if rcr.StructureVersion != other.StructureVersion {
		return false
	}

	if !rcr.RankingResult.Equals(other.RankingResult) {
		return false
	}

	if !rcr.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	if !rcr.ExpiredTime.Equals(other.ExpiredTime) {
		return false
	}

	return rcr.MaxLength.Equals(other.MaxLength)
}

// String returns the string representation of the RankingCachedResult
func (rcr *RankingCachedResult) String() string {
	return rcr.FormatToString(0)
}

// FormatToString pretty-prints the RankingCachedResult using the provided indentation level
func (rcr *RankingCachedResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingCachedResult{\n")
	b.WriteString(fmt.Sprintf("%sRankingResult (parent): %s,\n", indentationValues, rcr.RankingResult.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCreatedTime: %s,\n", indentationValues, rcr.CreatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sExpiredTime: %s,\n", indentationValues, rcr.ExpiredTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMaxLength: %s,\n", indentationValues, rcr.MaxLength))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingCachedResult returns a new RankingCachedResult
func NewRankingCachedResult() *RankingCachedResult {
	rcr := &RankingCachedResult{
		RankingResult: NewRankingResult(),
		CreatedTime:   types.NewDateTime(0),
		ExpiredTime:   types.NewDateTime(0),
		MaxLength:     types.NewPrimitiveU8(0),
	}

	return rcr
}
