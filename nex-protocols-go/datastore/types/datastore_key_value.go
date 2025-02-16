// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreKeyValue is a type within the DataStore protocol
type DataStoreKeyValue struct {
	types.Structure
	Key   *types.String
	Value *types.String
}

// WriteTo writes the DataStoreKeyValue to the given writable
func (dskv *DataStoreKeyValue) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dskv.Key.WriteTo(contentWritable)
	dskv.Value.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dskv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreKeyValue from the given readable
func (dskv *DataStoreKeyValue) ExtractFrom(readable types.Readable) error {
	var err error

	err = dskv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreKeyValue header. %s", err.Error())
	}

	err = dskv.Key.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreKeyValue.Key. %s", err.Error())
	}

	err = dskv.Value.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreKeyValue.Value. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreKeyValue
func (dskv *DataStoreKeyValue) Copy() types.RVType {
	copied := NewDataStoreKeyValue()

	copied.StructureVersion = dskv.StructureVersion
	copied.Key = dskv.Key.Copy().(*types.String)
	copied.Value = dskv.Value.Copy().(*types.String)

	return copied
}

// Equals checks if the given DataStoreKeyValue contains the same data as the current DataStoreKeyValue
func (dskv *DataStoreKeyValue) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreKeyValue); !ok {
		return false
	}

	other := o.(*DataStoreKeyValue)

	if dskv.StructureVersion != other.StructureVersion {
		return false
	}

	if !dskv.Key.Equals(other.Key) {
		return false
	}

	return dskv.Value.Equals(other.Value)
}

// String returns the string representation of the DataStoreKeyValue
func (dskv *DataStoreKeyValue) String() string {
	return dskv.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreKeyValue using the provided indentation level
func (dskv *DataStoreKeyValue) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreKeyValue{\n")
	b.WriteString(fmt.Sprintf("%sKey: %s,\n", indentationValues, dskv.Key))
	b.WriteString(fmt.Sprintf("%sValue: %s,\n", indentationValues, dskv.Value))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreKeyValue returns a new DataStoreKeyValue
func NewDataStoreKeyValue() *DataStoreKeyValue {
	dskv := &DataStoreKeyValue{
		Key:   types.NewString(""),
		Value: types.NewString(""),
	}

	return dskv
}
