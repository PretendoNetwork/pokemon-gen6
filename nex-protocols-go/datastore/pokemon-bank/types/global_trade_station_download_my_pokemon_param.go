// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationDownloadMyPokemonParam is a type within the DataStore protocol
type GlobalTradeStationDownloadMyPokemonParam struct {
	types.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
}

// WriteTo writes the GlobalTradeStationDownloadMyPokemonParam to the given writable
func (gtsdmpp *GlobalTradeStationDownloadMyPokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsdmpp.PrepareUploadKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsdmpp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationDownloadMyPokemonParam from the given readable
func (gtsdmpp *GlobalTradeStationDownloadMyPokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsdmpp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonParam header. %s", err.Error())
	}

	err = gtsdmpp.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonParam.PrepareUploadKey. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationDownloadMyPokemonParam
func (gtsdmpp *GlobalTradeStationDownloadMyPokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationDownloadMyPokemonParam()

	copied.StructureVersion = gtsdmpp.StructureVersion
	copied.PrepareUploadKey = gtsdmpp.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the given GlobalTradeStationDownloadMyPokemonParam contains the same data as the current GlobalTradeStationDownloadMyPokemonParam
func (gtsdmpp *GlobalTradeStationDownloadMyPokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationDownloadMyPokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationDownloadMyPokemonParam)

	if gtsdmpp.StructureVersion != other.StructureVersion {
		return false
	}

	return gtsdmpp.PrepareUploadKey.Equals(other.PrepareUploadKey)
}

// String returns the string representation of the GlobalTradeStationDownloadMyPokemonParam
func (gtsdmpp *GlobalTradeStationDownloadMyPokemonParam) String() string {
	return gtsdmpp.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationDownloadMyPokemonParam using the provided indentation level
func (gtsdmpp *GlobalTradeStationDownloadMyPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDownloadMyPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s,\n", indentationValues, gtsdmpp.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadMyPokemonParam returns a new GlobalTradeStationDownloadMyPokemonParam
func NewGlobalTradeStationDownloadMyPokemonParam() *GlobalTradeStationDownloadMyPokemonParam {
	gtsdmpp := &GlobalTradeStationDownloadMyPokemonParam{
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
	}

	return gtsdmpp
}
