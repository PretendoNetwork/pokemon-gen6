// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePrepareUpdateParam is a type within the DataStore protocol
type DataStorePrepareUpdateParam struct {
	types.Structure
	DataID         *types.PrimitiveU64
	Size           *types.PrimitiveU32
	UpdatePassword *types.PrimitiveU64        // * NEX v3.0.0
	ExtraData      *types.List[*types.String] // * NEX v3.5.0
}

// WriteTo writes the DataStorePrepareUpdateParam to the given writable
func (dspup *DataStorePrepareUpdateParam) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.DataStore

	contentWritable := writable.CopyNew()

	if libraryVersion.GreaterOrEqual("3.0.0") {
		dspup.DataID.WriteTo(contentWritable)
	} else {
		contentWritable.WritePrimitiveUInt32LE(uint32(dspup.DataID.Value))
	}

	dspup.Size.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("3.0.0") {
		dspup.UpdatePassword.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.5.0") {
		dspup.ExtraData.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dspup.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePrepareUpdateParam from the given readable
func (dspup *DataStorePrepareUpdateParam) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.DataStore

	var err error

	err = dspup.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam header. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		err = dspup.DataID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.DataID. %s", err.Error())
		}
	} else {
		dataID, err := readable.ReadPrimitiveUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.DataID. %s", err.Error())
		}

		dspup.DataID.Value = uint64(dataID)
	}

	err = dspup.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.Size. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		err = dspup.UpdatePassword.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.UpdatePassword. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.5.0") {
		err = dspup.ExtraData.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.ExtraData. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStorePrepareUpdateParam
func (dspup *DataStorePrepareUpdateParam) Copy() types.RVType {
	copied := NewDataStorePrepareUpdateParam()

	copied.StructureVersion = dspup.StructureVersion
	copied.DataID = dspup.DataID.Copy().(*types.PrimitiveU64)
	copied.Size = dspup.Size.Copy().(*types.PrimitiveU32)
	copied.UpdatePassword = dspup.UpdatePassword.Copy().(*types.PrimitiveU64)
	copied.ExtraData = dspup.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the given DataStorePrepareUpdateParam contains the same data as the current DataStorePrepareUpdateParam
func (dspup *DataStorePrepareUpdateParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePrepareUpdateParam); !ok {
		return false
	}

	other := o.(*DataStorePrepareUpdateParam)

	if dspup.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspup.DataID.Equals(other.DataID) {
		return false
	}

	if !dspup.Size.Equals(other.Size) {
		return false
	}

	if !dspup.UpdatePassword.Equals(other.UpdatePassword) {
		return false
	}

	return dspup.ExtraData.Equals(other.ExtraData)
}

// String returns the string representation of the DataStorePrepareUpdateParam
func (dspup *DataStorePrepareUpdateParam) String() string {
	return dspup.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePrepareUpdateParam using the provided indentation level
func (dspup *DataStorePrepareUpdateParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePrepareUpdateParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dspup.DataID))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dspup.Size))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %s,\n", indentationValues, dspup.UpdatePassword))
	b.WriteString(fmt.Sprintf("%sExtraData: %s,\n", indentationValues, dspup.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareUpdateParam returns a new DataStorePrepareUpdateParam
func NewDataStorePrepareUpdateParam() *DataStorePrepareUpdateParam {
	dspup := &DataStorePrepareUpdateParam{
		DataID:         types.NewPrimitiveU64(0),
		Size:           types.NewPrimitiveU32(0),
		UpdatePassword: types.NewPrimitiveU64(0),
		ExtraData:      types.NewList[*types.String](),
	}

	dspup.ExtraData.Type = types.NewString("")

	return dspup
}
