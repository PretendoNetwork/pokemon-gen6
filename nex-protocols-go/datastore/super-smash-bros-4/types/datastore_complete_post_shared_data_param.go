// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

// DataStoreCompletePostSharedDataParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreCompletePostSharedDataParam struct {
	types.Structure
	DataID        *types.PrimitiveU64
	CompleteParam *datastore_types.DataStoreCompletePostParam
	PrepareParam  *DataStorePreparePostSharedDataParam
}

// WriteTo writes the DataStoreCompletePostSharedDataParam to the given writable
func (dscpsdp *DataStoreCompletePostSharedDataParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscpsdp.DataID.WriteTo(contentWritable)
	dscpsdp.CompleteParam.WriteTo(contentWritable)
	dscpsdp.PrepareParam.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscpsdp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreCompletePostSharedDataParam from the given readable
func (dscpsdp *DataStoreCompletePostSharedDataParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dscpsdp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam header. %s", err.Error())
	}

	err = dscpsdp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.DataID. %s", err.Error())
	}

	err = dscpsdp.CompleteParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.CompleteParam. %s", err.Error())
	}

	err = dscpsdp.PrepareParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostSharedDataParam.PrepareParam. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreCompletePostSharedDataParam
func (dscpsdp *DataStoreCompletePostSharedDataParam) Copy() types.RVType {
	copied := NewDataStoreCompletePostSharedDataParam()

	copied.StructureVersion = dscpsdp.StructureVersion
	copied.DataID = dscpsdp.DataID.Copy().(*types.PrimitiveU64)
	copied.CompleteParam = dscpsdp.CompleteParam.Copy().(*datastore_types.DataStoreCompletePostParam)
	copied.PrepareParam = dscpsdp.PrepareParam.Copy().(*DataStorePreparePostSharedDataParam)

	return copied
}

// Equals checks if the given DataStoreCompletePostSharedDataParam contains the same data as the current DataStoreCompletePostSharedDataParam
func (dscpsdp *DataStoreCompletePostSharedDataParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCompletePostSharedDataParam); !ok {
		return false
	}

	other := o.(*DataStoreCompletePostSharedDataParam)

	if dscpsdp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscpsdp.DataID.Equals(other.DataID) {
		return false
	}

	if !dscpsdp.CompleteParam.Equals(other.CompleteParam) {
		return false
	}

	return dscpsdp.PrepareParam.Equals(other.PrepareParam)
}

// String returns the string representation of the DataStoreCompletePostSharedDataParam
func (dscpsdp *DataStoreCompletePostSharedDataParam) String() string {
	return dscpsdp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreCompletePostSharedDataParam using the provided indentation level
func (dscpsdp *DataStoreCompletePostSharedDataParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostSharedDataParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dscpsdp.DataID))
	b.WriteString(fmt.Sprintf("%sCompleteParam: %s,\n", indentationValues, dscpsdp.CompleteParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareParam: %s,\n", indentationValues, dscpsdp.PrepareParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostSharedDataParam returns a new DataStoreCompletePostSharedDataParam
func NewDataStoreCompletePostSharedDataParam() *DataStoreCompletePostSharedDataParam {
	dscpsdp := &DataStoreCompletePostSharedDataParam{
		DataID:        types.NewPrimitiveU64(0),
		CompleteParam: datastore_types.NewDataStoreCompletePostParam(),
		PrepareParam:  NewDataStorePreparePostSharedDataParam(),
	}

	return dscpsdp
}
