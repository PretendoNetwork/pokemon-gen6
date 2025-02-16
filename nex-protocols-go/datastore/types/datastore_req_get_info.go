// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreReqGetInfo is a type within the DataStore protocol
type DataStoreReqGetInfo struct {
	types.Structure
	URL            *types.String
	RequestHeaders *types.List[*DataStoreKeyValue]
	Size           *types.PrimitiveU32
	RootCACert     *types.Buffer
	DataID         *types.PrimitiveU64 // * NEX v3.5.0
}

// WriteTo writes the DataStoreReqGetInfo to the given writable
func (dsrgi *DataStoreReqGetInfo) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.DataStore

	contentWritable := writable.CopyNew()

	dsrgi.URL.WriteTo(contentWritable)
	dsrgi.RequestHeaders.WriteTo(contentWritable)
	dsrgi.Size.WriteTo(contentWritable)
	dsrgi.RootCACert.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("3.5.0") {
		dsrgi.DataID.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dsrgi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReqGetInfo from the given readable
func (dsrgi *DataStoreReqGetInfo) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.DataStore

	var err error

	err = dsrgi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfo header. %s", err.Error())
	}

	err = dsrgi.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfo.URL. %s", err.Error())
	}

	err = dsrgi.RequestHeaders.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfo.RequestHeaders. %s", err.Error())
	}

	err = dsrgi.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfo.Size. %s", err.Error())
	}

	err = dsrgi.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfo.RootCACert. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.5.0") {
		err = dsrgi.DataID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreReqGetInfo.DataID. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReqGetInfo
func (dsrgi *DataStoreReqGetInfo) Copy() types.RVType {
	copied := NewDataStoreReqGetInfo()

	copied.StructureVersion = dsrgi.StructureVersion
	copied.URL = dsrgi.URL.Copy().(*types.String)
	copied.RequestHeaders = dsrgi.RequestHeaders.Copy().(*types.List[*DataStoreKeyValue])
	copied.Size = dsrgi.Size.Copy().(*types.PrimitiveU32)
	copied.RootCACert = dsrgi.RootCACert.Copy().(*types.Buffer)
	copied.DataID = dsrgi.DataID.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given DataStoreReqGetInfo contains the same data as the current DataStoreReqGetInfo
func (dsrgi *DataStoreReqGetInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqGetInfo); !ok {
		return false
	}

	other := o.(*DataStoreReqGetInfo)

	if dsrgi.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrgi.URL.Equals(other.URL) {
		return false
	}

	if !dsrgi.RequestHeaders.Equals(other.RequestHeaders) {
		return false
	}

	if !dsrgi.Size.Equals(other.Size) {
		return false
	}

	if !dsrgi.RootCACert.Equals(other.RootCACert) {
		return false
	}

	return dsrgi.DataID.Equals(other.DataID)
}

// String returns the string representation of the DataStoreReqGetInfo
func (dsrgi *DataStoreReqGetInfo) String() string {
	return dsrgi.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreReqGetInfo using the provided indentation level
func (dsrgi *DataStoreReqGetInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetInfo{\n")
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dsrgi.URL))
	b.WriteString(fmt.Sprintf("%sRequestHeaders: %s,\n", indentationValues, dsrgi.RequestHeaders))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dsrgi.Size))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s,\n", indentationValues, dsrgi.RootCACert))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsrgi.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetInfo returns a new DataStoreReqGetInfo
func NewDataStoreReqGetInfo() *DataStoreReqGetInfo {
	dsrgi := &DataStoreReqGetInfo{
		URL:            types.NewString(""),
		RequestHeaders: types.NewList[*DataStoreKeyValue](),
		Size:           types.NewPrimitiveU32(0),
		RootCACert:     types.NewBuffer(nil),
		DataID:         types.NewPrimitiveU64(0),
	}

	dsrgi.RequestHeaders.Type = NewDataStoreKeyValue()

	return dsrgi
}
