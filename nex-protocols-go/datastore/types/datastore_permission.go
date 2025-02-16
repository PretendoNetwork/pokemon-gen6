// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePermission is a type within the DataStore protocol
type DataStorePermission struct {
	types.Structure
	Permission   *types.PrimitiveU8
	RecipientIDs *types.List[*types.PID]
}

// WriteTo writes the DataStorePermission to the given writable
func (dsp *DataStorePermission) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsp.Permission.WriteTo(contentWritable)
	dsp.RecipientIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePermission from the given readable
func (dsp *DataStorePermission) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePermission header. %s", err.Error())
	}

	err = dsp.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePermission.Permission. %s", err.Error())
	}

	err = dsp.RecipientIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePermission.RecipientIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePermission
func (dsp *DataStorePermission) Copy() types.RVType {
	copied := NewDataStorePermission()

	copied.StructureVersion = dsp.StructureVersion
	copied.Permission = dsp.Permission.Copy().(*types.PrimitiveU8)
	copied.RecipientIDs = dsp.RecipientIDs.Copy().(*types.List[*types.PID])

	return copied
}

// Equals checks if the given DataStorePermission contains the same data as the current DataStorePermission
func (dsp *DataStorePermission) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePermission); !ok {
		return false
	}

	other := o.(*DataStorePermission)

	if dsp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsp.Permission.Equals(other.Permission) {
		return false
	}

	return dsp.RecipientIDs.Equals(other.RecipientIDs)
}

// String returns the string representation of the DataStorePermission
func (dsp *DataStorePermission) String() string {
	return dsp.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePermission using the provided indentation level
func (dsp *DataStorePermission) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePermission{\n")
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dsp.Permission))
	b.WriteString(fmt.Sprintf("%sRecipientIDs: %s,\n", indentationValues, dsp.RecipientIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePermission returns a new DataStorePermission
func NewDataStorePermission() *DataStorePermission {
	dsp := &DataStorePermission{
		Permission:   types.NewPrimitiveU8(0),
		RecipientIDs: types.NewList[*types.PID](),
	}

	dsp.RecipientIDs.Type = types.NewPID(0)

	return dsp
}
