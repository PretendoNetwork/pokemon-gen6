// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetMetaParam is a type within the DataStore protocol
type DataStoreGetMetaParam struct {
	types.Structure
	DataID            *types.PrimitiveU64
	PersistenceTarget *DataStorePersistenceTarget
	ResultOption      *types.PrimitiveU8
	AccessPassword    *types.PrimitiveU64
}

// WriteTo writes the DataStoreGetMetaParam to the given writable
func (dsgmp *DataStoreGetMetaParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgmp.DataID.WriteTo(contentWritable)
	dsgmp.PersistenceTarget.WriteTo(contentWritable)
	dsgmp.ResultOption.WriteTo(contentWritable)
	dsgmp.AccessPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgmp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetMetaParam from the given readable
func (dsgmp *DataStoreGetMetaParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgmp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam header. %s", err.Error())
	}

	err = dsgmp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.DataID. %s", err.Error())
	}

	err = dsgmp.PersistenceTarget.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.PersistenceTarget. %s", err.Error())
	}

	err = dsgmp.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.ResultOption. %s", err.Error())
	}

	err = dsgmp.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetMetaParam
func (dsgmp *DataStoreGetMetaParam) Copy() types.RVType {
	copied := NewDataStoreGetMetaParam()

	copied.StructureVersion = dsgmp.StructureVersion
	copied.DataID = dsgmp.DataID.Copy().(*types.PrimitiveU64)
	copied.PersistenceTarget = dsgmp.PersistenceTarget.Copy().(*DataStorePersistenceTarget)
	copied.ResultOption = dsgmp.ResultOption.Copy().(*types.PrimitiveU8)
	copied.AccessPassword = dsgmp.AccessPassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given DataStoreGetMetaParam contains the same data as the current DataStoreGetMetaParam
func (dsgmp *DataStoreGetMetaParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetMetaParam); !ok {
		return false
	}

	other := o.(*DataStoreGetMetaParam)

	if dsgmp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsgmp.DataID.Equals(other.DataID) {
		return false
	}

	if !dsgmp.PersistenceTarget.Equals(other.PersistenceTarget) {
		return false
	}

	if !dsgmp.ResultOption.Equals(other.ResultOption) {
		return false
	}

	return dsgmp.AccessPassword.Equals(other.AccessPassword)
}

// String returns the string representation of the DataStoreGetMetaParam
func (dsgmp *DataStoreGetMetaParam) String() string {
	return dsgmp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetMetaParam using the provided indentation level
func (dsgmp *DataStoreGetMetaParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetMetaParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsgmp.DataID))
	b.WriteString(fmt.Sprintf("%sPersistenceTarget: %s,\n", indentationValues, dsgmp.PersistenceTarget.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dsgmp.ResultOption))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s,\n", indentationValues, dsgmp.AccessPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetMetaParam returns a new DataStoreGetMetaParam
func NewDataStoreGetMetaParam() *DataStoreGetMetaParam {
	dsgmp := &DataStoreGetMetaParam{
		DataID:            types.NewPrimitiveU64(0),
		PersistenceTarget: NewDataStorePersistenceTarget(),
		ResultOption:      types.NewPrimitiveU8(0),
		AccessPassword:    types.NewPrimitiveU64(0),
	}

	return dsgmp
}
