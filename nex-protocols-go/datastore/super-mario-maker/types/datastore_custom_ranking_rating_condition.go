// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreCustomRankingRatingCondition is a type within the DataStore protocol
type DataStoreCustomRankingRatingCondition struct {
	types.Structure
	Slot     *types.PrimitiveS8
	MinValue *types.PrimitiveS32
	MaxValue *types.PrimitiveS32
	MinCount *types.PrimitiveS32 // * Revision 1
	MaxCount *types.PrimitiveS32 // * Revision 1
}

// WriteTo writes the DataStoreCustomRankingRatingCondition to the given writable
func (dscrrc *DataStoreCustomRankingRatingCondition) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscrrc.Slot.WriteTo(contentWritable)
	dscrrc.MinValue.WriteTo(contentWritable)
	dscrrc.MaxValue.WriteTo(contentWritable)

	if dscrrc.StructureVersion >= 1 {
		dscrrc.MinCount.WriteTo(contentWritable)
	}

	if dscrrc.StructureVersion >= 1 {
		dscrrc.MaxCount.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dscrrc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreCustomRankingRatingCondition from the given readable
func (dscrrc *DataStoreCustomRankingRatingCondition) ExtractFrom(readable types.Readable) error {
	var err error

	err = dscrrc.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition header. %s", err.Error())
	}

	err = dscrrc.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.Slot. %s", err.Error())
	}

	err = dscrrc.MinValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.MinValue. %s", err.Error())
	}

	err = dscrrc.MaxValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.MaxValue. %s", err.Error())
	}

	if dscrrc.StructureVersion >= 1 {
		err = dscrrc.MinCount.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.MinCount. %s", err.Error())
		}
	}

	if dscrrc.StructureVersion >= 1 {
		err = dscrrc.MaxCount.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCustomRankingRatingCondition.MaxCount. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStoreCustomRankingRatingCondition
func (dscrrc *DataStoreCustomRankingRatingCondition) Copy() types.RVType {
	copied := NewDataStoreCustomRankingRatingCondition()

	copied.StructureVersion = dscrrc.StructureVersion
	copied.Slot = dscrrc.Slot.Copy().(*types.PrimitiveS8)
	copied.MinValue = dscrrc.MinValue.Copy().(*types.PrimitiveS32)
	copied.MaxValue = dscrrc.MaxValue.Copy().(*types.PrimitiveS32)
	copied.MinCount = dscrrc.MinCount.Copy().(*types.PrimitiveS32)
	copied.MaxCount = dscrrc.MaxCount.Copy().(*types.PrimitiveS32)

	return copied
}

// Equals checks if the given DataStoreCustomRankingRatingCondition contains the same data as the current DataStoreCustomRankingRatingCondition
func (dscrrc *DataStoreCustomRankingRatingCondition) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCustomRankingRatingCondition); !ok {
		return false
	}

	other := o.(*DataStoreCustomRankingRatingCondition)

	if dscrrc.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscrrc.Slot.Equals(other.Slot) {
		return false
	}

	if !dscrrc.MinValue.Equals(other.MinValue) {
		return false
	}

	if !dscrrc.MaxValue.Equals(other.MaxValue) {
		return false
	}

	if !dscrrc.MinCount.Equals(other.MinCount) {
		return false
	}

	return dscrrc.MaxCount.Equals(other.MaxCount)
}

// String returns the string representation of the DataStoreCustomRankingRatingCondition
func (dscrrc *DataStoreCustomRankingRatingCondition) String() string {
	return dscrrc.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreCustomRankingRatingCondition using the provided indentation level
func (dscrrc *DataStoreCustomRankingRatingCondition) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCustomRankingRatingCondition{\n")
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dscrrc.Slot))
	b.WriteString(fmt.Sprintf("%sMinValue: %s,\n", indentationValues, dscrrc.MinValue))
	b.WriteString(fmt.Sprintf("%sMaxValue: %s,\n", indentationValues, dscrrc.MaxValue))
	b.WriteString(fmt.Sprintf("%sMinCount: %s,\n", indentationValues, dscrrc.MinCount))
	b.WriteString(fmt.Sprintf("%sMaxCount: %s,\n", indentationValues, dscrrc.MaxCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCustomRankingRatingCondition returns a new DataStoreCustomRankingRatingCondition
func NewDataStoreCustomRankingRatingCondition() *DataStoreCustomRankingRatingCondition {
	dscrrc := &DataStoreCustomRankingRatingCondition{
		Slot:     types.NewPrimitiveS8(0),
		MinValue: types.NewPrimitiveS32(0),
		MaxValue: types.NewPrimitiveS32(0),
		MinCount: types.NewPrimitiveS32(0),
		MaxCount: types.NewPrimitiveS32(0),
	}

	return dscrrc
}
