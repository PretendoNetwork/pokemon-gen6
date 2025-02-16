// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetSpecificMetaParamV1 is a type within the DataStore protocol
type DataStoreGetSpecificMetaParamV1 struct {
	types.Structure
	DataIDs *types.List[*types.PrimitiveU32]
}

// WriteTo writes the DataStoreGetSpecificMetaParamV1 to the given writable
func (dsgsmpv *DataStoreGetSpecificMetaParamV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgsmpv.DataIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgsmpv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetSpecificMetaParamV1 from the given readable
func (dsgsmpv *DataStoreGetSpecificMetaParamV1) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgsmpv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetSpecificMetaParamV1 header. %s", err.Error())
	}

	err = dsgsmpv.DataIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetSpecificMetaParamV1.DataIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetSpecificMetaParamV1
func (dsgsmpv *DataStoreGetSpecificMetaParamV1) Copy() types.RVType {
	copied := NewDataStoreGetSpecificMetaParamV1()

	copied.StructureVersion = dsgsmpv.StructureVersion
	copied.DataIDs = dsgsmpv.DataIDs.Copy().(*types.List[*types.PrimitiveU32])

	return copied
}

// Equals checks if the given DataStoreGetSpecificMetaParamV1 contains the same data as the current DataStoreGetSpecificMetaParamV1
func (dsgsmpv *DataStoreGetSpecificMetaParamV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetSpecificMetaParamV1); !ok {
		return false
	}

	other := o.(*DataStoreGetSpecificMetaParamV1)

	if dsgsmpv.StructureVersion != other.StructureVersion {
		return false
	}

	return dsgsmpv.DataIDs.Equals(other.DataIDs)
}

// String returns the string representation of the DataStoreGetSpecificMetaParamV1
func (dsgsmpv *DataStoreGetSpecificMetaParamV1) String() string {
	return dsgsmpv.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetSpecificMetaParamV1 using the provided indentation level
func (dsgsmpv *DataStoreGetSpecificMetaParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetSpecificMetaParamV1{\n")
	b.WriteString(fmt.Sprintf("%sDataIDs: %s,\n", indentationValues, dsgsmpv.DataIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetSpecificMetaParamV1 returns a new DataStoreGetSpecificMetaParamV1
func NewDataStoreGetSpecificMetaParamV1() *DataStoreGetSpecificMetaParamV1 {
	dsgsmpv := &DataStoreGetSpecificMetaParamV1{
		DataIDs: types.NewList[*types.PrimitiveU32](),
	}

	dsgsmpv.DataIDs.Type = types.NewPrimitiveU32(0)

	return dsgsmpv
}
