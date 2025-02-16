// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreCompleteUpdateParam is a type within the DataStore protocol
type DataStoreCompleteUpdateParam struct {
	types.Structure
	DataID    *types.PrimitiveU64
	Version   *types.PrimitiveU32
	IsSuccess *types.PrimitiveBool
}

// WriteTo writes the DataStoreCompleteUpdateParam to the given writable
func (dscup *DataStoreCompleteUpdateParam) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.DataStore

	contentWritable := writable.CopyNew()

	if libraryVersion.GreaterOrEqual("3.0.0") {
		dscup.DataID.WriteTo(contentWritable)
	} else {
		contentWritable.WritePrimitiveUInt32LE(uint32(dscup.DataID.Value))
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		dscup.Version.WriteTo(contentWritable)
	} else {
		contentWritable.WritePrimitiveUInt16LE(uint16(dscup.Version.Value))
	}

	dscup.IsSuccess.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscup.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreCompleteUpdateParam from the given readable
func (dscup *DataStoreCompleteUpdateParam) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.DataStore

	var err error

	err = dscup.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam header. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		err = dscup.DataID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.DataID. %s", err.Error())
		}
	} else {
		dataID, err := readable.ReadPrimitiveUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.DataID. %s", err.Error())
		}

		dscup.DataID.Value = uint64(dataID)
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		err = dscup.Version.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.Version. %s", err.Error())
		}
	} else {
		version, err := readable.ReadPrimitiveUInt16LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.Version. %s", err.Error())
		}

		dscup.Version.Value = uint32(version)
	}

	err = dscup.IsSuccess.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.IsSuccess. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreCompleteUpdateParam
func (dscup *DataStoreCompleteUpdateParam) Copy() types.RVType {
	copied := NewDataStoreCompleteUpdateParam()

	copied.StructureVersion = dscup.StructureVersion
	copied.DataID = dscup.DataID.Copy().(*types.PrimitiveU64)
	copied.Version = dscup.Version.Copy().(*types.PrimitiveU32)
	copied.IsSuccess = dscup.IsSuccess.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the given DataStoreCompleteUpdateParam contains the same data as the current DataStoreCompleteUpdateParam
func (dscup *DataStoreCompleteUpdateParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCompleteUpdateParam); !ok {
		return false
	}

	other := o.(*DataStoreCompleteUpdateParam)

	if dscup.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscup.DataID.Equals(other.DataID) {
		return false
	}

	if !dscup.Version.Equals(other.Version) {
		return false
	}

	return dscup.IsSuccess.Equals(other.IsSuccess)
}

// String returns the string representation of the DataStoreCompleteUpdateParam
func (dscup *DataStoreCompleteUpdateParam) String() string {
	return dscup.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreCompleteUpdateParam using the provided indentation level
func (dscup *DataStoreCompleteUpdateParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompleteUpdateParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dscup.DataID))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, dscup.Version))
	b.WriteString(fmt.Sprintf("%sIsSuccess: %s,\n", indentationValues, dscup.IsSuccess))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompleteUpdateParam returns a new DataStoreCompleteUpdateParam
func NewDataStoreCompleteUpdateParam() *DataStoreCompleteUpdateParam {
	dscup := &DataStoreCompleteUpdateParam{
		DataID:    types.NewPrimitiveU64(0),
		Version:   types.NewPrimitiveU32(0),
		IsSuccess: types.NewPrimitiveBool(false),
	}

	return dscup
}
