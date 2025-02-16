// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationPrepareTradePokemonResult is a type within the DataStore protocol
type GlobalTradeStationPrepareTradePokemonResult struct {
	types.Structure
	Result          *GlobalTradeStationDownloadPokemonResult
	PrepareTradeKey *GlobalTradeStationRecordKey
}

// WriteTo writes the GlobalTradeStationPrepareTradePokemonResult to the given writable
func (gtsptpr *GlobalTradeStationPrepareTradePokemonResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsptpr.Result.WriteTo(contentWritable)
	gtsptpr.PrepareTradeKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsptpr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationPrepareTradePokemonResult from the given readable
func (gtsptpr *GlobalTradeStationPrepareTradePokemonResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsptpr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonResult header. %s", err.Error())
	}

	err = gtsptpr.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonResult.Result. %s", err.Error())
	}

	err = gtsptpr.PrepareTradeKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonResult.PrepareTradeKey. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationPrepareTradePokemonResult
func (gtsptpr *GlobalTradeStationPrepareTradePokemonResult) Copy() types.RVType {
	copied := NewGlobalTradeStationPrepareTradePokemonResult()

	copied.StructureVersion = gtsptpr.StructureVersion
	copied.Result = gtsptpr.Result.Copy().(*GlobalTradeStationDownloadPokemonResult)
	copied.PrepareTradeKey = gtsptpr.PrepareTradeKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the given GlobalTradeStationPrepareTradePokemonResult contains the same data as the current GlobalTradeStationPrepareTradePokemonResult
func (gtsptpr *GlobalTradeStationPrepareTradePokemonResult) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationPrepareTradePokemonResult); !ok {
		return false
	}

	other := o.(*GlobalTradeStationPrepareTradePokemonResult)

	if gtsptpr.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsptpr.Result.Equals(other.Result) {
		return false
	}

	return gtsptpr.PrepareTradeKey.Equals(other.PrepareTradeKey)
}

// String returns the string representation of the GlobalTradeStationPrepareTradePokemonResult
func (gtsptpr *GlobalTradeStationPrepareTradePokemonResult) String() string {
	return gtsptpr.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationPrepareTradePokemonResult using the provided indentation level
func (gtsptpr *GlobalTradeStationPrepareTradePokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationPrepareTradePokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sResult: %s,\n", indentationValues, gtsptpr.Result.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareTradeKey: %s,\n", indentationValues, gtsptpr.PrepareTradeKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationPrepareTradePokemonResult returns a new GlobalTradeStationPrepareTradePokemonResult
func NewGlobalTradeStationPrepareTradePokemonResult() *GlobalTradeStationPrepareTradePokemonResult {
	gtsptpr := &GlobalTradeStationPrepareTradePokemonResult{
		Result:          NewGlobalTradeStationDownloadPokemonResult(),
		PrepareTradeKey: NewGlobalTradeStationRecordKey(),
	}

	return gtsptpr
}
