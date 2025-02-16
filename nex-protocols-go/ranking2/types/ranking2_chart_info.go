// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2ChartInfo is a type within the Ranking2 protocol
type Ranking2ChartInfo struct {
	types.Structure
	CreateTime           *types.DateTime
	Index                *types.PrimitiveU32
	Category             *types.PrimitiveU32
	Season               *types.PrimitiveS32
	BinsSize             *types.PrimitiveU8
	SamplingRate         *types.PrimitiveU8
	ScoreOrder           *types.PrimitiveBool
	EstimateLength       *types.PrimitiveU32
	EstimateHighestScore *types.PrimitiveU32
	EstimateLowestScore  *types.PrimitiveU32
	EstimateMedianScore  *types.PrimitiveU32
	EstimateAverageScore *types.PrimitiveF64
	HighestBinsScore     *types.PrimitiveU32
	LowestBinsScore      *types.PrimitiveU32
	BinsWidth            *types.PrimitiveU32
	Attribute1           *types.PrimitiveU32
	Attribute2           *types.PrimitiveU32
	Quantities           *types.List[*types.PrimitiveU32]
}

// WriteTo writes the Ranking2ChartInfo to the given writable
func (rci *Ranking2ChartInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rci.CreateTime.WriteTo(contentWritable)
	rci.Index.WriteTo(contentWritable)
	rci.Category.WriteTo(contentWritable)
	rci.Season.WriteTo(contentWritable)
	rci.BinsSize.WriteTo(contentWritable)
	rci.SamplingRate.WriteTo(contentWritable)
	rci.ScoreOrder.WriteTo(contentWritable)
	rci.EstimateLength.WriteTo(contentWritable)
	rci.EstimateHighestScore.WriteTo(contentWritable)
	rci.EstimateLowestScore.WriteTo(contentWritable)
	rci.EstimateMedianScore.WriteTo(contentWritable)
	rci.EstimateAverageScore.WriteTo(contentWritable)
	rci.HighestBinsScore.WriteTo(contentWritable)
	rci.LowestBinsScore.WriteTo(contentWritable)
	rci.BinsWidth.WriteTo(contentWritable)
	rci.Attribute1.WriteTo(contentWritable)
	rci.Attribute2.WriteTo(contentWritable)
	rci.Quantities.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rci.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2ChartInfo from the given readable
func (rci *Ranking2ChartInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = rci.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo header. %s", err.Error())
	}

	err = rci.CreateTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.CreateTime. %s", err.Error())
	}

	err = rci.Index.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Index. %s", err.Error())
	}

	err = rci.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Category. %s", err.Error())
	}

	err = rci.Season.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Season. %s", err.Error())
	}

	err = rci.BinsSize.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.BinsSize. %s", err.Error())
	}

	err = rci.SamplingRate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.SamplingRate. %s", err.Error())
	}

	err = rci.ScoreOrder.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.ScoreOrder. %s", err.Error())
	}

	err = rci.EstimateLength.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateLength. %s", err.Error())
	}

	err = rci.EstimateHighestScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateHighestScore. %s", err.Error())
	}

	err = rci.EstimateLowestScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateLowestScore. %s", err.Error())
	}

	err = rci.EstimateMedianScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateMedianScore. %s", err.Error())
	}

	err = rci.EstimateAverageScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateAverageScore. %s", err.Error())
	}

	err = rci.HighestBinsScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.HighestBinsScore. %s", err.Error())
	}

	err = rci.LowestBinsScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.LowestBinsScore. %s", err.Error())
	}

	err = rci.BinsWidth.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.BinsWidth. %s", err.Error())
	}

	err = rci.Attribute1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Attribute1. %s", err.Error())
	}

	err = rci.Attribute2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Attribute2. %s", err.Error())
	}

	err = rci.Quantities.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Quantities. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2ChartInfo
func (rci *Ranking2ChartInfo) Copy() types.RVType {
	copied := NewRanking2ChartInfo()

	copied.StructureVersion = rci.StructureVersion
	copied.CreateTime = rci.CreateTime.Copy().(*types.DateTime)
	copied.Index = rci.Index.Copy().(*types.PrimitiveU32)
	copied.Category = rci.Category.Copy().(*types.PrimitiveU32)
	copied.Season = rci.Season.Copy().(*types.PrimitiveS32)
	copied.BinsSize = rci.BinsSize.Copy().(*types.PrimitiveU8)
	copied.SamplingRate = rci.SamplingRate.Copy().(*types.PrimitiveU8)
	copied.ScoreOrder = rci.ScoreOrder.Copy().(*types.PrimitiveBool)
	copied.EstimateLength = rci.EstimateLength.Copy().(*types.PrimitiveU32)
	copied.EstimateHighestScore = rci.EstimateHighestScore.Copy().(*types.PrimitiveU32)
	copied.EstimateLowestScore = rci.EstimateLowestScore.Copy().(*types.PrimitiveU32)
	copied.EstimateMedianScore = rci.EstimateMedianScore.Copy().(*types.PrimitiveU32)
	copied.EstimateAverageScore = rci.EstimateAverageScore.Copy().(*types.PrimitiveF64)
	copied.HighestBinsScore = rci.HighestBinsScore.Copy().(*types.PrimitiveU32)
	copied.LowestBinsScore = rci.LowestBinsScore.Copy().(*types.PrimitiveU32)
	copied.BinsWidth = rci.BinsWidth.Copy().(*types.PrimitiveU32)
	copied.Attribute1 = rci.Attribute1.Copy().(*types.PrimitiveU32)
	copied.Attribute2 = rci.Attribute2.Copy().(*types.PrimitiveU32)
	copied.Quantities = rci.Quantities.Copy().(*types.List[*types.PrimitiveU32])

	return copied
}

// Equals checks if the given Ranking2ChartInfo contains the same data as the current Ranking2ChartInfo
func (rci *Ranking2ChartInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2ChartInfo); !ok {
		return false
	}

	other := o.(*Ranking2ChartInfo)

	if rci.StructureVersion != other.StructureVersion {
		return false
	}

	if !rci.CreateTime.Equals(other.CreateTime) {
		return false
	}

	if !rci.Index.Equals(other.Index) {
		return false
	}

	if !rci.Category.Equals(other.Category) {
		return false
	}

	if !rci.Season.Equals(other.Season) {
		return false
	}

	if !rci.BinsSize.Equals(other.BinsSize) {
		return false
	}

	if !rci.SamplingRate.Equals(other.SamplingRate) {
		return false
	}

	if !rci.ScoreOrder.Equals(other.ScoreOrder) {
		return false
	}

	if !rci.EstimateLength.Equals(other.EstimateLength) {
		return false
	}

	if !rci.EstimateHighestScore.Equals(other.EstimateHighestScore) {
		return false
	}

	if !rci.EstimateLowestScore.Equals(other.EstimateLowestScore) {
		return false
	}

	if !rci.EstimateMedianScore.Equals(other.EstimateMedianScore) {
		return false
	}

	if !rci.EstimateAverageScore.Equals(other.EstimateAverageScore) {
		return false
	}

	if !rci.HighestBinsScore.Equals(other.HighestBinsScore) {
		return false
	}

	if !rci.LowestBinsScore.Equals(other.LowestBinsScore) {
		return false
	}

	if !rci.BinsWidth.Equals(other.BinsWidth) {
		return false
	}

	if !rci.Attribute1.Equals(other.Attribute1) {
		return false
	}

	if !rci.Attribute2.Equals(other.Attribute2) {
		return false
	}

	return rci.Quantities.Equals(other.Quantities)
}

// String returns the string representation of the Ranking2ChartInfo
func (rci *Ranking2ChartInfo) String() string {
	return rci.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2ChartInfo using the provided indentation level
func (rci *Ranking2ChartInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2ChartInfo{\n")
	b.WriteString(fmt.Sprintf("%sCreateTime: %s,\n", indentationValues, rci.CreateTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sIndex: %s,\n", indentationValues, rci.Index))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rci.Category))
	b.WriteString(fmt.Sprintf("%sSeason: %s,\n", indentationValues, rci.Season))
	b.WriteString(fmt.Sprintf("%sBinsSize: %s,\n", indentationValues, rci.BinsSize))
	b.WriteString(fmt.Sprintf("%sSamplingRate: %s,\n", indentationValues, rci.SamplingRate))
	b.WriteString(fmt.Sprintf("%sScoreOrder: %s,\n", indentationValues, rci.ScoreOrder))
	b.WriteString(fmt.Sprintf("%sEstimateLength: %s,\n", indentationValues, rci.EstimateLength))
	b.WriteString(fmt.Sprintf("%sEstimateHighestScore: %s,\n", indentationValues, rci.EstimateHighestScore))
	b.WriteString(fmt.Sprintf("%sEstimateLowestScore: %s,\n", indentationValues, rci.EstimateLowestScore))
	b.WriteString(fmt.Sprintf("%sEstimateMedianScore: %s,\n", indentationValues, rci.EstimateMedianScore))
	b.WriteString(fmt.Sprintf("%sEstimateAverageScore: %s,\n", indentationValues, rci.EstimateAverageScore))
	b.WriteString(fmt.Sprintf("%sHighestBinsScore: %s,\n", indentationValues, rci.HighestBinsScore))
	b.WriteString(fmt.Sprintf("%sLowestBinsScore: %s,\n", indentationValues, rci.LowestBinsScore))
	b.WriteString(fmt.Sprintf("%sBinsWidth: %s,\n", indentationValues, rci.BinsWidth))
	b.WriteString(fmt.Sprintf("%sAttribute1: %s,\n", indentationValues, rci.Attribute1))
	b.WriteString(fmt.Sprintf("%sAttribute2: %s,\n", indentationValues, rci.Attribute2))
	b.WriteString(fmt.Sprintf("%sQuantities: %s,\n", indentationValues, rci.Quantities))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2ChartInfo returns a new Ranking2ChartInfo
func NewRanking2ChartInfo() *Ranking2ChartInfo {
	rci := &Ranking2ChartInfo{
		CreateTime:           types.NewDateTime(0),
		Index:                types.NewPrimitiveU32(0),
		Category:             types.NewPrimitiveU32(0),
		Season:               types.NewPrimitiveS32(0),
		BinsSize:             types.NewPrimitiveU8(0),
		SamplingRate:         types.NewPrimitiveU8(0),
		ScoreOrder:           types.NewPrimitiveBool(false),
		EstimateLength:       types.NewPrimitiveU32(0),
		EstimateHighestScore: types.NewPrimitiveU32(0),
		EstimateLowestScore:  types.NewPrimitiveU32(0),
		EstimateMedianScore:  types.NewPrimitiveU32(0),
		EstimateAverageScore: types.NewPrimitiveF64(0),
		HighestBinsScore:     types.NewPrimitiveU32(0),
		LowestBinsScore:      types.NewPrimitiveU32(0),
		BinsWidth:            types.NewPrimitiveU32(0),
		Attribute1:           types.NewPrimitiveU32(0),
		Attribute2:           types.NewPrimitiveU32(0),
		Quantities:           types.NewList[*types.PrimitiveU32](),
	}

	rci.Quantities.Type = types.NewPrimitiveU32(0)

	return rci
}
