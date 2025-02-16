// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePersistenceInitParam is a type within the DataStore protocol
type DataStorePersistenceInitParam struct {
	types.Structure
	PersistenceSlotID *types.PrimitiveU16
	DeleteLastObject  *types.PrimitiveBool
}

// WriteTo writes the DataStorePersistenceInitParam to the given writable
func (dspip *DataStorePersistenceInitParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspip.PersistenceSlotID.WriteTo(contentWritable)
	dspip.DeleteLastObject.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePersistenceInitParam from the given readable
func (dspip *DataStorePersistenceInitParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspip.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInitParam header. %s", err.Error())
	}

	err = dspip.PersistenceSlotID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInitParam.PersistenceSlotID. %s", err.Error())
	}

	err = dspip.DeleteLastObject.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInitParam.DeleteLastObject. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePersistenceInitParam
func (dspip *DataStorePersistenceInitParam) Copy() types.RVType {
	copied := NewDataStorePersistenceInitParam()

	copied.StructureVersion = dspip.StructureVersion
	copied.PersistenceSlotID = dspip.PersistenceSlotID.Copy().(*types.PrimitiveU16)
	copied.DeleteLastObject = dspip.DeleteLastObject.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the given DataStorePersistenceInitParam contains the same data as the current DataStorePersistenceInitParam
func (dspip *DataStorePersistenceInitParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePersistenceInitParam); !ok {
		return false
	}

	other := o.(*DataStorePersistenceInitParam)

	if dspip.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspip.PersistenceSlotID.Equals(other.PersistenceSlotID) {
		return false
	}

	return dspip.DeleteLastObject.Equals(other.DeleteLastObject)
}

// String returns the string representation of the DataStorePersistenceInitParam
func (dspip *DataStorePersistenceInitParam) String() string {
	return dspip.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePersistenceInitParam using the provided indentation level
func (dspip *DataStorePersistenceInitParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePersistenceInitParam{\n")
	b.WriteString(fmt.Sprintf("%sPersistenceSlotID: %s,\n", indentationValues, dspip.PersistenceSlotID))
	b.WriteString(fmt.Sprintf("%sDeleteLastObject: %s,\n", indentationValues, dspip.DeleteLastObject))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePersistenceInitParam returns a new DataStorePersistenceInitParam
func NewDataStorePersistenceInitParam() *DataStorePersistenceInitParam {
	dspip := &DataStorePersistenceInitParam{
		PersistenceSlotID: types.NewPrimitiveU16(0),
		DeleteLastObject:  types.NewPrimitiveBool(false),
	}

	return dspip
}
