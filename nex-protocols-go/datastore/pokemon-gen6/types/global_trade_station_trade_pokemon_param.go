// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationTradePokemonParam is a type within the DataStore protocol
type GlobalTradeStationTradePokemonParam struct {
	types.Structure
	TradeKey         *GlobalTradeStationTradeKey
	PrepareTradeKey  *GlobalTradeStationRecordKey
	PrepareUploadKey *GlobalTradeStationRecordKey
	Period           *types.PrimitiveU16
	IndexData        *types.QBuffer
	PokemonData      *types.QBuffer
	Signature        *types.QBuffer
	NeedData         *types.PrimitiveBool
}

// WriteTo writes the GlobalTradeStationTradePokemonParam to the given writable
func (gtstpp *GlobalTradeStationTradePokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtstpp.TradeKey.WriteTo(contentWritable)
	gtstpp.PrepareTradeKey.WriteTo(contentWritable)
	gtstpp.PrepareUploadKey.WriteTo(contentWritable)
	gtstpp.Period.WriteTo(contentWritable)
	gtstpp.IndexData.WriteTo(contentWritable)
	gtstpp.PokemonData.WriteTo(contentWritable)
	gtstpp.Signature.WriteTo(contentWritable)
	gtstpp.NeedData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtstpp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationTradePokemonParam from the given readable
func (gtstpp *GlobalTradeStationTradePokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtstpp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam header. %s", err.Error())
	}

	err = gtstpp.TradeKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.TradeKey. %s", err.Error())
	}

	err = gtstpp.PrepareTradeKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.PrepareTradeKey. %s", err.Error())
	}

	err = gtstpp.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.PrepareUploadKey. %s", err.Error())
	}

	err = gtstpp.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.Period. %s", err.Error())
	}

	err = gtstpp.IndexData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.IndexData. %s", err.Error())
	}

	err = gtstpp.PokemonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.PokemonData. %s", err.Error())
	}

	err = gtstpp.Signature.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.Signature. %s", err.Error())
	}

	err = gtstpp.NeedData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.NeedData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationTradePokemonParam
func (gtstpp *GlobalTradeStationTradePokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationTradePokemonParam()

	copied.StructureVersion = gtstpp.StructureVersion
	copied.TradeKey = gtstpp.TradeKey.Copy().(*GlobalTradeStationTradeKey)
	copied.PrepareTradeKey = gtstpp.PrepareTradeKey.Copy().(*GlobalTradeStationRecordKey)
	copied.PrepareUploadKey = gtstpp.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.Period = gtstpp.Period.Copy().(*types.PrimitiveU16)
	copied.IndexData = gtstpp.IndexData.Copy().(*types.QBuffer)
	copied.PokemonData = gtstpp.PokemonData.Copy().(*types.QBuffer)
	copied.Signature = gtstpp.Signature.Copy().(*types.QBuffer)
	copied.NeedData = gtstpp.NeedData.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the given GlobalTradeStationTradePokemonParam contains the same data as the current GlobalTradeStationTradePokemonParam
func (gtstpp *GlobalTradeStationTradePokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationTradePokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationTradePokemonParam)

	if gtstpp.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtstpp.TradeKey.Equals(other.TradeKey) {
		return false
	}

	if !gtstpp.PrepareTradeKey.Equals(other.PrepareTradeKey) {
		return false
	}

	if !gtstpp.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if !gtstpp.Period.Equals(other.Period) {
		return false
	}

	if !gtstpp.IndexData.Equals(other.IndexData) {
		return false
	}

	if !gtstpp.PokemonData.Equals(other.PokemonData) {
		return false
	}

	if !gtstpp.Signature.Equals(other.Signature) {
		return false
	}

	return gtstpp.NeedData.Equals(other.NeedData)
}

// String returns the string representation of the GlobalTradeStationTradePokemonParam
func (gtstpp *GlobalTradeStationTradePokemonParam) String() string {
	return gtstpp.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationTradePokemonParam using the provided indentation level
func (gtstpp *GlobalTradeStationTradePokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationTradePokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sTradeKey: %s,\n", indentationValues, gtstpp.TradeKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareTradeKey: %s,\n", indentationValues, gtstpp.PrepareTradeKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s,\n", indentationValues, gtstpp.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, gtstpp.Period))
	b.WriteString(fmt.Sprintf("%sIndexData: %s,\n", indentationValues, gtstpp.IndexData))
	b.WriteString(fmt.Sprintf("%sPokemonData: %s,\n", indentationValues, gtstpp.PokemonData))
	b.WriteString(fmt.Sprintf("%sSignature: %s,\n", indentationValues, gtstpp.Signature))
	b.WriteString(fmt.Sprintf("%sNeedData: %s,\n", indentationValues, gtstpp.NeedData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationTradePokemonParam returns a new GlobalTradeStationTradePokemonParam
func NewGlobalTradeStationTradePokemonParam() *GlobalTradeStationTradePokemonParam {
	gtstpp := &GlobalTradeStationTradePokemonParam{
		TradeKey:         NewGlobalTradeStationTradeKey(),
		PrepareTradeKey:  NewGlobalTradeStationRecordKey(),
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
		Period:           types.NewPrimitiveU16(0),
		IndexData:        types.NewQBuffer(nil),
		PokemonData:      types.NewQBuffer(nil),
		Signature:        types.NewQBuffer(nil),
		NeedData:         types.NewPrimitiveBool(false),
	}

	return gtstpp
}
