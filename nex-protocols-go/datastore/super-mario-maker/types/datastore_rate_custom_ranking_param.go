// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreRateCustomRankingParam is a type within the DataStore protocol
type DataStoreRateCustomRankingParam struct {
	types.Structure
	DataID        *types.PrimitiveU64
	ApplicationID *types.PrimitiveU32
	Score         *types.PrimitiveU32
	Period        *types.PrimitiveU16
}

// WriteTo writes the DataStoreRateCustomRankingParam to the given writable
func (dsrcrp *DataStoreRateCustomRankingParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrcrp.DataID.WriteTo(contentWritable)
	dsrcrp.ApplicationID.WriteTo(contentWritable)
	dsrcrp.Score.WriteTo(contentWritable)
	dsrcrp.Period.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrcrp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRateCustomRankingParam from the given readable
func (dsrcrp *DataStoreRateCustomRankingParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrcrp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam header. %s", err.Error())
	}

	err = dsrcrp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.DataID. %s", err.Error())
	}

	err = dsrcrp.ApplicationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.ApplicationID. %s", err.Error())
	}

	err = dsrcrp.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Score. %s", err.Error())
	}

	err = dsrcrp.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Period. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRateCustomRankingParam
func (dsrcrp *DataStoreRateCustomRankingParam) Copy() types.RVType {
	copied := NewDataStoreRateCustomRankingParam()

	copied.StructureVersion = dsrcrp.StructureVersion
	copied.DataID = dsrcrp.DataID.Copy().(*types.PrimitiveU64)
	copied.ApplicationID = dsrcrp.ApplicationID.Copy().(*types.PrimitiveU32)
	copied.Score = dsrcrp.Score.Copy().(*types.PrimitiveU32)
	copied.Period = dsrcrp.Period.Copy().(*types.PrimitiveU16)

	return copied
}

// Equals checks if the given DataStoreRateCustomRankingParam contains the same data as the current DataStoreRateCustomRankingParam
func (dsrcrp *DataStoreRateCustomRankingParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRateCustomRankingParam); !ok {
		return false
	}

	other := o.(*DataStoreRateCustomRankingParam)

	if dsrcrp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrcrp.DataID.Equals(other.DataID) {
		return false
	}

	if !dsrcrp.ApplicationID.Equals(other.ApplicationID) {
		return false
	}

	if !dsrcrp.Score.Equals(other.Score) {
		return false
	}

	return dsrcrp.Period.Equals(other.Period)
}

// String returns the string representation of the DataStoreRateCustomRankingParam
func (dsrcrp *DataStoreRateCustomRankingParam) String() string {
	return dsrcrp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreRateCustomRankingParam using the provided indentation level
func (dsrcrp *DataStoreRateCustomRankingParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRateCustomRankingParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsrcrp.DataID))
	b.WriteString(fmt.Sprintf("%sApplicationID: %s,\n", indentationValues, dsrcrp.ApplicationID))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dsrcrp.Score))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dsrcrp.Period))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRateCustomRankingParam returns a new DataStoreRateCustomRankingParam
func NewDataStoreRateCustomRankingParam() *DataStoreRateCustomRankingParam {
	dsrcrp := &DataStoreRateCustomRankingParam{
		DataID:        types.NewPrimitiveU64(0),
		ApplicationID: types.NewPrimitiveU32(0),
		Score:         types.NewPrimitiveU32(0),
		Period:        types.NewPrimitiveU16(0),
	}

	return dsrcrp
}
