// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreReqGetInfoV1 is a type within the DataStore protocol
type DataStoreReqGetInfoV1 struct {
	types.Structure
	URL            *types.String
	RequestHeaders *types.List[*DataStoreKeyValue]
	Size           *types.PrimitiveU32
	RootCACert     *types.Buffer
}

// WriteTo writes the DataStoreReqGetInfoV1 to the given writable
func (dsrgiv *DataStoreReqGetInfoV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrgiv.URL.WriteTo(contentWritable)
	dsrgiv.RequestHeaders.WriteTo(contentWritable)
	dsrgiv.Size.WriteTo(contentWritable)
	dsrgiv.RootCACert.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrgiv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReqGetInfoV1 from the given readable
func (dsrgiv *DataStoreReqGetInfoV1) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrgiv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1 header. %s", err.Error())
	}

	err = dsrgiv.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.URL. %s", err.Error())
	}

	err = dsrgiv.RequestHeaders.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.RequestHeaders. %s", err.Error())
	}

	err = dsrgiv.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.Size. %s", err.Error())
	}

	err = dsrgiv.RootCACert.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.RootCACert. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReqGetInfoV1
func (dsrgiv *DataStoreReqGetInfoV1) Copy() types.RVType {
	copied := NewDataStoreReqGetInfoV1()

	copied.StructureVersion = dsrgiv.StructureVersion
	copied.URL = dsrgiv.URL.Copy().(*types.String)
	copied.RequestHeaders = dsrgiv.RequestHeaders.Copy().(*types.List[*DataStoreKeyValue])
	copied.Size = dsrgiv.Size.Copy().(*types.PrimitiveU32)
	copied.RootCACert = dsrgiv.RootCACert.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the given DataStoreReqGetInfoV1 contains the same data as the current DataStoreReqGetInfoV1
func (dsrgiv *DataStoreReqGetInfoV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReqGetInfoV1); !ok {
		return false
	}

	other := o.(*DataStoreReqGetInfoV1)

	if dsrgiv.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrgiv.URL.Equals(other.URL) {
		return false
	}

	if !dsrgiv.RequestHeaders.Equals(other.RequestHeaders) {
		return false
	}

	if !dsrgiv.Size.Equals(other.Size) {
		return false
	}

	return dsrgiv.RootCACert.Equals(other.RootCACert)
}

// String returns the string representation of the DataStoreReqGetInfoV1
func (dsrgiv *DataStoreReqGetInfoV1) String() string {
	return dsrgiv.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreReqGetInfoV1 using the provided indentation level
func (dsrgiv *DataStoreReqGetInfoV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetInfoV1{\n")
	b.WriteString(fmt.Sprintf("%sURL: %s,\n", indentationValues, dsrgiv.URL))
	b.WriteString(fmt.Sprintf("%sRequestHeaders: %s,\n", indentationValues, dsrgiv.RequestHeaders))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dsrgiv.Size))
	b.WriteString(fmt.Sprintf("%sRootCACert: %s,\n", indentationValues, dsrgiv.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetInfoV1 returns a new DataStoreReqGetInfoV1
func NewDataStoreReqGetInfoV1() *DataStoreReqGetInfoV1 {
	dsrgiv := &DataStoreReqGetInfoV1{
		URL:            types.NewString(""),
		RequestHeaders: types.NewList[*DataStoreKeyValue](),
		Size:           types.NewPrimitiveU32(0),
		RootCACert:     types.NewBuffer(nil),
	}

	dsrgiv.RequestHeaders.Type = NewDataStoreKeyValue()

	return dsrgiv
}
