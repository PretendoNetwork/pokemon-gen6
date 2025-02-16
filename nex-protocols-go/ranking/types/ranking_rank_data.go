// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RankingRankData is a type within the Ranking protocol
type RankingRankData struct {
	types.Structure
	PrincipalID *types.PID
	UniqueID    *types.PrimitiveU64
	Order       *types.PrimitiveU32
	Category    *types.PrimitiveU32
	Score       *types.PrimitiveU32
	Groups      *types.Buffer
	Param       *types.PrimitiveU64
	CommonData  *types.Buffer
	UpdateTime  *types.DateTime // * NEX v3.6.0
}

// WriteTo writes the RankingRankData to the given writable
func (rrd *RankingRankData) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.Ranking

	contentWritable := writable.CopyNew()

	rrd.PrincipalID.WriteTo(contentWritable)
	rrd.UniqueID.WriteTo(contentWritable)
	rrd.Order.WriteTo(contentWritable)
	rrd.Category.WriteTo(contentWritable)
	rrd.Score.WriteTo(contentWritable)
	rrd.Groups.WriteTo(contentWritable)
	rrd.Param.WriteTo(contentWritable)
	rrd.CommonData.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("3.6.0") {
		rrd.UpdateTime.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rrd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingRankData from the given readable
func (rrd *RankingRankData) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.Ranking

	var err error

	err = rrd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData header. %s", err.Error())
	}

	err = rrd.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.PrincipalID. %s", err.Error())
	}

	err = rrd.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.UniqueID. %s", err.Error())
	}

	err = rrd.Order.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Order. %s", err.Error())
	}

	err = rrd.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Category. %s", err.Error())
	}

	err = rrd.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Score. %s", err.Error())
	}

	err = rrd.Groups.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Groups. %s", err.Error())
	}

	err = rrd.Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Param. %s", err.Error())
	}

	err = rrd.CommonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.CommonData. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.6.0") {
		err = rrd.UpdateTime.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract RankingRankData.UpdateTime. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of RankingRankData
func (rrd *RankingRankData) Copy() types.RVType {
	copied := NewRankingRankData()

	copied.StructureVersion = rrd.StructureVersion
	copied.PrincipalID = rrd.PrincipalID.Copy().(*types.PID)
	copied.UniqueID = rrd.UniqueID.Copy().(*types.PrimitiveU64)
	copied.Order = rrd.Order.Copy().(*types.PrimitiveU32)
	copied.Category = rrd.Category.Copy().(*types.PrimitiveU32)
	copied.Score = rrd.Score.Copy().(*types.PrimitiveU32)
	copied.Groups = rrd.Groups.Copy().(*types.Buffer)
	copied.Param = rrd.Param.Copy().(*types.PrimitiveU64)
	copied.CommonData = rrd.CommonData.Copy().(*types.Buffer)
	copied.UpdateTime = rrd.UpdateTime.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given RankingRankData contains the same data as the current RankingRankData
func (rrd *RankingRankData) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingRankData); !ok {
		return false
	}

	other := o.(*RankingRankData)

	if rrd.StructureVersion != other.StructureVersion {
		return false
	}

	if !rrd.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !rrd.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !rrd.Order.Equals(other.Order) {
		return false
	}

	if !rrd.Category.Equals(other.Category) {
		return false
	}

	if !rrd.Score.Equals(other.Score) {
		return false
	}

	if !rrd.Groups.Equals(other.Groups) {
		return false
	}

	if !rrd.Param.Equals(other.Param) {
		return false
	}

	if !rrd.CommonData.Equals(other.CommonData) {
		return false
	}

	return rrd.UpdateTime.Equals(other.UpdateTime)
}

// String returns the string representation of the RankingRankData
func (rrd *RankingRankData) String() string {
	return rrd.FormatToString(0)
}

// FormatToString pretty-prints the RankingRankData using the provided indentation level
func (rrd *RankingRankData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingRankData{\n")
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, rrd.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, rrd.UniqueID))
	b.WriteString(fmt.Sprintf("%sOrder: %s,\n", indentationValues, rrd.Order))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rrd.Category))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, rrd.Score))
	b.WriteString(fmt.Sprintf("%sGroups: %s,\n", indentationValues, rrd.Groups))
	b.WriteString(fmt.Sprintf("%sParam: %s,\n", indentationValues, rrd.Param))
	b.WriteString(fmt.Sprintf("%sCommonData: %s,\n", indentationValues, rrd.CommonData))
	b.WriteString(fmt.Sprintf("%sUpdateTime: %s,\n", indentationValues, rrd.UpdateTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingRankData returns a new RankingRankData
func NewRankingRankData() *RankingRankData {
	rrd := &RankingRankData{
		PrincipalID: types.NewPID(0),
		UniqueID:    types.NewPrimitiveU64(0),
		Order:       types.NewPrimitiveU32(0),
		Category:    types.NewPrimitiveU32(0),
		Score:       types.NewPrimitiveU32(0),
		Groups:      types.NewBuffer(nil),
		Param:       types.NewPrimitiveU64(0),
		CommonData:  types.NewBuffer(nil),
		UpdateTime:  types.NewDateTime(0),
	}

	return rrd
}
