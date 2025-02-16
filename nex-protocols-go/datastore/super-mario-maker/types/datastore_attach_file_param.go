// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

// DataStoreAttachFileParam is a type within the DataStore protocol
type DataStoreAttachFileParam struct {
	types.Structure
	PostParam   *datastore_types.DataStorePreparePostParam
	ReferDataID *types.PrimitiveU64
	ContentType *types.String
}

// WriteTo writes the DataStoreAttachFileParam to the given writable
func (dsafp *DataStoreAttachFileParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsafp.PostParam.WriteTo(contentWritable)
	dsafp.ReferDataID.WriteTo(contentWritable)
	dsafp.ContentType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsafp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreAttachFileParam from the given readable
func (dsafp *DataStoreAttachFileParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsafp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam header. %s", err.Error())
	}

	err = dsafp.PostParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.PostParam. %s", err.Error())
	}

	err = dsafp.ReferDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.ReferDataID. %s", err.Error())
	}

	err = dsafp.ContentType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.ContentType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreAttachFileParam
func (dsafp *DataStoreAttachFileParam) Copy() types.RVType {
	copied := NewDataStoreAttachFileParam()

	copied.StructureVersion = dsafp.StructureVersion
	copied.PostParam = dsafp.PostParam.Copy().(*datastore_types.DataStorePreparePostParam)
	copied.ReferDataID = dsafp.ReferDataID.Copy().(*types.PrimitiveU64)
	copied.ContentType = dsafp.ContentType.Copy().(*types.String)

	return copied
}

// Equals checks if the given DataStoreAttachFileParam contains the same data as the current DataStoreAttachFileParam
func (dsafp *DataStoreAttachFileParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreAttachFileParam); !ok {
		return false
	}

	other := o.(*DataStoreAttachFileParam)

	if dsafp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsafp.PostParam.Equals(other.PostParam) {
		return false
	}

	if !dsafp.ReferDataID.Equals(other.ReferDataID) {
		return false
	}

	return dsafp.ContentType.Equals(other.ContentType)
}

// String returns the string representation of the DataStoreAttachFileParam
func (dsafp *DataStoreAttachFileParam) String() string {
	return dsafp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreAttachFileParam using the provided indentation level
func (dsafp *DataStoreAttachFileParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreAttachFileParam{\n")
	b.WriteString(fmt.Sprintf("%sPostParam: %s,\n", indentationValues, dsafp.PostParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sReferDataID: %s,\n", indentationValues, dsafp.ReferDataID))
	b.WriteString(fmt.Sprintf("%sContentType: %s,\n", indentationValues, dsafp.ContentType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreAttachFileParam returns a new DataStoreAttachFileParam
func NewDataStoreAttachFileParam() *DataStoreAttachFileParam {
	dsafp := &DataStoreAttachFileParam{
		PostParam:   datastore_types.NewDataStorePreparePostParam(),
		ReferDataID: types.NewPrimitiveU64(0),
		ContentType: types.NewString(""),
	}

	return dsafp
}
