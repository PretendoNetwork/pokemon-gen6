// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationSearchPokemonResult is a type within the DataStore protocol
type GlobalTradeStationSearchPokemonResult struct {
	types.Structure
	TotalCount     *types.PrimitiveU32
	Result         *types.List[*GlobalTradeStationData]
	TotalCountType *types.PrimitiveU8
}

// WriteTo writes the GlobalTradeStationSearchPokemonResult to the given writable
func (gtsspr *GlobalTradeStationSearchPokemonResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsspr.TotalCount.WriteTo(contentWritable)
	gtsspr.Result.WriteTo(contentWritable)
	gtsspr.TotalCountType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsspr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationSearchPokemonResult from the given readable
func (gtsspr *GlobalTradeStationSearchPokemonResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsspr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonResult header. %s", err.Error())
	}

	err = gtsspr.TotalCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonResult.TotalCount. %s", err.Error())
	}

	err = gtsspr.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonResult.Result. %s", err.Error())
	}

	err = gtsspr.TotalCountType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonResult.TotalCountType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationSearchPokemonResult
func (gtsspr *GlobalTradeStationSearchPokemonResult) Copy() types.RVType {
	copied := NewGlobalTradeStationSearchPokemonResult()

	copied.StructureVersion = gtsspr.StructureVersion
	copied.TotalCount = gtsspr.TotalCount.Copy().(*types.PrimitiveU32)
	copied.Result = gtsspr.Result.Copy().(*types.List[*GlobalTradeStationData])
	copied.TotalCountType = gtsspr.TotalCountType.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given GlobalTradeStationSearchPokemonResult contains the same data as the current GlobalTradeStationSearchPokemonResult
func (gtsspr *GlobalTradeStationSearchPokemonResult) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationSearchPokemonResult); !ok {
		return false
	}

	other := o.(*GlobalTradeStationSearchPokemonResult)

	if gtsspr.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsspr.TotalCount.Equals(other.TotalCount) {
		return false
	}

	if !gtsspr.Result.Equals(other.Result) {
		return false
	}

	return gtsspr.TotalCountType.Equals(other.TotalCountType)
}

// String returns the string representation of the GlobalTradeStationSearchPokemonResult
func (gtsspr *GlobalTradeStationSearchPokemonResult) String() string {
	return gtsspr.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationSearchPokemonResult using the provided indentation level
func (gtsspr *GlobalTradeStationSearchPokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationSearchPokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sTotalCount: %s,\n", indentationValues, gtsspr.TotalCount))
	b.WriteString(fmt.Sprintf("%sResult: %s,\n", indentationValues, gtsspr.Result))
	b.WriteString(fmt.Sprintf("%sTotalCountType: %s,\n", indentationValues, gtsspr.TotalCountType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationSearchPokemonResult returns a new GlobalTradeStationSearchPokemonResult
func NewGlobalTradeStationSearchPokemonResult() *GlobalTradeStationSearchPokemonResult {
	gtsspr := &GlobalTradeStationSearchPokemonResult{
		TotalCount:     types.NewPrimitiveU32(0),
		Result:         types.NewList[*GlobalTradeStationData](),
		TotalCountType: types.NewPrimitiveU8(0),
	}

	gtsspr.Result.Type = NewGlobalTradeStationData()

	return gtsspr
}
