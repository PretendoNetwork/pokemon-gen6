// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePasswordInfo is a type within the DataStore protocol
type DataStorePasswordInfo struct {
	types.Structure
	DataID         *types.PrimitiveU64
	AccessPassword *types.PrimitiveU64
	UpdatePassword *types.PrimitiveU64
}

// WriteTo writes the DataStorePasswordInfo to the given writable
func (dspi *DataStorePasswordInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspi.DataID.WriteTo(contentWritable)
	dspi.AccessPassword.WriteTo(contentWritable)
	dspi.UpdatePassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePasswordInfo from the given readable
func (dspi *DataStorePasswordInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePasswordInfo header. %s", err.Error())
	}

	err = dspi.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePasswordInfo.DataID. %s", err.Error())
	}

	err = dspi.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePasswordInfo.AccessPassword. %s", err.Error())
	}

	err = dspi.UpdatePassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePasswordInfo.UpdatePassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePasswordInfo
func (dspi *DataStorePasswordInfo) Copy() types.RVType {
	copied := NewDataStorePasswordInfo()

	copied.StructureVersion = dspi.StructureVersion
	copied.DataID = dspi.DataID.Copy().(*types.PrimitiveU64)
	copied.AccessPassword = dspi.AccessPassword.Copy().(*types.PrimitiveU64)
	copied.UpdatePassword = dspi.UpdatePassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given DataStorePasswordInfo contains the same data as the current DataStorePasswordInfo
func (dspi *DataStorePasswordInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePasswordInfo); !ok {
		return false
	}

	other := o.(*DataStorePasswordInfo)

	if dspi.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspi.DataID.Equals(other.DataID) {
		return false
	}

	if !dspi.AccessPassword.Equals(other.AccessPassword) {
		return false
	}

	return dspi.UpdatePassword.Equals(other.UpdatePassword)
}

// String returns the string representation of the DataStorePasswordInfo
func (dspi *DataStorePasswordInfo) String() string {
	return dspi.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePasswordInfo using the provided indentation level
func (dspi *DataStorePasswordInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePasswordInfo{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dspi.DataID))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s,\n", indentationValues, dspi.AccessPassword))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %s,\n", indentationValues, dspi.UpdatePassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePasswordInfo returns a new DataStorePasswordInfo
func NewDataStorePasswordInfo() *DataStorePasswordInfo {
	dspi := &DataStorePasswordInfo{
		DataID:         types.NewPrimitiveU64(0),
		AccessPassword: types.NewPrimitiveU64(0),
		UpdatePassword: types.NewPrimitiveU64(0),
	}

	return dspi
}
