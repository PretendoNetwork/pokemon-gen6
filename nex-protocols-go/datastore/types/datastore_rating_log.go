// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreRatingLog is a type within the DataStore protocol
type DataStoreRatingLog struct {
	types.Structure
	IsRated            *types.PrimitiveBool
	PID                *types.PID
	RatingValue        *types.PrimitiveS32
	LockExpirationTime *types.DateTime
}

// WriteTo writes the DataStoreRatingLog to the given writable
func (dsrl *DataStoreRatingLog) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrl.IsRated.WriteTo(contentWritable)
	dsrl.PID.WriteTo(contentWritable)
	dsrl.RatingValue.WriteTo(contentWritable)
	dsrl.LockExpirationTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrl.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRatingLog from the given readable
func (dsrl *DataStoreRatingLog) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrl.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog header. %s", err.Error())
	}

	err = dsrl.IsRated.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.IsRated. %s", err.Error())
	}

	err = dsrl.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.PID. %s", err.Error())
	}

	err = dsrl.RatingValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.RatingValue. %s", err.Error())
	}

	err = dsrl.LockExpirationTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingLog.LockExpirationTime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingLog
func (dsrl *DataStoreRatingLog) Copy() types.RVType {
	copied := NewDataStoreRatingLog()

	copied.StructureVersion = dsrl.StructureVersion
	copied.IsRated = dsrl.IsRated.Copy().(*types.PrimitiveBool)
	copied.PID = dsrl.PID.Copy().(*types.PID)
	copied.RatingValue = dsrl.RatingValue.Copy().(*types.PrimitiveS32)
	copied.LockExpirationTime = dsrl.LockExpirationTime.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given DataStoreRatingLog contains the same data as the current DataStoreRatingLog
func (dsrl *DataStoreRatingLog) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingLog); !ok {
		return false
	}

	other := o.(*DataStoreRatingLog)

	if dsrl.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrl.IsRated.Equals(other.IsRated) {
		return false
	}

	if !dsrl.PID.Equals(other.PID) {
		return false
	}

	if !dsrl.RatingValue.Equals(other.RatingValue) {
		return false
	}

	return dsrl.LockExpirationTime.Equals(other.LockExpirationTime)
}

// String returns the string representation of the DataStoreRatingLog
func (dsrl *DataStoreRatingLog) String() string {
	return dsrl.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreRatingLog using the provided indentation level
func (dsrl *DataStoreRatingLog) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingLog{\n")
	b.WriteString(fmt.Sprintf("%sIsRated: %s,\n", indentationValues, dsrl.IsRated))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, dsrl.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRatingValue: %s,\n", indentationValues, dsrl.RatingValue))
	b.WriteString(fmt.Sprintf("%sLockExpirationTime: %s,\n", indentationValues, dsrl.LockExpirationTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingLog returns a new DataStoreRatingLog
func NewDataStoreRatingLog() *DataStoreRatingLog {
	dsrl := &DataStoreRatingLog{
		IsRated:            types.NewPrimitiveBool(false),
		PID:                types.NewPID(0),
		RatingValue:        types.NewPrimitiveS32(0),
		LockExpirationTime: types.NewDateTime(0),
	}

	return dsrl
}
