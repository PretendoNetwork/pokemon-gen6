// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreReplayPlayer is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreReplayPlayer struct {
	types.Structure
	Fighter     *types.PrimitiveU8
	Health      *types.PrimitiveU8
	WinningRate *types.PrimitiveU16
	Color       *types.PrimitiveU8
	Color2      *types.PrimitiveU8
	PrincipalID *types.PrimitiveU32
	Country     *types.PrimitiveU32
	Region      *types.PrimitiveU8
	Number      *types.PrimitiveU8
}

// WriteTo writes the DataStoreReplayPlayer to the given writable
func (dsrp *DataStoreReplayPlayer) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrp.Fighter.WriteTo(contentWritable)
	dsrp.Health.WriteTo(contentWritable)
	dsrp.WinningRate.WriteTo(contentWritable)
	dsrp.Color.WriteTo(contentWritable)
	dsrp.Color2.WriteTo(contentWritable)
	dsrp.PrincipalID.WriteTo(contentWritable)
	dsrp.Country.WriteTo(contentWritable)
	dsrp.Region.WriteTo(contentWritable)
	dsrp.Number.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReplayPlayer from the given readable
func (dsrp *DataStoreReplayPlayer) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer header. %s", err.Error())
	}

	err = dsrp.Fighter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Fighter. %s", err.Error())
	}

	err = dsrp.Health.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Health. %s", err.Error())
	}

	err = dsrp.WinningRate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.WinningRate. %s", err.Error())
	}

	err = dsrp.Color.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Color. %s", err.Error())
	}

	err = dsrp.Color2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Color2. %s", err.Error())
	}

	err = dsrp.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.PrincipalID. %s", err.Error())
	}

	err = dsrp.Country.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Country. %s", err.Error())
	}

	err = dsrp.Region.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Region. %s", err.Error())
	}

	err = dsrp.Number.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Number. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReplayPlayer
func (dsrp *DataStoreReplayPlayer) Copy() types.RVType {
	copied := NewDataStoreReplayPlayer()

	copied.StructureVersion = dsrp.StructureVersion
	copied.Fighter = dsrp.Fighter.Copy().(*types.PrimitiveU8)
	copied.Health = dsrp.Health.Copy().(*types.PrimitiveU8)
	copied.WinningRate = dsrp.WinningRate.Copy().(*types.PrimitiveU16)
	copied.Color = dsrp.Color.Copy().(*types.PrimitiveU8)
	copied.Color2 = dsrp.Color2.Copy().(*types.PrimitiveU8)
	copied.PrincipalID = dsrp.PrincipalID.Copy().(*types.PrimitiveU32)
	copied.Country = dsrp.Country.Copy().(*types.PrimitiveU32)
	copied.Region = dsrp.Region.Copy().(*types.PrimitiveU8)
	copied.Number = dsrp.Number.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given DataStoreReplayPlayer contains the same data as the current DataStoreReplayPlayer
func (dsrp *DataStoreReplayPlayer) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReplayPlayer); !ok {
		return false
	}

	other := o.(*DataStoreReplayPlayer)

	if dsrp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrp.Fighter.Equals(other.Fighter) {
		return false
	}

	if !dsrp.Health.Equals(other.Health) {
		return false
	}

	if !dsrp.WinningRate.Equals(other.WinningRate) {
		return false
	}

	if !dsrp.Color.Equals(other.Color) {
		return false
	}

	if !dsrp.Color2.Equals(other.Color2) {
		return false
	}

	if !dsrp.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !dsrp.Country.Equals(other.Country) {
		return false
	}

	if !dsrp.Region.Equals(other.Region) {
		return false
	}

	return dsrp.Number.Equals(other.Number)
}

// String returns the string representation of the DataStoreReplayPlayer
func (dsrp *DataStoreReplayPlayer) String() string {
	return dsrp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreReplayPlayer using the provided indentation level
func (dsrp *DataStoreReplayPlayer) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReplayPlayer{\n")
	b.WriteString(fmt.Sprintf("%sFighter: %s,\n", indentationValues, dsrp.Fighter))
	b.WriteString(fmt.Sprintf("%sHealth: %s,\n", indentationValues, dsrp.Health))
	b.WriteString(fmt.Sprintf("%sWinningRate: %s,\n", indentationValues, dsrp.WinningRate))
	b.WriteString(fmt.Sprintf("%sColor: %s,\n", indentationValues, dsrp.Color))
	b.WriteString(fmt.Sprintf("%sColor2: %s,\n", indentationValues, dsrp.Color2))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, dsrp.PrincipalID))
	b.WriteString(fmt.Sprintf("%sCountry: %s,\n", indentationValues, dsrp.Country))
	b.WriteString(fmt.Sprintf("%sRegion: %s,\n", indentationValues, dsrp.Region))
	b.WriteString(fmt.Sprintf("%sNumber: %s,\n", indentationValues, dsrp.Number))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReplayPlayer returns a new DataStoreReplayPlayer
func NewDataStoreReplayPlayer() *DataStoreReplayPlayer {
	dsrp := &DataStoreReplayPlayer{
		Fighter:     types.NewPrimitiveU8(0),
		Health:      types.NewPrimitiveU8(0),
		WinningRate: types.NewPrimitiveU16(0),
		Color:       types.NewPrimitiveU8(0),
		Color2:      types.NewPrimitiveU8(0),
		PrincipalID: types.NewPrimitiveU32(0),
		Country:     types.NewPrimitiveU32(0),
		Region:      types.NewPrimitiveU8(0),
		Number:      types.NewPrimitiveU8(0),
	}

	return dsrp
}
