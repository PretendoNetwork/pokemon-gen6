// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreChangePlayablePlatformParam is a type within the DataStore protocol
type DataStoreChangePlayablePlatformParam struct {
	types.Structure
	DataID           *types.PrimitiveU64
	PlayablePlatform *types.PrimitiveU32
}

// WriteTo writes the DataStoreChangePlayablePlatformParam to the given writable
func (dscppp *DataStoreChangePlayablePlatformParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscppp.DataID.WriteTo(contentWritable)
	dscppp.PlayablePlatform.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscppp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreChangePlayablePlatformParam from the given readable
func (dscppp *DataStoreChangePlayablePlatformParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dscppp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangePlayablePlatformParam header. %s", err.Error())
	}

	err = dscppp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangePlayablePlatformParam.DataID. %s", err.Error())
	}

	err = dscppp.PlayablePlatform.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangePlayablePlatformParam.PlayablePlatform. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreChangePlayablePlatformParam
func (dscppp *DataStoreChangePlayablePlatformParam) Copy() types.RVType {
	copied := NewDataStoreChangePlayablePlatformParam()

	copied.StructureVersion = dscppp.StructureVersion
	copied.DataID = dscppp.DataID.Copy().(*types.PrimitiveU64)
	copied.PlayablePlatform = dscppp.PlayablePlatform.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given DataStoreChangePlayablePlatformParam contains the same data as the current DataStoreChangePlayablePlatformParam
func (dscppp *DataStoreChangePlayablePlatformParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreChangePlayablePlatformParam); !ok {
		return false
	}

	other := o.(*DataStoreChangePlayablePlatformParam)

	if dscppp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscppp.DataID.Equals(other.DataID) {
		return false
	}

	return dscppp.PlayablePlatform.Equals(other.PlayablePlatform)
}

// String returns the string representation of the DataStoreChangePlayablePlatformParam
func (dscppp *DataStoreChangePlayablePlatformParam) String() string {
	return dscppp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreChangePlayablePlatformParam using the provided indentation level
func (dscppp *DataStoreChangePlayablePlatformParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangePlayablePlatformParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dscppp.DataID))
	b.WriteString(fmt.Sprintf("%sPlayablePlatform: %s,\n", indentationValues, dscppp.PlayablePlatform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangePlayablePlatformParam returns a new DataStoreChangePlayablePlatformParam
func NewDataStoreChangePlayablePlatformParam() *DataStoreChangePlayablePlatformParam {
	dscppp := &DataStoreChangePlayablePlatformParam{
		DataID:           types.NewPrimitiveU64(0),
		PlayablePlatform: types.NewPrimitiveU32(0),
	}

	return dscppp
}
