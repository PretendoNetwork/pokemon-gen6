// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationData is a type within the DataStore protocol
type GlobalTradeStationData struct {
	types.Structure
	DataID      *types.PrimitiveU64
	OwnerID     *types.PID
	UpdatedTime *types.DateTime
	IndexData   *types.QBuffer
	Version     *types.PrimitiveU32
}

// WriteTo writes the GlobalTradeStationData to the given writable
func (gtsd *GlobalTradeStationData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsd.DataID.WriteTo(contentWritable)
	gtsd.OwnerID.WriteTo(contentWritable)
	gtsd.UpdatedTime.WriteTo(contentWritable)
	gtsd.IndexData.WriteTo(contentWritable)
	gtsd.Version.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationData from the given readable
func (gtsd *GlobalTradeStationData) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData header. %s", err.Error())
	}

	err = gtsd.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.DataID. %s", err.Error())
	}

	err = gtsd.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.OwnerID. %s", err.Error())
	}

	err = gtsd.UpdatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.UpdatedTime. %s", err.Error())
	}

	err = gtsd.IndexData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.IndexData. %s", err.Error())
	}

	err = gtsd.Version.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.Version. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationData
func (gtsd *GlobalTradeStationData) Copy() types.RVType {
	copied := NewGlobalTradeStationData()

	copied.StructureVersion = gtsd.StructureVersion
	copied.DataID = gtsd.DataID.Copy().(*types.PrimitiveU64)
	copied.OwnerID = gtsd.OwnerID.Copy().(*types.PID)
	copied.UpdatedTime = gtsd.UpdatedTime.Copy().(*types.DateTime)
	copied.IndexData = gtsd.IndexData.Copy().(*types.QBuffer)
	copied.Version = gtsd.Version.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given GlobalTradeStationData contains the same data as the current GlobalTradeStationData
func (gtsd *GlobalTradeStationData) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationData); !ok {
		return false
	}

	other := o.(*GlobalTradeStationData)

	if gtsd.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsd.DataID.Equals(other.DataID) {
		return false
	}

	if !gtsd.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !gtsd.UpdatedTime.Equals(other.UpdatedTime) {
		return false
	}

	if !gtsd.IndexData.Equals(other.IndexData) {
		return false
	}

	return gtsd.Version.Equals(other.Version)
}

// String returns the string representation of the GlobalTradeStationData
func (gtsd *GlobalTradeStationData) String() string {
	return gtsd.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationData using the provided indentation level
func (gtsd *GlobalTradeStationData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationData{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, gtsd.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, gtsd.OwnerID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUpdatedTime: %s,\n", indentationValues, gtsd.UpdatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sIndexData: %s,\n", indentationValues, gtsd.IndexData))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, gtsd.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationData returns a new GlobalTradeStationData
func NewGlobalTradeStationData() *GlobalTradeStationData {
	gtsd := &GlobalTradeStationData{
		DataID:      types.NewPrimitiveU64(0),
		OwnerID:     types.NewPID(0),
		UpdatedTime: types.NewDateTime(0),
		IndexData:   types.NewQBuffer(nil),
		Version:     types.NewPrimitiveU32(0),
	}

	return gtsd
}
