// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RankingResult is a type within the Ranking protocol
type RankingResult struct {
	types.Structure
	RankDataList *types.List[*RankingRankData]
	TotalCount   *types.PrimitiveU32
	SinceTime    *types.DateTime
}

// WriteTo writes the RankingResult to the given writable
func (rr *RankingResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rr.RankDataList.WriteTo(contentWritable)
	rr.TotalCount.WriteTo(contentWritable)
	rr.SinceTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingResult from the given readable
func (rr *RankingResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = rr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingResult header. %s", err.Error())
	}

	err = rr.RankDataList.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingResult.RankDataList. %s", err.Error())
	}

	err = rr.TotalCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingResult.TotalCount. %s", err.Error())
	}

	err = rr.SinceTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingResult.SinceTime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingResult
func (rr *RankingResult) Copy() types.RVType {
	copied := NewRankingResult()

	copied.StructureVersion = rr.StructureVersion
	copied.RankDataList = rr.RankDataList.Copy().(*types.List[*RankingRankData])
	copied.TotalCount = rr.TotalCount.Copy().(*types.PrimitiveU32)
	copied.SinceTime = rr.SinceTime.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given RankingResult contains the same data as the current RankingResult
func (rr *RankingResult) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingResult); !ok {
		return false
	}

	other := o.(*RankingResult)

	if rr.StructureVersion != other.StructureVersion {
		return false
	}

	if !rr.RankDataList.Equals(other.RankDataList) {
		return false
	}

	if !rr.TotalCount.Equals(other.TotalCount) {
		return false
	}

	return rr.SinceTime.Equals(other.SinceTime)
}

// String returns the string representation of the RankingResult
func (rr *RankingResult) String() string {
	return rr.FormatToString(0)
}

// FormatToString pretty-prints the RankingResult using the provided indentation level
func (rr *RankingResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingResult{\n")
	b.WriteString(fmt.Sprintf("%sRankDataList: %s,\n", indentationValues, rr.RankDataList))
	b.WriteString(fmt.Sprintf("%sTotalCount: %s,\n", indentationValues, rr.TotalCount))
	b.WriteString(fmt.Sprintf("%sSinceTime: %s,\n", indentationValues, rr.SinceTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingResult returns a new RankingResult
func NewRankingResult() *RankingResult {
	rr := &RankingResult{
		RankDataList: types.NewList[*RankingRankData](),
		TotalCount:   types.NewPrimitiveU32(0),
		SinceTime:    types.NewDateTime(0),
	}

	rr.RankDataList.Type = NewRankingRankData()

	return rr
}
