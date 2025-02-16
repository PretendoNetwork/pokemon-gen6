// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreRatingInitParamWithSlot is a type within the DataStore protocol
type DataStoreRatingInitParamWithSlot struct {
	types.Structure
	Slot  *types.PrimitiveS8
	Param *DataStoreRatingInitParam
}

// WriteTo writes the DataStoreRatingInitParamWithSlot to the given writable
func (dsripws *DataStoreRatingInitParamWithSlot) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsripws.Slot.WriteTo(contentWritable)
	dsripws.Param.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsripws.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRatingInitParamWithSlot from the given readable
func (dsripws *DataStoreRatingInitParamWithSlot) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsripws.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParamWithSlot header. %s", err.Error())
	}

	err = dsripws.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParamWithSlot.Slot. %s", err.Error())
	}

	err = dsripws.Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParamWithSlot.Param. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingInitParamWithSlot
func (dsripws *DataStoreRatingInitParamWithSlot) Copy() types.RVType {
	copied := NewDataStoreRatingInitParamWithSlot()

	copied.StructureVersion = dsripws.StructureVersion
	copied.Slot = dsripws.Slot.Copy().(*types.PrimitiveS8)
	copied.Param = dsripws.Param.Copy().(*DataStoreRatingInitParam)

	return copied
}

// Equals checks if the given DataStoreRatingInitParamWithSlot contains the same data as the current DataStoreRatingInitParamWithSlot
func (dsripws *DataStoreRatingInitParamWithSlot) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingInitParamWithSlot); !ok {
		return false
	}

	other := o.(*DataStoreRatingInitParamWithSlot)

	if dsripws.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsripws.Slot.Equals(other.Slot) {
		return false
	}

	return dsripws.Param.Equals(other.Param)
}

// String returns the string representation of the DataStoreRatingInitParamWithSlot
func (dsripws *DataStoreRatingInitParamWithSlot) String() string {
	return dsripws.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreRatingInitParamWithSlot using the provided indentation level
func (dsripws *DataStoreRatingInitParamWithSlot) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingInitParamWithSlot{\n")
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dsripws.Slot))
	b.WriteString(fmt.Sprintf("%sParam: %s,\n", indentationValues, dsripws.Param.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInitParamWithSlot returns a new DataStoreRatingInitParamWithSlot
func NewDataStoreRatingInitParamWithSlot() *DataStoreRatingInitParamWithSlot {
	dsripws := &DataStoreRatingInitParamWithSlot{
		Slot:  types.NewPrimitiveS8(0),
		Param: NewDataStoreRatingInitParam(),
	}

	return dsripws
}
