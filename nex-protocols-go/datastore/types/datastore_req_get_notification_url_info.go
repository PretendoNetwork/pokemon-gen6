// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreReqGetNotificationURLInfo is a type within the DataStore protocol
type DataStoreReqGetNotificationURLInfo struct {
	types.Structure
	URL        *types.String
	Key        *types.String
	Query      *types.String
	RootCACert *types.Buffer
}

// WriteTo writes the DataStoreReqGetNotificationURLInfo to the given writable
func (dsrgnurli *DataStoreReqGetNotificationURLInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrgnurli.URL.WriteTo(contentWritable)
	dsrgnurli.Key.WriteTo(contentWritable)
	dsrgnurli.Query.WriteTo(contentWritable)
	dsrgnurli.RootCACert.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrgnurli.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReqGetNotificationURLInfo from the given readable
func (dsrgnurli *DataStoreReqGetNotificationURLInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrgnurli.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo header. %s", err.Error())
	}

	err = dsrgnurli.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.URL. %s", err.Error())
	}

	err = dsrgnurli.Key.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.Key. %s", err.Error())
	}

	err = dsrgnurli.Query.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.Query. %s", err.Error())
	}

	err = dsrgnurli.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetNotificationURLInfo.RootCACert. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReqGetNotificationURLInfo
func (dsrgnurli *DataStoreReqGetNotificationURLInfo) Copy() types.RVType {
	copied := NewDataStoreReqGetNotificationURLInfo()

	copied.StructureVersion = dsrgnurli.StructureVersion
	copied.URL = dsrgnurli.URL.Copy().(*types.String)
	copied.Key = dsrgnurli.Key.Copy().(*types.String)
	copied.Query = dsrgnurli.Query.Copy().(*types.String)
	copied.RootCACert = dsrgnurli.RootCACert.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the given DataStoreReqGetNotificationURLInfo contains the same data as the current DataStoreReqGetNotificationURLInfo
func (dsrgnurli *DataStoreReqGetNotificationURLInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqGetNotificationURLInfo); !ok {
		return false
	}

	other := o.(*DataStoreReqGetNotificationURLInfo)

	if dsrgnurli.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrgnurli.URL.Equals(other.URL) {
		return false
	}

	if !dsrgnurli.Key.Equals(other.Key) {
		return false
	}

	if !dsrgnurli.Query.Equals(other.Query) {
		return false
	}

	return dsrgnurli.RootCACert.Equals(other.RootCACert)
}

// String returns the string representation of the DataStoreReqGetNotificationURLInfo
func (dsrgnurli *DataStoreReqGetNotificationURLInfo) String() string {
	return dsrgnurli.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreReqGetNotificationURLInfo using the provided indentation level
func (dsrgnurli *DataStoreReqGetNotificationURLInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetNotificationURLInfo{\n")
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dsrgnurli.URL))
	b.WriteString(fmt.Sprintf("%sKey: %s,\n", indentationValues, dsrgnurli.Key))
	b.WriteString(fmt.Sprintf("%sQuery: %s,\n", indentationValues, dsrgnurli.Query))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s,\n", indentationValues, dsrgnurli.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetNotificationURLInfo returns a new DataStoreReqGetNotificationURLInfo
func NewDataStoreReqGetNotificationURLInfo() *DataStoreReqGetNotificationURLInfo {
	dsrgnurli := &DataStoreReqGetNotificationURLInfo{
		URL:        types.NewString(""),
		Key:        types.NewString(""),
		Query:      types.NewString(""),
		RootCACert: types.NewBuffer(nil),
	}

	return dsrgnurli
}
