// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreSearchSharedDataParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreSearchSharedDataParam struct {
	types.Structure
	DataType    *types.PrimitiveU8
	Owner       *types.PrimitiveU32
	Region      *types.PrimitiveU8
	Attribute1  *types.PrimitiveU8
	Attribute2  *types.PrimitiveU8
	Fighter     *types.PrimitiveU8
	ResultRange *types.ResultRange
}

// WriteTo writes the DataStoreSearchSharedDataParam to the given writable
func (dsssdp *DataStoreSearchSharedDataParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsssdp.DataType.WriteTo(contentWritable)
	dsssdp.Owner.WriteTo(contentWritable)
	dsssdp.Region.WriteTo(contentWritable)
	dsssdp.Attribute1.WriteTo(contentWritable)
	dsssdp.Attribute2.WriteTo(contentWritable)
	dsssdp.Fighter.WriteTo(contentWritable)
	dsssdp.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsssdp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreSearchSharedDataParam from the given readable
func (dsssdp *DataStoreSearchSharedDataParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsssdp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam header. %s", err.Error())
	}

	err = dsssdp.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.DataType. %s", err.Error())
	}

	err = dsssdp.Owner.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Owner. %s", err.Error())
	}

	err = dsssdp.Region.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Region. %s", err.Error())
	}

	err = dsssdp.Attribute1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Attribute1. %s", err.Error())
	}

	err = dsssdp.Attribute2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Attribute2. %s", err.Error())
	}

	err = dsssdp.Fighter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.Fighter. %s", err.Error())
	}

	err = dsssdp.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchSharedDataParam.ResultRange. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSearchSharedDataParam
func (dsssdp *DataStoreSearchSharedDataParam) Copy() types.RVType {
	copied := NewDataStoreSearchSharedDataParam()

	copied.StructureVersion = dsssdp.StructureVersion
	copied.DataType = dsssdp.DataType.Copy().(*types.PrimitiveU8)
	copied.Owner = dsssdp.Owner.Copy().(*types.PrimitiveU32)
	copied.Region = dsssdp.Region.Copy().(*types.PrimitiveU8)
	copied.Attribute1 = dsssdp.Attribute1.Copy().(*types.PrimitiveU8)
	copied.Attribute2 = dsssdp.Attribute2.Copy().(*types.PrimitiveU8)
	copied.Fighter = dsssdp.Fighter.Copy().(*types.PrimitiveU8)
	copied.ResultRange = dsssdp.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the given DataStoreSearchSharedDataParam contains the same data as the current DataStoreSearchSharedDataParam
func (dsssdp *DataStoreSearchSharedDataParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSearchSharedDataParam); !ok {
		return false
	}

	other := o.(*DataStoreSearchSharedDataParam)

	if dsssdp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsssdp.DataType.Equals(other.DataType) {
		return false
	}

	if !dsssdp.Owner.Equals(other.Owner) {
		return false
	}

	if !dsssdp.Region.Equals(other.Region) {
		return false
	}

	if !dsssdp.Attribute1.Equals(other.Attribute1) {
		return false
	}

	if !dsssdp.Attribute2.Equals(other.Attribute2) {
		return false
	}

	if !dsssdp.Fighter.Equals(other.Fighter) {
		return false
	}

	return dsssdp.ResultRange.Equals(other.ResultRange)
}

// String returns the string representation of the DataStoreSearchSharedDataParam
func (dsssdp *DataStoreSearchSharedDataParam) String() string {
	return dsssdp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreSearchSharedDataParam using the provided indentation level
func (dsssdp *DataStoreSearchSharedDataParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchSharedDataParam{\n")
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dsssdp.DataType))
	b.WriteString(fmt.Sprintf("%sOwner: %s,\n", indentationValues, dsssdp.Owner))
	b.WriteString(fmt.Sprintf("%sRegion: %s,\n", indentationValues, dsssdp.Region))
	b.WriteString(fmt.Sprintf("%sAttribute1: %s,\n", indentationValues, dsssdp.Attribute1))
	b.WriteString(fmt.Sprintf("%sAttribute2: %s,\n", indentationValues, dsssdp.Attribute2))
	b.WriteString(fmt.Sprintf("%sFighter: %s,\n", indentationValues, dsssdp.Fighter))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dsssdp.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchSharedDataParam returns a new DataStoreSearchSharedDataParam
func NewDataStoreSearchSharedDataParam() *DataStoreSearchSharedDataParam {
	dsssdp := &DataStoreSearchSharedDataParam{
		DataType:    types.NewPrimitiveU8(0),
		Owner:       types.NewPrimitiveU32(0),
		Region:      types.NewPrimitiveU8(0),
		Attribute1:  types.NewPrimitiveU8(0),
		Attribute2:  types.NewPrimitiveU8(0),
		Fighter:     types.NewPrimitiveU8(0),
		ResultRange: types.NewResultRange(),
	}

	return dsssdp
}
