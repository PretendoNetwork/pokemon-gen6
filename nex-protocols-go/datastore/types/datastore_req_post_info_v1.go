// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreReqPostInfoV1 is a type within the DataStore protocol
type DataStoreReqPostInfoV1 struct {
	types.Structure
	DataID         *types.PrimitiveU32
	URL            *types.String
	RequestHeaders *types.List[*DataStoreKeyValue]
	FormFields     *types.List[*DataStoreKeyValue]
	RootCACert     *types.Buffer
}

// WriteTo writes the DataStoreReqPostInfoV1 to the given writable
func (dsrpiv *DataStoreReqPostInfoV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrpiv.DataID.WriteTo(contentWritable)
	dsrpiv.URL.WriteTo(contentWritable)
	dsrpiv.RequestHeaders.WriteTo(contentWritable)
	dsrpiv.FormFields.WriteTo(contentWritable)
	dsrpiv.RootCACert.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrpiv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReqPostInfoV1 from the given readable
func (dsrpiv *DataStoreReqPostInfoV1) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrpiv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1 header. %s", err.Error())
	}

	err = dsrpiv.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.DataID. %s", err.Error())
	}

	err = dsrpiv.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.URL. %s", err.Error())
	}

	err = dsrpiv.RequestHeaders.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.RequestHeaders. %s", err.Error())
	}

	err = dsrpiv.FormFields.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.FormFields. %s", err.Error())
	}

	err = dsrpiv.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.RootCACert. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReqPostInfoV1
func (dsrpiv *DataStoreReqPostInfoV1) Copy() types.RVType {
	copied := NewDataStoreReqPostInfoV1()

	copied.StructureVersion = dsrpiv.StructureVersion
	copied.DataID = dsrpiv.DataID.Copy().(*types.PrimitiveU32)
	copied.URL = dsrpiv.URL.Copy().(*types.String)
	copied.RequestHeaders = dsrpiv.RequestHeaders.Copy().(*types.List[*DataStoreKeyValue])
	copied.FormFields = dsrpiv.FormFields.Copy().(*types.List[*DataStoreKeyValue])
	copied.RootCACert = dsrpiv.RootCACert.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the given DataStoreReqPostInfoV1 contains the same data as the current DataStoreReqPostInfoV1
func (dsrpiv *DataStoreReqPostInfoV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqPostInfoV1); !ok {
		return false
	}

	other := o.(*DataStoreReqPostInfoV1)

	if dsrpiv.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrpiv.DataID.Equals(other.DataID) {
		return false
	}

	if !dsrpiv.URL.Equals(other.URL) {
		return false
	}

	if !dsrpiv.RequestHeaders.Equals(other.RequestHeaders) {
		return false
	}

	if !dsrpiv.FormFields.Equals(other.FormFields) {
		return false
	}

	return dsrpiv.RootCACert.Equals(other.RootCACert)
}

// String returns the string representation of the DataStoreReqPostInfoV1
func (dsrpiv *DataStoreReqPostInfoV1) String() string {
	return dsrpiv.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreReqPostInfoV1 using the provided indentation level
func (dsrpiv *DataStoreReqPostInfoV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqPostInfoV1{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsrpiv.DataID))
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dsrpiv.URL))
	b.WriteString(fmt.Sprintf("%sRequestHeaders: %s,\n", indentationValues, dsrpiv.RequestHeaders))
	b.WriteString(fmt.Sprintf("%sFormFields: %s,\n", indentationValues, dsrpiv.FormFields))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s,\n", indentationValues, dsrpiv.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqPostInfoV1 returns a new DataStoreReqPostInfoV1
func NewDataStoreReqPostInfoV1() *DataStoreReqPostInfoV1 {
	dsrpiv := &DataStoreReqPostInfoV1{
		DataID:         types.NewPrimitiveU32(0),
		URL:            types.NewString(""),
		RequestHeaders: types.NewList[*DataStoreKeyValue](),
		FormFields:     types.NewList[*DataStoreKeyValue](),
		RootCACert:     types.NewBuffer(nil),
	}

	dsrpiv.RequestHeaders.Type = NewDataStoreKeyValue()
	dsrpiv.FormFields.Type = NewDataStoreKeyValue()

	return dsrpiv
}
