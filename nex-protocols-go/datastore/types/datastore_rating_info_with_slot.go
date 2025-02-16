// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreRatingInfoWithSlot is a type within the DataStore protocol
type DataStoreRatingInfoWithSlot struct {
	types.Structure
	Slot   *types.PrimitiveS8
	Rating *DataStoreRatingInfo
}

// WriteTo writes the DataStoreRatingInfoWithSlot to the given writable
func (dsriws *DataStoreRatingInfoWithSlot) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsriws.Slot.WriteTo(contentWritable)
	dsriws.Rating.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsriws.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRatingInfoWithSlot from the given readable
func (dsriws *DataStoreRatingInfoWithSlot) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsriws.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfoWithSlot header. %s", err.Error())
	}

	err = dsriws.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfoWithSlot.Slot. %s", err.Error())
	}

	err = dsriws.Rating.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfoWithSlot.Rating. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingInfoWithSlot
func (dsriws *DataStoreRatingInfoWithSlot) Copy() types.RVType {
	copied := NewDataStoreRatingInfoWithSlot()

	copied.StructureVersion = dsriws.StructureVersion
	copied.Slot = dsriws.Slot.Copy().(*types.PrimitiveS8)
	copied.Rating = dsriws.Rating.Copy().(*DataStoreRatingInfo)

	return copied
}

// Equals checks if the given DataStoreRatingInfoWithSlot contains the same data as the current DataStoreRatingInfoWithSlot
func (dsriws *DataStoreRatingInfoWithSlot) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingInfoWithSlot); !ok {
		return false
	}

	other := o.(*DataStoreRatingInfoWithSlot)

	if dsriws.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsriws.Slot.Equals(other.Slot) {
		return false
	}

	return dsriws.Rating.Equals(other.Rating)
}

// String returns the string representation of the DataStoreRatingInfoWithSlot
func (dsriws *DataStoreRatingInfoWithSlot) String() string {
	return dsriws.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreRatingInfoWithSlot using the provided indentation level
func (dsriws *DataStoreRatingInfoWithSlot) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingInfoWithSlot{\n")
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dsriws.Slot))
	b.WriteString(fmt.Sprintf("%sRating: %s,\n", indentationValues, dsriws.Rating.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInfoWithSlot returns a new DataStoreRatingInfoWithSlot
func NewDataStoreRatingInfoWithSlot() *DataStoreRatingInfoWithSlot {
	dsriws := &DataStoreRatingInfoWithSlot{
		Slot:   types.NewPrimitiveS8(0),
		Rating: NewDataStoreRatingInfo(),
	}

	return dsriws
}
