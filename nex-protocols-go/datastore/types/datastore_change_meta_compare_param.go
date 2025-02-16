// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreChangeMetaCompareParam is a type within the DataStore protocol
type DataStoreChangeMetaCompareParam struct {
	types.Structure
	ComparisonFlag *types.PrimitiveU32
	Name           *types.String
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         *types.PrimitiveU16
	MetaBinary     *types.QBuffer
	Tags           *types.List[*types.String]
	ReferredCnt    *types.PrimitiveU32
	DataType       *types.PrimitiveU16
	Status         *types.PrimitiveU8
}

// WriteTo writes the DataStoreChangeMetaCompareParam to the given writable
func (dscmcp *DataStoreChangeMetaCompareParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscmcp.ComparisonFlag.WriteTo(contentWritable)
	dscmcp.Name.WriteTo(contentWritable)
	dscmcp.Permission.WriteTo(contentWritable)
	dscmcp.DelPermission.WriteTo(contentWritable)
	dscmcp.Period.WriteTo(contentWritable)
	dscmcp.MetaBinary.WriteTo(contentWritable)
	dscmcp.Tags.WriteTo(contentWritable)
	dscmcp.ReferredCnt.WriteTo(contentWritable)
	dscmcp.DataType.WriteTo(contentWritable)
	dscmcp.Status.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscmcp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreChangeMetaCompareParam from the given readable
func (dscmcp *DataStoreChangeMetaCompareParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dscmcp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam header. %s", err.Error())
	}

	err = dscmcp.ComparisonFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.ComparisonFlag. %s", err.Error())
	}

	err = dscmcp.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Name. %s", err.Error())
	}

	err = dscmcp.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Permission. %s", err.Error())
	}

	err = dscmcp.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.DelPermission. %s", err.Error())
	}

	err = dscmcp.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Period. %s", err.Error())
	}

	err = dscmcp.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.MetaBinary. %s", err.Error())
	}

	err = dscmcp.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Tags. %s", err.Error())
	}

	err = dscmcp.ReferredCnt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.ReferredCnt. %s", err.Error())
	}

	err = dscmcp.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.DataType. %s", err.Error())
	}

	err = dscmcp.Status.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Status. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreChangeMetaCompareParam
func (dscmcp *DataStoreChangeMetaCompareParam) Copy() types.RVType {
	copied := NewDataStoreChangeMetaCompareParam()

	copied.StructureVersion = dscmcp.StructureVersion
	copied.ComparisonFlag = dscmcp.ComparisonFlag.Copy().(*types.PrimitiveU32)
	copied.Name = dscmcp.Name.Copy().(*types.String)
	copied.Permission = dscmcp.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dscmcp.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dscmcp.Period.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dscmcp.MetaBinary.Copy().(*types.QBuffer)
	copied.Tags = dscmcp.Tags.Copy().(*types.List[*types.String])
	copied.ReferredCnt = dscmcp.ReferredCnt.Copy().(*types.PrimitiveU32)
	copied.DataType = dscmcp.DataType.Copy().(*types.PrimitiveU16)
	copied.Status = dscmcp.Status.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given DataStoreChangeMetaCompareParam contains the same data as the current DataStoreChangeMetaCompareParam
func (dscmcp *DataStoreChangeMetaCompareParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreChangeMetaCompareParam); !ok {
		return false
	}

	other := o.(*DataStoreChangeMetaCompareParam)

	if dscmcp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscmcp.ComparisonFlag.Equals(other.ComparisonFlag) {
		return false
	}

	if !dscmcp.Name.Equals(other.Name) {
		return false
	}

	if !dscmcp.Permission.Equals(other.Permission) {
		return false
	}

	if !dscmcp.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dscmcp.Period.Equals(other.Period) {
		return false
	}

	if !dscmcp.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dscmcp.Tags.Equals(other.Tags) {
		return false
	}

	if !dscmcp.ReferredCnt.Equals(other.ReferredCnt) {
		return false
	}

	if !dscmcp.DataType.Equals(other.DataType) {
		return false
	}

	return dscmcp.Status.Equals(other.Status)
}

// String returns the string representation of the DataStoreChangeMetaCompareParam
func (dscmcp *DataStoreChangeMetaCompareParam) String() string {
	return dscmcp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreChangeMetaCompareParam using the provided indentation level
func (dscmcp *DataStoreChangeMetaCompareParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangeMetaCompareParam{\n")
	b.WriteString(fmt.Sprintf("%sComparisonFlag: %s,\n", indentationValues, dscmcp.ComparisonFlag))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dscmcp.Name))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dscmcp.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dscmcp.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dscmcp.Period))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dscmcp.MetaBinary))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dscmcp.Tags))
	b.WriteString(fmt.Sprintf("%sReferredCnt: %s,\n", indentationValues, dscmcp.ReferredCnt))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dscmcp.DataType))
	b.WriteString(fmt.Sprintf("%sStatus: %s,\n", indentationValues, dscmcp.Status))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangeMetaCompareParam returns a new DataStoreChangeMetaCompareParam
func NewDataStoreChangeMetaCompareParam() *DataStoreChangeMetaCompareParam {
	dscmcp := &DataStoreChangeMetaCompareParam{
		ComparisonFlag: types.NewPrimitiveU32(0),
		Name:           types.NewString(""),
		Permission:     NewDataStorePermission(),
		DelPermission:  NewDataStorePermission(),
		Period:         types.NewPrimitiveU16(0),
		MetaBinary:     types.NewQBuffer(nil),
		Tags:           types.NewList[*types.String](),
		ReferredCnt:    types.NewPrimitiveU32(0),
		DataType:       types.NewPrimitiveU16(0),
		Status:         types.NewPrimitiveU8(0),
	}

	dscmcp.Tags.Type = types.NewString("")

	return dscmcp
}
