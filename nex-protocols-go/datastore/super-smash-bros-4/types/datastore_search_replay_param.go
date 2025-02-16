// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreSearchReplayParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreSearchReplayParam struct {
	types.Structure
	Mode        *types.PrimitiveU8
	Style       *types.PrimitiveU8
	Fighter     *types.PrimitiveU8
	ResultRange *types.ResultRange
}

// WriteTo writes the DataStoreSearchReplayParam to the given writable
func (dssrp *DataStoreSearchReplayParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dssrp.Mode.WriteTo(contentWritable)
	dssrp.Style.WriteTo(contentWritable)
	dssrp.Fighter.WriteTo(contentWritable)
	dssrp.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dssrp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreSearchReplayParam from the given readable
func (dssrp *DataStoreSearchReplayParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dssrp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam header. %s", err.Error())
	}

	err = dssrp.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Mode. %s", err.Error())
	}

	err = dssrp.Style.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Style. %s", err.Error())
	}

	err = dssrp.Fighter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Fighter. %s", err.Error())
	}

	err = dssrp.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.ResultRange. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSearchReplayParam
func (dssrp *DataStoreSearchReplayParam) Copy() types.RVType {
	copied := NewDataStoreSearchReplayParam()

	copied.StructureVersion = dssrp.StructureVersion
	copied.Mode = dssrp.Mode.Copy().(*types.PrimitiveU8)
	copied.Style = dssrp.Style.Copy().(*types.PrimitiveU8)
	copied.Fighter = dssrp.Fighter.Copy().(*types.PrimitiveU8)
	copied.ResultRange = dssrp.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the given DataStoreSearchReplayParam contains the same data as the current DataStoreSearchReplayParam
func (dssrp *DataStoreSearchReplayParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSearchReplayParam); !ok {
		return false
	}

	other := o.(*DataStoreSearchReplayParam)

	if dssrp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dssrp.Mode.Equals(other.Mode) {
		return false
	}

	if !dssrp.Style.Equals(other.Style) {
		return false
	}

	if !dssrp.Fighter.Equals(other.Fighter) {
		return false
	}

	return dssrp.ResultRange.Equals(other.ResultRange)
}

// String returns the string representation of the DataStoreSearchReplayParam
func (dssrp *DataStoreSearchReplayParam) String() string {
	return dssrp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreSearchReplayParam using the provided indentation level
func (dssrp *DataStoreSearchReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sMode: %s,\n", indentationValues, dssrp.Mode))
	b.WriteString(fmt.Sprintf("%sStyle: %s,\n", indentationValues, dssrp.Style))
	b.WriteString(fmt.Sprintf("%sFighter: %s,\n", indentationValues, dssrp.Fighter))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dssrp.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchReplayParam returns a new DataStoreSearchReplayParam
func NewDataStoreSearchReplayParam() *DataStoreSearchReplayParam {
	dssrp := &DataStoreSearchReplayParam{
		Mode:        types.NewPrimitiveU8(0),
		Style:       types.NewPrimitiveU8(0),
		Fighter:     types.NewPrimitiveU8(0),
		ResultRange: types.NewResultRange(),
	}

	return dssrp
}
