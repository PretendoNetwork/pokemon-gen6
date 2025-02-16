// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePersistenceInfo is a type within the DataStore protocol
type DataStorePersistenceInfo struct {
	types.Structure
	OwnerID           *types.PID
	PersistenceSlotID *types.PrimitiveU16
	DataID            *types.PrimitiveU64
}

// WriteTo writes the DataStorePersistenceInfo to the given writable
func (dspi *DataStorePersistenceInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspi.OwnerID.WriteTo(contentWritable)
	dspi.PersistenceSlotID.WriteTo(contentWritable)
	dspi.DataID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePersistenceInfo from the given readable
func (dspi *DataStorePersistenceInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo header. %s", err.Error())
	}

	err = dspi.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.OwnerID. %s", err.Error())
	}

	err = dspi.PersistenceSlotID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.PersistenceSlotID. %s", err.Error())
	}

	err = dspi.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.DataID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePersistenceInfo
func (dspi *DataStorePersistenceInfo) Copy() types.RVType {
	copied := NewDataStorePersistenceInfo()

	copied.StructureVersion = dspi.StructureVersion
	copied.OwnerID = dspi.OwnerID.Copy().(*types.PID)
	copied.PersistenceSlotID = dspi.PersistenceSlotID.Copy().(*types.PrimitiveU16)
	copied.DataID = dspi.DataID.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given DataStorePersistenceInfo contains the same data as the current DataStorePersistenceInfo
func (dspi *DataStorePersistenceInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePersistenceInfo); !ok {
		return false
	}

	other := o.(*DataStorePersistenceInfo)

	if dspi.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspi.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dspi.PersistenceSlotID.Equals(other.PersistenceSlotID) {
		return false
	}

	return dspi.DataID.Equals(other.DataID)
}

// String returns the string representation of the DataStorePersistenceInfo
func (dspi *DataStorePersistenceInfo) String() string {
	return dspi.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePersistenceInfo using the provided indentation level
func (dspi *DataStorePersistenceInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePersistenceInfo{\n")
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, dspi.OwnerID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPersistenceSlotID: %s,\n", indentationValues, dspi.PersistenceSlotID))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dspi.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePersistenceInfo returns a new DataStorePersistenceInfo
func NewDataStorePersistenceInfo() *DataStorePersistenceInfo {
	dspi := &DataStorePersistenceInfo{
		OwnerID:           types.NewPID(0),
		PersistenceSlotID: types.NewPrimitiveU16(0),
		DataID:            types.NewPrimitiveU64(0),
	}

	return dspi
}
