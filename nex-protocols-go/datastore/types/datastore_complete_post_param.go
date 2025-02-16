// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreCompletePostParam is a type within the DataStore protocol
type DataStoreCompletePostParam struct {
	types.Structure
	DataID    *types.PrimitiveU64
	IsSuccess *types.PrimitiveBool
}

// WriteTo writes the DataStoreCompletePostParam to the given writable
func (dscpp *DataStoreCompletePostParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscpp.DataID.WriteTo(contentWritable)
	dscpp.IsSuccess.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscpp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreCompletePostParam from the given readable
func (dscpp *DataStoreCompletePostParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dscpp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParam header. %s", err.Error())
	}

	err = dscpp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParam.DataID. %s", err.Error())
	}

	err = dscpp.IsSuccess.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParam.IsSuccess. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreCompletePostParam
func (dscpp *DataStoreCompletePostParam) Copy() types.RVType {
	copied := NewDataStoreCompletePostParam()

	copied.StructureVersion = dscpp.StructureVersion
	copied.DataID = dscpp.DataID.Copy().(*types.PrimitiveU64)
	copied.IsSuccess = dscpp.IsSuccess.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the given DataStoreCompletePostParam contains the same data as the current DataStoreCompletePostParam
func (dscpp *DataStoreCompletePostParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCompletePostParam); !ok {
		return false
	}

	other := o.(*DataStoreCompletePostParam)

	if dscpp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscpp.DataID.Equals(other.DataID) {
		return false
	}

	return dscpp.IsSuccess.Equals(other.IsSuccess)
}

// String returns the string representation of the DataStoreCompletePostParam
func (dscpp *DataStoreCompletePostParam) String() string {
	return dscpp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreCompletePostParam using the provided indentation level
func (dscpp *DataStoreCompletePostParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dscpp.DataID))
	b.WriteString(fmt.Sprintf("%sIsSuccess: %s,\n", indentationValues, dscpp.IsSuccess))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostParam returns a new DataStoreCompletePostParam
func NewDataStoreCompletePostParam() *DataStoreCompletePostParam {
	dscpp := &DataStoreCompletePostParam{
		DataID:    types.NewPrimitiveU64(0),
		IsSuccess: types.NewPrimitiveBool(false),
	}

	return dscpp
}
