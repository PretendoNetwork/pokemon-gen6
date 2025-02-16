// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RankingScoreData is a type within the Ranking protocol
type RankingScoreData struct {
	types.Structure
	Category   *types.PrimitiveU32
	Score      *types.PrimitiveU32
	OrderBy    *types.PrimitiveU8
	UpdateMode *types.PrimitiveU8
	Groups     *types.Buffer
	Param      *types.PrimitiveU64
}

// WriteTo writes the RankingScoreData to the given writable
func (rsd *RankingScoreData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rsd.Category.WriteTo(contentWritable)
	rsd.Score.WriteTo(contentWritable)
	rsd.OrderBy.WriteTo(contentWritable)
	rsd.UpdateMode.WriteTo(contentWritable)
	rsd.Groups.WriteTo(contentWritable)
	rsd.Param.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rsd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingScoreData from the given readable
func (rsd *RankingScoreData) ExtractFrom(readable types.Readable) error {
	var err error

	err = rsd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData header. %s", err.Error())
	}

	err = rsd.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Category. %s", err.Error())
	}

	err = rsd.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Score. %s", err.Error())
	}

	err = rsd.OrderBy.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.OrderBy. %s", err.Error())
	}

	err = rsd.UpdateMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.UpdateMode. %s", err.Error())
	}

	err = rsd.Groups.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Groups. %s", err.Error())
	}

	err = rsd.Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Param. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingScoreData
func (rsd *RankingScoreData) Copy() types.RVType {
	copied := NewRankingScoreData()

	copied.StructureVersion = rsd.StructureVersion
	copied.Category = rsd.Category.Copy().(*types.PrimitiveU32)
	copied.Score = rsd.Score.Copy().(*types.PrimitiveU32)
	copied.OrderBy = rsd.OrderBy.Copy().(*types.PrimitiveU8)
	copied.UpdateMode = rsd.UpdateMode.Copy().(*types.PrimitiveU8)
	copied.Groups = rsd.Groups.Copy().(*types.Buffer)
	copied.Param = rsd.Param.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given RankingScoreData contains the same data as the current RankingScoreData
func (rsd *RankingScoreData) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingScoreData); !ok {
		return false
	}

	other := o.(*RankingScoreData)

	if rsd.StructureVersion != other.StructureVersion {
		return false
	}

	if !rsd.Category.Equals(other.Category) {
		return false
	}

	if !rsd.Score.Equals(other.Score) {
		return false
	}

	if !rsd.OrderBy.Equals(other.OrderBy) {
		return false
	}

	if !rsd.UpdateMode.Equals(other.UpdateMode) {
		return false
	}

	if !rsd.Groups.Equals(other.Groups) {
		return false
	}

	return rsd.Param.Equals(other.Param)
}

// String returns the string representation of the RankingScoreData
func (rsd *RankingScoreData) String() string {
	return rsd.FormatToString(0)
}

// FormatToString pretty-prints the RankingScoreData using the provided indentation level
func (rsd *RankingScoreData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingScoreData{\n")
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rsd.Category))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, rsd.Score))
	b.WriteString(fmt.Sprintf("%sOrderBy: %s,\n", indentationValues, rsd.OrderBy))
	b.WriteString(fmt.Sprintf("%sUpdateMode: %s,\n", indentationValues, rsd.UpdateMode))
	b.WriteString(fmt.Sprintf("%sGroups: %s,\n", indentationValues, rsd.Groups))
	b.WriteString(fmt.Sprintf("%sParam: %s,\n", indentationValues, rsd.Param))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingScoreData returns a new RankingScoreData
func NewRankingScoreData() *RankingScoreData {
	rsd := &RankingScoreData{
		Category:   types.NewPrimitiveU32(0),
		Score:      types.NewPrimitiveU32(0),
		OrderBy:    types.NewPrimitiveU8(0),
		UpdateMode: types.NewPrimitiveU8(0),
		Groups:     types.NewBuffer(nil),
		Param:      types.NewPrimitiveU64(0),
	}

	return rsd
}
