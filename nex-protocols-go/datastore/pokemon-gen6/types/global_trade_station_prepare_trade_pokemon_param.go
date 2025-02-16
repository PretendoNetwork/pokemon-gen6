// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationPrepareTradePokemonParam is a type within the DataStore protocol
type GlobalTradeStationPrepareTradePokemonParam struct {
	types.Structure
	TradeKey         *GlobalTradeStationTradeKey
	PrepareUploadKey *GlobalTradeStationRecordKey
}

// WriteTo writes the GlobalTradeStationPrepareTradePokemonParam to the given writable
func (gtsptpp *GlobalTradeStationPrepareTradePokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsptpp.TradeKey.WriteTo(contentWritable)
	gtsptpp.PrepareUploadKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsptpp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationPrepareTradePokemonParam from the given readable
func (gtsptpp *GlobalTradeStationPrepareTradePokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsptpp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonParam header. %s", err.Error())
	}

	err = gtsptpp.TradeKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonParam.TradeKey. %s", err.Error())
	}

	err = gtsptpp.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonParam.PrepareUploadKey. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationPrepareTradePokemonParam
func (gtsptpp *GlobalTradeStationPrepareTradePokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationPrepareTradePokemonParam()

	copied.StructureVersion = gtsptpp.StructureVersion
	copied.TradeKey = gtsptpp.TradeKey.Copy().(*GlobalTradeStationTradeKey)
	copied.PrepareUploadKey = gtsptpp.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the given GlobalTradeStationPrepareTradePokemonParam contains the same data as the current GlobalTradeStationPrepareTradePokemonParam
func (gtsptpp *GlobalTradeStationPrepareTradePokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationPrepareTradePokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationPrepareTradePokemonParam)

	if gtsptpp.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsptpp.TradeKey.Equals(other.TradeKey) {
		return false
	}

	return gtsptpp.PrepareUploadKey.Equals(other.PrepareUploadKey)
}

// String returns the string representation of the GlobalTradeStationPrepareTradePokemonParam
func (gtsptpp *GlobalTradeStationPrepareTradePokemonParam) String() string {
	return gtsptpp.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationPrepareTradePokemonParam using the provided indentation level
func (gtsptpp *GlobalTradeStationPrepareTradePokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationPrepareTradePokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sTradeKey: %s,\n", indentationValues, gtsptpp.TradeKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s,\n", indentationValues, gtsptpp.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationPrepareTradePokemonParam returns a new GlobalTradeStationPrepareTradePokemonParam
func NewGlobalTradeStationPrepareTradePokemonParam() *GlobalTradeStationPrepareTradePokemonParam {
	gtsptpp := &GlobalTradeStationPrepareTradePokemonParam{
		TradeKey:         NewGlobalTradeStationTradeKey(),
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
	}

	return gtsptpp
}
