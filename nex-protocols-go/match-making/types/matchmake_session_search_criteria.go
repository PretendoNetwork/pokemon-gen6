// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeSessionSearchCriteria is a type within the Matchmaking protocol
type MatchmakeSessionSearchCriteria struct {
	types.Structure
	Attribs                  *types.List[*types.String]
	GameMode                 *types.String
	MinParticipants          *types.String
	MaxParticipants          *types.String
	MatchmakeSystemType      *types.String
	VacantOnly               *types.PrimitiveBool
	ExcludeLocked            *types.PrimitiveBool
	ExcludeNonHostPID        *types.PrimitiveBool
	SelectionMethod          *types.PrimitiveU32  // * NEX v3.0.0
	VacantParticipants       *types.PrimitiveU16  // * NEX v3.4.0
	MatchmakeParam           *MatchmakeParam      // * NEX v3.6.0
	ExcludeUserPasswordSet   *types.PrimitiveBool // * NEX v3.7.0
	ExcludeSystemPasswordSet *types.PrimitiveBool // * NEX v3.7.0
	ReferGID                 *types.PrimitiveU32  // * NEX v3.8.0
	CodeWord                 *types.String        // * NEX v4.0.0
	ResultRange              *types.ResultRange   // * NEX v4.0.0
}

// WriteTo writes the MatchmakeSessionSearchCriteria to the given writable
func (mssc *MatchmakeSessionSearchCriteria) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.MatchMaking

	contentWritable := writable.CopyNew()

	mssc.Attribs.WriteTo(contentWritable)
	mssc.GameMode.WriteTo(contentWritable)
	mssc.MinParticipants.WriteTo(contentWritable)
	mssc.MaxParticipants.WriteTo(contentWritable)
	mssc.MatchmakeSystemType.WriteTo(contentWritable)
	mssc.VacantOnly.WriteTo(contentWritable)
	mssc.ExcludeLocked.WriteTo(contentWritable)
	mssc.ExcludeNonHostPID.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("3.0.0") {
		mssc.SelectionMethod.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.4.0") {
		mssc.VacantParticipants.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.6.0") {
		mssc.MatchmakeParam.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.7.0") {
		mssc.ExcludeUserPasswordSet.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.7.0") {
		mssc.ExcludeSystemPasswordSet.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		mssc.ReferGID.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		mssc.CodeWord.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		mssc.ResultRange.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	mssc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeSessionSearchCriteria from the given readable
func (mssc *MatchmakeSessionSearchCriteria) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.MatchMaking

	var err error

	err = mssc.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria header. %s", err.Error())
	}

	err = mssc.Attribs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.Attribs. %s", err.Error())
	}

	err = mssc.GameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.GameMode. %s", err.Error())
	}

	err = mssc.MinParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MinParticipants. %s", err.Error())
	}

	err = mssc.MaxParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MaxParticipants. %s", err.Error())
	}

	err = mssc.MatchmakeSystemType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MatchmakeSystemType. %s", err.Error())
	}

	err = mssc.VacantOnly.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.VacantOnly. %s", err.Error())
	}

	err = mssc.ExcludeLocked.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeLocked. %s", err.Error())
	}

	err = mssc.ExcludeNonHostPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeNonHostPID. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		err = mssc.SelectionMethod.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.SelectionMethod. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.4.0") {
		err = mssc.VacantParticipants.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.VacantParticipants. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.6.0") {
		err = mssc.MatchmakeParam.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MatchmakeParam. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.7.0") {
		err = mssc.ExcludeUserPasswordSet.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeUserPasswordSet. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.7.0") {
		err = mssc.ExcludeSystemPasswordSet.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeSystemPasswordSet. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		err = mssc.ReferGID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ReferGID. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		err = mssc.CodeWord.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.CodeWord. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		err = mssc.ResultRange.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ResultRange. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeSessionSearchCriteria
func (mssc *MatchmakeSessionSearchCriteria) Copy() types.RVType {
	copied := NewMatchmakeSessionSearchCriteria()

	copied.StructureVersion = mssc.StructureVersion
	copied.Attribs = mssc.Attribs.Copy().(*types.List[*types.String])
	copied.GameMode = mssc.GameMode.Copy().(*types.String)
	copied.MinParticipants = mssc.MinParticipants.Copy().(*types.String)
	copied.MaxParticipants = mssc.MaxParticipants.Copy().(*types.String)
	copied.MatchmakeSystemType = mssc.MatchmakeSystemType.Copy().(*types.String)
	copied.VacantOnly = mssc.VacantOnly.Copy().(*types.PrimitiveBool)
	copied.ExcludeLocked = mssc.ExcludeLocked.Copy().(*types.PrimitiveBool)
	copied.ExcludeNonHostPID = mssc.ExcludeNonHostPID.Copy().(*types.PrimitiveBool)
	copied.SelectionMethod = mssc.SelectionMethod.Copy().(*types.PrimitiveU32)
	copied.VacantParticipants = mssc.VacantParticipants.Copy().(*types.PrimitiveU16)
	copied.MatchmakeParam = mssc.MatchmakeParam.Copy().(*MatchmakeParam)
	copied.ExcludeUserPasswordSet = mssc.ExcludeUserPasswordSet.Copy().(*types.PrimitiveBool)
	copied.ExcludeSystemPasswordSet = mssc.ExcludeSystemPasswordSet.Copy().(*types.PrimitiveBool)
	copied.ReferGID = mssc.ReferGID.Copy().(*types.PrimitiveU32)
	copied.CodeWord = mssc.CodeWord.Copy().(*types.String)
	copied.ResultRange = mssc.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the given MatchmakeSessionSearchCriteria contains the same data as the current MatchmakeSessionSearchCriteria
func (mssc *MatchmakeSessionSearchCriteria) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeSessionSearchCriteria); !ok {
		return false
	}

	other := o.(*MatchmakeSessionSearchCriteria)

	if mssc.StructureVersion != other.StructureVersion {
		return false
	}

	if !mssc.Attribs.Equals(other.Attribs) {
		return false
	}

	if !mssc.GameMode.Equals(other.GameMode) {
		return false
	}

	if !mssc.MinParticipants.Equals(other.MinParticipants) {
		return false
	}

	if !mssc.MaxParticipants.Equals(other.MaxParticipants) {
		return false
	}

	if !mssc.MatchmakeSystemType.Equals(other.MatchmakeSystemType) {
		return false
	}

	if !mssc.VacantOnly.Equals(other.VacantOnly) {
		return false
	}

	if !mssc.ExcludeLocked.Equals(other.ExcludeLocked) {
		return false
	}

	if !mssc.ExcludeNonHostPID.Equals(other.ExcludeNonHostPID) {
		return false
	}

	if !mssc.SelectionMethod.Equals(other.SelectionMethod) {
		return false
	}

	if !mssc.VacantParticipants.Equals(other.VacantParticipants) {
		return false
	}

	if !mssc.MatchmakeParam.Equals(other.MatchmakeParam) {
		return false
	}

	if !mssc.ExcludeUserPasswordSet.Equals(other.ExcludeUserPasswordSet) {
		return false
	}

	if !mssc.ExcludeSystemPasswordSet.Equals(other.ExcludeSystemPasswordSet) {
		return false
	}

	if !mssc.ReferGID.Equals(other.ReferGID) {
		return false
	}

	if !mssc.CodeWord.Equals(other.CodeWord) {
		return false
	}

	return mssc.ResultRange.Equals(other.ResultRange)
}

// String returns the string representation of the MatchmakeSessionSearchCriteria
func (mssc *MatchmakeSessionSearchCriteria) String() string {
	return mssc.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeSessionSearchCriteria using the provided indentation level
func (mssc *MatchmakeSessionSearchCriteria) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeSessionSearchCriteria{\n")
	b.WriteString(fmt.Sprintf("%sAttribs: %s,\n", indentationValues, mssc.Attribs))
	b.WriteString(fmt.Sprintf("%sGameMode: %s,\n", indentationValues, mssc.GameMode))
	b.WriteString(fmt.Sprintf("%sMinParticipants: %s,\n", indentationValues, mssc.MinParticipants))
	b.WriteString(fmt.Sprintf("%sMaxParticipants: %s,\n", indentationValues, mssc.MaxParticipants))
	b.WriteString(fmt.Sprintf("%sMatchmakeSystemType: %s,\n", indentationValues, mssc.MatchmakeSystemType))
	b.WriteString(fmt.Sprintf("%sVacantOnly: %s,\n", indentationValues, mssc.VacantOnly))
	b.WriteString(fmt.Sprintf("%sExcludeLocked: %s,\n", indentationValues, mssc.ExcludeLocked))
	b.WriteString(fmt.Sprintf("%sExcludeNonHostPID: %s,\n", indentationValues, mssc.ExcludeNonHostPID))
	b.WriteString(fmt.Sprintf("%sSelectionMethod: %s,\n", indentationValues, mssc.SelectionMethod))
	b.WriteString(fmt.Sprintf("%sVacantParticipants: %s,\n", indentationValues, mssc.VacantParticipants))
	b.WriteString(fmt.Sprintf("%sMatchmakeParam: %s,\n", indentationValues, mssc.MatchmakeParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sExcludeUserPasswordSet: %s,\n", indentationValues, mssc.ExcludeUserPasswordSet))
	b.WriteString(fmt.Sprintf("%sExcludeSystemPasswordSet: %s,\n", indentationValues, mssc.ExcludeSystemPasswordSet))
	b.WriteString(fmt.Sprintf("%sReferGID: %s,\n", indentationValues, mssc.ReferGID))
	b.WriteString(fmt.Sprintf("%sCodeWord: %s,\n", indentationValues, mssc.CodeWord))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, mssc.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeSessionSearchCriteria returns a new MatchmakeSessionSearchCriteria
func NewMatchmakeSessionSearchCriteria() *MatchmakeSessionSearchCriteria {
	mssc := &MatchmakeSessionSearchCriteria{
		Attribs:                  types.NewList[*types.String](),
		GameMode:                 types.NewString(""),
		MinParticipants:          types.NewString(""),
		MaxParticipants:          types.NewString(""),
		MatchmakeSystemType:      types.NewString(""),
		VacantOnly:               types.NewPrimitiveBool(false),
		ExcludeLocked:            types.NewPrimitiveBool(false),
		ExcludeNonHostPID:        types.NewPrimitiveBool(false),
		SelectionMethod:          types.NewPrimitiveU32(0),
		VacantParticipants:       types.NewPrimitiveU16(0),
		MatchmakeParam:           NewMatchmakeParam(),
		ExcludeUserPasswordSet:   types.NewPrimitiveBool(false),
		ExcludeSystemPasswordSet: types.NewPrimitiveBool(false),
		ReferGID:                 types.NewPrimitiveU32(0),
		CodeWord:                 types.NewString(""),
		ResultRange:              types.NewResultRange(),
	}

	mssc.Attribs.Type = types.NewString("")

	return mssc
}
