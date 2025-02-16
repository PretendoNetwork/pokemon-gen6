// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreProfileInfo is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreProfileInfo struct {
	types.Structure
	PID     *types.PID
	Profile *types.QBuffer
}

// WriteTo writes the DataStoreProfileInfo to the given writable
func (dspi *DataStoreProfileInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspi.PID.WriteTo(contentWritable)
	dspi.Profile.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreProfileInfo from the given readable
func (dspi *DataStoreProfileInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreProfileInfo header. %s", err.Error())
	}

	err = dspi.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreProfileInfo.PID. %s", err.Error())
	}

	err = dspi.Profile.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreProfileInfo.Profile. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreProfileInfo
func (dspi *DataStoreProfileInfo) Copy() types.RVType {
	copied := NewDataStoreProfileInfo()

	copied.StructureVersion = dspi.StructureVersion
	copied.PID = dspi.PID.Copy().(*types.PID)
	copied.Profile = dspi.Profile.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given DataStoreProfileInfo contains the same data as the current DataStoreProfileInfo
func (dspi *DataStoreProfileInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreProfileInfo); !ok {
		return false
	}

	other := o.(*DataStoreProfileInfo)

	if dspi.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspi.PID.Equals(other.PID) {
		return false
	}

	return dspi.Profile.Equals(other.Profile)
}

// String returns the string representation of the DataStoreProfileInfo
func (dspi *DataStoreProfileInfo) String() string {
	return dspi.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreProfileInfo using the provided indentation level
func (dspi *DataStoreProfileInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreProfileInfo{\n")
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, dspi.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sProfile: %s,\n", indentationValues, dspi.Profile))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreProfileInfo returns a new DataStoreProfileInfo
func NewDataStoreProfileInfo() *DataStoreProfileInfo {
	dspi := &DataStoreProfileInfo{
		PID:     types.NewPID(0),
		Profile: types.NewQBuffer(nil),
	}

	return dspi
}
