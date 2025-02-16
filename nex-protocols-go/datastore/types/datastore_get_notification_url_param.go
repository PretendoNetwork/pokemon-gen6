// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetNotificationURLParam is a type within the DataStore protocol
type DataStoreGetNotificationURLParam struct {
	types.Structure
	PreviousURL *types.String
}

// WriteTo writes the DataStoreGetNotificationURLParam to the given writable
func (dsgnurlp *DataStoreGetNotificationURLParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgnurlp.PreviousURL.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgnurlp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetNotificationURLParam from the given readable
func (dsgnurlp *DataStoreGetNotificationURLParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgnurlp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNotificationURLParam header. %s", err.Error())
	}

	err = dsgnurlp.PreviousURL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNotificationURLParam.PreviousURL. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetNotificationURLParam
func (dsgnurlp *DataStoreGetNotificationURLParam) Copy() types.RVType {
	copied := NewDataStoreGetNotificationURLParam()

	copied.StructureVersion = dsgnurlp.StructureVersion
	copied.PreviousURL = dsgnurlp.PreviousURL.Copy().(*types.String)

	return copied
}

// Equals checks if the given DataStoreGetNotificationURLParam contains the same data as the current DataStoreGetNotificationURLParam
func (dsgnurlp *DataStoreGetNotificationURLParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetNotificationURLParam); !ok {
		return false
	}

	other := o.(*DataStoreGetNotificationURLParam)

	if dsgnurlp.StructureVersion != other.StructureVersion {
		return false
	}

	return dsgnurlp.PreviousURL.Equals(other.PreviousURL)
}

// String returns the string representation of the DataStoreGetNotificationURLParam
func (dsgnurlp *DataStoreGetNotificationURLParam) String() string {
	return dsgnurlp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetNotificationURLParam using the provided indentation level
func (dsgnurlp *DataStoreGetNotificationURLParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetNotificationURLParam{\n")
	b.WriteString(fmt.Sprintf("%sPreviousURL: %s,\n", indentationValues, dsgnurlp.PreviousURL))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetNotificationURLParam returns a new DataStoreGetNotificationURLParam
func NewDataStoreGetNotificationURLParam() *DataStoreGetNotificationURLParam {
	dsgnurlp := &DataStoreGetNotificationURLParam{
		PreviousURL: types.NewString(""),
	}

	return dsgnurlp
}
