// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationDownloadOtherPokemonParam is a type within the DataStore protocol
type GlobalTradeStationDownloadOtherPokemonParam struct {
	types.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
}

// WriteTo writes the GlobalTradeStationDownloadOtherPokemonParam to the given writable
func (gtsdopp *GlobalTradeStationDownloadOtherPokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsdopp.PrepareUploadKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsdopp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationDownloadOtherPokemonParam from the given readable
func (gtsdopp *GlobalTradeStationDownloadOtherPokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsdopp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadOtherPokemonParam header. %s", err.Error())
	}

	err = gtsdopp.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadOtherPokemonParam.PrepareUploadKey. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationDownloadOtherPokemonParam
func (gtsdopp *GlobalTradeStationDownloadOtherPokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationDownloadOtherPokemonParam()

	copied.StructureVersion = gtsdopp.StructureVersion
	copied.PrepareUploadKey = gtsdopp.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the given GlobalTradeStationDownloadOtherPokemonParam contains the same data as the current GlobalTradeStationDownloadOtherPokemonParam
func (gtsdopp *GlobalTradeStationDownloadOtherPokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationDownloadOtherPokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationDownloadOtherPokemonParam)

	if gtsdopp.StructureVersion != other.StructureVersion {
		return false
	}

	return gtsdopp.PrepareUploadKey.Equals(other.PrepareUploadKey)
}

// String returns the string representation of the GlobalTradeStationDownloadOtherPokemonParam
func (gtsdopp *GlobalTradeStationDownloadOtherPokemonParam) String() string {
	return gtsdopp.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationDownloadOtherPokemonParam using the provided indentation level
func (gtsdopp *GlobalTradeStationDownloadOtherPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDownloadOtherPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s,\n", indentationValues, gtsdopp.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadOtherPokemonParam returns a new GlobalTradeStationDownloadOtherPokemonParam
func NewGlobalTradeStationDownloadOtherPokemonParam() *GlobalTradeStationDownloadOtherPokemonParam {
	gtsdopp := &GlobalTradeStationDownloadOtherPokemonParam{
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
	}

	return gtsdopp
}
