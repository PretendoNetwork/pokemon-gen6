// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationTradePokemonResult is a type within the DataStore protocol
type GlobalTradeStationTradePokemonResult struct {
	types.Structure
	Result   *GlobalTradeStationDownloadPokemonResult
	MyDataID *types.PrimitiveU64
}

// WriteTo writes the GlobalTradeStationTradePokemonResult to the given writable
func (gtstpr *GlobalTradeStationTradePokemonResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtstpr.Result.WriteTo(contentWritable)
	gtstpr.MyDataID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtstpr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationTradePokemonResult from the given readable
func (gtstpr *GlobalTradeStationTradePokemonResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtstpr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonResult header. %s", err.Error())
	}

	err = gtstpr.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonResult.Result. %s", err.Error())
	}

	err = gtstpr.MyDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonResult.MyDataID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationTradePokemonResult
func (gtstpr *GlobalTradeStationTradePokemonResult) Copy() types.RVType {
	copied := NewGlobalTradeStationTradePokemonResult()

	copied.StructureVersion = gtstpr.StructureVersion
	copied.Result = gtstpr.Result.Copy().(*GlobalTradeStationDownloadPokemonResult)
	copied.MyDataID = gtstpr.MyDataID.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given GlobalTradeStationTradePokemonResult contains the same data as the current GlobalTradeStationTradePokemonResult
func (gtstpr *GlobalTradeStationTradePokemonResult) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationTradePokemonResult); !ok {
		return false
	}

	other := o.(*GlobalTradeStationTradePokemonResult)

	if gtstpr.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtstpr.Result.Equals(other.Result) {
		return false
	}

	return gtstpr.MyDataID.Equals(other.MyDataID)
}

// String returns the string representation of the GlobalTradeStationTradePokemonResult
func (gtstpr *GlobalTradeStationTradePokemonResult) String() string {
	return gtstpr.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationTradePokemonResult using the provided indentation level
func (gtstpr *GlobalTradeStationTradePokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationTradePokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sResult: %s,\n", indentationValues, gtstpr.Result.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMyDataID: %s,\n", indentationValues, gtstpr.MyDataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationTradePokemonResult returns a new GlobalTradeStationTradePokemonResult
func NewGlobalTradeStationTradePokemonResult() *GlobalTradeStationTradePokemonResult {
	gtstpr := &GlobalTradeStationTradePokemonResult{
		Result:   NewGlobalTradeStationDownloadPokemonResult(),
		MyDataID: types.NewPrimitiveU64(0),
	}

	return gtstpr
}
