// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetSpecificMetaParam is a type within the DataStore protocol
type DataStoreGetSpecificMetaParam struct {
	types.Structure
	DataIDs *types.List[*types.PrimitiveU64]
}

// WriteTo writes the DataStoreGetSpecificMetaParam to the given writable
func (dsgsmp *DataStoreGetSpecificMetaParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgsmp.DataIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgsmp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetSpecificMetaParam from the given readable
func (dsgsmp *DataStoreGetSpecificMetaParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgsmp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetSpecificMetaParam header. %s", err.Error())
	}

	err = dsgsmp.DataIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetSpecificMetaParam.DataIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetSpecificMetaParam
func (dsgsmp *DataStoreGetSpecificMetaParam) Copy() types.RVType {
	copied := NewDataStoreGetSpecificMetaParam()

	copied.StructureVersion = dsgsmp.StructureVersion
	copied.DataIDs = dsgsmp.DataIDs.Copy().(*types.List[*types.PrimitiveU64])

	return copied
}

// Equals checks if the given DataStoreGetSpecificMetaParam contains the same data as the current DataStoreGetSpecificMetaParam
func (dsgsmp *DataStoreGetSpecificMetaParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetSpecificMetaParam); !ok {
		return false
	}

	other := o.(*DataStoreGetSpecificMetaParam)

	if dsgsmp.StructureVersion != other.StructureVersion {
		return false
	}

	return dsgsmp.DataIDs.Equals(other.DataIDs)
}

// String returns the string representation of the DataStoreGetSpecificMetaParam
func (dsgsmp *DataStoreGetSpecificMetaParam) String() string {
	return dsgsmp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetSpecificMetaParam using the provided indentation level
func (dsgsmp *DataStoreGetSpecificMetaParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetSpecificMetaParam{\n")
	b.WriteString(fmt.Sprintf("%sDataIDs: %s,\n", indentationValues, dsgsmp.DataIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetSpecificMetaParam returns a new DataStoreGetSpecificMetaParam
func NewDataStoreGetSpecificMetaParam() *DataStoreGetSpecificMetaParam {
	dsgsmp := &DataStoreGetSpecificMetaParam{
		DataIDs: types.NewList[*types.PrimitiveU64](),
	}

	dsgsmp.DataIDs.Type = types.NewPrimitiveU64(0)

	return dsgsmp
}
