// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreNotificationV1 is a type within the DataStore protocol
type DataStoreNotificationV1 struct {
	types.Structure
	NotificationID *types.PrimitiveU64
	DataID         *types.PrimitiveU32
}

// WriteTo writes the DataStoreNotificationV1 to the given writable
func (dsnv *DataStoreNotificationV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsnv.NotificationID.WriteTo(contentWritable)
	dsnv.DataID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsnv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreNotificationV1 from the given readable
func (dsnv *DataStoreNotificationV1) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsnv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotificationV1 header. %s", err.Error())
	}

	err = dsnv.NotificationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotificationV1.NotificationID. %s", err.Error())
	}

	err = dsnv.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotificationV1.DataID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreNotificationV1
func (dsnv *DataStoreNotificationV1) Copy() types.RVType {
	copied := NewDataStoreNotificationV1()

	copied.StructureVersion = dsnv.StructureVersion
	copied.NotificationID = dsnv.NotificationID.Copy().(*types.PrimitiveU64)
	copied.DataID = dsnv.DataID.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given DataStoreNotificationV1 contains the same data as the current DataStoreNotificationV1
func (dsnv *DataStoreNotificationV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreNotificationV1); !ok {
		return false
	}

	other := o.(*DataStoreNotificationV1)

	if dsnv.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsnv.NotificationID.Equals(other.NotificationID) {
		return false
	}

	return dsnv.DataID.Equals(other.DataID)
}

// String returns the string representation of the DataStoreNotificationV1
func (dsnv *DataStoreNotificationV1) String() string {
	return dsnv.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreNotificationV1 using the provided indentation level
func (dsnv *DataStoreNotificationV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreNotificationV1{\n")
	b.WriteString(fmt.Sprintf("%sNotificationID: %s,\n", indentationValues, dsnv.NotificationID))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsnv.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreNotificationV1 returns a new DataStoreNotificationV1
func NewDataStoreNotificationV1() *DataStoreNotificationV1 {
	dsnv := &DataStoreNotificationV1{
		NotificationID: types.NewPrimitiveU64(0),
		DataID:         types.NewPrimitiveU32(0),
	}

	return dsnv
}
