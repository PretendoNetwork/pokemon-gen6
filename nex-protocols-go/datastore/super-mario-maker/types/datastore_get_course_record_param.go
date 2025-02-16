// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetCourseRecordParam is a type within the DataStore protocol
type DataStoreGetCourseRecordParam struct {
	types.Structure
	DataID *types.PrimitiveU64
	Slot   *types.PrimitiveU8
}

// WriteTo writes the DataStoreGetCourseRecordParam to the given writable
func (dsgcrp *DataStoreGetCourseRecordParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgcrp.DataID.WriteTo(contentWritable)
	dsgcrp.Slot.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgcrp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetCourseRecordParam from the given readable
func (dsgcrp *DataStoreGetCourseRecordParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgcrp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam header. %s", err.Error())
	}

	err = dsgcrp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam.DataID. %s", err.Error())
	}

	err = dsgcrp.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam.Slot. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetCourseRecordParam
func (dsgcrp *DataStoreGetCourseRecordParam) Copy() types.RVType {
	copied := NewDataStoreGetCourseRecordParam()

	copied.StructureVersion = dsgcrp.StructureVersion
	copied.DataID = dsgcrp.DataID.Copy().(*types.PrimitiveU64)
	copied.Slot = dsgcrp.Slot.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given DataStoreGetCourseRecordParam contains the same data as the current DataStoreGetCourseRecordParam
func (dsgcrp *DataStoreGetCourseRecordParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetCourseRecordParam); !ok {
		return false
	}

	other := o.(*DataStoreGetCourseRecordParam)

	if dsgcrp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsgcrp.DataID.Equals(other.DataID) {
		return false
	}

	return dsgcrp.Slot.Equals(other.Slot)
}

// String returns the string representation of the DataStoreGetCourseRecordParam
func (dsgcrp *DataStoreGetCourseRecordParam) String() string {
	return dsgcrp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetCourseRecordParam using the provided indentation level
func (dsgcrp *DataStoreGetCourseRecordParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetCourseRecordParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsgcrp.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dsgcrp.Slot))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCourseRecordParam returns a new DataStoreGetCourseRecordParam
func NewDataStoreGetCourseRecordParam() *DataStoreGetCourseRecordParam {
	dsgcrp := &DataStoreGetCourseRecordParam{
		DataID: types.NewPrimitiveU64(0),
		Slot:   types.NewPrimitiveU8(0),
	}

	return dsgcrp
}
