// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreCompletePostParamV1 is a type within the DataStore protocol
type DataStoreCompletePostParamV1 struct {
	types.Structure
	DataID    *types.PrimitiveU32
	IsSuccess *types.PrimitiveBool
}

// WriteTo writes the DataStoreCompletePostParamV1 to the given writable
func (dscppv *DataStoreCompletePostParamV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscppv.DataID.WriteTo(contentWritable)
	dscppv.IsSuccess.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscppv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreCompletePostParamV1 from the given readable
func (dscppv *DataStoreCompletePostParamV1) ExtractFrom(readable types.Readable) error {
	var err error

	err = dscppv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParamV1 header. %s", err.Error())
	}

	err = dscppv.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParamV1.DataID. %s", err.Error())
	}

	err = dscppv.IsSuccess.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParamV1.IsSuccess. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreCompletePostParamV1
func (dscppv *DataStoreCompletePostParamV1) Copy() types.RVType {
	copied := NewDataStoreCompletePostParamV1()

	copied.StructureVersion = dscppv.StructureVersion
	copied.DataID = dscppv.DataID.Copy().(*types.PrimitiveU32)
	copied.IsSuccess = dscppv.IsSuccess.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the given DataStoreCompletePostParamV1 contains the same data as the current DataStoreCompletePostParamV1
func (dscppv *DataStoreCompletePostParamV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCompletePostParamV1); !ok {
		return false
	}

	other := o.(*DataStoreCompletePostParamV1)

	if dscppv.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscppv.DataID.Equals(other.DataID) {
		return false
	}

	return dscppv.IsSuccess.Equals(other.IsSuccess)
}

// String returns the string representation of the DataStoreCompletePostParamV1
func (dscppv *DataStoreCompletePostParamV1) String() string {
	return dscppv.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreCompletePostParamV1 using the provided indentation level
func (dscppv *DataStoreCompletePostParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostParamV1{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dscppv.DataID))
	b.WriteString(fmt.Sprintf("%sIsSuccess: %s,\n", indentationValues, dscppv.IsSuccess))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostParamV1 returns a new DataStoreCompletePostParamV1
func NewDataStoreCompletePostParamV1() *DataStoreCompletePostParamV1 {
	dscppv := &DataStoreCompletePostParamV1{
		DataID:    types.NewPrimitiveU32(0),
		IsSuccess: types.NewPrimitiveBool(false),
	}

	return dscppv
}
