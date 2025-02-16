// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// JoinMatchmakeSessionParam is a type within the Matchmaking protocol
type JoinMatchmakeSessionParam struct {
	types.Structure
	GID                          *types.PrimitiveU32
	AdditionalParticipants       *types.List[*types.PID]
	GIDForParticipationCheck     *types.PrimitiveU32
	JoinMatchmakeSessionOption   *types.PrimitiveU32
	JoinMatchmakeSessionBehavior *types.PrimitiveU8
	StrUserPassword              *types.String
	StrSystemPassword            *types.String
	JoinMessage                  *types.String
	ParticipationCount           *types.PrimitiveU16
	ExtraParticipants            *types.PrimitiveU16
	BlockListParam               *MatchmakeBlockListParam // * NEX v4.0
}

// WriteTo writes the JoinMatchmakeSessionParam to the given writable
func (jmsp *JoinMatchmakeSessionParam) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.MatchMaking

	contentWritable := writable.CopyNew()

	jmsp.GID.WriteTo(contentWritable)
	jmsp.AdditionalParticipants.WriteTo(contentWritable)
	jmsp.GIDForParticipationCheck.WriteTo(contentWritable)
	jmsp.JoinMatchmakeSessionOption.WriteTo(contentWritable)
	jmsp.JoinMatchmakeSessionBehavior.WriteTo(contentWritable)
	jmsp.StrUserPassword.WriteTo(contentWritable)
	jmsp.StrSystemPassword.WriteTo(contentWritable)
	jmsp.JoinMessage.WriteTo(contentWritable)
	jmsp.ParticipationCount.WriteTo(contentWritable)
	jmsp.ExtraParticipants.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("4.0") {
		jmsp.BlockListParam.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	jmsp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the JoinMatchmakeSessionParam from the given readable
func (jmsp *JoinMatchmakeSessionParam) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.MatchMaking

	var err error

	err = jmsp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam header. %s", err.Error())
	}

	err = jmsp.GID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.GID. %s", err.Error())
	}

	err = jmsp.AdditionalParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.AdditionalParticipants. %s", err.Error())
	}

	err = jmsp.GIDForParticipationCheck.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.GIDForParticipationCheck. %s", err.Error())
	}

	err = jmsp.JoinMatchmakeSessionOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.JoinMatchmakeSessionOption. %s", err.Error())
	}

	err = jmsp.JoinMatchmakeSessionBehavior.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.JoinMatchmakeSessionBehavior. %s", err.Error())
	}

	err = jmsp.StrUserPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.StrUserPassword. %s", err.Error())
	}

	err = jmsp.StrSystemPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.StrSystemPassword. %s", err.Error())
	}

	err = jmsp.JoinMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.JoinMessage. %s", err.Error())
	}

	err = jmsp.ParticipationCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.ParticipationCount. %s", err.Error())
	}

	err = jmsp.ExtraParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.ExtraParticipants. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("4.0") {
		err = jmsp.BlockListParam.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.BlockListParam. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of JoinMatchmakeSessionParam
func (jmsp *JoinMatchmakeSessionParam) Copy() types.RVType {
	copied := NewJoinMatchmakeSessionParam()

	copied.StructureVersion = jmsp.StructureVersion
	copied.GID = jmsp.GID.Copy().(*types.PrimitiveU32)
	copied.AdditionalParticipants = jmsp.AdditionalParticipants.Copy().(*types.List[*types.PID])
	copied.GIDForParticipationCheck = jmsp.GIDForParticipationCheck.Copy().(*types.PrimitiveU32)
	copied.JoinMatchmakeSessionOption = jmsp.JoinMatchmakeSessionOption.Copy().(*types.PrimitiveU32)
	copied.JoinMatchmakeSessionBehavior = jmsp.JoinMatchmakeSessionBehavior.Copy().(*types.PrimitiveU8)
	copied.StrUserPassword = jmsp.StrUserPassword.Copy().(*types.String)
	copied.StrSystemPassword = jmsp.StrSystemPassword.Copy().(*types.String)
	copied.JoinMessage = jmsp.JoinMessage.Copy().(*types.String)
	copied.ParticipationCount = jmsp.ParticipationCount.Copy().(*types.PrimitiveU16)
	copied.ExtraParticipants = jmsp.ExtraParticipants.Copy().(*types.PrimitiveU16)
	copied.BlockListParam = jmsp.BlockListParam.Copy().(*MatchmakeBlockListParam)

	return copied
}

// Equals checks if the given JoinMatchmakeSessionParam contains the same data as the current JoinMatchmakeSessionParam
func (jmsp *JoinMatchmakeSessionParam) Equals(o types.RVType) bool {
	if _, ok := o.(*JoinMatchmakeSessionParam); !ok {
		return false
	}

	other := o.(*JoinMatchmakeSessionParam)

	if jmsp.StructureVersion != other.StructureVersion {
		return false
	}

	if !jmsp.GID.Equals(other.GID) {
		return false
	}

	if !jmsp.AdditionalParticipants.Equals(other.AdditionalParticipants) {
		return false
	}

	if !jmsp.GIDForParticipationCheck.Equals(other.GIDForParticipationCheck) {
		return false
	}

	if !jmsp.JoinMatchmakeSessionOption.Equals(other.JoinMatchmakeSessionOption) {
		return false
	}

	if !jmsp.JoinMatchmakeSessionBehavior.Equals(other.JoinMatchmakeSessionBehavior) {
		return false
	}

	if !jmsp.StrUserPassword.Equals(other.StrUserPassword) {
		return false
	}

	if !jmsp.StrSystemPassword.Equals(other.StrSystemPassword) {
		return false
	}

	if !jmsp.JoinMessage.Equals(other.JoinMessage) {
		return false
	}

	if !jmsp.ParticipationCount.Equals(other.ParticipationCount) {
		return false
	}

	if !jmsp.ExtraParticipants.Equals(other.ExtraParticipants) {
		return false
	}

	return jmsp.BlockListParam.Equals(other.BlockListParam)
}

// String returns the string representation of the JoinMatchmakeSessionParam
func (jmsp *JoinMatchmakeSessionParam) String() string {
	return jmsp.FormatToString(0)
}

// FormatToString pretty-prints the JoinMatchmakeSessionParam using the provided indentation level
func (jmsp *JoinMatchmakeSessionParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("JoinMatchmakeSessionParam{\n")
	b.WriteString(fmt.Sprintf("%sGID: %s,\n", indentationValues, jmsp.GID))
	b.WriteString(fmt.Sprintf("%sAdditionalParticipants: %s,\n", indentationValues, jmsp.AdditionalParticipants))
	b.WriteString(fmt.Sprintf("%sGIDForParticipationCheck: %s,\n", indentationValues, jmsp.GIDForParticipationCheck))
	b.WriteString(fmt.Sprintf("%sJoinMatchmakeSessionOption: %s,\n", indentationValues, jmsp.JoinMatchmakeSessionOption))
	b.WriteString(fmt.Sprintf("%sJoinMatchmakeSessionBehavior: %s,\n", indentationValues, jmsp.JoinMatchmakeSessionBehavior))
	b.WriteString(fmt.Sprintf("%sStrUserPassword: %s,\n", indentationValues, jmsp.StrUserPassword))
	b.WriteString(fmt.Sprintf("%sStrSystemPassword: %s,\n", indentationValues, jmsp.StrSystemPassword))
	b.WriteString(fmt.Sprintf("%sJoinMessage: %s,\n", indentationValues, jmsp.JoinMessage))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %s,\n", indentationValues, jmsp.ParticipationCount))
	b.WriteString(fmt.Sprintf("%sExtraParticipants: %s,\n", indentationValues, jmsp.ExtraParticipants))
	b.WriteString(fmt.Sprintf("%sBlockListParam: %s,\n", indentationValues, jmsp.BlockListParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewJoinMatchmakeSessionParam returns a new JoinMatchmakeSessionParam
func NewJoinMatchmakeSessionParam() *JoinMatchmakeSessionParam {
	jmsp := &JoinMatchmakeSessionParam{
		GID:                          types.NewPrimitiveU32(0),
		AdditionalParticipants:       types.NewList[*types.PID](),
		GIDForParticipationCheck:     types.NewPrimitiveU32(0),
		JoinMatchmakeSessionOption:   types.NewPrimitiveU32(0),
		JoinMatchmakeSessionBehavior: types.NewPrimitiveU8(0),
		StrUserPassword:              types.NewString(""),
		StrSystemPassword:            types.NewString(""),
		JoinMessage:                  types.NewString(""),
		ParticipationCount:           types.NewPrimitiveU16(0),
		ExtraParticipants:            types.NewPrimitiveU16(0),
		BlockListParam:               NewMatchmakeBlockListParam(),
	}

	jmsp.AdditionalParticipants.Type = types.NewPID(0)

	return jmsp
}
