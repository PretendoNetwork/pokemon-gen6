// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePostProfileParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStorePostProfileParam struct {
	types.Structure
	Profile *types.QBuffer
}

// WriteTo writes the DataStorePostProfileParam to the given writable
func (dsppp *DataStorePostProfileParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsppp.Profile.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsppp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePostProfileParam from the given readable
func (dsppp *DataStorePostProfileParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsppp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostProfileParam header. %s", err.Error())
	}

	err = dsppp.Profile.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostProfileParam.Profile. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePostProfileParam
func (dsppp *DataStorePostProfileParam) Copy() types.RVType {
	copied := NewDataStorePostProfileParam()

	copied.StructureVersion = dsppp.StructureVersion
	copied.Profile = dsppp.Profile.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given DataStorePostProfileParam contains the same data as the current DataStorePostProfileParam
func (dsppp *DataStorePostProfileParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePostProfileParam); !ok {
		return false
	}

	other := o.(*DataStorePostProfileParam)

	if dsppp.StructureVersion != other.StructureVersion {
		return false
	}

	return dsppp.Profile.Equals(other.Profile)
}

// String returns the string representation of the DataStorePostProfileParam
func (dsppp *DataStorePostProfileParam) String() string {
	return dsppp.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePostProfileParam using the provided indentation level
func (dsppp *DataStorePostProfileParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePostProfileParam{\n")
	b.WriteString(fmt.Sprintf("%sProfile: %s,\n", indentationValues, dsppp.Profile))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePostProfileParam returns a new DataStorePostProfileParam
func NewDataStorePostProfileParam() *DataStorePostProfileParam {
	dsppp := &DataStorePostProfileParam{
		Profile: types.NewQBuffer(nil),
	}

	return dsppp
}
