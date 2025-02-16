// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreReqPostInfo is a type within the DataStore protocol
type DataStoreReqPostInfo struct {
	types.Structure
	DataID         *types.PrimitiveU64
	URL            *types.String
	RequestHeaders *types.List[*DataStoreKeyValue]
	FormFields     *types.List[*DataStoreKeyValue]
	RootCACert     *types.Buffer
}

// WriteTo writes the DataStoreReqPostInfo to the given writable
func (dsrpi *DataStoreReqPostInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrpi.DataID.WriteTo(contentWritable)
	dsrpi.URL.WriteTo(contentWritable)
	dsrpi.RequestHeaders.WriteTo(contentWritable)
	dsrpi.FormFields.WriteTo(contentWritable)
	dsrpi.RootCACert.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrpi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReqPostInfo from the given readable
func (dsrpi *DataStoreReqPostInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrpi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo header. %s", err.Error())
	}

	err = dsrpi.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo.DataID. %s", err.Error())
	}

	err = dsrpi.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo.URL. %s", err.Error())
	}

	err = dsrpi.RequestHeaders.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo.RequestHeaders. %s", err.Error())
	}

	err = dsrpi.FormFields.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo.FormFields. %s", err.Error())
	}

	err = dsrpi.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfo.RootCACert. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReqPostInfo
func (dsrpi *DataStoreReqPostInfo) Copy() types.RVType {
	copied := NewDataStoreReqPostInfo()

	copied.StructureVersion = dsrpi.StructureVersion
	copied.DataID = dsrpi.DataID.Copy().(*types.PrimitiveU64)
	copied.URL = dsrpi.URL.Copy().(*types.String)
	copied.RequestHeaders = dsrpi.RequestHeaders.Copy().(*types.List[*DataStoreKeyValue])
	copied.FormFields = dsrpi.FormFields.Copy().(*types.List[*DataStoreKeyValue])
	copied.RootCACert = dsrpi.RootCACert.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the given DataStoreReqPostInfo contains the same data as the current DataStoreReqPostInfo
func (dsrpi *DataStoreReqPostInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqPostInfo); !ok {
		return false
	}

	other := o.(*DataStoreReqPostInfo)

	if dsrpi.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrpi.DataID.Equals(other.DataID) {
		return false
	}

	if !dsrpi.URL.Equals(other.URL) {
		return false
	}

	if !dsrpi.RequestHeaders.Equals(other.RequestHeaders) {
		return false
	}

	if !dsrpi.FormFields.Equals(other.FormFields) {
		return false
	}

	return dsrpi.RootCACert.Equals(other.RootCACert)
}

// String returns the string representation of the DataStoreReqPostInfo
func (dsrpi *DataStoreReqPostInfo) String() string {
	return dsrpi.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreReqPostInfo using the provided indentation level
func (dsrpi *DataStoreReqPostInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqPostInfo{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsrpi.DataID))
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dsrpi.URL))
	b.WriteString(fmt.Sprintf("%sRequestHeaders: %s,\n", indentationValues, dsrpi.RequestHeaders))
	b.WriteString(fmt.Sprintf("%sFormFields: %s,\n", indentationValues, dsrpi.FormFields))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s,\n", indentationValues, dsrpi.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqPostInfo returns a new DataStoreReqPostInfo
func NewDataStoreReqPostInfo() *DataStoreReqPostInfo {
	dsrpi := &DataStoreReqPostInfo{
		DataID:         types.NewPrimitiveU64(0),
		URL:            types.NewString(""),
		RequestHeaders: types.NewList[*DataStoreKeyValue](),
		FormFields:     types.NewList[*DataStoreKeyValue](),
		RootCACert:     types.NewBuffer(nil),
	}

	dsrpi.RequestHeaders.Type = NewDataStoreKeyValue()
	dsrpi.FormFields.Type = NewDataStoreKeyValue()

	return dsrpi
}
