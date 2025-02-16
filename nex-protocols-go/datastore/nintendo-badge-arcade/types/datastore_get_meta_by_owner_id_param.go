// Package types implements all the types used by the DataStoreNintendoBadgeArcade protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetMetaByOwnerIDParam is a type within the DataStoreNintendoBadgeArcade protocol
type DataStoreGetMetaByOwnerIDParam struct {
	types.Structure
	OwnerIDs     *types.List[*types.PID]
	DataTypes    *types.List[*types.PrimitiveU16]
	ResultOption *types.PrimitiveU8
	ResultRange  *types.ResultRange
}

// WriteTo writes the DataStoreGetMetaByOwnerIDParam to the given writable
func (dsgmboidp *DataStoreGetMetaByOwnerIDParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgmboidp.OwnerIDs.WriteTo(contentWritable)
	dsgmboidp.DataTypes.WriteTo(contentWritable)
	dsgmboidp.ResultOption.WriteTo(contentWritable)
	dsgmboidp.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgmboidp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetMetaByOwnerIDParam from the given readable
func (dsgmboidp *DataStoreGetMetaByOwnerIDParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgmboidp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam header. %s", err.Error())
	}

	err = dsgmboidp.OwnerIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.OwnerIDs. %s", err.Error())
	}

	err = dsgmboidp.DataTypes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.DataTypes. %s", err.Error())
	}

	err = dsgmboidp.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultOption. %s", err.Error())
	}

	err = dsgmboidp.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultRange. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetMetaByOwnerIDParam
func (dsgmboidp *DataStoreGetMetaByOwnerIDParam) Copy() types.RVType {
	copied := NewDataStoreGetMetaByOwnerIDParam()

	copied.StructureVersion = dsgmboidp.StructureVersion
	copied.OwnerIDs = dsgmboidp.OwnerIDs.Copy().(*types.List[*types.PID])
	copied.DataTypes = dsgmboidp.DataTypes.Copy().(*types.List[*types.PrimitiveU16])
	copied.ResultOption = dsgmboidp.ResultOption.Copy().(*types.PrimitiveU8)
	copied.ResultRange = dsgmboidp.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the given DataStoreGetMetaByOwnerIDParam contains the same data as the current DataStoreGetMetaByOwnerIDParam
func (dsgmboidp *DataStoreGetMetaByOwnerIDParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetMetaByOwnerIDParam); !ok {
		return false
	}

	other := o.(*DataStoreGetMetaByOwnerIDParam)

	if dsgmboidp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsgmboidp.OwnerIDs.Equals(other.OwnerIDs) {
		return false
	}

	if !dsgmboidp.DataTypes.Equals(other.DataTypes) {
		return false
	}

	if !dsgmboidp.ResultOption.Equals(other.ResultOption) {
		return false
	}

	return dsgmboidp.ResultRange.Equals(other.ResultRange)
}

// String returns the string representation of the DataStoreGetMetaByOwnerIDParam
func (dsgmboidp *DataStoreGetMetaByOwnerIDParam) String() string {
	return dsgmboidp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetMetaByOwnerIDParam using the provided indentation level
func (dsgmboidp *DataStoreGetMetaByOwnerIDParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetMetaByOwnerIDParam{\n")
	b.WriteString(fmt.Sprintf("%sOwnerIDs: %s,\n", indentationValues, dsgmboidp.OwnerIDs))
	b.WriteString(fmt.Sprintf("%sDataTypes: %s,\n", indentationValues, dsgmboidp.DataTypes))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dsgmboidp.ResultOption))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dsgmboidp.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetMetaByOwnerIDParam returns a new DataStoreGetMetaByOwnerIDParam
func NewDataStoreGetMetaByOwnerIDParam() *DataStoreGetMetaByOwnerIDParam {
	dsgmboidp := &DataStoreGetMetaByOwnerIDParam{
		OwnerIDs:     types.NewList[*types.PID](),
		DataTypes:    types.NewList[*types.PrimitiveU16](),
		ResultOption: types.NewPrimitiveU8(0),
		ResultRange:  types.NewResultRange(),
	}

	dsgmboidp.OwnerIDs.Type = types.NewPID(0)
	dsgmboidp.DataTypes.Type = types.NewPrimitiveU16(0)

	return dsgmboidp
}
