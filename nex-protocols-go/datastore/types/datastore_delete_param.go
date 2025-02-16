// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreDeleteParam is a type within the DataStore protocol
type DataStoreDeleteParam struct {
	types.Structure
	DataID         *types.PrimitiveU64
	UpdatePassword *types.PrimitiveU64
}

// WriteTo writes the DataStoreDeleteParam to the given writable
func (dsdp *DataStoreDeleteParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsdp.DataID.WriteTo(contentWritable)
	dsdp.UpdatePassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsdp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreDeleteParam from the given readable
func (dsdp *DataStoreDeleteParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsdp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreDeleteParam header. %s", err.Error())
	}

	err = dsdp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreDeleteParam.DataID. %s", err.Error())
	}

	err = dsdp.UpdatePassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreDeleteParam.UpdatePassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreDeleteParam
func (dsdp *DataStoreDeleteParam) Copy() types.RVType {
	copied := NewDataStoreDeleteParam()

	copied.StructureVersion = dsdp.StructureVersion
	copied.DataID = dsdp.DataID.Copy().(*types.PrimitiveU64)
	copied.UpdatePassword = dsdp.UpdatePassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given DataStoreDeleteParam contains the same data as the current DataStoreDeleteParam
func (dsdp *DataStoreDeleteParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreDeleteParam); !ok {
		return false
	}

	other := o.(*DataStoreDeleteParam)

	if dsdp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsdp.DataID.Equals(other.DataID) {
		return false
	}

	return dsdp.UpdatePassword.Equals(other.UpdatePassword)
}

// String returns the string representation of the DataStoreDeleteParam
func (dsdp *DataStoreDeleteParam) String() string {
	return dsdp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreDeleteParam using the provided indentation level
func (dsdp *DataStoreDeleteParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreDeleteParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsdp.DataID))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %s,\n", indentationValues, dsdp.UpdatePassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreDeleteParam returns a new DataStoreDeleteParam
func NewDataStoreDeleteParam() *DataStoreDeleteParam {
	dsdp := &DataStoreDeleteParam{
		DataID:         types.NewPrimitiveU64(0),
		UpdatePassword: types.NewPrimitiveU64(0),
	}

	return dsdp
}
