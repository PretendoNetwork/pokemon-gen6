// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// AutoMatchmakeParam is a type within the Matchmaking protocol
type AutoMatchmakeParam struct {
	types.Structure
	SourceMatchmakeSession   *MatchmakeSession
	AdditionalParticipants   *types.List[*types.PID]
	GIDForParticipationCheck *types.PrimitiveU32
	AutoMatchmakeOption      *types.PrimitiveU32
	JoinMessage              *types.String
	ParticipationCount       *types.PrimitiveU16
	LstSearchCriteria        *types.List[*MatchmakeSessionSearchCriteria]
	TargetGIDs               *types.List[*types.PrimitiveU32]
}

// WriteTo writes the AutoMatchmakeParam to the given writable
func (amp *AutoMatchmakeParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	amp.SourceMatchmakeSession.WriteTo(contentWritable)
	amp.AdditionalParticipants.WriteTo(contentWritable)
	amp.GIDForParticipationCheck.WriteTo(contentWritable)
	amp.AutoMatchmakeOption.WriteTo(contentWritable)
	amp.JoinMessage.WriteTo(contentWritable)
	amp.ParticipationCount.WriteTo(contentWritable)
	amp.LstSearchCriteria.WriteTo(contentWritable)
	amp.TargetGIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	amp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the AutoMatchmakeParam from the given readable
func (amp *AutoMatchmakeParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = amp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam header. %s", err.Error())
	}

	err = amp.SourceMatchmakeSession.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.SourceMatchmakeSession. %s", err.Error())
	}

	err = amp.AdditionalParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.AdditionalParticipants. %s", err.Error())
	}

	err = amp.GIDForParticipationCheck.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.GIDForParticipationCheck. %s", err.Error())
	}

	err = amp.AutoMatchmakeOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.AutoMatchmakeOption. %s", err.Error())
	}

	err = amp.JoinMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.JoinMessage. %s", err.Error())
	}

	err = amp.ParticipationCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.ParticipationCount. %s", err.Error())
	}

	err = amp.LstSearchCriteria.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.LstSearchCriteria. %s", err.Error())
	}

	err = amp.TargetGIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.TargetGIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of AutoMatchmakeParam
func (amp *AutoMatchmakeParam) Copy() types.RVType {
	copied := NewAutoMatchmakeParam()

	copied.StructureVersion = amp.StructureVersion
	copied.SourceMatchmakeSession = amp.SourceMatchmakeSession.Copy().(*MatchmakeSession)
	copied.AdditionalParticipants = amp.AdditionalParticipants.Copy().(*types.List[*types.PID])
	copied.GIDForParticipationCheck = amp.GIDForParticipationCheck.Copy().(*types.PrimitiveU32)
	copied.AutoMatchmakeOption = amp.AutoMatchmakeOption.Copy().(*types.PrimitiveU32)
	copied.JoinMessage = amp.JoinMessage.Copy().(*types.String)
	copied.ParticipationCount = amp.ParticipationCount.Copy().(*types.PrimitiveU16)
	copied.LstSearchCriteria = amp.LstSearchCriteria.Copy().(*types.List[*MatchmakeSessionSearchCriteria])
	copied.TargetGIDs = amp.TargetGIDs.Copy().(*types.List[*types.PrimitiveU32])

	return copied
}

// Equals checks if the given AutoMatchmakeParam contains the same data as the current AutoMatchmakeParam
func (amp *AutoMatchmakeParam) Equals(o types.RVType) bool {
	if _, ok := o.(*AutoMatchmakeParam); !ok {
		return false
	}

	other := o.(*AutoMatchmakeParam)

	if amp.StructureVersion != other.StructureVersion {
		return false
	}

	if !amp.SourceMatchmakeSession.Equals(other.SourceMatchmakeSession) {
		return false
	}

	if !amp.AdditionalParticipants.Equals(other.AdditionalParticipants) {
		return false
	}

	if !amp.GIDForParticipationCheck.Equals(other.GIDForParticipationCheck) {
		return false
	}

	if !amp.AutoMatchmakeOption.Equals(other.AutoMatchmakeOption) {
		return false
	}

	if !amp.JoinMessage.Equals(other.JoinMessage) {
		return false
	}

	if !amp.ParticipationCount.Equals(other.ParticipationCount) {
		return false
	}

	if !amp.LstSearchCriteria.Equals(other.LstSearchCriteria) {
		return false
	}

	return amp.TargetGIDs.Equals(other.TargetGIDs)
}

// String returns the string representation of the AutoMatchmakeParam
func (amp *AutoMatchmakeParam) String() string {
	return amp.FormatToString(0)
}

// FormatToString pretty-prints the AutoMatchmakeParam using the provided indentation level
func (amp *AutoMatchmakeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("AutoMatchmakeParam{\n")
	b.WriteString(fmt.Sprintf("%sSourceMatchmakeSession: %s,\n", indentationValues, amp.SourceMatchmakeSession.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sAdditionalParticipants: %s,\n", indentationValues, amp.AdditionalParticipants))
	b.WriteString(fmt.Sprintf("%sGIDForParticipationCheck: %s,\n", indentationValues, amp.GIDForParticipationCheck))
	b.WriteString(fmt.Sprintf("%sAutoMatchmakeOption: %s,\n", indentationValues, amp.AutoMatchmakeOption))
	b.WriteString(fmt.Sprintf("%sJoinMessage: %s,\n", indentationValues, amp.JoinMessage))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %s,\n", indentationValues, amp.ParticipationCount))
	b.WriteString(fmt.Sprintf("%sLstSearchCriteria: %s,\n", indentationValues, amp.LstSearchCriteria))
	b.WriteString(fmt.Sprintf("%sTargetGIDs: %s,\n", indentationValues, amp.TargetGIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAutoMatchmakeParam returns a new AutoMatchmakeParam
func NewAutoMatchmakeParam() *AutoMatchmakeParam {
	amp := &AutoMatchmakeParam{
		SourceMatchmakeSession:   NewMatchmakeSession(),
		AdditionalParticipants:   types.NewList[*types.PID](),
		GIDForParticipationCheck: types.NewPrimitiveU32(0),
		AutoMatchmakeOption:      types.NewPrimitiveU32(0),
		JoinMessage:              types.NewString(""),
		ParticipationCount:       types.NewPrimitiveU16(0),
		LstSearchCriteria:        types.NewList[*MatchmakeSessionSearchCriteria](),
		TargetGIDs:               types.NewList[*types.PrimitiveU32](),
	}

	amp.AdditionalParticipants.Type = types.NewPID(0)
	amp.LstSearchCriteria.Type = NewMatchmakeSessionSearchCriteria()
	amp.TargetGIDs.Type = types.NewPrimitiveU32(0)

	return amp
}
