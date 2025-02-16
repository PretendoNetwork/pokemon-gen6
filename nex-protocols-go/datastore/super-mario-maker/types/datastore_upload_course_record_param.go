// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreUploadCourseRecordParam is a type within the DataStore protocol
type DataStoreUploadCourseRecordParam struct {
	types.Structure
	DataID *types.PrimitiveU64
	Slot   *types.PrimitiveU8
	Score  *types.PrimitiveS32
}

// WriteTo writes the DataStoreUploadCourseRecordParam to the given writable
func (dsucrp *DataStoreUploadCourseRecordParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsucrp.DataID.WriteTo(contentWritable)
	dsucrp.Slot.WriteTo(contentWritable)
	dsucrp.Score.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsucrp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreUploadCourseRecordParam from the given readable
func (dsucrp *DataStoreUploadCourseRecordParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsucrp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam header. %s", err.Error())
	}

	err = dsucrp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.DataID. %s", err.Error())
	}

	err = dsucrp.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.Slot. %s", err.Error())
	}

	err = dsucrp.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.Score. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreUploadCourseRecordParam
func (dsucrp *DataStoreUploadCourseRecordParam) Copy() types.RVType {
	copied := NewDataStoreUploadCourseRecordParam()

	copied.StructureVersion = dsucrp.StructureVersion
	copied.DataID = dsucrp.DataID.Copy().(*types.PrimitiveU64)
	copied.Slot = dsucrp.Slot.Copy().(*types.PrimitiveU8)
	copied.Score = dsucrp.Score.Copy().(*types.PrimitiveS32)

	return copied
}

// Equals checks if the given DataStoreUploadCourseRecordParam contains the same data as the current DataStoreUploadCourseRecordParam
func (dsucrp *DataStoreUploadCourseRecordParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreUploadCourseRecordParam); !ok {
		return false
	}

	other := o.(*DataStoreUploadCourseRecordParam)

	if dsucrp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsucrp.DataID.Equals(other.DataID) {
		return false
	}

	if !dsucrp.Slot.Equals(other.Slot) {
		return false
	}

	return dsucrp.Score.Equals(other.Score)
}

// String returns the string representation of the DataStoreUploadCourseRecordParam
func (dsucrp *DataStoreUploadCourseRecordParam) String() string {
	return dsucrp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreUploadCourseRecordParam using the provided indentation level
func (dsucrp *DataStoreUploadCourseRecordParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreUploadCourseRecordParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsucrp.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dsucrp.Slot))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dsucrp.Score))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreUploadCourseRecordParam returns a new DataStoreUploadCourseRecordParam
func NewDataStoreUploadCourseRecordParam() *DataStoreUploadCourseRecordParam {
	dsucrp := &DataStoreUploadCourseRecordParam{
		DataID: types.NewPrimitiveU64(0),
		Slot:   types.NewPrimitiveU8(0),
		Score:  types.NewPrimitiveS32(0),
	}

	return dsucrp
}
