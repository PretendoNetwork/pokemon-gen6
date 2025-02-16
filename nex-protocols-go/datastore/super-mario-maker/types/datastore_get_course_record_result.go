// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetCourseRecordResult is a type within the DataStore protocol
type DataStoreGetCourseRecordResult struct {
	types.Structure
	DataID      *types.PrimitiveU64
	Slot        *types.PrimitiveU8
	FirstPID    *types.PID
	BestPID     *types.PID
	BestScore   *types.PrimitiveS32
	CreatedTime *types.DateTime
	UpdatedTime *types.DateTime
}

// WriteTo writes the DataStoreGetCourseRecordResult to the given writable
func (dsgcrr *DataStoreGetCourseRecordResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgcrr.DataID.WriteTo(contentWritable)
	dsgcrr.Slot.WriteTo(contentWritable)
	dsgcrr.FirstPID.WriteTo(contentWritable)
	dsgcrr.BestPID.WriteTo(contentWritable)
	dsgcrr.BestScore.WriteTo(contentWritable)
	dsgcrr.CreatedTime.WriteTo(contentWritable)
	dsgcrr.UpdatedTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgcrr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetCourseRecordResult from the given readable
func (dsgcrr *DataStoreGetCourseRecordResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgcrr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult header. %s", err.Error())
	}

	err = dsgcrr.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.DataID. %s", err.Error())
	}

	err = dsgcrr.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.Slot. %s", err.Error())
	}

	err = dsgcrr.FirstPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.FirstPID. %s", err.Error())
	}

	err = dsgcrr.BestPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.BestPID. %s", err.Error())
	}

	err = dsgcrr.BestScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.BestScore. %s", err.Error())
	}

	err = dsgcrr.CreatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.CreatedTime. %s", err.Error())
	}

	err = dsgcrr.UpdatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordResult.UpdatedTime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetCourseRecordResult
func (dsgcrr *DataStoreGetCourseRecordResult) Copy() types.RVType {
	copied := NewDataStoreGetCourseRecordResult()

	copied.StructureVersion = dsgcrr.StructureVersion
	copied.DataID = dsgcrr.DataID.Copy().(*types.PrimitiveU64)
	copied.Slot = dsgcrr.Slot.Copy().(*types.PrimitiveU8)
	copied.FirstPID = dsgcrr.FirstPID.Copy().(*types.PID)
	copied.BestPID = dsgcrr.BestPID.Copy().(*types.PID)
	copied.BestScore = dsgcrr.BestScore.Copy().(*types.PrimitiveS32)
	copied.CreatedTime = dsgcrr.CreatedTime.Copy().(*types.DateTime)
	copied.UpdatedTime = dsgcrr.UpdatedTime.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given DataStoreGetCourseRecordResult contains the same data as the current DataStoreGetCourseRecordResult
func (dsgcrr *DataStoreGetCourseRecordResult) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetCourseRecordResult); !ok {
		return false
	}

	other := o.(*DataStoreGetCourseRecordResult)

	if dsgcrr.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsgcrr.DataID.Equals(other.DataID) {
		return false
	}

	if !dsgcrr.Slot.Equals(other.Slot) {
		return false
	}

	if !dsgcrr.FirstPID.Equals(other.FirstPID) {
		return false
	}

	if !dsgcrr.BestPID.Equals(other.BestPID) {
		return false
	}

	if !dsgcrr.BestScore.Equals(other.BestScore) {
		return false
	}

	if !dsgcrr.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	return dsgcrr.UpdatedTime.Equals(other.UpdatedTime)
}

// String returns the string representation of the DataStoreGetCourseRecordResult
func (dsgcrr *DataStoreGetCourseRecordResult) String() string {
	return dsgcrr.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetCourseRecordResult using the provided indentation level
func (dsgcrr *DataStoreGetCourseRecordResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetCourseRecordResult{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsgcrr.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, dsgcrr.Slot))
	b.WriteString(fmt.Sprintf("%sFirstPID: %s,\n", indentationValues, dsgcrr.FirstPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sBestPID: %s,\n", indentationValues, dsgcrr.BestPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sBestScore: %s,\n", indentationValues, dsgcrr.BestScore))
	b.WriteString(fmt.Sprintf("%sCreatedTime: %s,\n", indentationValues, dsgcrr.CreatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUpdatedTime: %s,\n", indentationValues, dsgcrr.UpdatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCourseRecordResult returns a new DataStoreGetCourseRecordResult
func NewDataStoreGetCourseRecordResult() *DataStoreGetCourseRecordResult {
	dsgcrr := &DataStoreGetCourseRecordResult{
		DataID:      types.NewPrimitiveU64(0),
		Slot:        types.NewPrimitiveU8(0),
		FirstPID:    types.NewPID(0),
		BestPID:     types.NewPID(0),
		BestScore:   types.NewPrimitiveS32(0),
		CreatedTime: types.NewDateTime(0),
		UpdatedTime: types.NewDateTime(0),
	}

	return dsgcrr
}
