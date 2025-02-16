// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreSpecificMetaInfoV1 is a type within the DataStore protocol
type DataStoreSpecificMetaInfoV1 struct {
	types.Structure
	DataID   *types.PrimitiveU32
	OwnerID  *types.PID
	Size     *types.PrimitiveU32
	DataType *types.PrimitiveU16
	Version  *types.PrimitiveU16
}

// WriteTo writes the DataStoreSpecificMetaInfoV1 to the given writable
func (dssmiv *DataStoreSpecificMetaInfoV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dssmiv.DataID.WriteTo(contentWritable)
	dssmiv.OwnerID.WriteTo(contentWritable)
	dssmiv.Size.WriteTo(contentWritable)
	dssmiv.DataType.WriteTo(contentWritable)
	dssmiv.Version.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dssmiv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreSpecificMetaInfoV1 from the given readable
func (dssmiv *DataStoreSpecificMetaInfoV1) ExtractFrom(readable types.Readable) error {
	var err error

	err = dssmiv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1 header. %s", err.Error())
	}

	err = dssmiv.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1.DataID. %s", err.Error())
	}

	err = dssmiv.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1.OwnerID. %s", err.Error())
	}

	err = dssmiv.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1.Size. %s", err.Error())
	}

	err = dssmiv.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1.DataType. %s", err.Error())
	}

	err = dssmiv.Version.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfoV1.Version. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSpecificMetaInfoV1
func (dssmiv *DataStoreSpecificMetaInfoV1) Copy() types.RVType {
	copied := NewDataStoreSpecificMetaInfoV1()

	copied.StructureVersion = dssmiv.StructureVersion
	copied.DataID = dssmiv.DataID.Copy().(*types.PrimitiveU32)
	copied.OwnerID = dssmiv.OwnerID.Copy().(*types.PID)
	copied.Size = dssmiv.Size.Copy().(*types.PrimitiveU32)
	copied.DataType = dssmiv.DataType.Copy().(*types.PrimitiveU16)
	copied.Version = dssmiv.Version.Copy().(*types.PrimitiveU16)

	return copied
}

// Equals checks if the given DataStoreSpecificMetaInfoV1 contains the same data as the current DataStoreSpecificMetaInfoV1
func (dssmiv *DataStoreSpecificMetaInfoV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSpecificMetaInfoV1); !ok {
		return false
	}

	other := o.(*DataStoreSpecificMetaInfoV1)

	if dssmiv.StructureVersion != other.StructureVersion {
		return false
	}

	if !dssmiv.DataID.Equals(other.DataID) {
		return false
	}

	if !dssmiv.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dssmiv.Size.Equals(other.Size) {
		return false
	}

	if !dssmiv.DataType.Equals(other.DataType) {
		return false
	}

	return dssmiv.Version.Equals(other.Version)
}

// String returns the string representation of the DataStoreSpecificMetaInfoV1
func (dssmiv *DataStoreSpecificMetaInfoV1) String() string {
	return dssmiv.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreSpecificMetaInfoV1 using the provided indentation level
func (dssmiv *DataStoreSpecificMetaInfoV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSpecificMetaInfoV1{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dssmiv.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, dssmiv.OwnerID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dssmiv.Size))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dssmiv.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, dssmiv.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSpecificMetaInfoV1 returns a new DataStoreSpecificMetaInfoV1
func NewDataStoreSpecificMetaInfoV1() *DataStoreSpecificMetaInfoV1 {
	dssmiv := &DataStoreSpecificMetaInfoV1{
		DataID:   types.NewPrimitiveU32(0),
		OwnerID:  types.NewPID(0),
		Size:     types.NewPrimitiveU32(0),
		DataType: types.NewPrimitiveU16(0),
		Version:  types.NewPrimitiveU16(0),
	}

	return dssmiv
}
