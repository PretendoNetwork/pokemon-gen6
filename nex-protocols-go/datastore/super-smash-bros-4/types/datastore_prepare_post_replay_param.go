// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePreparePostReplayParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStorePreparePostReplayParam struct {
	types.Structure
	Size          *types.PrimitiveU32
	Mode          *types.PrimitiveU8
	Style         *types.PrimitiveU8
	Rule          *types.PrimitiveU8
	Stage         *types.PrimitiveU8
	ReplayType    *types.PrimitiveU8
	CompetitionID *types.PrimitiveU64
	Score         *types.PrimitiveS32
	Players       *types.List[*DataStoreReplayPlayer]
	Winners       *types.List[*types.PrimitiveU32]
	KeyVersion    *types.PrimitiveU16
	ExtraData     *types.List[*types.String]
}

// WriteTo writes the DataStorePreparePostReplayParam to the given writable
func (dspprp *DataStorePreparePostReplayParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspprp.Size.WriteTo(contentWritable)
	dspprp.Mode.WriteTo(contentWritable)
	dspprp.Style.WriteTo(contentWritable)
	dspprp.Rule.WriteTo(contentWritable)
	dspprp.Stage.WriteTo(contentWritable)
	dspprp.ReplayType.WriteTo(contentWritable)
	dspprp.CompetitionID.WriteTo(contentWritable)
	dspprp.Score.WriteTo(contentWritable)
	dspprp.Players.WriteTo(contentWritable)
	dspprp.Winners.WriteTo(contentWritable)
	dspprp.KeyVersion.WriteTo(contentWritable)
	dspprp.ExtraData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspprp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePreparePostReplayParam from the given readable
func (dspprp *DataStorePreparePostReplayParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspprp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam header. %s", err.Error())
	}

	err = dspprp.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Size. %s", err.Error())
	}

	err = dspprp.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Mode. %s", err.Error())
	}

	err = dspprp.Style.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Style. %s", err.Error())
	}

	err = dspprp.Rule.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Rule. %s", err.Error())
	}

	err = dspprp.Stage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Stage. %s", err.Error())
	}

	err = dspprp.ReplayType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.ReplayType. %s", err.Error())
	}

	err = dspprp.CompetitionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.CompetitionID. %s", err.Error())
	}

	err = dspprp.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Score. %s", err.Error())
	}

	err = dspprp.Players.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Players. %s", err.Error())
	}

	err = dspprp.Winners.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Winners. %s", err.Error())
	}

	err = dspprp.KeyVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.KeyVersion. %s", err.Error())
	}

	err = dspprp.ExtraData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.ExtraData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePreparePostReplayParam
func (dspprp *DataStorePreparePostReplayParam) Copy() types.RVType {
	copied := NewDataStorePreparePostReplayParam()

	copied.StructureVersion = dspprp.StructureVersion
	copied.Size = dspprp.Size.Copy().(*types.PrimitiveU32)
	copied.Mode = dspprp.Mode.Copy().(*types.PrimitiveU8)
	copied.Style = dspprp.Style.Copy().(*types.PrimitiveU8)
	copied.Rule = dspprp.Rule.Copy().(*types.PrimitiveU8)
	copied.Stage = dspprp.Stage.Copy().(*types.PrimitiveU8)
	copied.ReplayType = dspprp.ReplayType.Copy().(*types.PrimitiveU8)
	copied.CompetitionID = dspprp.CompetitionID.Copy().(*types.PrimitiveU64)
	copied.Score = dspprp.Score.Copy().(*types.PrimitiveS32)
	copied.Players = dspprp.Players.Copy().(*types.List[*DataStoreReplayPlayer])
	copied.Winners = dspprp.Winners.Copy().(*types.List[*types.PrimitiveU32])
	copied.KeyVersion = dspprp.KeyVersion.Copy().(*types.PrimitiveU16)
	copied.ExtraData = dspprp.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the given DataStorePreparePostReplayParam contains the same data as the current DataStorePreparePostReplayParam
func (dspprp *DataStorePreparePostReplayParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePreparePostReplayParam); !ok {
		return false
	}

	other := o.(*DataStorePreparePostReplayParam)

	if dspprp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspprp.Size.Equals(other.Size) {
		return false
	}

	if !dspprp.Mode.Equals(other.Mode) {
		return false
	}

	if !dspprp.Style.Equals(other.Style) {
		return false
	}

	if !dspprp.Rule.Equals(other.Rule) {
		return false
	}

	if !dspprp.Stage.Equals(other.Stage) {
		return false
	}

	if !dspprp.ReplayType.Equals(other.ReplayType) {
		return false
	}

	if !dspprp.CompetitionID.Equals(other.CompetitionID) {
		return false
	}

	if !dspprp.Score.Equals(other.Score) {
		return false
	}

	if !dspprp.Players.Equals(other.Players) {
		return false
	}

	if !dspprp.Winners.Equals(other.Winners) {
		return false
	}

	if !dspprp.KeyVersion.Equals(other.KeyVersion) {
		return false
	}

	return dspprp.ExtraData.Equals(other.ExtraData)
}

// String returns the string representation of the DataStorePreparePostReplayParam
func (dspprp *DataStorePreparePostReplayParam) String() string {
	return dspprp.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePreparePostReplayParam using the provided indentation level
func (dspprp *DataStorePreparePostReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dspprp.Size))
	b.WriteString(fmt.Sprintf("%sMode: %s,\n", indentationValues, dspprp.Mode))
	b.WriteString(fmt.Sprintf("%sStyle: %s,\n", indentationValues, dspprp.Style))
	b.WriteString(fmt.Sprintf("%sRule: %s,\n", indentationValues, dspprp.Rule))
	b.WriteString(fmt.Sprintf("%sStage: %s,\n", indentationValues, dspprp.Stage))
	b.WriteString(fmt.Sprintf("%sReplayType: %s,\n", indentationValues, dspprp.ReplayType))
	b.WriteString(fmt.Sprintf("%sCompetitionID: %s,\n", indentationValues, dspprp.CompetitionID))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dspprp.Score))
	b.WriteString(fmt.Sprintf("%sPlayers: %s,\n", indentationValues, dspprp.Players))
	b.WriteString(fmt.Sprintf("%sWinners: %s,\n", indentationValues, dspprp.Winners))
	b.WriteString(fmt.Sprintf("%sKeyVersion: %s,\n", indentationValues, dspprp.KeyVersion))
	b.WriteString(fmt.Sprintf("%sExtraData: %s,\n", indentationValues, dspprp.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostReplayParam returns a new DataStorePreparePostReplayParam
func NewDataStorePreparePostReplayParam() *DataStorePreparePostReplayParam {
	dspprp := &DataStorePreparePostReplayParam{
		Size:          types.NewPrimitiveU32(0),
		Mode:          types.NewPrimitiveU8(0),
		Style:         types.NewPrimitiveU8(0),
		Rule:          types.NewPrimitiveU8(0),
		Stage:         types.NewPrimitiveU8(0),
		ReplayType:    types.NewPrimitiveU8(0),
		CompetitionID: types.NewPrimitiveU64(0),
		Score:         types.NewPrimitiveS32(0),
		Players:       types.NewList[*DataStoreReplayPlayer](),
		Winners:       types.NewList[*types.PrimitiveU32](),
		KeyVersion:    types.NewPrimitiveU16(0),
		ExtraData:     types.NewList[*types.String](),
	}

	dspprp.Players.Type = NewDataStoreReplayPlayer()
	dspprp.Winners.Type = types.NewPrimitiveU32(0)
	dspprp.ExtraData.Type = types.NewString("")

	return dspprp
}
