// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2RankData is a type within the Ranking2 protocol
type Ranking2RankData struct {
	types.Structure
	Misc        *types.PrimitiveU64
	NexUniqueID *types.PrimitiveU64
	PrincipalID *types.PID
	Rank        *types.PrimitiveU32
	Score       *types.PrimitiveU32
	CommonData  *Ranking2CommonData
}

// WriteTo writes the Ranking2RankData to the given writable
func (rrd *Ranking2RankData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rrd.Misc.WriteTo(contentWritable)
	rrd.NexUniqueID.WriteTo(contentWritable)
	rrd.PrincipalID.WriteTo(contentWritable)
	rrd.Rank.WriteTo(contentWritable)
	rrd.Score.WriteTo(contentWritable)
	rrd.CommonData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rrd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2RankData from the given readable
func (rrd *Ranking2RankData) ExtractFrom(readable types.Readable) error {
	var err error

	err = rrd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData header. %s", err.Error())
	}

	err = rrd.Misc.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.Misc. %s", err.Error())
	}

	err = rrd.NexUniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.NexUniqueID. %s", err.Error())
	}

	err = rrd.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.PrincipalID. %s", err.Error())
	}

	err = rrd.Rank.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.Rank. %s", err.Error())
	}

	err = rrd.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.Score. %s", err.Error())
	}

	err = rrd.CommonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.CommonData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2RankData
func (rrd *Ranking2RankData) Copy() types.RVType {
	copied := NewRanking2RankData()

	copied.StructureVersion = rrd.StructureVersion
	copied.Misc = rrd.Misc.Copy().(*types.PrimitiveU64)
	copied.NexUniqueID = rrd.NexUniqueID.Copy().(*types.PrimitiveU64)
	copied.PrincipalID = rrd.PrincipalID.Copy().(*types.PID)
	copied.Rank = rrd.Rank.Copy().(*types.PrimitiveU32)
	copied.Score = rrd.Score.Copy().(*types.PrimitiveU32)
	copied.CommonData = rrd.CommonData.Copy().(*Ranking2CommonData)

	return copied
}

// Equals checks if the given Ranking2RankData contains the same data as the current Ranking2RankData
func (rrd *Ranking2RankData) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2RankData); !ok {
		return false
	}

	other := o.(*Ranking2RankData)

	if rrd.StructureVersion != other.StructureVersion {
		return false
	}

	if !rrd.Misc.Equals(other.Misc) {
		return false
	}

	if !rrd.NexUniqueID.Equals(other.NexUniqueID) {
		return false
	}

	if !rrd.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !rrd.Rank.Equals(other.Rank) {
		return false
	}

	if !rrd.Score.Equals(other.Score) {
		return false
	}

	return rrd.CommonData.Equals(other.CommonData)
}

// String returns the string representation of the Ranking2RankData
func (rrd *Ranking2RankData) String() string {
	return rrd.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2RankData using the provided indentation level
func (rrd *Ranking2RankData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2RankData{\n")
	b.WriteString(fmt.Sprintf("%sMisc: %s,\n", indentationValues, rrd.Misc))
	b.WriteString(fmt.Sprintf("%sNexUniqueID: %s,\n", indentationValues, rrd.NexUniqueID))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, rrd.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRank: %s,\n", indentationValues, rrd.Rank))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, rrd.Score))
	b.WriteString(fmt.Sprintf("%sCommonData: %s,\n", indentationValues, rrd.CommonData.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2RankData returns a new Ranking2RankData
func NewRanking2RankData() *Ranking2RankData {
	rrd := &Ranking2RankData{
		Misc:        types.NewPrimitiveU64(0),
		NexUniqueID: types.NewPrimitiveU64(0),
		PrincipalID: types.NewPID(0),
		Rank:        types.NewPrimitiveU32(0),
		Score:       types.NewPrimitiveU32(0),
		CommonData:  NewRanking2CommonData(),
	}

	return rrd
}
