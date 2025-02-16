// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePersistenceTarget is a type within the DataStore protocol
type DataStorePersistenceTarget struct {
	types.Structure
	OwnerID           *types.PID
	PersistenceSlotID *types.PrimitiveU16
}

// WriteTo writes the DataStorePersistenceTarget to the given writable
func (dspt *DataStorePersistenceTarget) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspt.OwnerID.WriteTo(contentWritable)
	dspt.PersistenceSlotID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspt.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePersistenceTarget from the given readable
func (dspt *DataStorePersistenceTarget) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspt.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceTarget header. %s", err.Error())
	}

	err = dspt.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceTarget.OwnerID. %s", err.Error())
	}

	err = dspt.PersistenceSlotID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceTarget.PersistenceSlotID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePersistenceTarget
func (dspt *DataStorePersistenceTarget) Copy() types.RVType {
	copied := NewDataStorePersistenceTarget()

	copied.StructureVersion = dspt.StructureVersion
	copied.OwnerID = dspt.OwnerID.Copy().(*types.PID)
	copied.PersistenceSlotID = dspt.PersistenceSlotID.Copy().(*types.PrimitiveU16)

	return copied
}

// Equals checks if the given DataStorePersistenceTarget contains the same data as the current DataStorePersistenceTarget
func (dspt *DataStorePersistenceTarget) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePersistenceTarget); !ok {
		return false
	}

	other := o.(*DataStorePersistenceTarget)

	if dspt.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspt.OwnerID.Equals(other.OwnerID) {
		return false
	}

	return dspt.PersistenceSlotID.Equals(other.PersistenceSlotID)
}

// String returns the string representation of the DataStorePersistenceTarget
func (dspt *DataStorePersistenceTarget) String() string {
	return dspt.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePersistenceTarget using the provided indentation level
func (dspt *DataStorePersistenceTarget) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePersistenceTarget{\n")
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, dspt.OwnerID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPersistenceSlotID: %s,\n", indentationValues, dspt.PersistenceSlotID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePersistenceTarget returns a new DataStorePersistenceTarget
func NewDataStorePersistenceTarget() *DataStorePersistenceTarget {
	dspt := &DataStorePersistenceTarget{
		OwnerID:           types.NewPID(0),
		PersistenceSlotID: types.NewPrimitiveU16(0),
	}

	return dspt
}
