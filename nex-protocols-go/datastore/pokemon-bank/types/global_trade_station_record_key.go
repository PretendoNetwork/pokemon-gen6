// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationRecordKey is a type within the DataStore protocol
type GlobalTradeStationRecordKey struct {
	types.Structure
	DataID   *types.PrimitiveU64
	Password *types.PrimitiveU64
}

// WriteTo writes the GlobalTradeStationRecordKey to the given writable
func (gtsrk *GlobalTradeStationRecordKey) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsrk.DataID.WriteTo(contentWritable)
	gtsrk.Password.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsrk.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationRecordKey from the given readable
func (gtsrk *GlobalTradeStationRecordKey) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsrk.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationRecordKey header. %s", err.Error())
	}

	err = gtsrk.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationRecordKey.DataID. %s", err.Error())
	}

	err = gtsrk.Password.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationRecordKey.Password. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationRecordKey
func (gtsrk *GlobalTradeStationRecordKey) Copy() types.RVType {
	copied := NewGlobalTradeStationRecordKey()

	copied.StructureVersion = gtsrk.StructureVersion
	copied.DataID = gtsrk.DataID.Copy().(*types.PrimitiveU64)
	copied.Password = gtsrk.Password.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given GlobalTradeStationRecordKey contains the same data as the current GlobalTradeStationRecordKey
func (gtsrk *GlobalTradeStationRecordKey) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationRecordKey); !ok {
		return false
	}

	other := o.(*GlobalTradeStationRecordKey)

	if gtsrk.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsrk.DataID.Equals(other.DataID) {
		return false
	}

	return gtsrk.Password.Equals(other.Password)
}

// String returns the string representation of the GlobalTradeStationRecordKey
func (gtsrk *GlobalTradeStationRecordKey) String() string {
	return gtsrk.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationRecordKey using the provided indentation level
func (gtsrk *GlobalTradeStationRecordKey) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationRecordKey{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, gtsrk.DataID))
	b.WriteString(fmt.Sprintf("%sPassword: %s,\n", indentationValues, gtsrk.Password))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationRecordKey returns a new GlobalTradeStationRecordKey
func NewGlobalTradeStationRecordKey() *GlobalTradeStationRecordKey {
	gtsrk := &GlobalTradeStationRecordKey{
		DataID:   types.NewPrimitiveU64(0),
		Password: types.NewPrimitiveU64(0),
	}

	return gtsrk
}
