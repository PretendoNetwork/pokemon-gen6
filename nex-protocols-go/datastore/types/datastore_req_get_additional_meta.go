// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreReqGetAdditionalMeta is a type within the DataStore protocol
type DataStoreReqGetAdditionalMeta struct {
	types.Structure
	OwnerID    *types.PID
	DataType   *types.PrimitiveU16
	Version    *types.PrimitiveU16
	MetaBinary *types.QBuffer
}

// WriteTo writes the DataStoreReqGetAdditionalMeta to the given writable
func (dsrgam *DataStoreReqGetAdditionalMeta) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrgam.OwnerID.WriteTo(contentWritable)
	dsrgam.DataType.WriteTo(contentWritable)
	dsrgam.Version.WriteTo(contentWritable)
	dsrgam.MetaBinary.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrgam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReqGetAdditionalMeta from the given readable
func (dsrgam *DataStoreReqGetAdditionalMeta) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrgam.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta header. %s", err.Error())
	}

	err = dsrgam.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.OwnerID. %s", err.Error())
	}

	err = dsrgam.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.DataType. %s", err.Error())
	}

	err = dsrgam.Version.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.Version. %s", err.Error())
	}

	err = dsrgam.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.MetaBinary. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReqGetAdditionalMeta
func (dsrgam *DataStoreReqGetAdditionalMeta) Copy() types.RVType {
	copied := NewDataStoreReqGetAdditionalMeta()

	copied.StructureVersion = dsrgam.StructureVersion
	copied.OwnerID = dsrgam.OwnerID.Copy().(*types.PID)
	copied.DataType = dsrgam.DataType.Copy().(*types.PrimitiveU16)
	copied.Version = dsrgam.Version.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dsrgam.MetaBinary.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given DataStoreReqGetAdditionalMeta contains the same data as the current DataStoreReqGetAdditionalMeta
func (dsrgam *DataStoreReqGetAdditionalMeta) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqGetAdditionalMeta); !ok {
		return false
	}

	other := o.(*DataStoreReqGetAdditionalMeta)

	if dsrgam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrgam.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dsrgam.DataType.Equals(other.DataType) {
		return false
	}

	if !dsrgam.Version.Equals(other.Version) {
		return false
	}

	return dsrgam.MetaBinary.Equals(other.MetaBinary)
}

// String returns the string representation of the DataStoreReqGetAdditionalMeta
func (dsrgam *DataStoreReqGetAdditionalMeta) String() string {
	return dsrgam.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreReqGetAdditionalMeta using the provided indentation level
func (dsrgam *DataStoreReqGetAdditionalMeta) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetAdditionalMeta{\n")
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, dsrgam.OwnerID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dsrgam.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, dsrgam.Version))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dsrgam.MetaBinary))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetAdditionalMeta returns a new DataStoreReqGetAdditionalMeta
func NewDataStoreReqGetAdditionalMeta() *DataStoreReqGetAdditionalMeta {
	dsrgam := &DataStoreReqGetAdditionalMeta{
		OwnerID:    types.NewPID(0),
		DataType:   types.NewPrimitiveU16(0),
		Version:    types.NewPrimitiveU16(0),
		MetaBinary: types.NewQBuffer(nil),
	}

	return dsrgam
}
