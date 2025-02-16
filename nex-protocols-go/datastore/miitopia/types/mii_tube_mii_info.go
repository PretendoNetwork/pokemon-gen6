// Package types implements all the types used by the DataStoreMiitopia protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

// MiiTubeMiiInfo is a type within the DataStoreMiitopia protocol
type MiiTubeMiiInfo struct {
	types.Structure
	MetaInfo    *datastore_types.DataStoreMetaInfo
	Category    *types.PrimitiveU8
	RankingType *types.PrimitiveU8
}

// WriteTo writes the MiiTubeMiiInfo to the given writable
func (mtmi *MiiTubeMiiInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	mtmi.MetaInfo.WriteTo(contentWritable)
	mtmi.Category.WriteTo(contentWritable)
	mtmi.RankingType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mtmi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MiiTubeMiiInfo from the given readable
func (mtmi *MiiTubeMiiInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = mtmi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeMiiInfo header. %s", err.Error())
	}

	err = mtmi.MetaInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeMiiInfo.MetaInfo. %s", err.Error())
	}

	err = mtmi.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeMiiInfo.Category. %s", err.Error())
	}

	err = mtmi.RankingType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeMiiInfo.RankingType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MiiTubeMiiInfo
func (mtmi *MiiTubeMiiInfo) Copy() types.RVType {
	copied := NewMiiTubeMiiInfo()

	copied.MetaInfo = mtmi.MetaInfo
	copied.Category = mtmi.Category
	copied.RankingType = mtmi.RankingType

	return copied
}

// Equals checks if the given MiiTubeMiiInfo contains the same data as the current MiiTubeMiiInfo
func (mtmi *MiiTubeMiiInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*MiiTubeMiiInfo); !ok {
		return false
	}

	other := o.(*MiiTubeMiiInfo)

	if !mtmi.MetaInfo.Equals(other.MetaInfo) {
		return false
	}

	if !mtmi.Category.Equals(other.Category) {
		return false
	}

	return mtmi.RankingType.Equals(other.RankingType)
}

// String returns the string representation of the MiiTubeMiiInfo
func (mtmi *MiiTubeMiiInfo) String() string {
	return mtmi.FormatToString(0)
}

// FormatToString pretty-prints the MiiTubeMiiInfo using the provided indentation level
func (mtmi *MiiTubeMiiInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MiiTubeMiiInfo{\n")
	b.WriteString(fmt.Sprintf("%sMetaInfo: %s,\n", indentationValues, mtmi.MetaInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, mtmi.Category))
	b.WriteString(fmt.Sprintf("%sRankingType: %s,\n", indentationValues, mtmi.RankingType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMiiTubeMiiInfo returns a new MiiTubeMiiInfo
func NewMiiTubeMiiInfo() *MiiTubeMiiInfo {
	mtmi := &MiiTubeMiiInfo{
		MetaInfo:    datastore_types.NewDataStoreMetaInfo(),
		Category:    types.NewPrimitiveU8(0),
		RankingType: types.NewPrimitiveU8(0),
	}

	return mtmi
}
