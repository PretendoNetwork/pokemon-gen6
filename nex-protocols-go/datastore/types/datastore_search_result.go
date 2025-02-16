// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreSearchResult is a type within the DataStore protocol
type DataStoreSearchResult struct {
	types.Structure
	TotalCount     *types.PrimitiveU32
	Result         *types.List[*DataStoreMetaInfo]
	TotalCountType *types.PrimitiveU8
}

// WriteTo writes the DataStoreSearchResult to the given writable
func (dssr *DataStoreSearchResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dssr.TotalCount.WriteTo(contentWritable)
	dssr.Result.WriteTo(contentWritable)
	dssr.TotalCountType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dssr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreSearchResult from the given readable
func (dssr *DataStoreSearchResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = dssr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchResult header. %s", err.Error())
	}

	err = dssr.TotalCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchResult.TotalCount. %s", err.Error())
	}

	err = dssr.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchResult.Result. %s", err.Error())
	}

	err = dssr.TotalCountType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchResult.TotalCountType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSearchResult
func (dssr *DataStoreSearchResult) Copy() types.RVType {
	copied := NewDataStoreSearchResult()

	copied.StructureVersion = dssr.StructureVersion
	copied.TotalCount = dssr.TotalCount.Copy().(*types.PrimitiveU32)
	copied.Result = dssr.Result.Copy().(*types.List[*DataStoreMetaInfo])
	copied.TotalCountType = dssr.TotalCountType.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given DataStoreSearchResult contains the same data as the current DataStoreSearchResult
func (dssr *DataStoreSearchResult) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSearchResult); !ok {
		return false
	}

	other := o.(*DataStoreSearchResult)

	if dssr.StructureVersion != other.StructureVersion {
		return false
	}

	if !dssr.TotalCount.Equals(other.TotalCount) {
		return false
	}

	if !dssr.Result.Equals(other.Result) {
		return false
	}

	return dssr.TotalCountType.Equals(other.TotalCountType)
}

// String returns the string representation of the DataStoreSearchResult
func (dssr *DataStoreSearchResult) String() string {
	return dssr.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreSearchResult using the provided indentation level
func (dssr *DataStoreSearchResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchResult{\n")
	b.WriteString(fmt.Sprintf("%sTotalCount: %s,\n", indentationValues, dssr.TotalCount))
	b.WriteString(fmt.Sprintf("%sResult: %s,\n", indentationValues, dssr.Result))
	b.WriteString(fmt.Sprintf("%sTotalCountType: %s,\n", indentationValues, dssr.TotalCountType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchResult returns a new DataStoreSearchResult
func NewDataStoreSearchResult() *DataStoreSearchResult {
	dssr := &DataStoreSearchResult{
		TotalCount:     types.NewPrimitiveU32(0),
		Result:         types.NewList[*DataStoreMetaInfo](),
		TotalCountType: types.NewPrimitiveU8(0),
	}

	dssr.Result.Type = NewDataStoreMetaInfo()

	return dssr
}
