// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2ChartInfoInput is a type within the Ranking2 protocol
type Ranking2ChartInfoInput struct {
	types.Structure
	ChartIndex         *types.PrimitiveU32
	NumSeasonsToGoBack *types.PrimitiveU8
}

// WriteTo writes the Ranking2ChartInfoInput to the given writable
func (rcii *Ranking2ChartInfoInput) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rcii.ChartIndex.WriteTo(contentWritable)
	rcii.NumSeasonsToGoBack.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rcii.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2ChartInfoInput from the given readable
func (rcii *Ranking2ChartInfoInput) ExtractFrom(readable types.Readable) error {
	var err error

	err = rcii.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfoInput header. %s", err.Error())
	}

	err = rcii.ChartIndex.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfoInput.ChartIndex. %s", err.Error())
	}

	err = rcii.NumSeasonsToGoBack.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfoInput.NumSeasonsToGoBack. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2ChartInfoInput
func (rcii *Ranking2ChartInfoInput) Copy() types.RVType {
	copied := NewRanking2ChartInfoInput()

	copied.StructureVersion = rcii.StructureVersion
	copied.ChartIndex = rcii.ChartIndex.Copy().(*types.PrimitiveU32)
	copied.NumSeasonsToGoBack = rcii.NumSeasonsToGoBack.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given Ranking2ChartInfoInput contains the same data as the current Ranking2ChartInfoInput
func (rcii *Ranking2ChartInfoInput) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2ChartInfoInput); !ok {
		return false
	}

	other := o.(*Ranking2ChartInfoInput)

	if rcii.StructureVersion != other.StructureVersion {
		return false
	}

	if !rcii.ChartIndex.Equals(other.ChartIndex) {
		return false
	}

	return rcii.NumSeasonsToGoBack.Equals(other.NumSeasonsToGoBack)
}

// String returns the string representation of the Ranking2ChartInfoInput
func (rcii *Ranking2ChartInfoInput) String() string {
	return rcii.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2ChartInfoInput using the provided indentation level
func (rcii *Ranking2ChartInfoInput) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2ChartInfoInput{\n")
	b.WriteString(fmt.Sprintf("%sChartIndex: %s,\n", indentationValues, rcii.ChartIndex))
	b.WriteString(fmt.Sprintf("%sNumSeasonsToGoBack: %s,\n", indentationValues, rcii.NumSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2ChartInfoInput returns a new Ranking2ChartInfoInput
func NewRanking2ChartInfoInput() *Ranking2ChartInfoInput {
	rcii := &Ranking2ChartInfoInput{
		ChartIndex:         types.NewPrimitiveU32(0),
		NumSeasonsToGoBack: types.NewPrimitiveU8(0),
	}

	return rcii
}
