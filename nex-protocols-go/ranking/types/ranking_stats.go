// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RankingStats is a type within the Ranking protocol
type RankingStats struct {
	types.Structure
	StatsList *types.List[*types.PrimitiveF64]
}

// WriteTo writes the RankingStats to the given writable
func (rs *RankingStats) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rs.StatsList.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingStats from the given readable
func (rs *RankingStats) ExtractFrom(readable types.Readable) error {
	var err error

	err = rs.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingStats header. %s", err.Error())
	}

	err = rs.StatsList.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingStats.StatsList. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingStats
func (rs *RankingStats) Copy() types.RVType {
	copied := NewRankingStats()

	copied.StructureVersion = rs.StructureVersion
	copied.StatsList = rs.StatsList.Copy().(*types.List[*types.PrimitiveF64])

	return copied
}

// Equals checks if the given RankingStats contains the same data as the current RankingStats
func (rs *RankingStats) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingStats); !ok {
		return false
	}

	other := o.(*RankingStats)

	if rs.StructureVersion != other.StructureVersion {
		return false
	}

	return rs.StatsList.Equals(other.StatsList)
}

// String returns the string representation of the RankingStats
func (rs *RankingStats) String() string {
	return rs.FormatToString(0)
}

// FormatToString pretty-prints the RankingStats using the provided indentation level
func (rs *RankingStats) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingStats{\n")
	b.WriteString(fmt.Sprintf("%sStatsList: %s,\n", indentationValues, rs.StatsList))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingStats returns a new RankingStats
func NewRankingStats() *RankingStats {
	rs := &RankingStats{
		StatsList: types.NewList[*types.PrimitiveF64](),
	}

	rs.StatsList.Type = types.NewPrimitiveF64(0)

	return rs
}
