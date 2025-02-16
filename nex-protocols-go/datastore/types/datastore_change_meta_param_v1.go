// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreChangeMetaParamV1 is a type within the DataStore protocol
type DataStoreChangeMetaParamV1 struct {
	types.Structure
	DataID         *types.PrimitiveU64
	ModifiesFlag   *types.PrimitiveU32
	Name           *types.String
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         *types.PrimitiveU16
	MetaBinary     *types.QBuffer
	Tags           *types.List[*types.String]
	UpdatePassword *types.PrimitiveU64
}

// WriteTo writes the DataStoreChangeMetaParamV1 to the given writable
func (dscmpv *DataStoreChangeMetaParamV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscmpv.DataID.WriteTo(contentWritable)
	dscmpv.ModifiesFlag.WriteTo(contentWritable)
	dscmpv.Name.WriteTo(contentWritable)
	dscmpv.Permission.WriteTo(contentWritable)
	dscmpv.DelPermission.WriteTo(contentWritable)
	dscmpv.Period.WriteTo(contentWritable)
	dscmpv.MetaBinary.WriteTo(contentWritable)
	dscmpv.Tags.WriteTo(contentWritable)
	dscmpv.UpdatePassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscmpv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreChangeMetaParamV1 from the given readable
func (dscmpv *DataStoreChangeMetaParamV1) ExtractFrom(readable types.Readable) error {
	var err error

	err = dscmpv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1 header. %s", err.Error())
	}

	err = dscmpv.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.DataID. %s", err.Error())
	}

	err = dscmpv.ModifiesFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.ModifiesFlag. %s", err.Error())
	}

	err = dscmpv.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Name. %s", err.Error())
	}

	err = dscmpv.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Permission. %s", err.Error())
	}

	err = dscmpv.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.DelPermission. %s", err.Error())
	}

	err = dscmpv.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Period. %s", err.Error())
	}

	err = dscmpv.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.MetaBinary. %s", err.Error())
	}

	err = dscmpv.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Tags. %s", err.Error())
	}

	err = dscmpv.UpdatePassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.UpdatePassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreChangeMetaParamV1
func (dscmpv *DataStoreChangeMetaParamV1) Copy() types.RVType {
	copied := NewDataStoreChangeMetaParamV1()

	copied.StructureVersion = dscmpv.StructureVersion
	copied.DataID = dscmpv.DataID.Copy().(*types.PrimitiveU64)
	copied.ModifiesFlag = dscmpv.ModifiesFlag.Copy().(*types.PrimitiveU32)
	copied.Name = dscmpv.Name.Copy().(*types.String)
	copied.Permission = dscmpv.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dscmpv.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dscmpv.Period.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dscmpv.MetaBinary.Copy().(*types.QBuffer)
	copied.Tags = dscmpv.Tags.Copy().(*types.List[*types.String])
	copied.UpdatePassword = dscmpv.UpdatePassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given DataStoreChangeMetaParamV1 contains the same data as the current DataStoreChangeMetaParamV1
func (dscmpv *DataStoreChangeMetaParamV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreChangeMetaParamV1); !ok {
		return false
	}

	other := o.(*DataStoreChangeMetaParamV1)

	if dscmpv.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscmpv.DataID.Equals(other.DataID) {
		return false
	}

	if !dscmpv.ModifiesFlag.Equals(other.ModifiesFlag) {
		return false
	}

	if !dscmpv.Name.Equals(other.Name) {
		return false
	}

	if !dscmpv.Permission.Equals(other.Permission) {
		return false
	}

	if !dscmpv.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dscmpv.Period.Equals(other.Period) {
		return false
	}

	if !dscmpv.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dscmpv.Tags.Equals(other.Tags) {
		return false
	}

	return dscmpv.UpdatePassword.Equals(other.UpdatePassword)
}

// String returns the string representation of the DataStoreChangeMetaParamV1
func (dscmpv *DataStoreChangeMetaParamV1) String() string {
	return dscmpv.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreChangeMetaParamV1 using the provided indentation level
func (dscmpv *DataStoreChangeMetaParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangeMetaParamV1{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dscmpv.DataID))
	b.WriteString(fmt.Sprintf("%sModifiesFlag: %s,\n", indentationValues, dscmpv.ModifiesFlag))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dscmpv.Name))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dscmpv.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dscmpv.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dscmpv.Period))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dscmpv.MetaBinary))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dscmpv.Tags))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %s,\n", indentationValues, dscmpv.UpdatePassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangeMetaParamV1 returns a new DataStoreChangeMetaParamV1
func NewDataStoreChangeMetaParamV1() *DataStoreChangeMetaParamV1 {
	dscmpv := &DataStoreChangeMetaParamV1{
		DataID:         types.NewPrimitiveU64(0),
		ModifiesFlag:   types.NewPrimitiveU32(0),
		Name:           types.NewString(""),
		Permission:     NewDataStorePermission(),
		DelPermission:  NewDataStorePermission(),
		Period:         types.NewPrimitiveU16(0),
		MetaBinary:     types.NewQBuffer(nil),
		Tags:           types.NewList[*types.String](),
		UpdatePassword: types.NewPrimitiveU64(0),
	}

	dscmpv.Tags.Type = types.NewString("")

	return dscmpv
}
