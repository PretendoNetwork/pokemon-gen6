// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// CompetitionRankingInfo is a type within the Ranking protocol
type CompetitionRankingInfo struct {
	types.Structure
	Unknown  *types.PrimitiveU32
	Unknown2 *types.PrimitiveU32
	Unknown3 *types.List[*types.PrimitiveU32]
}

// WriteTo writes the CompetitionRankingInfo to the given writable
func (cri *CompetitionRankingInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	cri.Unknown.WriteTo(contentWritable)
	cri.Unknown2.WriteTo(contentWritable)
	cri.Unknown3.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	cri.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the CompetitionRankingInfo from the given readable
func (cri *CompetitionRankingInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = cri.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfo header. %s", err.Error())
	}

	err = cri.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfo.Unknown. %s", err.Error())
	}

	err = cri.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfo.Unknown2. %s", err.Error())
	}

	err = cri.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfo.Unknown3. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CompetitionRankingInfo
func (cri *CompetitionRankingInfo) Copy() types.RVType {
	copied := NewCompetitionRankingInfo()

	copied.StructureVersion = cri.StructureVersion
	copied.Unknown = cri.Unknown.Copy().(*types.PrimitiveU32)
	copied.Unknown2 = cri.Unknown2.Copy().(*types.PrimitiveU32)
	copied.Unknown3 = cri.Unknown3.Copy().(*types.List[*types.PrimitiveU32])

	return copied
}

// Equals checks if the given CompetitionRankingInfo contains the same data as the current CompetitionRankingInfo
func (cri *CompetitionRankingInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*CompetitionRankingInfo); !ok {
		return false
	}

	other := o.(*CompetitionRankingInfo)

	if cri.StructureVersion != other.StructureVersion {
		return false
	}

	if !cri.Unknown.Equals(other.Unknown) {
		return false
	}

	if !cri.Unknown2.Equals(other.Unknown2) {
		return false
	}

	return cri.Unknown3.Equals(other.Unknown3)
}

// String returns the string representation of the CompetitionRankingInfo
func (cri *CompetitionRankingInfo) String() string {
	return cri.FormatToString(0)
}

// FormatToString pretty-prints the CompetitionRankingInfo using the provided indentation level
func (cri *CompetitionRankingInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CompetitionRankingInfo{\n")
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, cri.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, cri.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, cri.Unknown3))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewCompetitionRankingInfo returns a new CompetitionRankingInfo
func NewCompetitionRankingInfo() *CompetitionRankingInfo {
	cri := &CompetitionRankingInfo{
		Unknown:  types.NewPrimitiveU32(0),
		Unknown2: types.NewPrimitiveU32(0),
		Unknown3: types.NewList[*types.PrimitiveU32](),
	}

	cri.Unknown3.Type = types.NewPrimitiveU32(0)

	return cri
}
