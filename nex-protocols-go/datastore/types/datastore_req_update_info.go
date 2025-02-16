// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreReqUpdateInfo is a type within the DataStore protocol
type DataStoreReqUpdateInfo struct {
	types.Structure
	Version        *types.PrimitiveU32
	URL            *types.String
	RequestHeaders *types.List[*DataStoreKeyValue]
	FormFields     *types.List[*DataStoreKeyValue]
	RootCACert     *types.Buffer
}

// WriteTo writes the DataStoreReqUpdateInfo to the given writable
func (dsrui *DataStoreReqUpdateInfo) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.DataStore

	contentWritable := writable.CopyNew()

	if libraryVersion.GreaterOrEqual("3.0.0") {
		dsrui.Version.WriteTo(contentWritable)
	} else {
		contentWritable.WritePrimitiveUInt16LE(uint16(dsrui.Version.Value))
	}

	dsrui.URL.WriteTo(contentWritable)
	dsrui.RequestHeaders.WriteTo(contentWritable)
	dsrui.FormFields.WriteTo(contentWritable)
	dsrui.RootCACert.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrui.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReqUpdateInfo from the given readable
func (dsrui *DataStoreReqUpdateInfo) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.DataStore

	var err error

	err = dsrui.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo header. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		err = dsrui.Version.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.Version. %s", err.Error())
		}
	} else {
		version, err := readable.ReadPrimitiveUInt16LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.Version. %s", err.Error())
		}

		dsrui.Version.Value = uint32(version)
	}

	err = dsrui.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.URL. %s", err.Error())
	}

	err = dsrui.RequestHeaders.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.RequestHeaders. %s", err.Error())
	}

	err = dsrui.FormFields.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.FormFields. %s", err.Error())
	}

	err = dsrui.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.RootCACert. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReqUpdateInfo
func (dsrui *DataStoreReqUpdateInfo) Copy() types.RVType {
	copied := NewDataStoreReqUpdateInfo()

	copied.StructureVersion = dsrui.StructureVersion
	copied.Version = dsrui.Version.Copy().(*types.PrimitiveU32)
	copied.URL = dsrui.URL.Copy().(*types.String)
	copied.RequestHeaders = dsrui.RequestHeaders.Copy().(*types.List[*DataStoreKeyValue])
	copied.FormFields = dsrui.FormFields.Copy().(*types.List[*DataStoreKeyValue])
	copied.RootCACert = dsrui.RootCACert.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the given DataStoreReqUpdateInfo contains the same data as the current DataStoreReqUpdateInfo
func (dsrui *DataStoreReqUpdateInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqUpdateInfo); !ok {
		return false
	}

	other := o.(*DataStoreReqUpdateInfo)

	if dsrui.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrui.Version.Equals(other.Version) {
		return false
	}

	if !dsrui.URL.Equals(other.URL) {
		return false
	}

	if !dsrui.RequestHeaders.Equals(other.RequestHeaders) {
		return false
	}

	if !dsrui.FormFields.Equals(other.FormFields) {
		return false
	}

	return dsrui.RootCACert.Equals(other.RootCACert)
}

// String returns the string representation of the DataStoreReqUpdateInfo
func (dsrui *DataStoreReqUpdateInfo) String() string {
	return dsrui.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreReqUpdateInfo using the provided indentation level
func (dsrui *DataStoreReqUpdateInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqUpdateInfo{\n")
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, dsrui.Version))
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dsrui.URL))
	b.WriteString(fmt.Sprintf("%sRequestHeaders: %s,\n", indentationValues, dsrui.RequestHeaders))
	b.WriteString(fmt.Sprintf("%sFormFields: %s,\n", indentationValues, dsrui.FormFields))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s,\n", indentationValues, dsrui.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqUpdateInfo returns a new DataStoreReqUpdateInfo
func NewDataStoreReqUpdateInfo() *DataStoreReqUpdateInfo {
	dsrui := &DataStoreReqUpdateInfo{
		Version:        types.NewPrimitiveU32(0),
		URL:            types.NewString(""),
		RequestHeaders: types.NewList[*DataStoreKeyValue](),
		FormFields:     types.NewList[*DataStoreKeyValue](),
		RootCACert:     types.NewBuffer(nil),
	}

	dsrui.RequestHeaders.Type = NewDataStoreKeyValue()
	dsrui.FormFields.Type = NewDataStoreKeyValue()

	return dsrui
}
