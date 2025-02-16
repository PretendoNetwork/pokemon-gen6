// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreChangeMetaParam is a type within the DataStore protocol
type DataStoreChangeMetaParam struct {
	types.Structure
	DataID            *types.PrimitiveU64
	ModifiesFlag      *types.PrimitiveU32
	Name              *types.String
	Permission        *DataStorePermission
	DelPermission     *DataStorePermission
	Period            *types.PrimitiveU16
	MetaBinary        *types.QBuffer
	Tags              *types.List[*types.String]
	UpdatePassword    *types.PrimitiveU64
	ReferredCnt       *types.PrimitiveU32
	DataType          *types.PrimitiveU16
	Status            *types.PrimitiveU8
	CompareParam      *DataStoreChangeMetaCompareParam
	PersistenceTarget *DataStorePersistenceTarget
}

// WriteTo writes the DataStoreChangeMetaParam to the given writable
func (dscmp *DataStoreChangeMetaParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscmp.DataID.WriteTo(contentWritable)
	dscmp.ModifiesFlag.WriteTo(contentWritable)
	dscmp.Name.WriteTo(contentWritable)
	dscmp.Permission.WriteTo(contentWritable)
	dscmp.DelPermission.WriteTo(contentWritable)
	dscmp.Period.WriteTo(contentWritable)
	dscmp.MetaBinary.WriteTo(contentWritable)
	dscmp.Tags.WriteTo(contentWritable)
	dscmp.UpdatePassword.WriteTo(contentWritable)
	dscmp.ReferredCnt.WriteTo(contentWritable)
	dscmp.DataType.WriteTo(contentWritable)
	dscmp.Status.WriteTo(contentWritable)
	dscmp.CompareParam.WriteTo(contentWritable)
	dscmp.PersistenceTarget.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscmp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreChangeMetaParam from the given readable
func (dscmp *DataStoreChangeMetaParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dscmp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam header. %s", err.Error())
	}

	err = dscmp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.DataID. %s", err.Error())
	}

	err = dscmp.ModifiesFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.ModifiesFlag. %s", err.Error())
	}

	err = dscmp.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Name. %s", err.Error())
	}

	err = dscmp.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Permission. %s", err.Error())
	}

	err = dscmp.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.DelPermission. %s", err.Error())
	}

	err = dscmp.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Period. %s", err.Error())
	}

	err = dscmp.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.MetaBinary. %s", err.Error())
	}

	err = dscmp.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Tags. %s", err.Error())
	}

	err = dscmp.UpdatePassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.UpdatePassword. %s", err.Error())
	}

	err = dscmp.ReferredCnt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.ReferredCnt. %s", err.Error())
	}

	err = dscmp.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.DataType. %s", err.Error())
	}

	err = dscmp.Status.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Status. %s", err.Error())
	}

	err = dscmp.CompareParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.CompareParam. %s", err.Error())
	}

	err = dscmp.PersistenceTarget.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.PersistenceTarget. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreChangeMetaParam
func (dscmp *DataStoreChangeMetaParam) Copy() types.RVType {
	copied := NewDataStoreChangeMetaParam()

	copied.StructureVersion = dscmp.StructureVersion
	copied.DataID = dscmp.DataID.Copy().(*types.PrimitiveU64)
	copied.ModifiesFlag = dscmp.ModifiesFlag.Copy().(*types.PrimitiveU32)
	copied.Name = dscmp.Name.Copy().(*types.String)
	copied.Permission = dscmp.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dscmp.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dscmp.Period.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dscmp.MetaBinary.Copy().(*types.QBuffer)
	copied.Tags = dscmp.Tags.Copy().(*types.List[*types.String])
	copied.UpdatePassword = dscmp.UpdatePassword.Copy().(*types.PrimitiveU64)
	copied.ReferredCnt = dscmp.ReferredCnt.Copy().(*types.PrimitiveU32)
	copied.DataType = dscmp.DataType.Copy().(*types.PrimitiveU16)
	copied.Status = dscmp.Status.Copy().(*types.PrimitiveU8)
	copied.CompareParam = dscmp.CompareParam.Copy().(*DataStoreChangeMetaCompareParam)
	copied.PersistenceTarget = dscmp.PersistenceTarget.Copy().(*DataStorePersistenceTarget)

	return copied
}

// Equals checks if the given DataStoreChangeMetaParam contains the same data as the current DataStoreChangeMetaParam
func (dscmp *DataStoreChangeMetaParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreChangeMetaParam); !ok {
		return false
	}

	other := o.(*DataStoreChangeMetaParam)

	if dscmp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscmp.DataID.Equals(other.DataID) {
		return false
	}

	if !dscmp.ModifiesFlag.Equals(other.ModifiesFlag) {
		return false
	}

	if !dscmp.Name.Equals(other.Name) {
		return false
	}

	if !dscmp.Permission.Equals(other.Permission) {
		return false
	}

	if !dscmp.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dscmp.Period.Equals(other.Period) {
		return false
	}

	if !dscmp.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dscmp.Tags.Equals(other.Tags) {
		return false
	}

	if !dscmp.UpdatePassword.Equals(other.UpdatePassword) {
		return false
	}

	if !dscmp.ReferredCnt.Equals(other.ReferredCnt) {
		return false
	}

	if !dscmp.DataType.Equals(other.DataType) {
		return false
	}

	if !dscmp.Status.Equals(other.Status) {
		return false
	}

	if !dscmp.CompareParam.Equals(other.CompareParam) {
		return false
	}

	return dscmp.PersistenceTarget.Equals(other.PersistenceTarget)
}

// String returns the string representation of the DataStoreChangeMetaParam
func (dscmp *DataStoreChangeMetaParam) String() string {
	return dscmp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreChangeMetaParam using the provided indentation level
func (dscmp *DataStoreChangeMetaParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangeMetaParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dscmp.DataID))
	b.WriteString(fmt.Sprintf("%sModifiesFlag: %s,\n", indentationValues, dscmp.ModifiesFlag))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dscmp.Name))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dscmp.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dscmp.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dscmp.Period))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dscmp.MetaBinary))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dscmp.Tags))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %s,\n", indentationValues, dscmp.UpdatePassword))
	b.WriteString(fmt.Sprintf("%sReferredCnt: %s,\n", indentationValues, dscmp.ReferredCnt))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dscmp.DataType))
	b.WriteString(fmt.Sprintf("%sStatus: %s,\n", indentationValues, dscmp.Status))
	b.WriteString(fmt.Sprintf("%sCompareParam: %s,\n", indentationValues, dscmp.CompareParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPersistenceTarget: %s,\n", indentationValues, dscmp.PersistenceTarget.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangeMetaParam returns a new DataStoreChangeMetaParam
func NewDataStoreChangeMetaParam() *DataStoreChangeMetaParam {
	dscmp := &DataStoreChangeMetaParam{
		DataID:            types.NewPrimitiveU64(0),
		ModifiesFlag:      types.NewPrimitiveU32(0),
		Name:              types.NewString(""),
		Permission:        NewDataStorePermission(),
		DelPermission:     NewDataStorePermission(),
		Period:            types.NewPrimitiveU16(0),
		MetaBinary:        types.NewQBuffer(nil),
		Tags:              types.NewList[*types.String](),
		UpdatePassword:    types.NewPrimitiveU64(0),
		ReferredCnt:       types.NewPrimitiveU32(0),
		DataType:          types.NewPrimitiveU16(0),
		Status:            types.NewPrimitiveU8(0),
		CompareParam:      NewDataStoreChangeMetaCompareParam(),
		PersistenceTarget: NewDataStorePersistenceTarget(),
	}

	dscmp.Tags.Type = types.NewString("")

	return dscmp
}
