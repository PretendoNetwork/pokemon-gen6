// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationDownloadPokemonResult is a type within the DataStore protocol
type GlobalTradeStationDownloadPokemonResult struct {
	types.Structure
	DataID      *types.PrimitiveU64
	IndexData   *types.QBuffer
	PokemonData *types.QBuffer
}

// WriteTo writes the GlobalTradeStationDownloadPokemonResult to the given writable
func (gtsdpr *GlobalTradeStationDownloadPokemonResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsdpr.DataID.WriteTo(contentWritable)
	gtsdpr.IndexData.WriteTo(contentWritable)
	gtsdpr.PokemonData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsdpr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationDownloadPokemonResult from the given readable
func (gtsdpr *GlobalTradeStationDownloadPokemonResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsdpr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadPokemonResult header. %s", err.Error())
	}

	err = gtsdpr.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadPokemonResult.DataID. %s", err.Error())
	}

	err = gtsdpr.IndexData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadPokemonResult.IndexData. %s", err.Error())
	}

	err = gtsdpr.PokemonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadPokemonResult.PokemonData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationDownloadPokemonResult
func (gtsdpr *GlobalTradeStationDownloadPokemonResult) Copy() types.RVType {
	copied := NewGlobalTradeStationDownloadPokemonResult()

	copied.StructureVersion = gtsdpr.StructureVersion
	copied.DataID = gtsdpr.DataID.Copy().(*types.PrimitiveU64)
	copied.IndexData = gtsdpr.IndexData.Copy().(*types.QBuffer)
	copied.PokemonData = gtsdpr.PokemonData.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given GlobalTradeStationDownloadPokemonResult contains the same data as the current GlobalTradeStationDownloadPokemonResult
func (gtsdpr *GlobalTradeStationDownloadPokemonResult) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationDownloadPokemonResult); !ok {
		return false
	}

	other := o.(*GlobalTradeStationDownloadPokemonResult)

	if gtsdpr.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsdpr.DataID.Equals(other.DataID) {
		return false
	}

	if !gtsdpr.IndexData.Equals(other.IndexData) {
		return false
	}

	return gtsdpr.PokemonData.Equals(other.PokemonData)
}

// String returns the string representation of the GlobalTradeStationDownloadPokemonResult
func (gtsdpr *GlobalTradeStationDownloadPokemonResult) String() string {
	return gtsdpr.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationDownloadPokemonResult using the provided indentation level
func (gtsdpr *GlobalTradeStationDownloadPokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDownloadPokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, gtsdpr.DataID))
	b.WriteString(fmt.Sprintf("%sIndexData: %s,\n", indentationValues, gtsdpr.IndexData))
	b.WriteString(fmt.Sprintf("%sPokemonData: %s,\n", indentationValues, gtsdpr.PokemonData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadPokemonResult returns a new GlobalTradeStationDownloadPokemonResult
func NewGlobalTradeStationDownloadPokemonResult() *GlobalTradeStationDownloadPokemonResult {
	gtsdpr := &GlobalTradeStationDownloadPokemonResult{
		DataID:      types.NewPrimitiveU64(0),
		IndexData:   types.NewQBuffer(nil),
		PokemonData: types.NewQBuffer(nil),
	}

	return gtsdpr
}
