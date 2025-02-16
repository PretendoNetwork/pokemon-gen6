// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

// DataStoreFileServerObjectInfo is a type within the DataStore protocol
type DataStoreFileServerObjectInfo struct {
	types.Structure
	DataID  *types.PrimitiveU64
	GetInfo *datastore_types.DataStoreReqGetInfo
}

// WriteTo writes the DataStoreFileServerObjectInfo to the given writable
func (dsfsoi *DataStoreFileServerObjectInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsfsoi.DataID.WriteTo(contentWritable)
	dsfsoi.GetInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsfsoi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreFileServerObjectInfo from the given readable
func (dsfsoi *DataStoreFileServerObjectInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsfsoi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFileServerObjectInfo header. %s", err.Error())
	}

	err = dsfsoi.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFileServerObjectInfo.DataID. %s", err.Error())
	}

	err = dsfsoi.GetInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFileServerObjectInfo.GetInfo. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreFileServerObjectInfo
func (dsfsoi *DataStoreFileServerObjectInfo) Copy() types.RVType {
	copied := NewDataStoreFileServerObjectInfo()

	copied.StructureVersion = dsfsoi.StructureVersion
	copied.DataID = dsfsoi.DataID.Copy().(*types.PrimitiveU64)
	copied.GetInfo = dsfsoi.GetInfo.Copy().(*datastore_types.DataStoreReqGetInfo)

	return copied
}

// Equals checks if the given DataStoreFileServerObjectInfo contains the same data as the current DataStoreFileServerObjectInfo
func (dsfsoi *DataStoreFileServerObjectInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreFileServerObjectInfo); !ok {
		return false
	}

	other := o.(*DataStoreFileServerObjectInfo)

	if dsfsoi.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsfsoi.DataID.Equals(other.DataID) {
		return false
	}

	return dsfsoi.GetInfo.Equals(other.GetInfo)
}

// String returns the string representation of the DataStoreFileServerObjectInfo
func (dsfsoi *DataStoreFileServerObjectInfo) String() string {
	return dsfsoi.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreFileServerObjectInfo using the provided indentation level
func (dsfsoi *DataStoreFileServerObjectInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreFileServerObjectInfo{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsfsoi.DataID))
	b.WriteString(fmt.Sprintf("%sGetInfo: %s,\n", indentationValues, dsfsoi.GetInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreFileServerObjectInfo returns a new DataStoreFileServerObjectInfo
func NewDataStoreFileServerObjectInfo() *DataStoreFileServerObjectInfo {
	dsfsoi := &DataStoreFileServerObjectInfo{
		DataID:  types.NewPrimitiveU64(0),
		GetInfo: datastore_types.NewDataStoreReqGetInfo(),
	}

	return dsfsoi
}
