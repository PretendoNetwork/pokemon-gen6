// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// CreateMatchmakeSessionParam is a type within the Matchmaking protocol
type CreateMatchmakeSessionParam struct {
	types.Structure
	SourceMatchmakeSession       *MatchmakeSession
	AdditionalParticipants       *types.List[*types.PID]
	GIDForParticipationCheck     *types.PrimitiveU32
	CreateMatchmakeSessionOption *types.PrimitiveU32
	JoinMessage                  *types.String
	ParticipationCount           *types.PrimitiveU16
}

// WriteTo writes the CreateMatchmakeSessionParam to the given writable
func (cmsp *CreateMatchmakeSessionParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	cmsp.SourceMatchmakeSession.WriteTo(contentWritable)
	cmsp.AdditionalParticipants.WriteTo(contentWritable)
	cmsp.GIDForParticipationCheck.WriteTo(contentWritable)
	cmsp.CreateMatchmakeSessionOption.WriteTo(contentWritable)
	cmsp.JoinMessage.WriteTo(contentWritable)
	cmsp.ParticipationCount.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	cmsp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the CreateMatchmakeSessionParam from the given readable
func (cmsp *CreateMatchmakeSessionParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = cmsp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam header. %s", err.Error())
	}

	err = cmsp.SourceMatchmakeSession.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.SourceMatchmakeSession. %s", err.Error())
	}

	err = cmsp.AdditionalParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.AdditionalParticipants. %s", err.Error())
	}

	err = cmsp.GIDForParticipationCheck.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.GIDForParticipationCheck. %s", err.Error())
	}

	err = cmsp.CreateMatchmakeSessionOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.CreateMatchmakeSessionOption. %s", err.Error())
	}

	err = cmsp.JoinMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.JoinMessage. %s", err.Error())
	}

	err = cmsp.ParticipationCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.ParticipationCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CreateMatchmakeSessionParam
func (cmsp *CreateMatchmakeSessionParam) Copy() types.RVType {
	copied := NewCreateMatchmakeSessionParam()

	copied.StructureVersion = cmsp.StructureVersion
	copied.SourceMatchmakeSession = cmsp.SourceMatchmakeSession.Copy().(*MatchmakeSession)
	copied.AdditionalParticipants = cmsp.AdditionalParticipants.Copy().(*types.List[*types.PID])
	copied.GIDForParticipationCheck = cmsp.GIDForParticipationCheck.Copy().(*types.PrimitiveU32)
	copied.CreateMatchmakeSessionOption = cmsp.CreateMatchmakeSessionOption.Copy().(*types.PrimitiveU32)
	copied.JoinMessage = cmsp.JoinMessage.Copy().(*types.String)
	copied.ParticipationCount = cmsp.ParticipationCount.Copy().(*types.PrimitiveU16)

	return copied
}

// Equals checks if the given CreateMatchmakeSessionParam contains the same data as the current CreateMatchmakeSessionParam
func (cmsp *CreateMatchmakeSessionParam) Equals(o types.RVType) bool {
	if _, ok := o.(*CreateMatchmakeSessionParam); !ok {
		return false
	}

	other := o.(*CreateMatchmakeSessionParam)

	if cmsp.StructureVersion != other.StructureVersion {
		return false
	}

	if !cmsp.SourceMatchmakeSession.Equals(other.SourceMatchmakeSession) {
		return false
	}

	if !cmsp.AdditionalParticipants.Equals(other.AdditionalParticipants) {
		return false
	}

	if !cmsp.GIDForParticipationCheck.Equals(other.GIDForParticipationCheck) {
		return false
	}

	if !cmsp.CreateMatchmakeSessionOption.Equals(other.CreateMatchmakeSessionOption) {
		return false
	}

	if !cmsp.JoinMessage.Equals(other.JoinMessage) {
		return false
	}

	return cmsp.ParticipationCount.Equals(other.ParticipationCount)
}

// String returns the string representation of the CreateMatchmakeSessionParam
func (cmsp *CreateMatchmakeSessionParam) String() string {
	return cmsp.FormatToString(0)
}

// FormatToString pretty-prints the CreateMatchmakeSessionParam using the provided indentation level
func (cmsp *CreateMatchmakeSessionParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CreateMatchmakeSessionParam{\n")
	b.WriteString(fmt.Sprintf("%sSourceMatchmakeSession: %s,\n", indentationValues, cmsp.SourceMatchmakeSession.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sAdditionalParticipants: %s,\n", indentationValues, cmsp.AdditionalParticipants))
	b.WriteString(fmt.Sprintf("%sGIDForParticipationCheck: %s,\n", indentationValues, cmsp.GIDForParticipationCheck))
	b.WriteString(fmt.Sprintf("%sCreateMatchmakeSessionOption: %s,\n", indentationValues, cmsp.CreateMatchmakeSessionOption))
	b.WriteString(fmt.Sprintf("%sJoinMessage: %s,\n", indentationValues, cmsp.JoinMessage))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %s,\n", indentationValues, cmsp.ParticipationCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewCreateMatchmakeSessionParam returns a new CreateMatchmakeSessionParam
func NewCreateMatchmakeSessionParam() *CreateMatchmakeSessionParam {
	cmsp := &CreateMatchmakeSessionParam{
		SourceMatchmakeSession:       NewMatchmakeSession(),
		AdditionalParticipants:       types.NewList[*types.PID](),
		GIDForParticipationCheck:     types.NewPrimitiveU32(0),
		CreateMatchmakeSessionOption: types.NewPrimitiveU32(0),
		JoinMessage:                  types.NewString(""),
		ParticipationCount:           types.NewPrimitiveU16(0),
	}

	cmsp.AdditionalParticipants.Type = types.NewPID(0)

	return cmsp
}
