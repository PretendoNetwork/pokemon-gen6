// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreRatingInfo is a type within the DataStore protocol
type DataStoreRatingInfo struct {
	types.Structure
	TotalValue   *types.PrimitiveS64
	Count        *types.PrimitiveU32
	InitialValue *types.PrimitiveS64
}

// WriteTo writes the DataStoreRatingInfo to the given writable
func (dsri *DataStoreRatingInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsri.TotalValue.WriteTo(contentWritable)
	dsri.Count.WriteTo(contentWritable)
	dsri.InitialValue.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsri.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRatingInfo from the given readable
func (dsri *DataStoreRatingInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsri.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfo header. %s", err.Error())
	}

	err = dsri.TotalValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfo.TotalValue. %s", err.Error())
	}

	err = dsri.Count.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfo.Count. %s", err.Error())
	}

	err = dsri.InitialValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfo.InitialValue. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingInfo
func (dsri *DataStoreRatingInfo) Copy() types.RVType {
	copied := NewDataStoreRatingInfo()

	copied.StructureVersion = dsri.StructureVersion
	copied.TotalValue = dsri.TotalValue.Copy().(*types.PrimitiveS64)
	copied.Count = dsri.Count.Copy().(*types.PrimitiveU32)
	copied.InitialValue = dsri.InitialValue.Copy().(*types.PrimitiveS64)

	return copied
}

// Equals checks if the given DataStoreRatingInfo contains the same data as the current DataStoreRatingInfo
func (dsri *DataStoreRatingInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingInfo); !ok {
		return false
	}

	other := o.(*DataStoreRatingInfo)

	if dsri.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsri.TotalValue.Equals(other.TotalValue) {
		return false
	}

	if !dsri.Count.Equals(other.Count) {
		return false
	}

	return dsri.InitialValue.Equals(other.InitialValue)
}

// String returns the string representation of the DataStoreRatingInfo
func (dsri *DataStoreRatingInfo) String() string {
	return dsri.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreRatingInfo using the provided indentation level
func (dsri *DataStoreRatingInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingInfo{\n")
	b.WriteString(fmt.Sprintf("%sTotalValue: %s,\n", indentationValues, dsri.TotalValue))
	b.WriteString(fmt.Sprintf("%sCount: %s,\n", indentationValues, dsri.Count))
	b.WriteString(fmt.Sprintf("%sInitialValue: %s,\n", indentationValues, dsri.InitialValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInfo returns a new DataStoreRatingInfo
func NewDataStoreRatingInfo() *DataStoreRatingInfo {
	dsri := &DataStoreRatingInfo{
		TotalValue:   types.NewPrimitiveS64(0),
		Count:        types.NewPrimitiveU32(0),
		InitialValue: types.NewPrimitiveS64(0),
	}

	return dsri
}
