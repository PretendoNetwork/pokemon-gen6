// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreFightingPowerScore is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreFightingPowerScore struct {
	types.Structure
	Score *types.PrimitiveU32
	Rank  *types.PrimitiveU32
}

// WriteTo writes the DataStoreFightingPowerScore to the given writable
func (dsfps *DataStoreFightingPowerScore) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsfps.Score.WriteTo(contentWritable)
	dsfps.Rank.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsfps.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreFightingPowerScore from the given readable
func (dsfps *DataStoreFightingPowerScore) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsfps.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerScore header. %s", err.Error())
	}

	err = dsfps.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerScore.Score. %s", err.Error())
	}

	err = dsfps.Rank.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerScore.Rank. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreFightingPowerScore
func (dsfps *DataStoreFightingPowerScore) Copy() types.RVType {
	copied := NewDataStoreFightingPowerScore()

	copied.StructureVersion = dsfps.StructureVersion
	copied.Score = dsfps.Score.Copy().(*types.PrimitiveU32)
	copied.Rank = dsfps.Rank.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given DataStoreFightingPowerScore contains the same data as the current DataStoreFightingPowerScore
func (dsfps *DataStoreFightingPowerScore) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreFightingPowerScore); !ok {
		return false
	}

	other := o.(*DataStoreFightingPowerScore)

	if dsfps.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsfps.Score.Equals(other.Score) {
		return false
	}

	return dsfps.Rank.Equals(other.Rank)
}

// String returns the string representation of the DataStoreFightingPowerScore
func (dsfps *DataStoreFightingPowerScore) String() string {
	return dsfps.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreFightingPowerScore using the provided indentation level
func (dsfps *DataStoreFightingPowerScore) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreFightingPowerScore{\n")
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dsfps.Score))
	b.WriteString(fmt.Sprintf("%sRank: %s,\n", indentationValues, dsfps.Rank))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreFightingPowerScore returns a new DataStoreFightingPowerScore
func NewDataStoreFightingPowerScore() *DataStoreFightingPowerScore {
	dsfps := &DataStoreFightingPowerScore{
		Score: types.NewPrimitiveU32(0),
		Rank:  types.NewPrimitiveU32(0),
	}

	return dsfps
}
