// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationUploadPokemonParam is a type within the DataStore protocol
type GlobalTradeStationUploadPokemonParam struct {
	types.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
	Period           *types.PrimitiveU16
	IndexData        *types.QBuffer
	PokemonData      *types.QBuffer
	Signature        *types.QBuffer
}

// WriteTo writes the GlobalTradeStationUploadPokemonParam to the given writable
func (gtsupp *GlobalTradeStationUploadPokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsupp.PrepareUploadKey.WriteTo(contentWritable)
	gtsupp.Period.WriteTo(contentWritable)
	gtsupp.IndexData.WriteTo(contentWritable)
	gtsupp.PokemonData.WriteTo(contentWritable)
	gtsupp.Signature.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsupp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationUploadPokemonParam from the given readable
func (gtsupp *GlobalTradeStationUploadPokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsupp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam header. %s", err.Error())
	}

	err = gtsupp.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.PrepareUploadKey. %s", err.Error())
	}

	err = gtsupp.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.Period. %s", err.Error())
	}

	err = gtsupp.IndexData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.IndexData. %s", err.Error())
	}

	err = gtsupp.PokemonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.PokemonData. %s", err.Error())
	}

	err = gtsupp.Signature.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.Signature. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationUploadPokemonParam
func (gtsupp *GlobalTradeStationUploadPokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationUploadPokemonParam()

	copied.StructureVersion = gtsupp.StructureVersion
	copied.PrepareUploadKey = gtsupp.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.Period = gtsupp.Period.Copy().(*types.PrimitiveU16)
	copied.IndexData = gtsupp.IndexData.Copy().(*types.QBuffer)
	copied.PokemonData = gtsupp.PokemonData.Copy().(*types.QBuffer)
	copied.Signature = gtsupp.Signature.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given GlobalTradeStationUploadPokemonParam contains the same data as the current GlobalTradeStationUploadPokemonParam
func (gtsupp *GlobalTradeStationUploadPokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationUploadPokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationUploadPokemonParam)

	if gtsupp.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsupp.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if !gtsupp.Period.Equals(other.Period) {
		return false
	}

	if !gtsupp.IndexData.Equals(other.IndexData) {
		return false
	}

	if !gtsupp.PokemonData.Equals(other.PokemonData) {
		return false
	}

	return gtsupp.Signature.Equals(other.Signature)
}

// String returns the string representation of the GlobalTradeStationUploadPokemonParam
func (gtsupp *GlobalTradeStationUploadPokemonParam) String() string {
	return gtsupp.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationUploadPokemonParam using the provided indentation level
func (gtsupp *GlobalTradeStationUploadPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationUploadPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s,\n", indentationValues, gtsupp.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, gtsupp.Period))
	b.WriteString(fmt.Sprintf("%sIndexData: %s,\n", indentationValues, gtsupp.IndexData))
	b.WriteString(fmt.Sprintf("%sPokemonData: %s,\n", indentationValues, gtsupp.PokemonData))
	b.WriteString(fmt.Sprintf("%sSignature: %s,\n", indentationValues, gtsupp.Signature))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationUploadPokemonParam returns a new GlobalTradeStationUploadPokemonParam
func NewGlobalTradeStationUploadPokemonParam() *GlobalTradeStationUploadPokemonParam {
	gtsupp := &GlobalTradeStationUploadPokemonParam{
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
		Period:           types.NewPrimitiveU16(0),
		IndexData:        types.NewQBuffer(nil),
		PokemonData:      types.NewQBuffer(nil),
		Signature:        types.NewQBuffer(nil),
	}

	return gtsupp
}
