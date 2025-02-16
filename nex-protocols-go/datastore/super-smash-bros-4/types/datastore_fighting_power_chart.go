// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreFightingPowerChart is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreFightingPowerChart struct {
	types.Structure
	UserNum *types.PrimitiveU32
	Chart   *types.List[*DataStoreFightingPowerScore]
}

// WriteTo writes the DataStoreFightingPowerChart to the given writable
func (dsfpc *DataStoreFightingPowerChart) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsfpc.UserNum.WriteTo(contentWritable)
	dsfpc.Chart.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsfpc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreFightingPowerChart from the given readable
func (dsfpc *DataStoreFightingPowerChart) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsfpc.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerChart header. %s", err.Error())
	}

	err = dsfpc.UserNum.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerChart.UserNum. %s", err.Error())
	}

	err = dsfpc.Chart.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerChart.Chart. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreFightingPowerChart
func (dsfpc *DataStoreFightingPowerChart) Copy() types.RVType {
	copied := NewDataStoreFightingPowerChart()

	copied.StructureVersion = dsfpc.StructureVersion
	copied.UserNum = dsfpc.UserNum.Copy().(*types.PrimitiveU32)
	copied.Chart = dsfpc.Chart.Copy().(*types.List[*DataStoreFightingPowerScore])

	return copied
}

// Equals checks if the given DataStoreFightingPowerChart contains the same data as the current DataStoreFightingPowerChart
func (dsfpc *DataStoreFightingPowerChart) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreFightingPowerChart); !ok {
		return false
	}

	other := o.(*DataStoreFightingPowerChart)

	if dsfpc.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsfpc.UserNum.Equals(other.UserNum) {
		return false
	}

	return dsfpc.Chart.Equals(other.Chart)
}

// String returns the string representation of the DataStoreFightingPowerChart
func (dsfpc *DataStoreFightingPowerChart) String() string {
	return dsfpc.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreFightingPowerChart using the provided indentation level
func (dsfpc *DataStoreFightingPowerChart) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreFightingPowerChart{\n")
	b.WriteString(fmt.Sprintf("%sUserNum: %s,\n", indentationValues, dsfpc.UserNum))
	b.WriteString(fmt.Sprintf("%sChart: %s,\n", indentationValues, dsfpc.Chart))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreFightingPowerChart returns a new DataStoreFightingPowerChart
func NewDataStoreFightingPowerChart() *DataStoreFightingPowerChart {
	dsfpc := &DataStoreFightingPowerChart{
		UserNum: types.NewPrimitiveU32(0),
		Chart:   types.NewList[*DataStoreFightingPowerScore](),
	}

	dsfpc.Chart.Type = NewDataStoreFightingPowerScore()

	return dsfpc
}
