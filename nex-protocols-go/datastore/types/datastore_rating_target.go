// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreRatingTarget is a type within the DataStore protocol
type DataStoreRatingTarget struct {
	types.Structure
	DataID *types.PrimitiveU64
	Slot   *types.PrimitiveU8
}

// WriteTo writes the DataStoreRatingTarget to the given writable
func (dsrt *DataStoreRatingTarget) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrt.DataID.WriteTo(contentWritable)
	dsrt.Slot.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrt.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRatingTarget from the given readable
func (dsrt *DataStoreRatingTarget) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrt.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingTarget header. %s", err.Error())
	}

	err = dsrt.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingTarget.DataID. %s", err.Error())
	}

	err = dsrt.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingTarget.Slot. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingTarget
func (dsrt *DataStoreRatingTarget) Copy() types.RVType {
	copied := NewDataStoreRatingTarget()

	copied.StructureVersion = dsrt.StructureVersion
	copied.DataID = dsrt.DataID.Copy().(*types.PrimitiveU64)
	copied.Slot = dsrt.Slot.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given DataStoreRatingTarget contains the same data as the current DataStoreRatingTarget
func (dsrt *DataStoreRatingTarget) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingTarget); !ok {
		return false
	}

	other := o.(*DataStoreRatingTarget)

	if dsrt.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrt.DataID.Equals(other.DataID) {
		return false
	}

	return dsrt.Slot.Equals(other.Slot)
}

// String returns the string representation of the DataStoreRatingTarget
func (dsrt *DataStoreRatingTarget) String() string {
	return dsrt.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreRatingTarget using the provided indentation level
func (dsrt *DataStoreRatingTarget) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingTarget{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsrt.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dsrt.Slot))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingTarget returns a new DataStoreRatingTarget
func NewDataStoreRatingTarget() *DataStoreRatingTarget {
	dsrt := &DataStoreRatingTarget{
		DataID: types.NewPrimitiveU64(0),
		Slot:   types.NewPrimitiveU8(0),
	}

	return dsrt
}
