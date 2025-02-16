// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationSearchPokemonParam is a type within the DataStore protocol
type GlobalTradeStationSearchPokemonParam struct {
	types.Structure
	PrepareUploadKey  *GlobalTradeStationRecordKey
	Conditions        *types.List[*types.PrimitiveU32]
	ResultOrderColumn *types.PrimitiveU8
	ResultOrder       *types.PrimitiveU8
	UploadedAfter     *types.DateTime
	UploadedBefore    *types.DateTime
	ResultRange       *types.ResultRange
}

// WriteTo writes the GlobalTradeStationSearchPokemonParam to the given writable
func (gtsspp *GlobalTradeStationSearchPokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsspp.PrepareUploadKey.WriteTo(contentWritable)
	gtsspp.Conditions.WriteTo(contentWritable)
	gtsspp.ResultOrderColumn.WriteTo(contentWritable)
	gtsspp.ResultOrder.WriteTo(contentWritable)
	gtsspp.UploadedAfter.WriteTo(contentWritable)
	gtsspp.UploadedBefore.WriteTo(contentWritable)
	gtsspp.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsspp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationSearchPokemonParam from the given readable
func (gtsspp *GlobalTradeStationSearchPokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsspp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam header. %s", err.Error())
	}

	err = gtsspp.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.PrepareUploadKey. %s", err.Error())
	}

	err = gtsspp.Conditions.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.Conditions. %s", err.Error())
	}

	err = gtsspp.ResultOrderColumn.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.ResultOrderColumn. %s", err.Error())
	}

	err = gtsspp.ResultOrder.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.ResultOrder. %s", err.Error())
	}

	err = gtsspp.UploadedAfter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.UploadedAfter. %s", err.Error())
	}

	err = gtsspp.UploadedBefore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.UploadedBefore. %s", err.Error())
	}

	err = gtsspp.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.ResultRange. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationSearchPokemonParam
func (gtsspp *GlobalTradeStationSearchPokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationSearchPokemonParam()

	copied.StructureVersion = gtsspp.StructureVersion
	copied.PrepareUploadKey = gtsspp.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.Conditions = gtsspp.Conditions.Copy().(*types.List[*types.PrimitiveU32])
	copied.ResultOrderColumn = gtsspp.ResultOrderColumn.Copy().(*types.PrimitiveU8)
	copied.ResultOrder = gtsspp.ResultOrder.Copy().(*types.PrimitiveU8)
	copied.UploadedAfter = gtsspp.UploadedAfter.Copy().(*types.DateTime)
	copied.UploadedBefore = gtsspp.UploadedBefore.Copy().(*types.DateTime)
	copied.ResultRange = gtsspp.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the given GlobalTradeStationSearchPokemonParam contains the same data as the current GlobalTradeStationSearchPokemonParam
func (gtsspp *GlobalTradeStationSearchPokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationSearchPokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationSearchPokemonParam)

	if gtsspp.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsspp.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if !gtsspp.Conditions.Equals(other.Conditions) {
		return false
	}

	if !gtsspp.ResultOrderColumn.Equals(other.ResultOrderColumn) {
		return false
	}

	if !gtsspp.ResultOrder.Equals(other.ResultOrder) {
		return false
	}

	if !gtsspp.UploadedAfter.Equals(other.UploadedAfter) {
		return false
	}

	if !gtsspp.UploadedBefore.Equals(other.UploadedBefore) {
		return false
	}

	return gtsspp.ResultRange.Equals(other.ResultRange)
}

// String returns the string representation of the GlobalTradeStationSearchPokemonParam
func (gtsspp *GlobalTradeStationSearchPokemonParam) String() string {
	return gtsspp.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationSearchPokemonParam using the provided indentation level
func (gtsspp *GlobalTradeStationSearchPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationSearchPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s,\n", indentationValues, gtsspp.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sConditions: %s,\n", indentationValues, gtsspp.Conditions))
	b.WriteString(fmt.Sprintf("%sResultOrderColumn: %s,\n", indentationValues, gtsspp.ResultOrderColumn))
	b.WriteString(fmt.Sprintf("%sResultOrder: %s,\n", indentationValues, gtsspp.ResultOrder))
	b.WriteString(fmt.Sprintf("%sUploadedAfter: %s,\n", indentationValues, gtsspp.UploadedAfter.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUploadedBefore: %s,\n", indentationValues, gtsspp.UploadedBefore.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, gtsspp.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationSearchPokemonParam returns a new GlobalTradeStationSearchPokemonParam
func NewGlobalTradeStationSearchPokemonParam() *GlobalTradeStationSearchPokemonParam {
	gtsspp := &GlobalTradeStationSearchPokemonParam{
		PrepareUploadKey:  NewGlobalTradeStationRecordKey(),
		Conditions:        types.NewList[*types.PrimitiveU32](),
		ResultOrderColumn: types.NewPrimitiveU8(0),
		ResultOrder:       types.NewPrimitiveU8(0),
		UploadedAfter:     types.NewDateTime(0),
		UploadedBefore:    types.NewDateTime(0),
		ResultRange:       types.NewResultRange(),
	}

	gtsspp.Conditions.Type = types.NewPrimitiveU32(0)

	return gtsspp
}
