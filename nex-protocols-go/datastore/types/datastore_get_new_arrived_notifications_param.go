// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetNewArrivedNotificationsParam is a type within the DataStore protocol
type DataStoreGetNewArrivedNotificationsParam struct {
	types.Structure
	LastNotificationID *types.PrimitiveU64
	Limit              *types.PrimitiveU16
}

// WriteTo writes the DataStoreGetNewArrivedNotificationsParam to the given writable
func (dsgnanp *DataStoreGetNewArrivedNotificationsParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgnanp.LastNotificationID.WriteTo(contentWritable)
	dsgnanp.Limit.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgnanp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetNewArrivedNotificationsParam from the given readable
func (dsgnanp *DataStoreGetNewArrivedNotificationsParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgnanp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNewArrivedNotificationsParam header. %s", err.Error())
	}

	err = dsgnanp.LastNotificationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNewArrivedNotificationsParam.LastNotificationID. %s", err.Error())
	}

	err = dsgnanp.Limit.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetNewArrivedNotificationsParam.Limit. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetNewArrivedNotificationsParam
func (dsgnanp *DataStoreGetNewArrivedNotificationsParam) Copy() types.RVType {
	copied := NewDataStoreGetNewArrivedNotificationsParam()

	copied.StructureVersion = dsgnanp.StructureVersion
	copied.LastNotificationID = dsgnanp.LastNotificationID.Copy().(*types.PrimitiveU64)
	copied.Limit = dsgnanp.Limit.Copy().(*types.PrimitiveU16)

	return copied
}

// Equals checks if the given DataStoreGetNewArrivedNotificationsParam contains the same data as the current DataStoreGetNewArrivedNotificationsParam
func (dsgnanp *DataStoreGetNewArrivedNotificationsParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetNewArrivedNotificationsParam); !ok {
		return false
	}

	other := o.(*DataStoreGetNewArrivedNotificationsParam)

	if dsgnanp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsgnanp.LastNotificationID.Equals(other.LastNotificationID) {
		return false
	}

	return dsgnanp.Limit.Equals(other.Limit)
}

// String returns the string representation of the DataStoreGetNewArrivedNotificationsParam
func (dsgnanp *DataStoreGetNewArrivedNotificationsParam) String() string {
	return dsgnanp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetNewArrivedNotificationsParam using the provided indentation level
func (dsgnanp *DataStoreGetNewArrivedNotificationsParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetNewArrivedNotificationsParam{\n")
	b.WriteString(fmt.Sprintf("%sLastNotificationID: %s,\n", indentationValues, dsgnanp.LastNotificationID))
	b.WriteString(fmt.Sprintf("%sLimit: %s,\n", indentationValues, dsgnanp.Limit))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetNewArrivedNotificationsParam returns a new DataStoreGetNewArrivedNotificationsParam
func NewDataStoreGetNewArrivedNotificationsParam() *DataStoreGetNewArrivedNotificationsParam {
	dsgnanp := &DataStoreGetNewArrivedNotificationsParam{
		LastNotificationID: types.NewPrimitiveU64(0),
		Limit:              types.NewPrimitiveU16(0),
	}

	return dsgnanp
}
