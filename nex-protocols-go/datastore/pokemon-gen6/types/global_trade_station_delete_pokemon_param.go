// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationDeletePokemonParam is a type within the DataStore protocol
type GlobalTradeStationDeletePokemonParam struct {
	types.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
	DeleteFlag       *types.PrimitiveU8
}

// WriteTo writes the GlobalTradeStationDeletePokemonParam to the given writable
func (gtsdpp *GlobalTradeStationDeletePokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsdpp.PrepareUploadKey.WriteTo(contentWritable)
	gtsdpp.DeleteFlag.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsdpp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationDeletePokemonParam from the given readable
func (gtsdpp *GlobalTradeStationDeletePokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsdpp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDeletePokemonParam header. %s", err.Error())
	}

	err = gtsdpp.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDeletePokemonParam.PrepareUploadKey. %s", err.Error())
	}

	err = gtsdpp.DeleteFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDeletePokemonParam.DeleteFlag. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationDeletePokemonParam
func (gtsdpp *GlobalTradeStationDeletePokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationDeletePokemonParam()

	copied.StructureVersion = gtsdpp.StructureVersion
	copied.PrepareUploadKey = gtsdpp.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.DeleteFlag = gtsdpp.DeleteFlag.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given GlobalTradeStationDeletePokemonParam contains the same data as the current GlobalTradeStationDeletePokemonParam
func (gtsdpp *GlobalTradeStationDeletePokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationDeletePokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationDeletePokemonParam)

	if gtsdpp.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsdpp.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	return gtsdpp.DeleteFlag.Equals(other.DeleteFlag)
}

// String returns the string representation of the GlobalTradeStationDeletePokemonParam
func (gtsdpp *GlobalTradeStationDeletePokemonParam) String() string {
	return gtsdpp.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationDeletePokemonParam using the provided indentation level
func (gtsdpp *GlobalTradeStationDeletePokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDeletePokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s,\n", indentationValues, gtsdpp.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDeleteFlag: %s,\n", indentationValues, gtsdpp.DeleteFlag))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDeletePokemonParam returns a new GlobalTradeStationDeletePokemonParam
func NewGlobalTradeStationDeletePokemonParam() *GlobalTradeStationDeletePokemonParam {
	gtsdpp := &GlobalTradeStationDeletePokemonParam{
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
		DeleteFlag:       types.NewPrimitiveU8(0),
	}

	return gtsdpp
}
